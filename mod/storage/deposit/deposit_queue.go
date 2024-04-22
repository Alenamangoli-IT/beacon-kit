// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package deposit

import (
	"context"
	"sync"

	sdkcollections "cosmossdk.io/collections"
	sdkcodec "cosmossdk.io/collections/codec"
	"github.com/berachain/beacon-kit/mod/primitives"
)

// Queue is a simple queue implementation that uses a map and two sequences.
// TODO: Check atomicity of write operations.
type Queue struct {
	// container is a map that holds the queue elements.
	container sdkcollections.Map[uint64, *primitives.Deposit]
	// headSeq is a sequence that points to the head of the queue.
	headSeq sdkcollections.Sequence // inclusive
	// tailSeq is a sequence that points to the tail of the queue.
	tailSeq sdkcollections.Sequence // exclusive
	// mu is a mutex that protects the queue.
	mu sync.RWMutex
}

// NewQueue creates a new queue with the provided prefix and name.
func NewQueue(
	schema *sdkcollections.SchemaBuilder, name string,
	valueCodec sdkcodec.ValueCodec[*primitives.Deposit],
) *Queue {
	var (
		queueName   = name + "_queue"
		headSeqName = name + "_head"
		tailSeqName = name + "_tail"
	)
	return &Queue{
		container: sdkcollections.NewMap(
			schema,
			sdkcollections.NewPrefix(queueName),
			queueName, sdkcollections.Uint64Key, valueCodec,
		),
		headSeq: sdkcollections.NewSequence(
			schema, sdkcollections.NewPrefix(headSeqName), headSeqName),
		tailSeq: sdkcollections.NewSequence(
			schema, sdkcollections.NewPrefix(tailSeqName), tailSeqName),
	}
}

func (q *Queue) Init(ctx context.Context) error {
	if err := q.headSeq.Set(ctx, 0); err != nil {
		return err
	}

	// The tail sequence should be set to the head sequence.
	return q.tailSeq.Set(ctx, 0)
}

// Peek wraps the peek method with a read lock.
func (q *Queue) Peek(ctx context.Context) (*primitives.Deposit, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.UnsafePeek(ctx)
}

// UnsafePeek returns the top element of the queue without removing it.
// It is unsafe to call this method without acquiring the read lock.
func (q *Queue) UnsafePeek(
	ctx context.Context,
) (*primitives.Deposit, error) {
	var (
		v                *primitives.Deposit
		headIdx, tailIdx uint64
		err              error
	)
	if headIdx, err = q.headSeq.Peek(ctx); err != nil {
		return v, err
	} else if tailIdx, err = q.tailSeq.Peek(ctx); err != nil {
		return v, err
	} else if headIdx >= tailIdx {
		return v, sdkcollections.ErrNotFound
	}
	return q.container.Get(ctx, headIdx)
}

// Pop returns the top element of the queue and removes it from the queue.
func (q *Queue) Pop(ctx context.Context) (*primitives.Deposit, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	var (
		v       *primitives.Deposit
		headIdx uint64
		err     error
	)

	if v, err = q.UnsafePeek(ctx); err != nil {
		return v, err
	} else if headIdx, err = q.headSeq.Next(ctx); err != nil {
		return v, err
	}
	err = q.container.Remove(ctx, headIdx)
	return v, err
}

// PeekMulti returns the top n elements of the queue.
func (q *Queue) PeekMulti(
	ctx context.Context,
	n uint64,
) ([]*primitives.Deposit, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	var err error

	headIdx, err := q.headSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}
	tailIdx, err := q.tailSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}
	endIdx := min(tailIdx, headIdx+n)
	ranger := new(sdkcollections.Range[uint64]).
		StartInclusive(headIdx).EndExclusive(endIdx)
	iter, err := q.container.Iterate(ctx, ranger)
	if err != nil {
		return nil, err
	}
	// iter.Values already closes the iterator.
	values, err := iter.Values()
	if err != nil {
		return nil, err
	}

	return values, nil
}

// PopMulti returns the top n elements of the queue and removes them from the
// queue.
func (q *Queue) PopMulti(
	ctx context.Context,
	n uint64,
) ([]*primitives.Deposit, error) {
	if n == 0 {
		return nil, nil
	}
	q.mu.Lock()
	defer q.mu.Unlock()

	var err error

	headIdx, err := q.headSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}
	tailIdx, err := q.tailSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}
	endIdx := min(tailIdx, headIdx+n)
	ranger := new(sdkcollections.Range[uint64]).
		StartInclusive(headIdx).EndExclusive(endIdx)
	iter, err := q.container.Iterate(ctx, ranger)
	if err != nil {
		return nil, err
	}

	// iter.Values already closes the iterator.
	values, err := iter.Values()
	if err != nil {
		return nil, err
	}

	// Clear the range (in batch) after the iteration is done.
	err = q.container.Clear(ctx, ranger)
	if err != nil {
		return nil, err
	}
	err = q.headSeq.Set(ctx, endIdx)
	if err != nil {
		return nil, err
	}
	return values, nil
}

// Push adds a new element to the queue.
func (q *Queue) Push(
	ctx context.Context,
	value *primitives.Deposit,
) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if err := q.container.Set(ctx, value.Index, value); err != nil {
		return err
	}
	return q.tailSeq.Set(ctx, value.Index)
}

// PushMulti adds multiple new elements to the queue.
func (q *Queue) PushMulti(
	ctx context.Context,
	values []*primitives.Deposit,
) error {
	if len(values) == 0 {
		return nil
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	var idx uint64

	for _, value := range values {
		if err := q.container.Set(ctx, value.Index, value); err != nil {
			return err
		}
		idx = value.Index
	}

	return q.tailSeq.Set(ctx, idx)
}

// Len returns the length of the queue.
func (q *Queue) Len(ctx context.Context) (uint64, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	var (
		headIdx, tailIdx uint64
		err              error
	)

	if headIdx, err = q.headSeq.Peek(ctx); err != nil {
		return 0, err
	} else if tailIdx, err = q.tailSeq.Peek(ctx); err != nil {
		return 0, err
	}
	return tailIdx - headIdx, nil
}

// Container returns the underlying map container of the queue.
func (q *Queue) Container() sdkcollections.Map[
	uint64, *primitives.Deposit,
] {
	return q.container
}

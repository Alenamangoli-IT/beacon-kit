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

package runtime

import (
	"context"

	"github.com/berachain/beacon-kit/mod/core"
	"github.com/berachain/beacon-kit/mod/core/state"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/consensus"
)

type DepositStore interface {
	ExpectedDeposits(
		numView uint64,
	) ([]*consensus.Deposit, error)
	EnqueueDeposit(deposit *consensus.Deposit) error
	// EnqueueDeposits pushes multiple deposits to the queue.
	EnqueueDeposits(deposits []*consensus.Deposit) error

	// DequeueDeposits returns the first numDequeue deposits in the queue.
	DequeueDeposits(
		numDequeue uint64,
	) ([]*consensus.Deposit, error)

	// PruneToIndex removes all deposits up to the given index.
	PruneToIndex(
		index uint64,
	) error
}

// BeaconStorageBackend is an interface that provides the
// beacon state to the runtime.
type BeaconStorageBackend[
	BlobSidecarsT any,
	DepositStoreT DepositStore,
	ReadOnlyBeaconBlockT any,
] interface {
	AvailabilityStore(
		ctx context.Context,
	) core.AvailabilityStore[ReadOnlyBeaconBlockT, BlobSidecarsT]
	BeaconState(ctx context.Context) state.BeaconState
	DepositStore(ctx context.Context) DepositStoreT
}

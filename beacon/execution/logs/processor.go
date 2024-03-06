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

package logs

import (
	"context"

	"cosmossdk.io/errors"
	ethcommon "github.com/ethereum/go-ethereum/common"
	engineclient "github.com/itsdevbear/bolaris/engine/client"
)

const (
	// DefaultBatchSize is the default batch size for processing logs.
	DefaultBatchSize = 1000
)

type Processor struct {
	// engine gives the access to the Engine API
	// of the execution client.
	engine engineclient.Caller

	// factory is for creating objects from Ethereum logs.
	factory LogFactory

	sigToCache map[ethcommon.Hash]LogCache
}

// NewProcessor creates a new Processor.
func NewProcessor(opts ...Option[Processor]) (*Processor, error) {
	p := &Processor{}
	for _, opt := range opts {
		if err := opt(p); err != nil {
			return nil, err
		}
	}
	return p, nil
}

// ProcessEth1Block processes the logs in the given Ethereum block.
func (p *Processor) ProcessEth1Block(
	ctx context.Context,
	blockHash ethcommon.Hash,
) error {
	header, err := p.engine.HeaderByHash(ctx, blockHash)
	if err != nil {
		return errors.Wrapf(err, "failed to get block header")
	}
	blockNum := header.Number.Uint64()
	return p.ProcessBlocks(ctx, blockNum, blockNum)
}

// ProcessBlocks processes the logs in the range
// from fromBlock (inclusive) to toBlock (inclusive).
func (p *Processor) ProcessBlocks(
	ctx context.Context,
	fromBlock uint64,
	toBlock uint64,
) error {
	// Get the registered addresses for the logs.
	registeredAddresses := p.factory.GetRegisteredAddresses()

	curr := fromBlock
	for curr <= toBlock {
		// Process the logs in batches.
		end, err := p.processBlocksInBatch(
			ctx,
			curr,
			toBlock,
			DefaultBatchSize,
			registeredAddresses,
		)
		if err != nil {
			return errors.Wrapf(err, "failed to process logs in batch")
		}
		curr = end + 1
	}
	return nil
}

// processBlocksInBatch processes the logs in the range
// from fromBlock (inclusive)
// to min(fromBlock + batchSize - 1, toBlock) (inclusive).
func (p *Processor) processBlocksInBatch(
	ctx context.Context,
	fromBlock uint64,
	toBlock uint64,
	batchSize uint64,
	registeredAddresses []ethcommon.Address,
) (uint64, error) {
	// Gather all the logs corresponding to
	// the registered addresses in the given range.
	start := fromBlock
	end := start + batchSize - 1
	if end > toBlock {
		end = toBlock
	}

	logs, err := p.engine.GetLogs(
		ctx,
		start,
		end,
		registeredAddresses,
	)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to get logs")
	}

	containers, err := p.factory.ProcessLogs(logs)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to process logs")
	}
	for _, container := range containers {
		if cache, ok := p.sigToCache[container.Signature()]; ok {
			err = cache.Insert(container)
			if err != nil {
				return 0, errors.Wrapf(err, "failed to insert log into cache")
			}
		}
	}

	return end, nil
}

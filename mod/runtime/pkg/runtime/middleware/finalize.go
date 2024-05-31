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

package middleware

import (
	"context"
	"encoding/json"
	"time"

	appmodulev2 "cosmossdk.io/core/appmodule/v2"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/genesis"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/ssz"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/encoding"
	cometabci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sourcegraph/conc/iter"
)

// FinalizeBlockMiddleware is a struct that encapsulates the necessary
// components to handle
// the proposal processes.
type FinalizeBlockMiddleware[
	BeaconBlockT interface {
		ssz.Marshallable
		NewFromSSZ([]byte, uint32) (BeaconBlockT, error)
	},
	BeaconStateT any,
	BlobsSidecarsT ssz.Marshallable,
] struct {
	// chainSpec is the chain specification.
	chainSpec primitives.ChainSpec
	// chainService represents the blockchain service.
	chainService BlockchainService[BeaconBlockT, BlobsSidecarsT]
	// metrics is the metrics for the middleware.
	metrics *finalizeMiddlewareMetrics
	// valUpdates caches the validator updates as they are produced.
	valUpdates []*transition.ValidatorUpdate
}

// NewFinalizeBlockMiddleware creates a new instance of the Handler struct.
func NewFinalizeBlockMiddleware[
	BeaconBlockT interface {
		ssz.Marshallable
		NewFromSSZ([]byte, uint32) (BeaconBlockT, error)
	},
	BeaconStateT any, BlobsSidecarsT ssz.Marshallable,
](
	chainSpec primitives.ChainSpec,
	chainService BlockchainService[BeaconBlockT, BlobsSidecarsT],
	telemetrySink TelemetrySink,
) *FinalizeBlockMiddleware[BeaconBlockT, BeaconStateT, BlobsSidecarsT] {
	// This is just for nilaway, TODO: remove later.
	if chainService == nil {
		panic("chain service is nil")
	}

	return &FinalizeBlockMiddleware[BeaconBlockT, BeaconStateT, BlobsSidecarsT]{
		chainSpec:    chainSpec,
		chainService: chainService,
		metrics:      newFinalizeMiddlewareMetrics(telemetrySink),
	}
}

// InitGenesis is called by the base app to initialize the state of the.
func (h *FinalizeBlockMiddleware[
	BeaconBlockT, BeaconStateT, BlobsSidecarsT,
]) InitGenesis(
	ctx context.Context,
	bz []byte,
) ([]appmodulev2.ValidatorUpdate, error) {
	data := new(
		genesis.Genesis[*types.Deposit, *types.ExecutionPayloadHeaderDeneb],
	)
	if err := json.Unmarshal(bz, data); err != nil {
		return nil, err
	}
	updates, err := h.chainService.ProcessGenesisData(
		ctx,
		data,
	)
	if err != nil {
		return nil, err
	}

	// Convert updates into the Cosmos SDK format.
	return iter.MapErr(updates, convertValidatorUpdate)
}

// PreBlock is called by the base app before the block is finalized. It
// is responsible for aggregating oracle data from each validator and writing
// the oracle data to the store.
func (h *FinalizeBlockMiddleware[
	BeaconBlockT, BeaconStateT, BlobsSidecarsT,
]) PreBlock(
	ctx sdk.Context, req *cometabci.FinalizeBlockRequest,
) error {
	startTime := time.Now()
	defer h.metrics.measureEndBlockDuration(startTime)

	blk, blobs, err := encoding.
		ExtractBlobsAndBlockFromRequest[BeaconBlockT, BlobsSidecarsT](req,
		BeaconBlockTxIndex,
		BlobSidecarsTxIndex,
		h.chainSpec.ActiveForkVersionForSlot(
			math.Slot(req.Height),
		))
	if err != nil {
		return err
	}

	// Process the state transition and produce the required delta from
	// the sync committee.
	h.valUpdates, err = h.chainService.ProcessBlockAndBlobs(
		ctx, blk, blobs, req.SyncingToHeight == req.Height,
	)
	return err
}

// EndBlock returns the validator set updates from the beacon state.
func (h FinalizeBlockMiddleware[
	BeaconBlockT, BeaconStateT, BlobsSidecarsT,
]) EndBlock(
	context.Context,
) ([]appmodulev2.ValidatorUpdate, error) {
	return iter.MapErr(h.valUpdates, convertValidatorUpdate)
}

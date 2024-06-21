// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package blockchain

import (
	"context"
	"sync"

	asynctypes "github.com/berachain/beacon-kit/mod/async/pkg/types"
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
)

// Service is the blockchain service.
type Service[
	AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT],
	BeaconBlockT BeaconBlock[BeaconBlockBodyT, ExecutionPayloadT],
	BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT],
	BeaconBlockHeaderT BeaconBlockHeader,
	BeaconStateT ReadOnlyBeaconState[
		BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT,
	],
	BlobSidecarsT BlobSidecars,
	DepositT any,
	ExecutionPayloadT ExecutionPayload,
	ExecutionPayloadHeaderT ExecutionPayloadHeader,
	GenesisT Genesis[DepositT, ExecutionPayloadHeaderT],
] struct {
	// sb represents the backend storage for beacon states and associated
	// sidecars.
	sb StorageBackend[
		AvailabilityStoreT,
		BeaconBlockBodyT,
		BeaconStateT,
		BlobSidecarsT,
	]
	// logger is used for logging messages in the service.
	logger log.Logger[any]
	// cs holds the chain specifications.
	cs common.ChainSpec
	// ee is the execution engine responsible for processing execution payloads.
	//nolint:lll // annoying linter.
	ee ExecutionEngine[*engineprimitives.PayloadAttributes[*engineprimitives.Withdrawal]]
	// lb is a local builder for constructing new beacon states.
	lb LocalBuilder[BeaconStateT]
	// bp is the blob processor for processing incoming blobs.
	bp BlobProcessor[
		AvailabilityStoreT, BeaconBlockBodyT, BlobSidecarsT, ExecutionPayloadT,
	]
	// sp is the state processor for beacon blocks and states.
	sp StateProcessor[
		BeaconBlockT,
		BeaconStateT,
		BlobSidecarsT,
		*transition.Context,
		DepositT,
		ExecutionPayloadHeaderT,
	]
	// metrics is the metrics for the service.
	metrics *chainMetrics
	// blockFeed is the event feed for new blocks.
	blockFeed EventFeed[*asynctypes.Event[BeaconBlockT]]
	// optimisticPayloadBuilds is a flag used when the optimistic payload
	// builder is enabled.
	optimisticPayloadBuilds bool
	// forceStartupSyncOnce is used to force a sync of the startup head.
	forceStartupSyncOnce *sync.Once
}

// NewService creates a new validator service.
func NewService[
	AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT],
	BeaconBlockT BeaconBlock[BeaconBlockBodyT, ExecutionPayloadT],
	BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT],
	BeaconBlockHeaderT BeaconBlockHeader,
	BeaconStateT ReadOnlyBeaconState[
		BeaconStateT, BeaconBlockHeaderT,
		ExecutionPayloadHeaderT,
	],
	BlobSidecarsT BlobSidecars,
	DepositT any,
	ExecutionPayloadT ExecutionPayload,
	ExecutionPayloadHeaderT ExecutionPayloadHeader,
	GenesisT Genesis[DepositT, ExecutionPayloadHeaderT],
](
	sb StorageBackend[
		AvailabilityStoreT,
		BeaconBlockBodyT,
		BeaconStateT,
		BlobSidecarsT,
	],
	logger log.Logger[any],
	cs common.ChainSpec,
	//nolint:lll // annoying formatter.
	ee ExecutionEngine[*engineprimitives.PayloadAttributes[*engineprimitives.Withdrawal]],
	lb LocalBuilder[BeaconStateT],
	bp BlobProcessor[
		AvailabilityStoreT, BeaconBlockBodyT, BlobSidecarsT, ExecutionPayloadT,
	],
	sp StateProcessor[
		BeaconBlockT,
		BeaconStateT,
		BlobSidecarsT,
		*transition.Context,
		DepositT,
		ExecutionPayloadHeaderT,
	],
	ts TelemetrySink,
	blockFeed EventFeed[*asynctypes.Event[BeaconBlockT]],
	optimisticPayloadBuilds bool,
) *Service[
	AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT, BeaconBlockHeaderT,
	BeaconStateT, BlobSidecarsT, DepositT, ExecutionPayloadT,
	ExecutionPayloadHeaderT, GenesisT,
] {
	return &Service[
		AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT, BeaconBlockHeaderT,
		BeaconStateT, BlobSidecarsT, DepositT, ExecutionPayloadT,
		ExecutionPayloadHeaderT, GenesisT,
	]{
		sb:                      sb,
		logger:                  logger,
		cs:                      cs,
		ee:                      ee,
		lb:                      lb,
		bp:                      bp,
		sp:                      sp,
		metrics:                 newChainMetrics(ts),
		blockFeed:               blockFeed,
		optimisticPayloadBuilds: optimisticPayloadBuilds,
		forceStartupSyncOnce:    new(sync.Once),
	}
}

// Name returns the name of the service.
func (s *Service[
	AvailabilityStoreT,
	BeaconBlockT,
	BeaconBlockBodyT,
	BeaconBlockHeaderT,
	BeaconStateT,
	BlobSidecarsT,
	DepositT,
	ExecutionPayloadT,
	ExecutionPayloadHeaderT,
	GenesisT,
]) Name() string {
	return "blockchain"
}

func (s *Service[
	AvailabilityStoreT,
	BeaconBlockT,
	BeaconBlockBodyT,
	BeaconBlockHeaderT,
	BeaconStateT,
	BlobSidecarsT,
	DepositT,
	ExecutionPayloadT,
	ExecutionPayloadHeaderT,
	GenesisT,
]) Start(
	context.Context,
) error {
	return nil
}

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

package cometbft

import (
	"context"

	"cosmossdk.io/core/transaction"
	"cosmossdk.io/server/v2/cometbft/handlers"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	cmtabci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/sourcegraph/conc/iter"
)

// Consensus is used to decouple the Comet consensus engine from the Cosmos SDK.
// Right now, it is very coupled to the sdk base app and we will
// eventually fully decouple this.
type ConsensusEngine[T transaction.Tx, ValidatorUpdateT any] struct {
	Middleware
	txCodec transaction.Codec[T]
}

// NewConsensusEngine returns a new consensus middleware.
func NewConsensusEngine[T transaction.Tx, ValidatorUpdateT any](
	txCodec transaction.Codec[T],
	m Middleware,
) *ConsensusEngine[T, ValidatorUpdateT] {
	return &ConsensusEngine[T, ValidatorUpdateT]{
		Middleware: m,
		txCodec:    txCodec,
	}
}

func (c *ConsensusEngine[T, ValidatorUpdateT]) InitGenesis(
	ctx context.Context,
	genesisBz []byte,
) ([]ValidatorUpdateT, error) {
	updates, err := c.Middleware.InitGenesis(ctx, genesisBz)
	if err != nil {
		return nil, err
	}
	// Convert updates into the Cosmos SDK format.
	return iter.MapErr(updates, convertValidatorUpdate[ValidatorUpdateT])
}

// TODO: Decouple Comet Types
func (c *ConsensusEngine[T, ValidatorUpdateT]) PrepareProposal(
	ctx context.Context,
	_ handlers.AppManager[T],
	txs []T,
	req proto.Message,
) ([]T, error) {
	abciReq, ok := req.(*cmtabci.PrepareProposalRequest)
	if !ok {
		return nil, ErrInvalidPrepareProposalRequest
	}
	slot := math.Slot(abciReq.Height)
	blkBz, sidecarsBz, err := c.Middleware.PrepareProposal(ctx, slot)
	if err != nil {
		return nil, err
	}

	return iter.MapErr([][]byte{blkBz, sidecarsBz}, c.convertTx)
}

// TODO: Decouple Comet Types
func (c *ConsensusEngine[T, ValidatorUpdateT]) ProcessProposal(
	ctx context.Context,
	_ handlers.AppManager[T],
	txs []T,
	req proto.Message,
) error {
	return c.Middleware.ProcessProposal(ctx, req)
}

// TODO: Decouple Comet Types
// TODO: reimplement pre block
// func (c *ConsensusEngine[T, ValidatorUpdateT]) PreBlock(
// 	ctx context.Context,
// ) error {
// 	return c.Middleware.PreBlock(ctx, req)
// }

func (c *ConsensusEngine[T, ValidatorUpdateT]) EndBlock(
	ctx context.Context,
) ([]ValidatorUpdateT, error) {
	updates, err := c.Middleware.EndBlock(ctx)
	if err != nil {
		return nil, err
	}
	return iter.MapErr(updates, convertValidatorUpdate[ValidatorUpdateT])
}

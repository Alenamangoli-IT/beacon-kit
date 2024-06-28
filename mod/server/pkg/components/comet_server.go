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

package components

import (
	"cosmossdk.io/core/transaction"
	"cosmossdk.io/depinject"
	cmtconsensus "github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft"
	nodecomponents "github.com/berachain/beacon-kit/mod/node-core/pkg/components"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/types"
	"github.com/berachain/beacon-kit/mod/server/pkg/components/cometbft"
)

// CometServerInput is the input for the CometServer for the depinject package.
type CometServerInput[T transaction.Tx, ValidatorUpdateT any] struct {
	depinject.In

	TxCodec         *nodecomponents.TxCodec[T]
	ConsensusEngine *cmtconsensus.ConsensusEngine[T, ValidatorUpdateT]
}

// ProvideCometServer is the provider for the CometServer for the depinject
// package.
func ProvideCometServer[
	NodeT types.Node[T], T transaction.Tx, ValidatorUpdateT any,
](
	in CometServerInput[T, ValidatorUpdateT],
) *cometbft.Server[NodeT, T, ValidatorUpdateT] {
	return cometbft.NewServer[NodeT, T, ValidatorUpdateT](
		in.TxCodec,
		in.ConsensusEngine,
	)
}

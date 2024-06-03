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
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
)

func GetPruneRangeFn[
	BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT],
	BeaconBlockT BeaconBlock[DepositT, BeaconBlockBodyT, ExecutionPayloadT],
	BlockEventT BlockEvent[
		DepositT, BeaconBlockBodyT, BeaconBlockT, ExecutionPayloadT,
	],
	DepositT Deposit[DepositT, WithdrawalCredentialsT],
	ExecutionPayloadT interface {
		GetNumber() math.U64
	},
	WithdrawalCredentialsT any,

](cs primitives.ChainSpec) func(BlockEventT) (uint64, uint64) {
	return func(event BlockEventT) (uint64, uint64) {
		blk := event.Block()
		deposits := blk.GetBody().GetDeposits()
		if len(deposits) == 0 {
			return 0, 0
		}
		index := deposits[len(deposits)-1].GetIndex()

		end := min(index, cs.MaxDepositsPerBlock())
		if index < cs.MaxDepositsPerBlock() {
			return 0, end
		}

		return index - cs.MaxDepositsPerBlock(), end
	}
}

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

package beacon

// Collection prefixes.
const (
	depositQueuePrefix                     = "deposit_queue"
	withdrawalQueuePrefix                  = "withdrawal_queue"
	randaoMixPrefix                        = "randao_mix"
	slashingsPrefix                        = "slashings"
	totalSlashingPrefix                    = "total_slashing"
	validatorIndexPrefix                   = "val_idx"
	blockRootsPrefix                       = "block_roots"
	stateRootsPrefix                       = "state_roots"
	validatorByIndexPrefix                 = "val_idx_to_pk"
	validatorPubkeyToIndexPrefix           = "val_pk_to_idx"
	validatorConsAddrToIndexPrefix         = "val_cons_addr_to_idx"
	validatorEffectiveBalanceToIndexPrefix = "val_eff_bal_to_idx"
	latestBeaconBlockHeaderPrefix          = "latest_beacon_block_header"
	slotPrefix                             = "slot"
	balancesPrefix                         = "balances"
	eth1BlockHashPrefix                    = "eth1_block_hash"
	eth1DepositIndexPrefix                 = "eth1_deposit_index"
	genesisValidatorsRootPrefix            = "genesis_validators_root"
)

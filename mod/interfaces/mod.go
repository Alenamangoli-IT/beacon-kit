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

package interfaces

import (
	"encoding/json"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/constraints"
)

// Deposit is an interface for deposits.
type Deposit[
	BLSPubkeyT any,
	BLSSignatureT any,
	DepositT any,
	U64T ~uint64,
	WithdrawalCredentialsT any,
] interface {
	// New creates a new deposit.
	New(
		BLSPubkeyT,
		WithdrawalCredentialsT,
		U64T,
		BLSSignatureT,
		uint64,
	) DepositT
	// GetIndex returns the index of the deposit.
	GetIndex() uint64
}

// -------------------------- ExecutionPayload --------------------------

// ExecutionPayloadBody is the interface for the execution data of a block.
// It contains all the fields that are part of both an execution payload header
// and a full execution payload.
type ExecutionPayloadBody[
	ExecutionAddressT any,
	ExecutionHashT any,
	Bytes32T any,
	U64T ~uint64,
	U256T any,
] interface {
	constraints.SSZMarshallable
	json.Marshaler
	json.Unmarshaler
	IsNil() bool
	Version() uint32
	GetPrevRandao() Bytes32T
	GetBlockHash() ExecutionHashT
	GetParentHash() ExecutionHashT
	GetNumber() U64T
	GetGasLimit() U64T
	GetGasUsed() U64T
	GetTimestamp() U64T
	GetExtraData() []byte
	GetBaseFeePerGas() U256T
	GetFeeRecipient() ExecutionAddressT
	GetStateRoot() Bytes32T
	GetReceiptsRoot() Bytes32T
	GetLogsBloom() []byte
	GetBlobGasUsed() U64T
	GetExcessBlobGas() U64T
}

// ExecutionPayloadWithTransactions is the interface for the execution data of a
// block that includes transactions.
type ExecutionPayloadWithTransactions[
	ExecutionAddressT any,
	ExecutionHashT any,
	Bytes32T any,
	U64T ~uint64,
	U256T any,
	TransactionsT any,
] interface {
	ExecutionPayloadBody[
		ExecutionAddressT,
		ExecutionHashT,
		Bytes32T,
		U64T,
		U256T,
	]
	GetTransactions() []TransactionsT
}

// ExecutionPayloadWithWithdrawals is the interface for the execution data of a
// block that includes withdrawals.
type ExecutionPayloadWithWithdrawals[
	ExecutionAddressT any,
	ExecutionHashT any,
	Bytes32T any,
	U64T ~uint64,
	U256T any,
	WithdrawalsT any,
] interface {
	ExecutionPayloadBody[
		ExecutionAddressT,
		ExecutionHashT,
		Bytes32T,
		U64T,
		U256T,
	]
	GetTransactions() []WithdrawalsT
}

// ExecutionPayload represents the execution data of a block.
type ExecutionPayload[
	ExecutionPayloadT any,
	ExecutionAddressT any,
	ExecutionHashT any,
	Bytes32T any,
	U64T ~uint64,
	U256T any,
	TransactionsT any,
	WithdrawalT any,
] interface {
	ExecutionPayloadBody[
		ExecutionAddressT,
		ExecutionHashT,
		Bytes32T,
		U64T,
		U256T,
	]
	Empty(uint32) ExecutionPayloadT
	GetTransactions() []TransactionsT
	GetWithdrawals() []WithdrawalT
}

// ExecutionPayloadHeader represents the execution header of a block.
type ExecutionPayloadHeader[
	ExecutionAddressT any,
	ExecutionHashT any,
	Bytes32T any,
	RootT any,
	U64T ~uint64,
	U256T any,
	WithdrawalT any,
] interface {
	ExecutionPayloadBody[
		ExecutionAddressT,
		ExecutionHashT,
		Bytes32T,
		U64T,
		U256T,
	]
	GetTransactionsRoot() RootT
	GetWithdrawalsRoot() RootT
}

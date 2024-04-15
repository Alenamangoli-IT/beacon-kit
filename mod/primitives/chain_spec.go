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

package primitives

// ChainSpec defines an interface for accessing chain-specific parameters.
type ChainSpec interface {
	// Gwei value constants.
	//
	// MinDepositAmount returns the minimum amount of Gwei required for a
	// deposit.
	MinDepositAmount() uint64
	// MaxEffectiveBalance returns the maximum balance counted in rewards
	// calculations in Gwei.
	MaxEffectiveBalance() uint64
	// EjectionBalance returns the balance below which a validator is ejected.
	EjectionBalance() uint64
	// EffectiveBalanceIncrement returns the increment of balance used in reward
	// calculations.
	EffectiveBalanceIncrement() uint64

	// Time parameters constants.
	//
	// SlotsPerEpoch returns the number of slots in an epoch.
	SlotsPerEpoch() uint64
	// SlotsPerHistoricalRoot returns the number of slots per historical root.
	SlotsPerHistoricalRoot() uint64

	// Eth1-related values.
	//
	// DepositContractAddress returns the deposit contract address.
	DepositContractAddress() ExecutionAddress

	// Fork-related values.
	//
	// ElectraForkEpoch returns the epoch at which the Electra fork takes
	// effect.
	ElectraForkEpoch() Epoch

	// State list lengths
	//
	// EpochsPerHistoricalVector returns the length of the historical vector.
	EpochsPerHistoricalVector() uint64
	// EpochsPerSlashingsVector returns the length of the slashing vector.
	EpochsPerSlashingsVector() uint64
	// HistoricalRootsLimit returns the maximum number of historical root
	// entries.
	HistoricalRootsLimit() uint64
	// ValidatorRegistryLimit returns the maximum number of validators in the
	// registry.
	ValidatorRegistryLimit() uint64

	// MaxDepositsPerBlock returns the maximum number of deposit operations per
	// block.
	MaxDepositsPerBlock() uint64
	// ProportionalSlashingMultiplier returns the multiplier for calculating
	// slashing penalties.
	ProportionalSlashingMultiplier() uint64

	// Capella Values
	//
	// MaxWithdrawalsPerPayload returns the maximum number of withdrawals per
	// payload.
	MaxWithdrawalsPerPayload() uint64
	// MaxValidatorsPerWithdrawalsSweep returns the maximum number of validators
	// per withdrawal sweep.
	MaxValidatorsPerWithdrawalsSweep() uint64

	// Deneb Values
	//
	// MinEpochsForBlobsSidecarsRequest returns the minimum number of epochs for
	// blob sidecar requests.
	MinEpochsForBlobsSidecarsRequest() uint64
	// MaxBlobCommitmentsPerBlock returns the maximum number of blob commitments
	// per block.
	MaxBlobCommitmentsPerBlock() uint64
	// MaxBlobsPerBlock returns the maximum number of blobs per block.
	MaxBlobsPerBlock() uint64
	// FieldElementsPerBlob returns the number of field elements per blob.
	FieldElementsPerBlob() uint64
	// BytesPerBlob returns the number of bytes per blob.
	BytesPerBlob() uint64

	// Helpers for ChainSpecData
	//
	// ActiveForkVersionForSlot returns the active fork version for a given
	// slot.
	ActiveForkVersionForSlot(slot Slot) uint32
	// ActiveForkVersionForEpoch returns the active fork version for a given
	// epoch.
	ActiveForkVersionForEpoch(epoch Epoch) uint32
	// SlotToEpoch converts a slot number to an epoch number.
	SlotToEpoch(slot Slot) Epoch
	// WithinDAPeriod checks if a given block slot is within the data
	// availability period relative to the current slot.
	WithinDAPeriod(block, current Slot) bool
}

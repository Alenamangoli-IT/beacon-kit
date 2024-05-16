package backend

import (
	"context"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives-engine"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
)

type StateDB interface {
	GetGenesisValidatorsRoot() (primitives.Root, error)
	GetSlot() (math.Slot, error)
	GetLatestExecutionPayloadHeader() (
		engineprimitives.ExecutionPayloadHeader, error,
	)
	SetLatestExecutionPayloadHeader(
		payloadHeader engineprimitives.ExecutionPayloadHeader,
	) error
	GetEth1DepositIndex() (uint64, error)
	SetEth1DepositIndex(
		index uint64,
	) error
	GetBalance(idx math.ValidatorIndex) (math.Gwei, error)
	SetBalance(idx math.ValidatorIndex, balance math.Gwei) error
	SetSlot(slot math.Slot) error
	GetFork() (*types.Fork, error)
	SetFork(fork *types.Fork) error
	GetLatestBlockHeader() (*types.BeaconBlockHeader, error)
	SetLatestBlockHeader(header *types.BeaconBlockHeader) error
	GetBlockRootAtIndex(index uint64) (primitives.Root, error)
	StateRootAtIndex(index uint64) (primitives.Root, error)
	GetEth1Data() (*types.Eth1Data, error)
	SetEth1Data(data *types.Eth1Data) error
	GetValidators() ([]*types.Validator, error)
	GetBalances() ([]uint64, error)
	GetNextWithdrawalIndex() (uint64, error)
	SetNextWithdrawalIndex(index uint64) error
	GetNextWithdrawalValidatorIndex() (math.ValidatorIndex, error)
	SetNextWithdrawalValidatorIndex(index math.ValidatorIndex) error
	GetTotalSlashing() (math.Gwei, error)
	SetTotalSlashing(total math.Gwei) error
	GetRandaoMixAtIndex(index uint64) (primitives.Bytes32, error)
	GetSlashings() ([]uint64, error)
	SetSlashingAtIndex(index uint64, amount math.Gwei) error
	GetSlashingAtIndex(index uint64) (math.Gwei, error)
	GetTotalValidators() (uint64, error)
	GetTotalActiveBalances(uint64) (math.Gwei, error)
	ValidatorByIndex(index math.ValidatorIndex) (*types.Validator, error)
	UpdateBlockRootAtIndex(index uint64, root primitives.Root) error
	UpdateStateRootAtIndex(index uint64, root primitives.Root) error
	UpdateRandaoMixAtIndex(index uint64, mix primitives.Bytes32) error
	UpdateValidatorAtIndex(
		index math.ValidatorIndex,
		validator *types.Validator,
	) error
	ValidatorIndexByPubkey(pubkey crypto.BLSPubkey) (math.ValidatorIndex, error)
	AddValidator(
		val *types.Validator,
	) error
	GetValidatorsByEffectiveBalance() ([]*types.Validator, error)
	RemoveValidatorAtIndex(idx math.ValidatorIndex) error
}

type Backend struct {
	getNewStateDB func(context.Context) StateDB
}

func New(getNewStateDB func(context.Context) StateDB) *Backend {
	return &Backend{
		getNewStateDB: getNewStateDB,
	}
}

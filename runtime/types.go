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

package runtime

import (
	"context"

	"github.com/berachain/beacon-kit/beacon/core/state"
	beacontypes "github.com/berachain/beacon-kit/beacon/core/types"
	"github.com/berachain/beacon-kit/beacon/forkchoice/ssf"
	bls12381 "github.com/berachain/beacon-kit/crypto/bls12-381"
	"github.com/berachain/beacon-kit/primitives"
)

// AppOptions is an interface that provides the ability to
// retrieve options from the application.
type AppOptions interface {
	Get(string) interface{}
}

// BeaconStorageBackend is an interface that provides the
// beacon state to the runtime.
type BeaconStorageBackend interface {
	AvailabilityStore(ctx context.Context) state.AvailabilityStore
	BeaconState(ctx context.Context) state.BeaconState
	// TODO: Decouple from the Specific SingleSlotFinalityStore Impl.
	ForkchoiceStore(ctx context.Context) ssf.SingleSlotFinalityStore
}

// ValsetUpdater is an interface that provides the
// ability to apply changes to the validator set.
type ValsetUpdater interface {
	IncreaseConsensusPower(
		ctx context.Context,
		delegator beacontypes.DepositCredentials,
		pubkey [bls12381.PubKeyLength]byte,
		amount uint64,
		signature []byte,
		index uint64,
	) error

	RedirectConsensusPower(
		ctx context.Context,
		delegator beacontypes.DepositCredentials,
		pubkey [bls12381.PubKeyLength]byte,
		newPubkey [bls12381.PubKeyLength]byte,
		amount uint64,
		signature []byte,
		index uint64,
	) error

	DecreaseConsensusPower(
		ctx context.Context,
		delegator primitives.ExecutionAddress,
		pubkey [bls12381.PubKeyLength]byte,
		amount uint64,
	) error
}

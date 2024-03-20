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

package evm

import (
	"context"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/registry"
	"github.com/berachain/beacon-kit/runtime/modules/beacon/keeper"
	"github.com/berachain/beacon-kit/runtime/modules/beacon/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"google.golang.org/grpc"
)

// ConsensusVersion defines the current x/beacon module consensus version.
const ConsensusVersion = 1

var (
	_ appmodule.HasServices        = AppModule{}
	_ appmodule.AppModule          = AppModule{}
	_ module.HasRegisterInterfaces = AppModule{}
	_ appmodule.HasServices        = AppModule{}
	_ module.HasABCIGenesis        = AppModule{}
	_ module.HasABCIEndBlock       = AppModule{}
)

// AppModule implements an application module for the evm module.
type AppModule struct {
	keeper *keeper.Keeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(
	keeper *keeper.Keeper,
) AppModule {
	return AppModule{
		keeper: keeper,
	}
}

// Name is the name of this module.
func (am AppModule) Name() string {
	return types.ModuleName
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(registrar grpc.ServiceRegistrar) error {
	types.RegisterMsgServer(registrar, am.keeper)
	return nil
}

func (am AppModule) RegisterInterfaces(r registry.LegacyRegistry) {
	types.RegisterInterfaces(r)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

func (am AppModule) EndBlock(
	ctx context.Context,
) ([]abci.ValidatorUpdate, error) {
	// Get the public key of the validator
	// pk, err := am.keeper.BeaconState(ctx).ValidatorPubKeyByIndex(0)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
	// // Return the validator update
	// return []abci.ValidatorUpdate{
	// 	{
	// 		PubKey: crypto.PublicKey{
	// 			Sum: &crypto.PublicKey_Bls12381{Bls12381: pk},
	// 		},
	// 		Power: 696969969696,
	// 	},
	// }, nil
}

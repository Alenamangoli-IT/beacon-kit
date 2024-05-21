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

package components

import (
	"cosmossdk.io/depinject"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/components/signer/privval"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
)

// BlsSignerInput is the input for the dep inject framework.
type BlsSignerInput struct {
	depinject.In
	AppOpts servertypes.AppOptions
}

// ProvideBlsSigner is a function that provides the module to the application.
// func ProvideBlsSigner(in BlsSignerInput) (crypto.BLSSigner, error) {
// 	// TODO: Fix
// 	//#nosec:G703 // temp.
// 	s, _ := signer.NewFromCometBFTNodeKey(
// 		cast.ToString(in.AppOpts.Get(flags.FlagHome)) +
// 			"/config/priv_validator_key.json",
// 	)
// 	return s, nil
// }

func ProvideBlsSigner(in BlsSignerInput) (crypto.BLSSigner, error) {
	keyFilePath := cast.ToString(in.AppOpts.Get(flags.FlagHome)) +
		"/config/priv_validator_key.json"
	stateFilePath := cast.ToString(in.AppOpts.Get(flags.FlagHome)) +
		"/config/priv_validator_state.json"
	s := privval.NewFileBLSSigner(keyFilePath, stateFilePath)
	return s, nil
}

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

package core

import (
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
)

// TODO: THIS FUNCTION NEEDS TO BE INTEGRATED BETTER WITH WITHDRAWALS AND
// SLASHING, IT IS A HACKY TEMPORARY WAY TO GET THE VALIDATOR SET UPDATING
// NICELY.
func (sp *StateProcessor[
	BeaconBlockT, BeaconBlockBodyT, BeaconStateT,
	BlobSidecarsT, ContextT,
]) processSyncCommitteeUpdates(
	st BeaconStateT,
) ([]*transition.ValidatorUpdate, error) {
	var (
		validatorUpdates = make([]*transition.ValidatorUpdate, 0)
	)

	// Get the public key of the validator
	vals, err := st.GetValidatorsByEffectiveBalance()
	if err != nil {
		panic(err)
	}

	for _, val := range vals {
		// If the validator is in the committee but below the ejection
		// balance
		// then they get ejected.
		validatorUpdates = append(validatorUpdates, &transition.ValidatorUpdate{
			Pubkey: val.Pubkey,
		})
	}

	return validatorUpdates, nil
}

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

package middleware

import (
	"sort"

	appmodulev2 "cosmossdk.io/core/appmodule/v2"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/transition"
	"github.com/sourcegraph/conc/iter"
)

// handleValUpdateConversion converts the validator updates from the
// transition package to the appmodulev2.ValidatorUpdate type.
func handleValUpdateConversion(
	valUpdates []*transition.ValidatorUpdate,
) ([]appmodulev2.ValidatorUpdate, error) {
	valUpdatesMap := make(map[string]*transition.ValidatorUpdate)
	for _, update := range valUpdates {
		pubKey := string(update.Pubkey[:])
		valUpdatesMap[pubKey] = update
	}

	// Convert map back to slice and sort by pubkey
	dedupedValUpdates := make(
		[]*transition.ValidatorUpdate,
		0,
		len(valUpdatesMap),
	)
	for _, update := range valUpdatesMap {
		dedupedValUpdates = append(dedupedValUpdates, update)
	}
	sort.Slice(dedupedValUpdates, func(i, j int) bool {
		return string(
			dedupedValUpdates[i].Pubkey[:],
		) < string(
			dedupedValUpdates[j].Pubkey[:],
		)
	})
	return iter.MapErr(dedupedValUpdates, convertValidatorUpdate)
}

// convertValidatorUpdate abstracts the conversion of a
// transition.ValidatorUpdate to an appmodulev2.ValidatorUpdate.
func convertValidatorUpdate(
	u **transition.ValidatorUpdate,
) (appmodulev2.ValidatorUpdate, error) {
	update := *u
	if update == nil {
		return appmodulev2.ValidatorUpdate{},
			ErrUndefinedValidatorUpdate
	}
	return appmodulev2.ValidatorUpdate{
		PubKey:     update.Pubkey[:],
		PubKeyType: crypto.CometBLSType,
		//#nosec:G701 // this is safe.
		Power: int64(update.EffectiveBalance.Unwrap()),
	}, nil
}

// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
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

package builder

import "github.com/berachain/beacon-kit/mod/errors"

var (
	// ErrPayloadBuilderDisabled is returned when the payload builder is
	// disabled.
	ErrPayloadBuilderDisabled = errors.New("payload builder is disabled")

	// ErrNilPayloadOnValidResponse is returned when a nil payload ID is
	// received on a VALID engine response.
	ErrNilPayloadOnValidResponse = errors.New(
		"received nil payload ID on VALID engine response",
	)

	// ErrNilPayloadID is returned when a nil payload ID is received.
	ErrNilPayloadID = errors.New("received nil payload ID")

	// ErrPayloadIDNotFound is returned when a payload ID is not found in the
	// cache.
	ErrPayloadIDNotFound = errors.New("unable to find payload ID in cache")

	// ErrCachedPayloadNotFoundOnExecutionClient is returned when a cached
	// payloadID is not found on the execution client.
	ErrCachedPayloadNotFoundOnExecutionClient = errors.New(
		"cached payload ID could not be resolved on execution client",
	)

	// ErrLocalBuildingDisabled is returned when local building is disabled.
	ErrLocalBuildingDisabled = errors.New("local building is disabled")

	// ErrNilPayloadEnvelope is returned when a nil payload envelope is
	// received.
	ErrNilPayloadEnvelope = errors.New("received nil payload envelope")

	// ErrNilPayload is returned when a nil payload envelope is
	// received.
	ErrNilPayload = errors.New("received nil payload envelope")
)

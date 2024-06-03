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

package encoding

import "github.com/berachain/beacon-kit/mod/errors"

var (
	// ErrNilBeaconBlockInRequest is an error for when
	// the beacon block in an abci request is nil.
	ErrNilBeaconBlockInRequest = errors.New("nil beacon block in abci request")

	// ErrNoBeaconBlockInRequest is an error for when
	// there is no beacon block in an abci request.
	ErrNoBeaconBlockInRequest = errors.New("no beacon block in abci request")

	// ErrBzIndexOutOfBounds is an error for when the index
	// is out of bounds.
	ErrBzIndexOutOfBounds = errors.New("bzIndex out of bounds")

	// ErrNilABCIRequest is an error for when the abci request
	// is nil.
	ErrNilABCIRequest = errors.New("nil abci request")

	// ErrInvalidType is an error for when the type is invalid.
	ErrInvalidType = errors.New("invalid type")

	// ErrNilBlobSidecarsInRequest is an error for when
	// the blob sidecars in an abci request is nil.
	ErrNilBlobSidecarsInRequest = errors.New(
		"nil blob sidecars in abci request",
	)
)

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

package verification

import (
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives"
	engineprimitives "github.com/berachain/beacon-kit/mod/primitives-engine"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core/state"
)

// PayloadVerifier is responsible for verifying incoming execution
// payloads to ensure they are valid.
type PayloadVerifier struct {
	cs     primitives.ChainSpec
	logger log.Logger[any]
}

// NewPayloadVerifier creates a new payload validator.
func NewPayloadVerifier(
	cs primitives.ChainSpec, logger log.Logger[any],
) *PayloadVerifier {
	return &PayloadVerifier{
		cs:     cs,
		logger: logger,
	}
}

// VerifyPayload verifies the incoming payload.
func (pv *PayloadVerifier) VerifyPayload(
	st state.BeaconState,
	payload engineprimitives.ExecutionPayload,
) error {
	return nil
}

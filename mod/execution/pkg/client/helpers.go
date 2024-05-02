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

package client

import (
	"fmt"
	"time"

	"github.com/berachain/beacon-kit/mod/primitives"
	engineprimitives "github.com/berachain/beacon-kit/mod/primitives-engine"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/net/jwt"
	gjwt "github.com/golang-jwt/jwt/v5"
)

// processPayloadStatusResult processes the payload status result and
// returns the latest valid hash or an error.
func processPayloadStatusResult(
	result *engineprimitives.PayloadStatus,
) (*primitives.ExecutionHash, error) {
	switch result.Status {
	case engineprimitives.PayloadStatusAccepted:
		return nil, ErrAcceptedPayloadStatus
	case engineprimitives.PayloadStatusSyncing:
		return nil, ErrSyncingPayloadStatus
	case engineprimitives.PayloadStatusInvalid:
		return result.LatestValidHash, ErrInvalidPayloadStatus
	case engineprimitives.PayloadStatusValid:
		return result.LatestValidHash, nil
	default:
		return nil, ErrUnknownPayloadStatus
	}
}

// buildSignedJWT builds a signed JWT from the provided JWT secret.
func buildSignedJWT(s *jwt.Secret) (string, error) {
	token := gjwt.NewWithClaims(gjwt.SigningMethodHS256, gjwt.MapClaims{
		"iat": &gjwt.NumericDate{Time: time.Now()},
	})
	str, err := token.SignedString(s[:])
	if err != nil {
		return "", fmt.Errorf("failed to create JWT token: %w", err)
	}
	return str, nil
}

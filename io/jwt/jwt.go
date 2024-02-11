// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
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

package jwt

import (
	"os"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// HexRegexp is a regular expression to match hexadecimal characters.
var HexRegexp = regexp.MustCompile(`^(?:0x)?[0-9a-fA-F]*$`)

// JWTLength defines the length of the JWT byte array to be 32 bytes as
// defined the Engine API specification.
// https://github.com/ethereum/execution-apis/blob/main/src/engine/authentication.md
const EthereumJWTLength = 32

// Secret represents a JSON Web Token as a fixed-size byte array.
type Secret [EthereumJWTLength]byte

// NewFromFile reads the JWT secret from a file and returns it.
func NewFromFile(filepath string) (Secret, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		// Return an error if the file cannot be read.
		return [EthereumJWTLength]byte{}, err
	}
	return NewFromHex(strings.TrimSpace(string(data)))
}

// NewFromHex creates a new JWT secret from a hexadecimal string.
func NewFromHex(hexStr string) (Secret, error) {
	// Ensure the hex string contains only hexadecimal characters.
	if !HexRegexp.MatchString(hexStr) {
		return [EthereumJWTLength]byte{}, ErrContainsIllegalCharacter
	}

	// Convert the hex string to a byte array.
	bz := common.FromHex(hexStr)
	if bz == nil || len(bz) != EthereumJWTLength {
		return [EthereumJWTLength]byte{}, ErrLengthMismatch
	}

	return Secret(bz), nil
}

// String returns the JWT secret as a string with the first 8 characters
// visible and the rest masked out for security.
func (s Secret) String() string {
	secret := hexutil.Encode(s[:])
	return secret[:8] + strings.Repeat("*", len(secret[8:]))
}

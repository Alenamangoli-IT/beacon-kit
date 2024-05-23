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

package validator

const (
	// defaultGraffiti is the default graffiti string.
	defaultGraffiti = ""
	// defaultSignerType is the default signer type.
	defaultSignerType = "bls"
	// defaultPrivKey is the default private key for the legacy signer.
	defaultPrivKey = ""
)

// Config is the validator configuration.
type Config struct {
	// Graffiti is the string that will be included in the
	// graffiti field of the beacon block.
	Graffiti string `mapstructure:"graffiti"`

	// SignerType is the type of the signer to use.
	// It can be either "bls" or "legacy" (or "remote" in the future).
	SignerType string `mapstructure:"signer-type"`

	// privKey is the private key for the legacy signer.
	PrivKey string
}

// DefaultConfig returns the default fork configuration.
func DefaultConfig() Config {
	return Config{
		Graffiti:   defaultGraffiti,
		SignerType: defaultSignerType,
		PrivKey:    defaultPrivKey,
	}
}

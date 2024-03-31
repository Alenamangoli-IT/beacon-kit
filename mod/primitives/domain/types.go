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

package domain

import "github.com/berachain/beacon-kit/mod/primitives"

// Domain constants as defined in the Ethereum 2.0 specification:
// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#domain-types
//
//nolint:lll,gochecknoglobals // Spec url is long, global vars are needed.
var (
	TypeProposer          = primitives.DomainType{0x00, 0x00, 0x00, 0x00}
	TypeAttester          = primitives.DomainType{0x01, 0x00, 0x00, 0x00}
	TypeRandao            = primitives.DomainType{0x02, 0x00, 0x00, 0x00}
	TypeDeposit           = primitives.DomainType{0x03, 0x00, 0x00, 0x00}
	TypeVoluntaryExit     = primitives.DomainType{0x04, 0x00, 0x00, 0x00}
	TypeSelectionProof    = primitives.DomainType{0x05, 0x00, 0x00, 0x00}
	TypeAggregateAndProof = primitives.DomainType{0x06, 0x00, 0x00, 0x00}
	TypeApplicationMask   = primitives.DomainType{0x00, 0x00, 0x00, 0x01}
)

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

package params

import "github.com/berachain/beacon-kit/mod/primitives"

// This file contains various constants as defined:
// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/beacon-chain.md#misc
//
//nolint:lll
const (
	// Uint64Max represents the maximum value for a uint64.
	Uint64Max = uint64(18446744073709551615)
	// Uint64MaxSqrt represents the square root of the maximum value for a
	// uint64.
	Uint64MaxSqrt = uint64(4294967295)
	// GenesisSlot represents the initial slot in the system.
	GenesisSlot = primitives.Slot(0)
	// GenesisEpoch represents the initial epoch in the system.
	GenesisEpoch = primitives.Epoch(0)
	// FarFutureEpoch represents a far future epoch value.
	FarFutureEpoch = primitives.Epoch(Uint64Max)
)

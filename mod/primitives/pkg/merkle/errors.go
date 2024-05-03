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

package merkle

import "github.com/berachain/beacon-kit/mod/errors"

var (
	// ErrNegativeIndex indicates that a negative index was provided.
	ErrNegativeIndex = errors.New("negative index provided")

	// ErrEmptyLeaves indicates that no items were provided to generate a Merkle
	// tree.
	ErrEmptyLeaves = errors.New("no items provided to generate Merkle tree")

	// ErrInsufficientDepthForLeaves indicates that the depth provided for the
	// Merkle tree is insufficient to store the provided leaves.
	ErrInsufficientDepthForLeaves = errors.New(
		"insufficient depth to store leaves",
	)

	// ErrZeroDepth indicates that the depth provided for the Merkle tree is
	// zero, which is invalid.
	ErrZeroDepth = errors.New("depth must be greater than 0")

	// ErrExceededDepth indicates that the provided depth exceeds the supported
	// maximum depth for a Merkle tree.
	ErrExceededDepth = errors.New("supported merkle tree depth exceeded")

	// ErrOddLengthTreeRoots is returned when the input list length must be
	// even.
	ErrOddLengthTreeRoots = errors.New("input list length must be even")

	// ErrMaxRootsExceeded is returned when the number of roots exceeds the
	// maximum allowed.
	ErrMaxRootsExceeded = errors.New(
		"number of roots exceeds the maximum allowed",
	)
)

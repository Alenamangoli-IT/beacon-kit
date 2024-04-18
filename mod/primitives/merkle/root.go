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

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/berachain/beacon-kit/mod/primitives/merkle/zero"
	"github.com/prysmaticlabs/gohashtree"
	"golang.org/x/sync/errgroup"
)

type U64[T ~uint64] interface {
	~uint64
	Unwrap() uint64
	NextPowerOfTwo() T
	ILog2Ceil() uint8
}

const (
	// MinParallelizationSize is the minimum size of the input list that
	// should be hashed using the default method. If the input list is smaller
	// than this size, the overhead of parallelizing the hashing process is.
	//
	// TODO: This value is arbitrary and should be benchmarked to find the
	// optimal value.
	MinParallelizationSize = 5000
	// two is a constant to make the linter happy.
	two = 2
)

// NewRootWithMaxLeaves constructs a Merkle tree root from a set of.
func NewRootWithMaxLeaves[U64T U64[U64T], LeafT, RootT ~[32]byte](
	leaves []LeafT,
	length uint64,
) (RootT, error) {
	return NewRootWithDepth[LeafT, RootT](
		leaves, U64T(length).NextPowerOfTwo().ILog2Ceil(),
	)
}

// NewRootWithDepth constructs a Merkle tree root from a set of leaves.
func NewRootWithDepth[LeafT, RootT ~[32]byte](
	leaves []LeafT,
	depth uint8,
) (RootT, error) {

	fmt.Println("DEPTH", depth)
	// Return zerohash at depth
	if len(leaves) == 0 {
		return zero.Hashes[depth], nil
	}

	for i := range depth {
		layerLen := len(leaves)
		oddNodeLength := layerLen%two == 1
		if oddNodeLength {
			zerohash := zero.Hashes[i]
			leaves = append(leaves, zerohash)
		}
		var err error
		leaves, err = BuildParentTreeRoots[LeafT, LeafT](leaves)
		if err != nil {
			return zero.Hashes[depth], err
		}
		fmt.Println("LEAVES", leaves)
	}
	if len(leaves) != 1 {
		return zero.Hashes[depth], nil
	}
	fmt.Println("LEAVES", leaves[0])
	return RootT(leaves[0]), nil
}

// BuildParentTreeRoots calls BuildParentTreeRootsWithNRoutines with the
// number of routines set to runtime.GOMAXPROCS(0)-1.
func BuildParentTreeRoots[LeafT, RootT ~[32]byte](
	inputList []LeafT,
) ([]RootT, error) {
	return BuildParentTreeRootsWithNRoutines[LeafT, RootT](
		inputList, runtime.GOMAXPROCS(0)-1,
	)
}

// BuildParentTreeRootsWithNRoutines optimizes hashing of a list of roots
// using CPU-specific vector instructions and parallel processing. This
// method adapts to the host machine's hardware for potential performance
// gains over sequential hashing.
func BuildParentTreeRootsWithNRoutines[LeafT, RootT ~[32]byte](
	inputList []LeafT, n int,
) ([]RootT, error) {
	// Validate input list length.
	inputLength := len(inputList)
	if inputLength%2 != 0 {
		return nil, ErrOddLengthTreeRoots
	}
	// Build output variables
	outputLength := inputLength / two
	outputList := make([]RootT, outputLength)

	// If the input list is small, hash it using the default method since
	// the overhead of parallelizing the hashing process is not worth it.
	if inputLength < MinParallelizationSize {
		return outputList, gohashtree.Hash(
			//#nosec:G103 // used of unsafe calls should be audited.
			*(*[][32]byte)(unsafe.Pointer(&outputList)),
			//#nosec:G103 // used of unsafe calls should be audited.
			*(*[][32]byte)(unsafe.Pointer(&inputList)))
	}

	// Otherwise parallelize the hashing process for large inputs.
	// Take the max(n, 1) to prevent division by 0.
	groupSize := inputLength / (two * max(n, 1))
	twiceGroupSize := two * groupSize
	eg := new(errgroup.Group)

	// if n is 0 the parallelization is disabled and the whole inputList is
	// hashed in the main goroutine at the end of this function.
	for j := 0; j <= n; j++ {
		// Define the segment of the inputList each goroutine will process.
		segmentStart := j * twiceGroupSize
		segmentEnd := min((j+1)*twiceGroupSize, inputLength)

		// inputList:  [---------------------2*groupSize---------------------]
		//              ^                    ^                    ^          ^
		//              |                    |                    |          |
		// j*2*groupSize   (j+1)*2*groupSize    (j+2)*2*groupSize  End
		//
		// outputList: [---------groupSize---------]
		//              ^                         ^
		//              |                         |
		//             j*groupSize         (j+1)*groupSize
		//
		// Each goroutine processes a segment of inputList that is twice as
		// large as the segment it fills in outputList. This is because the hash
		// operation reduces the
		// size of the input by half.
		eg.Go(func() error {
			return gohashtree.Hash(
				//#nosec:G103 // used of unsafe calls should be audited.
				(*(*[][32]byte)(
					unsafe.Pointer(
						&outputList,
					),
				))[j*groupSize:min((j+1)*groupSize, outputLength)],
				//#nosec:G103 // used of unsafe calls should be audited.
				(*(*[][32]byte)(
					unsafe.Pointer(
						&inputList,
					),
				))[segmentStart:segmentEnd],
			)
		})
	}

	// Wait for all goroutines to complete.
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return outputList, nil
}

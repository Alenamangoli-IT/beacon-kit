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

package sha256_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/itsdevbear/bolaris/crypto/sha256"
	"github.com/prysmaticlabs/gohashtree"
	"github.com/stretchr/testify/require"
)

func Test_VectorizedSha256(t *testing.T) {
	// Test with slices of varying sizes to ensure robustness across different conditions
	sliceSizes := []int{16, 32, 64}
	for _, size := range sliceSizes {
		t.Run(fmt.Sprintf("Size%d", size*sha256.MinParallelizationSize), func(t *testing.T) {
			largeSlice := make([][32]byte, size*sha256.MinParallelizationSize)
			secondLargeSlice := make([][32]byte, size*sha256.MinParallelizationSize)
			// Assuming hash reduces size by half
			hash1 := make([][32]byte, size*sha256.MinParallelizationSize/2)
			var hash2 [][32]byte
			var err error

			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				var tempHash [][32]byte
				tempHash, err = sha256.VectorizedSha256(largeSlice)
				copy(hash1, tempHash)
			}()
			wg.Wait()
			require.NoError(t, err)

			hash2, err = sha256.VectorizedSha256(secondLargeSlice)
			require.NoError(t, err)

			require.Equal(t, len(hash1), len(hash2), "Hash lengths should be equal")
			for i, r := range hash1 {
				require.Equal(t, r, hash2[i], fmt.Sprintf("Hash mismatch at index %d", i))
			}
		})
	}
}

func Test_GoHashTreeHashConformance(t *testing.T) {
	// Define a test table with various input sizes,
	// including ones above and below MinParallelizationSize
	testCases := []struct {
		name string
		size int
	}{
		{"BelowMinParallelizationSize", sha256.MinParallelizationSize / 2},
		{"AtMinParallelizationSize", sha256.MinParallelizationSize},
		{"AboveMinParallelizationSize", sha256.MinParallelizationSize * 2},
		{"SmallSize", 16},
		{"MediumSize", 64},
		{"LargeSize", 128},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputList := make([][32]byte, tc.size)
			outputList2 := make([][32]byte, tc.size/2)

			// Fill inputList with pseudo-random data
			randSource := rand.NewSource(time.Now().UnixNano())
			randGen := rand.New(randSource)
			for i := range inputList {
				for j := range inputList[i] {
					inputList[i][j] = byte(randGen.Intn(256))
				}
			}

			// Use VectorizedSha256 for outputList1
			outputList1, err := sha256.VectorizedSha256(inputList)
			require.NoError(t, err)

			// Use gohashtree.Hash directly for outputList2
			err = gohashtree.Hash(outputList2, inputList)
			require.NoError(t, err)

			// Use gohashtree.Hash directly for outputList2
			err = gohashtree.Hash(outputList2, inputList)
			require.NoError(t, err)

			// Compare the outputs
			require.Equal(
				t, outputList1, outputList2,
				"Outputs from VectorizedSha256 and gohashtree.Hash should be identical")
		})
	}
}

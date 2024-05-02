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

package types

import (
	"errors"

	"github.com/sourcegraph/conc/iter"
)

// SideCars is a slice of blob side cars to be included in the block.
//
//go:generate go run github.com/ferranbt/fastssz/sszgen -path ./sidecars.go -objs BlobSidecars -include ../../../primitives/pkg/consensus,../../../primitives/pkg/bytes,./sidecar.go,../../../primitives/pkg/math,../../../primitives/mod.go,../../../primitives/pkg/eip4844,$GETH_PKG_INCLUDE/common,$GETH_PKG_INCLUDE/common/hexutil -output sidecars.ssz.go
type BlobSidecars struct {
	// Sidecars is a slice of blob side cars to be included in the block.
	Sidecars []*BlobSidecar `ssz-max:"6"`
}

// ValidateBlockRoots checks to make sure that
// all blobs in the sidecar are from the same block.
func (bs *BlobSidecars) ValidateBlockRoots() error {
	// We only need to check if there is more than
	// a single blob in the sidecar.
	if sc := bs.Sidecars; len(sc) > 1 {
		firstHtr, err := sc[0].BeaconBlockHeader.HashTreeRoot()
		if err != nil {
			return err
		}

		var nextHtr [32]byte
		for i := 1; i < len(sc); i++ {
			nextHtr, err = sc[i].BeaconBlockHeader.HashTreeRoot()
			if err != nil {
				return err
			}
			if firstHtr != nextHtr {
				return ErrSidecarContainsDifferingBlockRoots
			}
		}
	}
	return nil
}

// VerifyInclusionProofs verifies the inclusion proofs for all sidecars.
func (bs *BlobSidecars) VerifyInclusionProofs(
	kzgOffset uint64,
) error {
	return errors.Join(iter.Map(
		bs.Sidecars,
		func(sidecar **BlobSidecar) error {
			sc := *sidecar
			if sc == nil {
				return ErrAttemptedToVerifyNilSidecar
			}

			// Verify the KZG inclusion proof.
			if !sc.HasValidInclusionProof(kzgOffset) {
				return ErrInvalidInclusionProof
			}
			return nil
		},
	)...)
}

// Len returns the number of sidecars in the sidecar.
func (bs *BlobSidecars) Len() int {
	return len(bs.Sidecars)
}

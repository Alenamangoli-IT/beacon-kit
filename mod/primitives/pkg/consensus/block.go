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

package consensus

import (
	"errors"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/version"
)

// EmptyBeaconBlock assembles a new beacon block from
// the given slot, time, execution data, and version. It
// returns an error if the fork version is not supported.
func EmptyBeaconBlock[
	SlotT, ProposerIndexT ~uint64,
	ParentBlockRootT, StateRootT ~[32]byte](
	slot SlotT,
	proposerIndex ProposerIndexT,
	parentBlockRoot ParentBlockRootT,
	stateRoot StateRootT,
	forkVersion uint32,
) (BeaconBlock, error) {
	var block BeaconBlock
	switch forkVersion {
	case version.Deneb:
		block = &BeaconBlockDeneb{
			BeaconBlockHeaderBase: BeaconBlockHeaderBase{
				Slot:            uint64(slot),
				ProposerIndex:   uint64(proposerIndex),
				ParentBlockRoot: bytes.B32(parentBlockRoot),
				StateRoot:       bytes.B32(stateRoot),
			},
			Body: &BeaconBlockBodyDeneb{},
		}
	default:
		return block, errors.New("unsupported fork version")
	}
	return block, nil
}

// BeaconBlockFromSSZ assembles a new beacon block
// from the given SSZ bytes and fork version.
func BeaconBlockFromSSZ(
	bz []byte,
	forkVersion uint32,
) (BeaconBlock, error) {
	// TODO: switch is fugazi atm.
	var block BeaconBlockDeneb
	switch forkVersion {
	case version.Deneb:
		_ = block
	default:
		return &block, ErrForkVersionNotSupported
	}

	if err := block.UnmarshalSSZ(bz); err != nil {
		return &block, err
	}
	return &block, nil
}

// BeaconBlockDeneb represents a block in the beacon chain during
// the Deneb fork.
//
//go:generate go run github.com/ferranbt/fastssz/sszgen --path block.go -objs BeaconBlockDeneb -include ../common,../crypto,../math,../../primitives.go,./header.go,./withdrawal_credentials.go,./withdrawal.go,./deposit.go,./payload.go,./deposit.go,../eip4844,../bytes,./eth1data.go,../math,../common,./body.go,$GETH_PKG_INCLUDE/common,$GETH_PKG_INCLUDE/common/hexutil -output block.ssz.go
type BeaconBlockDeneb struct {
	// BeaconBlockHeaderBase is the base of the BeaconBlockDeneb.
	BeaconBlockHeaderBase
	// Body is the body of the BeaconBlockDeneb, containing the block's
	// operations.
	Body *BeaconBlockBodyDeneb
}

// Version identifies the version of the BeaconBlockDeneb.
func (b *BeaconBlockDeneb) Version() uint32 {
	return version.Deneb
}

// IsNil checks if the BeaconBlockDeneb instance is nil.
func (b *BeaconBlockDeneb) IsNil() bool {
	return b == nil
}

// GetBody retrieves the body of the BeaconBlockDeneb.
func (b *BeaconBlockDeneb) GetBody() BeaconBlockBody {
	return b.Body
}

// GetHeader builds a BeaconBlockHeader from the BeaconBlockDeneb.
func (b BeaconBlockDeneb) GetHeader() *BeaconBlockHeader {
	bodyRoot, err := b.GetBody().HashTreeRoot()
	if err != nil {
		return nil
	}

	return &BeaconBlockHeader{
		BeaconBlockHeaderBase: BeaconBlockHeaderBase{
			Slot:            b.Slot,
			ProposerIndex:   b.ProposerIndex,
			ParentBlockRoot: b.ParentBlockRoot,
			StateRoot:       b.StateRoot,
		},
		BodyRoot: bodyRoot,
	}
}

// Code generated by fastssz. DO NOT EDIT.
// Hash: 9f4e1ffc6b2e734752c9ca5f453a62b708e3d6d345a77f4e393cc8670b30d87b
// Version: 0.1.3
package consensus

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconBlockHeader object
func (b *BeaconBlockHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockHeader object to a target array
func (b *BeaconBlockHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, b.Slot)

	// Field (1) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, b.ProposerIndex)

	// Field (2) 'ParentBlockRoot'
	dst = append(dst, b.ParentBlockRoot[:]...)

	// Field (3) 'StateRoot'
	dst = append(dst, b.StateRoot[:]...)

	// Field (4) 'BodyRoot'
	dst = append(dst, b.BodyRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockHeader object
func (b *BeaconBlockHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 112 {
		return ssz.ErrSize
	}

	// Field (0) 'Slot'
	b.Slot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ProposerIndex'
	b.ProposerIndex = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'ParentBlockRoot'
	copy(b.ParentBlockRoot[:], buf[16:48])

	// Field (3) 'StateRoot'
	copy(b.StateRoot[:], buf[48:80])

	// Field (4) 'BodyRoot'
	copy(b.BodyRoot[:], buf[80:112])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockHeader object
func (b *BeaconBlockHeader) SizeSSZ() (size int) {
	size = 112
	return
}

// HashTreeRoot ssz hashes the BeaconBlockHeader object
func (b *BeaconBlockHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockHeader object with a hasher
func (b *BeaconBlockHeader) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(b.Slot)

	// Field (1) 'ProposerIndex'
	hh.PutUint64(b.ProposerIndex)

	// Field (2) 'ParentBlockRoot'
	hh.PutBytes(b.ParentBlockRoot[:])

	// Field (3) 'StateRoot'
	hh.PutBytes(b.StateRoot[:])

	// Field (4) 'BodyRoot'
	hh.PutBytes(b.BodyRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconBlockHeader object
func (b *BeaconBlockHeader) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

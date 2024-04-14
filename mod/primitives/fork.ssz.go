// Code generated by fastssz. DO NOT EDIT.
// Hash: 9451543d627fdb018cb5826f7f92ff60f6b1d29792d88817eb7d65e3278e3e08
// Version: 0.1.3
package primitives

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Fork object
func (f *Fork) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the Fork object to a target array
func (f *Fork) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'PreviousVersion'
	dst = append(dst, f.PreviousVersion[:]...)

	// Field (1) 'CurrentVersion'
	dst = append(dst, f.CurrentVersion[:]...)

	// Field (2) 'Epoch'
	dst = ssz.MarshalUint64(dst, uint64(f.Epoch))

	return
}

// UnmarshalSSZ ssz unmarshals the Fork object
func (f *Fork) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'PreviousVersion'
	copy(f.PreviousVersion[:], buf[0:4])

	// Field (1) 'CurrentVersion'
	copy(f.CurrentVersion[:], buf[4:8])

	// Field (2) 'Epoch'
	f.Epoch = Epoch(ssz.UnmarshallUint64(buf[8:16]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Fork object
func (f *Fork) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the Fork object
func (f *Fork) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the Fork object with a hasher
func (f *Fork) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'PreviousVersion'
	hh.PutBytes(f.PreviousVersion[:])

	// Field (1) 'CurrentVersion'
	hh.PutBytes(f.CurrentVersion[:])

	// Field (2) 'Epoch'
	hh.PutUint64(uint64(f.Epoch))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Fork object
func (f *Fork) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(f)
}

// Code generated by fastssz. DO NOT EDIT.
// Hash: d3bdeb1bc6e9260faab76460cfd9529701b5b40b06576853892190f4ede20f2e
// Version: 0.1.3
package signature

import (
	"github.com/berachain/beacon-kit/mod/primitives"
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
	if size != 72 {
		return ssz.ErrSize
	}

	// Field (0) 'PreviousVersion'
	copy(f.PreviousVersion[:], buf[0:32])

	// Field (1) 'CurrentVersion'
	copy(f.CurrentVersion[:], buf[32:64])

	// Field (2) 'Epoch'
	f.Epoch = primitives.Epoch(ssz.UnmarshallUint64(buf[64:72]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Fork object
func (f *Fork) SizeSSZ() (size int) {
	size = 72
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

// MarshalSSZ ssz marshals the ForkData object
func (f *ForkData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the ForkData object to a target array
func (f *ForkData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'CurrentVersion'
	dst = append(dst, f.CurrentVersion[:]...)

	// Field (1) 'GenesisValidatorsRoot'
	dst = append(dst, f.GenesisValidatorsRoot[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the ForkData object
func (f *ForkData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 36 {
		return ssz.ErrSize
	}

	// Field (0) 'CurrentVersion'
	copy(f.CurrentVersion[:], buf[0:4])

	// Field (1) 'GenesisValidatorsRoot'
	copy(f.GenesisValidatorsRoot[:], buf[4:36])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ForkData object
func (f *ForkData) SizeSSZ() (size int) {
	size = 36
	return
}

// HashTreeRoot ssz hashes the ForkData object
func (f *ForkData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the ForkData object with a hasher
func (f *ForkData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentVersion'
	hh.PutBytes(f.CurrentVersion[:])

	// Field (1) 'GenesisValidatorsRoot'
	hh.PutBytes(f.GenesisValidatorsRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ForkData object
func (f *ForkData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(f)
}

// MarshalSSZ ssz marshals the SigningData object
func (s *SigningData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SigningData object to a target array
func (s *SigningData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ObjectRoot'
	dst = append(dst, s.ObjectRoot[:]...)

	// Field (1) 'Domain'
	dst = append(dst, s.Domain[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SigningData object
func (s *SigningData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'ObjectRoot'
	copy(s.ObjectRoot[:], buf[0:32])

	// Field (1) 'Domain'
	copy(s.Domain[:], buf[32:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SigningData object
func (s *SigningData) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the SigningData object
func (s *SigningData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SigningData object with a hasher
func (s *SigningData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'ObjectRoot'
	hh.PutBytes(s.ObjectRoot[:])

	// Field (1) 'Domain'
	hh.PutBytes(s.Domain[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SigningData object
func (s *SigningData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}

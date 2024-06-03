// Code generated by fastssz. DO NOT EDIT.
// Hash: 9de2b816183341cbb8f3b95065f296e6fa828de38646620dcfa50817b661bfe2
// Version: 0.1.3
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Eth1Data object
func (e *Eth1Data) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(e)
}

// MarshalSSZTo ssz marshals the Eth1Data object to a target array
func (e *Eth1Data) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'DepositRoot'
	dst = append(dst, e.DepositRoot[:]...)

	// Field (1) 'DepositCount'
	dst = ssz.MarshalUint64(dst, e.DepositCount)

	// Field (2) 'BlockHash'
	dst = append(dst, e.BlockHash[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the Eth1Data object
func (e *Eth1Data) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 72 {
		return ssz.ErrSize
	}

	// Field (0) 'DepositRoot'
	copy(e.DepositRoot[:], buf[0:32])

	// Field (1) 'DepositCount'
	e.DepositCount = ssz.UnmarshallUint64(buf[32:40])

	// Field (2) 'BlockHash'
	copy(e.BlockHash[:], buf[40:72])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Eth1Data object
func (e *Eth1Data) SizeSSZ() (size int) {
	size = 72
	return
}

// HashTreeRoot ssz hashes the Eth1Data object
func (e *Eth1Data) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the Eth1Data object with a hasher
func (e *Eth1Data) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'DepositRoot'
	hh.PutBytes(e.DepositRoot[:])

	// Field (1) 'DepositCount'
	hh.PutUint64(e.DepositCount)

	// Field (2) 'BlockHash'
	hh.PutBytes(e.BlockHash[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Eth1Data object
func (e *Eth1Data) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}

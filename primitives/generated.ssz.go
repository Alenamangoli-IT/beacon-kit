// Code generated by fastssz. DO NOT EDIT.
// Hash: 39b9fbbd293c4718227441b5e5511b0747a5208654414a9036846f96b5a90867
package primitives

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the SigningData object
func (s *SigningData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SigningData object to a target array
func (s *SigningData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ObjectRoot'
	if size := len(s.ObjectRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ObjectRoot", size, 32)
		return
	}
	dst = append(dst, s.ObjectRoot...)

	// Field (1) 'Domain'
	if size := len(s.Domain); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Domain", size, 32)
		return
	}
	dst = append(dst, s.Domain...)

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
	if cap(s.ObjectRoot) == 0 {
		s.ObjectRoot = make([]byte, 0, len(buf[0:32]))
	}
	s.ObjectRoot = append(s.ObjectRoot, buf[0:32]...)

	// Field (1) 'Domain'
	if cap(s.Domain) == 0 {
		s.Domain = make([]byte, 0, len(buf[32:64]))
	}
	s.Domain = append(s.Domain, buf[32:64]...)

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
func (s *SigningData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ObjectRoot'
	if size := len(s.ObjectRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ObjectRoot", size, 32)
		return
	}
	hh.PutBytes(s.ObjectRoot)

	// Field (1) 'Domain'
	if size := len(s.Domain); size != 32 {
		err = ssz.ErrBytesLengthFn("--.Domain", size, 32)
		return
	}
	hh.PutBytes(s.Domain)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

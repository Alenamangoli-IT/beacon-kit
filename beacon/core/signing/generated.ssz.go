// Code generated by fastssz. DO NOT EDIT.
// Hash: 61fed304c802a625ac035c1c2b90549cedfc88582bb5c584afee95588b0c55e6
package signing

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the ForkData object
func (f *ForkData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the ForkData object to a target array
func (f *ForkData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Field (0) 'CurrentVersion'
	dst = append(dst, f.CurrentVersion[:]...)

	// Offset (1) 'ChainID'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(f.ChainID)

	// Field (1) 'ChainID'
	if size := len(f.ChainID); size > 50 {
		err = ssz.ErrBytesLengthFn("--.ChainID", size, 50)
		return
	}
	dst = append(dst, f.ChainID...)

	return
}

// UnmarshalSSZ ssz unmarshals the ForkData object
func (f *ForkData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'CurrentVersion'
	copy(f.CurrentVersion[:], buf[0:4])

	// Offset (1) 'ChainID'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'ChainID'
	{
		buf = tail[o1:]
		if len(buf) > 50 {
			return ssz.ErrBytesLength
		}
		if cap(f.ChainID) == 0 {
			f.ChainID = make([]byte, 0, len(buf))
		}
		f.ChainID = append(f.ChainID, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ForkData object
func (f *ForkData) SizeSSZ() (size int) {
	size = 8

	// Field (1) 'ChainID'
	size += len(f.ChainID)

	return
}

// HashTreeRoot ssz hashes the ForkData object
func (f *ForkData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the ForkData object with a hasher
func (f *ForkData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentVersion'
	hh.PutBytes(f.CurrentVersion[:])

	// Field (1) 'ChainID'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(f.ChainID))
		if byteLen > 50 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(f.ChainID)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (50+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (50+31)/32)
		}
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Data object
func (d *Data) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Data object to a target array
func (d *Data) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ObjectRoot'
	dst = append(dst, d.ObjectRoot[:]...)

	// Field (1) 'Domain'
	dst = append(dst, d.Domain[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the Data object
func (d *Data) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'ObjectRoot'
	copy(d.ObjectRoot[:], buf[0:32])

	// Field (1) 'Domain'
	copy(d.Domain[:], buf[32:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Data object
func (d *Data) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the Data object
func (d *Data) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Data object with a hasher
func (d *Data) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ObjectRoot'
	hh.PutBytes(d.ObjectRoot[:])

	// Field (1) 'Domain'
	hh.PutBytes(d.Domain[:])

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

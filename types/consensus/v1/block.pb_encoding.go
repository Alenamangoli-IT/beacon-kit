// Code generated by fastssz. DO NOT EDIT.
// Hash: b8a21210216b2b09b56beb4f139a0a7673c32f373e445a4703275914f7cd76e7
package v1

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the BeaconKitBlock object
func (b *BeaconKitBlock) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconKitBlock object to a target array
func (b *BeaconKitBlock) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(44)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, b.Slot)

	// Offset (1) 'BlockBodyGeneric'
	dst = ssz.WriteOffset(dst, offset)
	if b.BlockBodyGeneric == nil {
		b.BlockBodyGeneric = new(BeaconBlockBody)
	}
	offset += b.BlockBodyGeneric.SizeSSZ()

	// Field (2) 'PayloadValue'
	if size := len(b.PayloadValue); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PayloadValue", size, 32)
		return
	}
	dst = append(dst, b.PayloadValue...)

	// Field (1) 'BlockBodyGeneric'
	if dst, err = b.BlockBodyGeneric.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconKitBlock object
func (b *BeaconKitBlock) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 44 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'Slot'
	b.Slot = ssz.UnmarshallUint64(buf[0:8])

	// Offset (1) 'BlockBodyGeneric'
	if o1 = ssz.ReadOffset(buf[8:12]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 44 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'PayloadValue'
	if cap(b.PayloadValue) == 0 {
		b.PayloadValue = make([]byte, 0, len(buf[12:44]))
	}
	b.PayloadValue = append(b.PayloadValue, buf[12:44]...)

	// Field (1) 'BlockBodyGeneric'
	{
		buf = tail[o1:]
		if b.BlockBodyGeneric == nil {
			b.BlockBodyGeneric = new(BeaconBlockBody)
		}
		if err = b.BlockBodyGeneric.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconKitBlock object
func (b *BeaconKitBlock) SizeSSZ() (size int) {
	size = 44

	// Field (1) 'BlockBodyGeneric'
	if b.BlockBodyGeneric == nil {
		b.BlockBodyGeneric = new(BeaconBlockBody)
	}
	size += b.BlockBodyGeneric.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BeaconKitBlock object
func (b *BeaconKitBlock) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconKitBlock object with a hasher
func (b *BeaconKitBlock) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(b.Slot)

	// Field (1) 'BlockBodyGeneric'
	if err = b.BlockBodyGeneric.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'PayloadValue'
	if size := len(b.PayloadValue); size != 32 {
		err = ssz.ErrBytesLengthFn("--.PayloadValue", size, 32)
		return
	}
	hh.PutBytes(b.PayloadValue)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

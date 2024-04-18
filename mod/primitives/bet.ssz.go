// Code generated by fastssz. DO NOT EDIT.
// Hash: e155f580133326a18afad62432ff7d2a6855e6bab5dc6ae1f150ece64ed77781
// Version: 0.1.3
package primitives

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the U64List2 object
func (u *U64List2) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(u)
}

// MarshalSSZTo ssz marshals the U64List2 object to a target array
func (u *U64List2) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Data'
	dst = ssz.WriteOffset(dst, offset)

	// Field (0) 'Data'
	if size := len(u.Data); size > 16 {
		err = ssz.ErrListTooBigFn("U64List2.Data", size, 16)
		return
	}
	for ii := 0; ii < len(u.Data); ii++ {
		dst = ssz.MarshalUint64(dst, u.Data[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the U64List2 object
func (u *U64List2) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Data'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Data'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 8, 16)
		if err != nil {
			return err
		}
		u.Data = ssz.ExtendUint64(u.Data, num)
		for ii := 0; ii < num; ii++ {
			u.Data[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the U64List2 object
func (u *U64List2) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Data'
	size += len(u.Data) * 8

	return
}

// HashTreeRoot ssz hashes the U64List2 object
func (u *U64List2) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(u)
}

// HashTreeRootWith ssz hashes the U64List2 object with a hasher
func (u *U64List2) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Data'
	{
		if size := len(u.Data); size > 16 {
			err = ssz.ErrListTooBigFn("U64List2.Data", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range u.Data {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(u.Data))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(16, numItems, 8))
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the U64List2 object
func (u *U64List2) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(u)
}

// MarshalSSZ ssz marshals the U64Container2 object
func (u *U64Container2) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(u)
}

// MarshalSSZTo ssz marshals the U64Container2 object to a target array
func (u *U64Container2) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'Field2'
	dst = ssz.WriteOffset(dst, offset)

	// Field (1) 'Field1'
	dst = ssz.MarshalUint64(dst, uint64(u.Field1))

	// Field (0) 'Field2'
	if size := len(u.Field2); size > 16 {
		err = ssz.ErrListTooBigFn("U64Container2.Field2", size, 16)
		return
	}
	for ii := 0; ii < len(u.Field2); ii++ {
		dst = ssz.MarshalUint64(dst, u.Field2[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the U64Container2 object
func (u *U64Container2) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Field2'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Field1'
	u.Field1 = U64(ssz.UnmarshallUint64(buf[4:12]))

	// Field (0) 'Field2'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 8, 16)
		if err != nil {
			return err
		}
		u.Field2 = ssz.ExtendUint64(u.Field2, num)
		for ii := 0; ii < num; ii++ {
			u.Field2[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the U64Container2 object
func (u *U64Container2) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'Field2'
	size += len(u.Field2) * 8

	return
}

// HashTreeRoot ssz hashes the U64Container2 object
func (u *U64Container2) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(u)
}

// HashTreeRootWith ssz hashes the U64Container2 object with a hasher
func (u *U64Container2) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Field2'
	{
		if size := len(u.Field2); size > 16 {
			err = ssz.ErrListTooBigFn("U64Container2.Field2", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range u.Field2 {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(u.Field2))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(16, numItems, 8))
	}

	// Field (1) 'Field1'
	hh.PutUint64(uint64(u.Field1))
	hh.FillUpTo32()

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the U64Container2 object
func (u *U64Container2) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(u)
}

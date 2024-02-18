// Code generated by fastssz. DO NOT EDIT.
// Hash: 637cd600e2ebcc6e4b997c2e6a1a959db7e0408527b1acd2d7229d38cbf3551a
package consensusv1

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the Deposit object
func (d *Deposit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Deposit object to a target array
func (d *Deposit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Data'
	dst = ssz.WriteOffset(dst, offset)
	if d.Data == nil {
		d.Data = new(Deposit_Data)
	}
	offset += d.Data.SizeSSZ()

	// Field (0) 'Data'
	if dst, err = d.Data.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Deposit object
func (d *Deposit) UnmarshalSSZ(buf []byte) error {
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
		if d.Data == nil {
			d.Data = new(Deposit_Data)
		}
		if err = d.Data.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Deposit object
func (d *Deposit) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Data'
	if d.Data == nil {
		d.Data = new(Deposit_Data)
	}
	size += d.Data.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Deposit object
func (d *Deposit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Deposit object with a hasher
func (d *Deposit) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Data'
	if err = d.Data.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Deposit_Data object
func (d *Deposit_Data) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Deposit_Data object to a target array
func (d *Deposit_Data) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(32)

	// Offset (0) 'Pubkey'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(d.Pubkey)

	// Field (1) 'WithdrawalCredentials'
	if size := len(d.WithdrawalCredentials); size != 20 {
		err = ssz.ErrBytesLengthFn("--.WithdrawalCredentials", size, 20)
		return
	}
	dst = append(dst, d.WithdrawalCredentials...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, d.Amount)

	// Field (0) 'Pubkey'
	if size := len(d.Pubkey); size > 96 {
		err = ssz.ErrBytesLengthFn("--.Pubkey", size, 96)
		return
	}
	dst = append(dst, d.Pubkey...)

	return
}

// UnmarshalSSZ ssz unmarshals the Deposit_Data object
func (d *Deposit_Data) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 32 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Pubkey'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 32 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'WithdrawalCredentials'
	if cap(d.WithdrawalCredentials) == 0 {
		d.WithdrawalCredentials = make([]byte, 0, len(buf[4:24]))
	}
	d.WithdrawalCredentials = append(d.WithdrawalCredentials, buf[4:24]...)

	// Field (2) 'Amount'
	d.Amount = ssz.UnmarshallUint64(buf[24:32])

	// Field (0) 'Pubkey'
	{
		buf = tail[o0:]
		if len(buf) > 96 {
			return ssz.ErrBytesLength
		}
		if cap(d.Pubkey) == 0 {
			d.Pubkey = make([]byte, 0, len(buf))
		}
		d.Pubkey = append(d.Pubkey, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Deposit_Data object
func (d *Deposit_Data) SizeSSZ() (size int) {
	size = 32

	// Field (0) 'Pubkey'
	size += len(d.Pubkey)

	return
}

// HashTreeRoot ssz hashes the Deposit_Data object
func (d *Deposit_Data) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Deposit_Data object with a hasher
func (d *Deposit_Data) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkey'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(d.Pubkey))
		if byteLen > 96 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(d.Pubkey)
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(elemIndx, byteLen, (96+31)/32)
		} else {
			hh.MerkleizeWithMixin(elemIndx, byteLen, (96+31)/32)
		}
	}

	// Field (1) 'WithdrawalCredentials'
	if size := len(d.WithdrawalCredentials); size != 20 {
		err = ssz.ErrBytesLengthFn("--.WithdrawalCredentials", size, 20)
		return
	}
	hh.PutBytes(d.WithdrawalCredentials)

	// Field (2) 'Amount'
	hh.PutUint64(d.Amount)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

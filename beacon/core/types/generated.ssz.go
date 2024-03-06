// Code generated by fastssz. DO NOT EDIT.
// Hash: d9cf62554f32cd0a165bd8af2e7026047059352f504233f005934c59a7df2c1d
package types

import (
	beacontypesv1 "github.com/itsdevbear/bolaris/beacon/core/types/v1"
	enginetypes "github.com/itsdevbear/bolaris/engine/types"
	"github.com/itsdevbear/bolaris/primitives"
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockDeneb object to a target array
func (b *BeaconBlockDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(76)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (1) 'ParentBlockRoot'
	dst = append(dst, b.ParentBlockRoot[:]...)

	// Offset (2) 'Body'
	dst = ssz.WriteOffset(dst, offset)
	if b.Body == nil {
		b.Body = new(BeaconBlockBodyDeneb)
	}
	offset += b.Body.SizeSSZ()

	// Field (3) 'PayloadValue'
	dst = append(dst, b.PayloadValue[:]...)

	// Field (2) 'Body'
	if dst, err = b.Body.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 76 {
		return ssz.ErrSize
	}

	tail := buf
	var o2 uint64

	// Field (0) 'Slot'
	b.Slot = primitives.Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'ParentBlockRoot'
	copy(b.ParentBlockRoot[:], buf[8:40])

	// Offset (2) 'Body'
	if o2 = ssz.ReadOffset(buf[40:44]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 76 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (3) 'PayloadValue'
	copy(b.PayloadValue[:], buf[44:76])

	// Field (2) 'Body'
	{
		buf = tail[o2:]
		if b.Body == nil {
			b.Body = new(BeaconBlockBodyDeneb)
		}
		if err = b.Body.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) SizeSSZ() (size int) {
	size = 76

	// Field (2) 'Body'
	if b.Body == nil {
		b.Body = new(BeaconBlockBodyDeneb)
	}
	size += b.Body.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BeaconBlockDeneb object
func (b *BeaconBlockDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockDeneb object with a hasher
func (b *BeaconBlockDeneb) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (1) 'ParentBlockRoot'
	hh.PutBytes(b.ParentBlockRoot[:])

	// Field (2) 'Body'
	if err = b.Body.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'PayloadValue'
	hh.PutBytes(b.PayloadValue[:])

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconBlockBodyDeneb object to a target array
func (b *BeaconBlockBodyDeneb) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(140)

	// Field (0) 'RandaoReveal'
	dst = append(dst, b.RandaoReveal[:]...)

	// Field (1) 'Graffiti'
	dst = append(dst, b.Graffiti[:]...)

	// Offset (2) 'Deposits'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.Deposits); ii++ {
		offset += 4
		offset += b.Deposits[ii].SizeSSZ()
	}

	// Offset (3) 'ExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(enginetypes.ExecutableDataDeneb)
	}
	offset += b.ExecutionPayload.SizeSSZ()

	// Offset (4) 'BlobKzgCommitments'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.BlobKzgCommitments) * 48

	// Field (2) 'Deposits'
	if size := len(b.Deposits); size > 16 {
		err = ssz.ErrListTooBigFn("--.Deposits", size, 16)
		return
	}
	{
		offset = 4 * len(b.Deposits)
		for ii := 0; ii < len(b.Deposits); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.Deposits[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.Deposits); ii++ {
		if dst, err = b.Deposits[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (3) 'ExecutionPayload'
	if dst, err = b.ExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'BlobKzgCommitments'
	if size := len(b.BlobKzgCommitments); size > 16 {
		err = ssz.ErrListTooBigFn("--.BlobKzgCommitments", size, 16)
		return
	}
	for ii := 0; ii < len(b.BlobKzgCommitments); ii++ {
		dst = append(dst, b.BlobKzgCommitments[ii][:]...)
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 140 {
		return ssz.ErrSize
	}

	tail := buf
	var o2, o3, o4 uint64

	// Field (0) 'RandaoReveal'
	copy(b.RandaoReveal[:], buf[0:96])

	// Field (1) 'Graffiti'
	copy(b.Graffiti[:], buf[96:128])

	// Offset (2) 'Deposits'
	if o2 = ssz.ReadOffset(buf[128:132]); o2 > size {
		return ssz.ErrOffset
	}

	if o2 < 140 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (3) 'ExecutionPayload'
	if o3 = ssz.ReadOffset(buf[132:136]); o3 > size || o2 > o3 {
		return ssz.ErrOffset
	}

	// Offset (4) 'BlobKzgCommitments'
	if o4 = ssz.ReadOffset(buf[136:140]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Field (2) 'Deposits'
	{
		buf = tail[o2:o3]
		num, err := ssz.DecodeDynamicLength(buf, 16)
		if err != nil {
			return err
		}
		b.Deposits = make([]*beacontypesv1.Deposit, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.Deposits[indx] == nil {
				b.Deposits[indx] = new(beacontypesv1.Deposit)
			}
			if err = b.Deposits[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (3) 'ExecutionPayload'
	{
		buf = tail[o3:o4]
		if b.ExecutionPayload == nil {
			b.ExecutionPayload = new(enginetypes.ExecutableDataDeneb)
		}
		if err = b.ExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (4) 'BlobKzgCommitments'
	{
		buf = tail[o4:]
		num, err := ssz.DivideInt2(len(buf), 48, 16)
		if err != nil {
			return err
		}
		b.BlobKzgCommitments = make([][48]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(b.BlobKzgCommitments[ii][:], buf[ii*48:(ii+1)*48])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) SizeSSZ() (size int) {
	size = 140

	// Field (2) 'Deposits'
	for ii := 0; ii < len(b.Deposits); ii++ {
		size += 4
		size += b.Deposits[ii].SizeSSZ()
	}

	// Field (3) 'ExecutionPayload'
	if b.ExecutionPayload == nil {
		b.ExecutionPayload = new(enginetypes.ExecutableDataDeneb)
	}
	size += b.ExecutionPayload.SizeSSZ()

	// Field (4) 'BlobKzgCommitments'
	size += len(b.BlobKzgCommitments) * 48

	return
}

// HashTreeRoot ssz hashes the BeaconBlockBodyDeneb object
func (b *BeaconBlockBodyDeneb) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconBlockBodyDeneb object with a hasher
func (b *BeaconBlockBodyDeneb) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'RandaoReveal'
	hh.PutBytes(b.RandaoReveal[:])

	// Field (1) 'Graffiti'
	hh.PutBytes(b.Graffiti[:])

	// Field (2) 'Deposits'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Deposits))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Deposits {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, num, 16)
		} else {
			hh.MerkleizeWithMixin(subIndx, num, 16)
		}
	}

	// Field (3) 'ExecutionPayload'
	if err = b.ExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'BlobKzgCommitments'
	{
		if size := len(b.BlobKzgCommitments); size > 16 {
			err = ssz.ErrListTooBigFn("--.BlobKzgCommitments", size, 16)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlobKzgCommitments {
			hh.PutBytes(i[:])
		}

		numItems := uint64(len(b.BlobKzgCommitments))
		if ssz.EnableVectorizedHTR {
			hh.MerkleizeWithMixinVectorizedHTR(subIndx, numItems, 16)
		} else {
			hh.MerkleizeWithMixin(subIndx, numItems, 16)
		}
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

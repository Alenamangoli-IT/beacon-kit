// Code generated by fastssz. DO NOT EDIT.
// Hash: 7178e492efc2d9fad59d406a2a867e3de0460031ad638d747291ae5016caff9a
// Version: 0.1.3
package deneb

import (
	"github.com/berachain/beacon-kit/mod/primitives"
	engineprimitives "github.com/berachain/beacon-kit/mod/primitives-engine"
	"github.com/berachain/beacon-kit/mod/primitives/math"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconState object
func (b *BeaconState) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconState object to a target array
func (b *BeaconState) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(300)

	// Field (0) 'GenesisValidatorsRoot'
	dst = append(dst, b.GenesisValidatorsRoot[:]...)

	// Field (1) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (2) 'Fork'
	if b.Fork == nil {
		b.Fork = new(primitives.Fork)
	}
	if dst, err = b.Fork.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(primitives.BeaconBlockHeader)
	}
	if dst, err = b.LatestBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (4) 'BlockRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.BlockRoots) * 32

	// Offset (5) 'StateRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.StateRoots) * 32

	// Offset (6) 'LatestExecutionPayload'
	dst = ssz.WriteOffset(dst, offset)
	if b.LatestExecutionPayload == nil {
		b.LatestExecutionPayload = new(engineprimitives.ExecutableDataDeneb)
	}
	offset += b.LatestExecutionPayload.SizeSSZ()

	// Field (7) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(primitives.Eth1Data)
	}
	if dst, err = b.Eth1Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (8) 'Eth1DepositIndex'
	dst = ssz.MarshalUint64(dst, b.Eth1DepositIndex)

	// Offset (9) 'Validators'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Validators) * 121

	// Offset (10) 'Balances'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Balances) * 8

	// Offset (11) 'RandaoMixes'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.RandaoMixes) * 32

	// Field (12) 'NextWithdrawalIndex'
	dst = ssz.MarshalUint64(dst, b.NextWithdrawalIndex)

	// Field (13) 'NextWithdrawalValidatorIndex'
	dst = ssz.MarshalUint64(dst, uint64(b.NextWithdrawalValidatorIndex))

	// Offset (14) 'Slashings'
	dst = ssz.WriteOffset(dst, offset)

	// Field (15) 'TotalSlashing'
	dst = ssz.MarshalUint64(dst, uint64(b.TotalSlashing))

	// Field (4) 'BlockRoots'
	if size := len(b.BlockRoots); size > 8192 {
		err = ssz.ErrListTooBigFn("BeaconState.BlockRoots", size, 8192)
		return
	}
	for ii := 0; ii < len(b.BlockRoots); ii++ {
		dst = append(dst, b.BlockRoots[ii][:]...)
	}

	// Field (5) 'StateRoots'
	if size := len(b.StateRoots); size > 8192 {
		err = ssz.ErrListTooBigFn("BeaconState.StateRoots", size, 8192)
		return
	}
	for ii := 0; ii < len(b.StateRoots); ii++ {
		dst = append(dst, b.StateRoots[ii][:]...)
	}

	// Field (6) 'LatestExecutionPayload'
	if dst, err = b.LatestExecutionPayload.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (9) 'Validators'
	if size := len(b.Validators); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.Validators", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Validators); ii++ {
		if dst, err = b.Validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (10) 'Balances'
	if size := len(b.Balances); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.Balances", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Balances); ii++ {
		dst = ssz.MarshalUint64(dst, b.Balances[ii])
	}

	// Field (11) 'RandaoMixes'
	if size := len(b.RandaoMixes); size > 65536 {
		err = ssz.ErrListTooBigFn("BeaconState.RandaoMixes", size, 65536)
		return
	}
	for ii := 0; ii < len(b.RandaoMixes); ii++ {
		dst = append(dst, b.RandaoMixes[ii][:]...)
	}

	// Field (14) 'Slashings'
	if size := len(b.Slashings); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.Slashings", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Slashings); ii++ {
		dst = ssz.MarshalUint64(dst, b.Slashings[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconState object
func (b *BeaconState) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 300 {
		return ssz.ErrSize
	}

	tail := buf
	var o4, o5, o6, o9, o10, o11, o14 uint64

	// Field (0) 'GenesisValidatorsRoot'
	copy(b.GenesisValidatorsRoot[:], buf[0:32])

	// Field (1) 'Slot'
	b.Slot = math.Slot(ssz.UnmarshallUint64(buf[32:40]))

	// Field (2) 'Fork'
	if b.Fork == nil {
		b.Fork = new(primitives.Fork)
	}
	if err = b.Fork.UnmarshalSSZ(buf[40:56]); err != nil {
		return err
	}

	// Field (3) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(primitives.BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.UnmarshalSSZ(buf[56:168]); err != nil {
		return err
	}

	// Offset (4) 'BlockRoots'
	if o4 = ssz.ReadOffset(buf[168:172]); o4 > size {
		return ssz.ErrOffset
	}

	if o4 < 300 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (5) 'StateRoots'
	if o5 = ssz.ReadOffset(buf[172:176]); o5 > size || o4 > o5 {
		return ssz.ErrOffset
	}

	// Offset (6) 'LatestExecutionPayload'
	if o6 = ssz.ReadOffset(buf[176:180]); o6 > size || o5 > o6 {
		return ssz.ErrOffset
	}

	// Field (7) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(primitives.Eth1Data)
	}
	if err = b.Eth1Data.UnmarshalSSZ(buf[180:252]); err != nil {
		return err
	}

	// Field (8) 'Eth1DepositIndex'
	b.Eth1DepositIndex = ssz.UnmarshallUint64(buf[252:260])

	// Offset (9) 'Validators'
	if o9 = ssz.ReadOffset(buf[260:264]); o9 > size || o6 > o9 {
		return ssz.ErrOffset
	}

	// Offset (10) 'Balances'
	if o10 = ssz.ReadOffset(buf[264:268]); o10 > size || o9 > o10 {
		return ssz.ErrOffset
	}

	// Offset (11) 'RandaoMixes'
	if o11 = ssz.ReadOffset(buf[268:272]); o11 > size || o10 > o11 {
		return ssz.ErrOffset
	}

	// Field (12) 'NextWithdrawalIndex'
	b.NextWithdrawalIndex = ssz.UnmarshallUint64(buf[272:280])

	// Field (13) 'NextWithdrawalValidatorIndex'
	b.NextWithdrawalValidatorIndex = math.ValidatorIndex(ssz.UnmarshallUint64(buf[280:288]))

	// Offset (14) 'Slashings'
	if o14 = ssz.ReadOffset(buf[288:292]); o14 > size || o11 > o14 {
		return ssz.ErrOffset
	}

	// Field (15) 'TotalSlashing'
	b.TotalSlashing = math.Gwei(ssz.UnmarshallUint64(buf[292:300]))

	// Field (4) 'BlockRoots'
	{
		buf = tail[o4:o5]
		num, err := ssz.DivideInt2(len(buf), 32, 8192)
		if err != nil {
			return err
		}
		b.BlockRoots = make([]primitives.Root, num)
		for ii := 0; ii < num; ii++ {
			copy(b.BlockRoots[ii][:], buf[ii*32:(ii+1)*32])
		}
	}

	// Field (5) 'StateRoots'
	{
		buf = tail[o5:o6]
		num, err := ssz.DivideInt2(len(buf), 32, 8192)
		if err != nil {
			return err
		}
		b.StateRoots = make([]primitives.Root, num)
		for ii := 0; ii < num; ii++ {
			copy(b.StateRoots[ii][:], buf[ii*32:(ii+1)*32])
		}
	}

	// Field (6) 'LatestExecutionPayload'
	{
		buf = tail[o6:o9]
		if b.LatestExecutionPayload == nil {
			b.LatestExecutionPayload = new(engineprimitives.ExecutableDataDeneb)
		}
		if err = b.LatestExecutionPayload.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (9) 'Validators'
	{
		buf = tail[o9:o10]
		num, err := ssz.DivideInt2(len(buf), 121, 1099511627776)
		if err != nil {
			return err
		}
		b.Validators = make([]*primitives.Validator, num)
		for ii := 0; ii < num; ii++ {
			if b.Validators[ii] == nil {
				b.Validators[ii] = new(primitives.Validator)
			}
			if err = b.Validators[ii].UnmarshalSSZ(buf[ii*121 : (ii+1)*121]); err != nil {
				return err
			}
		}
	}

	// Field (10) 'Balances'
	{
		buf = tail[o10:o11]
		num, err := ssz.DivideInt2(len(buf), 8, 1099511627776)
		if err != nil {
			return err
		}
		b.Balances = ssz.ExtendUint64(b.Balances, num)
		for ii := 0; ii < num; ii++ {
			b.Balances[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}

	// Field (11) 'RandaoMixes'
	{
		buf = tail[o11:o14]
		num, err := ssz.DivideInt2(len(buf), 32, 65536)
		if err != nil {
			return err
		}
		b.RandaoMixes = make([]primitives.Bytes32, num)
		for ii := 0; ii < num; ii++ {
			copy(b.RandaoMixes[ii][:], buf[ii*32:(ii+1)*32])
		}
	}

	// Field (14) 'Slashings'
	{
		buf = tail[o14:]
		num, err := ssz.DivideInt2(len(buf), 8, 1099511627776)
		if err != nil {
			return err
		}
		b.Slashings = ssz.ExtendUint64(b.Slashings, num)
		for ii := 0; ii < num; ii++ {
			b.Slashings[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconState object
func (b *BeaconState) SizeSSZ() (size int) {
	size = 300

	// Field (4) 'BlockRoots'
	size += len(b.BlockRoots) * 32

	// Field (5) 'StateRoots'
	size += len(b.StateRoots) * 32

	// Field (6) 'LatestExecutionPayload'
	if b.LatestExecutionPayload == nil {
		b.LatestExecutionPayload = new(engineprimitives.ExecutableDataDeneb)
	}
	size += b.LatestExecutionPayload.SizeSSZ()

	// Field (9) 'Validators'
	size += len(b.Validators) * 121

	// Field (10) 'Balances'
	size += len(b.Balances) * 8

	// Field (11) 'RandaoMixes'
	size += len(b.RandaoMixes) * 32

	// Field (14) 'Slashings'
	size += len(b.Slashings) * 8

	return
}

// HashTreeRoot ssz hashes the BeaconState object
func (b *BeaconState) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconState object with a hasher
func (b *BeaconState) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'GenesisValidatorsRoot'
	hh.PutBytes(b.GenesisValidatorsRoot[:])

	// Field (1) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (2) 'Fork'
	if b.Fork == nil {
		b.Fork = new(primitives.Fork)
	}
	if err = b.Fork.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (3) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(primitives.BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'BlockRoots'
	{
		if size := len(b.BlockRoots); size > 8192 {
			err = ssz.ErrListTooBigFn("BeaconState.BlockRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlockRoots {
			hh.Append(i[:])
		}
		numItems := uint64(len(b.BlockRoots))
		hh.MerkleizeWithMixin(subIndx, numItems, 8192)
	}

	// Field (5) 'StateRoots'
	{
		if size := len(b.StateRoots); size > 8192 {
			err = ssz.ErrListTooBigFn("BeaconState.StateRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.StateRoots {
			hh.Append(i[:])
		}
		numItems := uint64(len(b.StateRoots))
		hh.MerkleizeWithMixin(subIndx, numItems, 8192)
	}

	// Field (6) 'LatestExecutionPayload'
	if err = b.LatestExecutionPayload.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (7) 'Eth1Data'
	if b.Eth1Data == nil {
		b.Eth1Data = new(primitives.Eth1Data)
	}
	if err = b.Eth1Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (8) 'Eth1DepositIndex'
	hh.PutUint64(b.Eth1DepositIndex)

	// Field (9) 'Validators'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Validators))
		if num > 1099511627776 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Validators {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1099511627776)
	}

	// Field (10) 'Balances'
	{
		if size := len(b.Balances); size > 1099511627776 {
			err = ssz.ErrListTooBigFn("BeaconState.Balances", size, 1099511627776)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Balances {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.Balances))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(1099511627776, numItems, 8))
	}

	// Field (11) 'RandaoMixes'
	{
		if size := len(b.RandaoMixes); size > 65536 {
			err = ssz.ErrListTooBigFn("BeaconState.RandaoMixes", size, 65536)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.RandaoMixes {
			hh.Append(i[:])
		}
		numItems := uint64(len(b.RandaoMixes))
		hh.MerkleizeWithMixin(subIndx, numItems, 65536)
	}

	// Field (12) 'NextWithdrawalIndex'
	hh.PutUint64(b.NextWithdrawalIndex)

	// Field (13) 'NextWithdrawalValidatorIndex'
	hh.PutUint64(uint64(b.NextWithdrawalValidatorIndex))

	// Field (14) 'Slashings'
	{
		if size := len(b.Slashings); size > 1099511627776 {
			err = ssz.ErrListTooBigFn("BeaconState.Slashings", size, 1099511627776)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Slashings {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.Slashings))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(1099511627776, numItems, 8))
	}

	// Field (15) 'TotalSlashing'
	hh.PutUint64(uint64(b.TotalSlashing))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconState object
func (b *BeaconState) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

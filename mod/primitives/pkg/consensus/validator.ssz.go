// Code generated by fastssz. DO NOT EDIT.
// Hash: 7e2f6d13a5f863ba763ba81c898e262278973e416996def3c348aac9445f4100
// Version: 0.1.3
package consensus

import (
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Validator object
func (v *Validator) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the Validator object to a target array
func (v *Validator) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Pubkey'
	dst = append(dst, v.Pubkey[:]...)

	// Field (1) 'WithdrawalCredentials'
	dst = append(dst, v.WithdrawalCredentials[:]...)

	// Field (2) 'EffectiveBalance'
	dst = ssz.MarshalUint64(dst, uint64(v.EffectiveBalance))

	// Field (3) 'Slashed'
	dst = ssz.MarshalBool(dst, v.Slashed)

	// Field (4) 'ActivationEligibilityEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.ActivationEligibilityEpoch))

	// Field (5) 'ActivationEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.ActivationEpoch))

	// Field (6) 'ExitEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.ExitEpoch))

	// Field (7) 'WithdrawableEpoch'
	dst = ssz.MarshalUint64(dst, uint64(v.WithdrawableEpoch))

	return
}

// UnmarshalSSZ ssz unmarshals the Validator object
func (v *Validator) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 121 {
		return ssz.ErrSize
	}

	// Field (0) 'Pubkey'
	copy(v.Pubkey[:], buf[0:48])

	// Field (1) 'WithdrawalCredentials'
	copy(v.WithdrawalCredentials[:], buf[48:80])

	// Field (2) 'EffectiveBalance'
	v.EffectiveBalance = math.Gwei(ssz.UnmarshallUint64(buf[80:88]))

	// Field (3) 'Slashed'
	v.Slashed = ssz.UnmarshalBool(buf[88:89])

	// Field (4) 'ActivationEligibilityEpoch'
	v.ActivationEligibilityEpoch = math.Epoch(ssz.UnmarshallUint64(buf[89:97]))

	// Field (5) 'ActivationEpoch'
	v.ActivationEpoch = math.Epoch(ssz.UnmarshallUint64(buf[97:105]))

	// Field (6) 'ExitEpoch'
	v.ExitEpoch = math.Epoch(ssz.UnmarshallUint64(buf[105:113]))

	// Field (7) 'WithdrawableEpoch'
	v.WithdrawableEpoch = math.Epoch(ssz.UnmarshallUint64(buf[113:121]))

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Validator object
func (v *Validator) SizeSSZ() (size int) {
	size = 121
	return
}

// HashTreeRoot ssz hashes the Validator object
func (v *Validator) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the Validator object with a hasher
func (v *Validator) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkey'
	hh.PutBytes(v.Pubkey[:])

	// Field (1) 'WithdrawalCredentials'
	hh.PutBytes(v.WithdrawalCredentials[:])

	// Field (2) 'EffectiveBalance'
	hh.PutUint64(uint64(v.EffectiveBalance))

	// Field (3) 'Slashed'
	hh.PutBool(v.Slashed)

	// Field (4) 'ActivationEligibilityEpoch'
	hh.PutUint64(uint64(v.ActivationEligibilityEpoch))

	// Field (5) 'ActivationEpoch'
	hh.PutUint64(uint64(v.ActivationEpoch))

	// Field (6) 'ExitEpoch'
	hh.PutUint64(uint64(v.ExitEpoch))

	// Field (7) 'WithdrawableEpoch'
	hh.PutUint64(uint64(v.WithdrawableEpoch))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Validator object
func (v *Validator) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(v)
}

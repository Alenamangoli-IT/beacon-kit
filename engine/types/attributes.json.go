// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package enginetypes

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	enginev1 "github.com/prysmaticlabs/prysm/v5/proto/engine/v1"
)

var _ = (*payloadAttributesJSONMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (p PayloadAttributes) MarshalJSON() ([]byte, error) {
	type PayloadAttributes struct {
		Timestamp             hexutil.Uint64         `json:"timestamp"             gencodec:"required"`
		PrevRandao            hexutil.Bytes          `json:"prevRandao"            gencodec:"required"`
		SuggestedFeeRecipient common.Address         `json:"suggestedFeeRecipient" gencodec:"required"`
		Withdrawals           []*enginev1.Withdrawal `json:"withdrawals"`
		ParentBeaconBlockRoot *common.Hash           `json:"parentBeaconBlockRoot"`
	}
	var enc PayloadAttributes
	enc.Timestamp = hexutil.Uint64(p.Timestamp)
	enc.PrevRandao = p.PrevRandao[:]
	enc.SuggestedFeeRecipient = p.SuggestedFeeRecipient
	enc.Withdrawals = p.Withdrawals
	enc.ParentBeaconBlockRoot = (*common.Hash)(p.ParentBeaconBlockRoot)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (p *PayloadAttributes) UnmarshalJSON(input []byte) error {
	type PayloadAttributes struct {
		Timestamp             *hexutil.Uint64        `json:"timestamp"             gencodec:"required"`
		PrevRandao            *hexutil.Bytes         `json:"prevRandao"            gencodec:"required"`
		SuggestedFeeRecipient *common.Address        `json:"suggestedFeeRecipient" gencodec:"required"`
		Withdrawals           []*enginev1.Withdrawal `json:"withdrawals"`
		ParentBeaconBlockRoot *common.Hash           `json:"parentBeaconBlockRoot"`
	}
	var dec PayloadAttributes
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for PayloadAttributes")
	}
	p.Timestamp = uint64(*dec.Timestamp)
	if dec.PrevRandao == nil {
		return errors.New("missing required field 'prevRandao' for PayloadAttributes")
	}
	if len(*dec.PrevRandao) != len(p.PrevRandao) {
		return errors.New("field 'prevRandao' has wrong length, need 32 items")
	}
	copy(p.PrevRandao[:], *dec.PrevRandao)
	if dec.SuggestedFeeRecipient == nil {
		return errors.New("missing required field 'suggestedFeeRecipient' for PayloadAttributes")
	}
	p.SuggestedFeeRecipient = *dec.SuggestedFeeRecipient
	if dec.Withdrawals != nil {
		p.Withdrawals = dec.Withdrawals
	}
	if dec.ParentBeaconBlockRoot != nil {
		p.ParentBeaconBlockRoot = (*[32]byte)(dec.ParentBeaconBlockRoot)
	}
	return nil
}

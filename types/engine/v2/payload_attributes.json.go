// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package v2

import (
	"encoding/json"

	v1 "github.com/itsdevbear/bolaris/types/engine/v1"
)

// MarshalJSON marshals as JSON.
func (p PayloadAttributesV2) MarshalJSON() ([]byte, error) {
	type PayloadAttributesV2 struct {
		Timestamp             uint64           `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
		PrevRandao            []byte           `protobuf:"bytes,2,opt,name=prev_randao,json=prevRandao,proto3" json:"prevRandao,omitempty" ssz-size:"32"`
		SuggestedFeeRecipient []byte           `protobuf:"bytes,3,opt,name=suggested_fee_recipient,json=suggestedFeeRecipient,proto3" json:"suggestedFeeRecipient,omitempty" ssz-size:"20"`
		Withdrawals           []*v1.Withdrawal `protobuf:"bytes,4,rep,name=withdrawals,proto3" json:"withdrawals,omitempty" ssz-max:"4"`
	}
	var enc PayloadAttributesV2
	enc.Timestamp = p.Timestamp
	enc.PrevRandao = p.PrevRandao
	enc.SuggestedFeeRecipient = p.SuggestedFeeRecipient
	enc.Withdrawals = p.Withdrawals
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (p *PayloadAttributesV2) UnmarshalJSON(input []byte) error {
	type PayloadAttributesV2 struct {
		Timestamp             *uint64          `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
		PrevRandao            []byte           `protobuf:"bytes,2,opt,name=prev_randao,json=prevRandao,proto3" json:"prevRandao,omitempty" ssz-size:"32"`
		SuggestedFeeRecipient []byte           `protobuf:"bytes,3,opt,name=suggested_fee_recipient,json=suggestedFeeRecipient,proto3" json:"suggestedFeeRecipient,omitempty" ssz-size:"20"`
		Withdrawals           []*v1.Withdrawal `protobuf:"bytes,4,rep,name=withdrawals,proto3" json:"withdrawals,omitempty" ssz-max:"4"`
	}
	var dec PayloadAttributesV2
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Timestamp != nil {
		p.Timestamp = *dec.Timestamp
	}
	if dec.PrevRandao != nil {
		p.PrevRandao = dec.PrevRandao
	}
	if dec.SuggestedFeeRecipient != nil {
		p.SuggestedFeeRecipient = dec.SuggestedFeeRecipient
	}
	if dec.Withdrawals != nil {
		p.Withdrawals = dec.Withdrawals
	}
	return nil
}

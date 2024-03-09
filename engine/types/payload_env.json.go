// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package enginetypes

import (
	"encoding/json"
	"github.com/cockroachdb/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*executionPayloadEnvelopeMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (e ExecutionPayloadEnvelopeDeneb) MarshalJSON() ([]byte, error) {
	type ExecutionPayloadEnvelopeDeneb struct {
		ExecutionPayload *ExecutableDataDeneb  `json:"executionPayload"      gencodec:"required"`
		BlockValue       *hexutil.Big          `json:"blockValue"            gencodec:"required"`
		BlobsBundle      *engine.BlobsBundleV1 `json:"blobsBundle"`
		Override         bool                  `json:"shouldOverrideBuilder"`
	}
	var enc ExecutionPayloadEnvelopeDeneb
	enc.ExecutionPayload = e.ExecutionPayload
	enc.BlockValue = (*hexutil.Big)(e.BlockValue)
	enc.BlobsBundle = e.BlobsBundle
	enc.Override = e.Override
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (e *ExecutionPayloadEnvelopeDeneb) UnmarshalJSON(input []byte) error {
	type ExecutionPayloadEnvelopeDeneb struct {
		ExecutionPayload *ExecutableDataDeneb  `json:"executionPayload"      gencodec:"required"`
		BlockValue       *hexutil.Big          `json:"blockValue"            gencodec:"required"`
		BlobsBundle      *engine.BlobsBundleV1 `json:"blobsBundle"`
		Override         *bool                 `json:"shouldOverrideBuilder"`
	}
	var dec ExecutionPayloadEnvelopeDeneb
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ExecutionPayload == nil {
		return errors.New("missing required field 'executionPayload' for ExecutionPayloadEnvelopeDeneb")
	}
	e.ExecutionPayload = dec.ExecutionPayload
	if dec.BlockValue == nil {
		return errors.New("missing required field 'blockValue' for ExecutionPayloadEnvelopeDeneb")
	}
	e.BlockValue = (*big.Int)(dec.BlockValue)
	if dec.BlobsBundle != nil {
		e.BlobsBundle = dec.BlobsBundle
	}
	if dec.Override != nil {
		e.Override = *dec.Override
	}
	return nil
}

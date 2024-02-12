// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package engine

import (
	"context"
	"time"

	"cosmossdk.io/log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/itsdevbear/bolaris/config"
	eth "github.com/itsdevbear/bolaris/execution/engine/ethclient"
	"github.com/itsdevbear/bolaris/types/consensus/primitives"
	"github.com/itsdevbear/bolaris/types/consensus/version"
	"github.com/itsdevbear/bolaris/types/engine"
	"github.com/itsdevbear/bolaris/types/engine/interfaces"
	"github.com/pkg/errors"
	enginev1 "github.com/prysmaticlabs/prysm/v4/proto/engine/v1"
)

// Caller is implemented by engineClient.
var _ Caller = (*engineClient)(nil)

// engineClient is a struct that holds a pointer to an Eth1Client.
type engineClient struct {
	*eth.Eth1Client
	engineTimeout time.Duration
	beaconCfg     *config.Beacon
	logger        log.Logger
}

// NewClient creates a new engine client engineClient.
// It takes an Eth1Client as an argument and returns a pointer to an engineClient.
func NewClient(opts ...Option) Caller {
	ec := &engineClient{}
	for _, opt := range opts {
		if err := opt(ec); err != nil {
			panic(err)
		}
	}

	return ec
}

// NewPayload calls the engine_newPayloadVX method via JSON-RPC.
func (s *engineClient) NewPayload(
	ctx context.Context, payload interfaces.ExecutionPayload,
	versionedHashes []common.Hash, parentBlockRoot *common.Hash,
) ([]byte, error) {
	dctx, cancel := context.WithTimeout(ctx, s.engineTimeout)
	defer cancel()

	payloadPb, err := s.getPayloadProto(payload)
	if err != nil {
		return nil, err
	}

	result, err := s.callNewPayloadRPC(dctx, payloadPb, versionedHashes, parentBlockRoot)
	if err != nil {
		return nil, err
	}

	// This case is only true when the payload is invalid, so `processPayloadStatusResult` below
	// will return an error.
	if validationErr := result.GetValidationError(); validationErr != "" {
		s.logger.Error("Got a validation error in newPayload", "err", errors.New(validationErr))
	}

	return processPayloadStatusResult(result)
}

// getPayloadProto returns the payload proto from the execution data.
func (s *engineClient) getPayloadProto(payload interfaces.ExecutionPayload) (interface{}, error) {
	switch payloadPb := payload.ToProto().(type) {
	case *enginev1.ExecutionPayloadCapella:
		return payloadPb, nil
	case *enginev1.ExecutionPayloadDeneb:
		return payloadPb, nil
	default:
		return nil, errors.New("unknown execution data type")
	}
}

// callNewPayloadRPC calls the engine_newPayloadVX method via JSON-RPC.
func (s *engineClient) callNewPayloadRPC(
	ctx context.Context, payloadPb interface{},
	versionedHashes []common.Hash, parentBlockRoot *common.Hash,
) (*enginev1.PayloadStatus, error) {
	switch payloadPb := payloadPb.(type) {
	case *enginev1.ExecutionPayloadCapella:
		return s.NewPayloadV2(ctx, payloadPb)
	case *enginev1.ExecutionPayloadDeneb:
		return s.NewPayloadV3(ctx, payloadPb, versionedHashes, parentBlockRoot)
	default:
		return nil, errors.New("invalid payload type for RPC call")
	}
}

// ForkchoiceUpdated calls the engine_forkchoiceUpdatedV1 method via JSON-RPC.
func (s *engineClient) ForkchoiceUpdated(
	ctx context.Context, state *enginev1.ForkchoiceState, attrs interfaces.PayloadAttributer,
) (*enginev1.PayloadIDBytes, []byte, error) {
	dctx, cancel := context.WithTimeout(ctx, s.engineTimeout)
	defer cancel()

	result, err := s.updateForkChoiceByVersion(dctx, state, attrs)
	if err != nil {
		return nil, nil, err
	}

	lastestValidHash, err := processPayloadStatusResult(result.Status)
	if err != nil {
		return nil, lastestValidHash, err
	}
	return result.PayloadID, lastestValidHash, nil
}

// updateForkChoiceByVersion calls the engine_forkchoiceUpdatedVX method via JSON-RPC.
func (s *engineClient) updateForkChoiceByVersion(
	ctx context.Context, state *enginev1.ForkchoiceState,
	attrProto interfaces.PayloadAttributer,
) (*eth.ForkchoiceUpdatedResponse, error) {
	switch v := attrProto.ToProto().(type) {
	case *enginev1.PayloadAttributesV3:
		return s.ForkchoiceUpdatedV3(ctx, state, v)
	case *enginev1.PayloadAttributesV2:
		return s.ForkchoiceUpdatedV2(ctx, state, v)
	default:
		return nil, errors.New("invalid payload attribute version")
	}
}

// GetPayload calls the engine_getPayloadVX method via JSON-RPC.
// It returns the execution data as well as the blobs bundle.
func (s *engineClient) GetPayload(
	ctx context.Context, payloadID primitives.PayloadID, slot primitives.Slot,
) (interfaces.ExecutionPayload, *enginev1.BlobsBundle, bool, error) {
	dctx, cancel := context.WithTimeout(ctx, s.engineTimeout)
	defer cancel()

	switch s.beaconCfg.ActiveForkVersion(primitives.Epoch(slot)) {
	case version.Deneb:
		return s.getPayloadDeneb(dctx, payloadID)
	case version.Capella:
		return s.getPayloadCapella(ctx, payloadID)
	default:
		return nil, nil, false, errors.New("unknown fork version")
	}
}

// handleDenebFork processes the Deneb fork version.
func (s *engineClient) getPayloadDeneb(
	ctx context.Context, payloadID primitives.PayloadID,
) (interfaces.ExecutionPayload, *enginev1.BlobsBundle, bool, error) {
	result, err := s.GetPayloadV3(ctx, enginev1.PayloadIDBytes(payloadID))
	if err != nil {
		return nil, nil, false, err
	}
	ed, err := engine.WrappedExecutionPayloadDeneb(
		result.GetPayload(), PayloadValueToWei(result.GetValue()),
	)
	if err != nil {
		return nil, nil, false, err
	}

	return ed, result.GetBlobsBundle(), result.GetShouldOverrideBuilder(), nil
}

// handleCapellaFork processes the Capella fork version.
func (s *engineClient) getPayloadCapella(
	ctx context.Context, payloadID primitives.PayloadID,
) (interfaces.ExecutionPayload, *enginev1.BlobsBundle, bool, error) {
	result, err := s.GetPayloadV2(ctx, enginev1.PayloadIDBytes(payloadID))
	if err != nil {
		return nil, nil, false, err
	}
	ed, err := engine.WrappedExecutionPayloadCapella(
		result.GetPayload(), PayloadValueToWei(result.GetValue()),
	)
	if err != nil {
		return nil, nil, false, err
	}

	return ed, nil, false, nil
}

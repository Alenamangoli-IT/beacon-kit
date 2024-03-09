// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
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

package types

import (
	"context"
	"github.com/cockroachdb/errors"
	"math/big"

	"cosmossdk.io/log"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/kurtosis-tech/kurtosis/api/golang/core/lib/services"
)

// ExecutionClient represents an execution client.
type ExecutionClient struct {
	*services.ServiceContext
	*JSONRPCConnection
	logger log.Logger
}

// NewExecutionClientFromServiceCtx creates a new execution client from a
// service context.
func NewExecutionClientFromServiceCtx(
	serviceCtx *services.ServiceContext,
	logger log.Logger,
) (*ExecutionClient, error) {
	jsonRPCConn, err := NewJSONRPCConnection(serviceCtx)
	if err != nil {
		return nil, err
	}

	return &ExecutionClient{
		ServiceContext:    serviceCtx,
		JSONRPCConnection: jsonRPCConn,
		logger: logger.With(
			"client-name",
			serviceCtx.GetServiceName(),
		),
	}, nil
}

// IsValidator returns true if the execution client is a validator.
// TODO: All nodes are validators rn.
func (c *ExecutionClient) IsValidator() bool {
	return true
}

// WaitForLatestBlockNumber waits for the head block number to reach the target.
func (c *ExecutionClient) WaitForLatestBlockNumber(
	ctx context.Context,
	target uint64,
) error {
	if !c.JSONRPCConnection.isWebSocket {
		return errors.New(
			"cannot wait for block number on non-websocket connection",
		)
	}

	ch := make(chan *coretypes.Header)
	sub, err := c.SubscribeNewHead(ctx, ch)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			sub.Unsubscribe()
			return ctx.Err()
		case header := <-ch:
			c.logger.Info(
				"received new head block",
				"number",
				header.Number.Uint64(),
				"",
			)
			if header.Number.Uint64() >= target {
				c.logger.Info("reached target block number 🎉", "target", target)
				sub.Unsubscribe()
				return nil
			}
		}
	}
}

// WaitForFinalizedBlockNumber waits for the finalized block number to reach the
// target block number.
func (c *ExecutionClient) WaitForFinalizedBlockNumber(
	ctx context.Context,
	target uint64,
) error {
	finalQuery := big.NewInt(int64(rpc.FinalizedBlockNumber))
retry:
	// Retry until the head block number is at least the target.
	if err := c.WaitForLatestBlockNumber(ctx, target+1); err != nil {
		return err
	}

	// Just to be safe, check the finalized block number again.
	finalized, err := c.BlockByNumber(ctx, finalQuery)
	if err != nil {
		return err
	}

	// If the finalized block number is less than the target, retry.
	if finalized.Number().Uint64() < target {
		goto retry
	}
	return nil
}

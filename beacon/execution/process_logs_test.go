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

package execution_test

import (
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/itsdevbear/bolaris/beacon/execution"
	"github.com/itsdevbear/bolaris/beacon/execution/logs"
	logmocks "github.com/itsdevbear/bolaris/beacon/execution/logs/mocks"
	"github.com/itsdevbear/bolaris/beacon/execution/mocks"
	"github.com/itsdevbear/bolaris/runtime/service"
	consensusv1 "github.com/itsdevbear/bolaris/types/consensus/v1"
	"github.com/stretchr/testify/require"
)

func TestProcessLogs(t *testing.T) {
	contractAddress := ethcommon.HexToAddress("0x1234")

	mockLogService := &logmocks.Service{}
	logFactory, err := logs.NewFactory(
		logs.WithRequestsFrom(mockLogService),
	)
	require.NoError(t, err)

	executionService := service.New[execution.Service](
		execution.WithLogFactory(logFactory),
	)

	blkNum := uint64(100)
	depositFactor := 3
	numDepositLogs := 10
	logs, err := mocks.CreateDepositLogs(
		numDepositLogs,
		depositFactor,
		contractAddress,
		blkNum,
	)
	require.NoError(t, err)

	vals, err := executionService.ProcessLogs(logs, blkNum)
	require.NoError(t, err)
	require.Len(t, vals, numDepositLogs)

	// Check if the values are returned in the correct order.
	for i, val := range vals {
		processedDeposit, ok := val.Interface().(*consensusv1.Deposit)
		require.True(t, ok)
		require.Equal(t, uint64(i*depositFactor), processedDeposit.GetAmount())
	}
}

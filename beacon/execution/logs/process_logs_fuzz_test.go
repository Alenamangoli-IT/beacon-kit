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

package logs_test

import (
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	beacontypesv1 "github.com/itsdevbear/bolaris/beacon/core/types/v1"
	executionlogs "github.com/itsdevbear/bolaris/beacon/execution/logs"
	stakinglogs "github.com/itsdevbear/bolaris/beacon/staking/logs"
	logmocks "github.com/itsdevbear/bolaris/beacon/staking/logs/mocks"
	"github.com/stretchr/testify/require"
)

func FuzzProcessStakingLogs(f *testing.F) {
	contractAddress := ethcommon.HexToAddress("0x1234")

	stakingLogRequest, err := stakinglogs.NewStakingRequest(
		contractAddress,
	)
	require.NoError(f, err)
	logFactory, err := executionlogs.NewFactory(
		executionlogs.WithRequest(stakingLogRequest),
	)
	require.NoError(f, err)

	f.Add(uint64(100), 3, 10)
	f.Fuzz(
		func(t *testing.T, blkNum uint64, depositFactor, numDepositLogs int) {
			if depositFactor <= 0 || numDepositLogs <= 0 {
				t.Skip()
			}

			var logs []ethtypes.Log
			logs, err = logmocks.CreateDepositLogs(
				numDepositLogs,
				depositFactor,
				contractAddress,
				blkNum,
			)
			require.NoError(t, err)

			var containers []executionlogs.LogContainer
			containers, err = logFactory.ProcessLogs(logs, blkNum)
			require.NoError(t, err)
			require.Len(t, containers, numDepositLogs)

			// Check if the returned values are correct.
			for i, container := range containers {
				val := container.Value()
				processedDeposit, ok := val.Interface().(*beacontypesv1.Deposit)
				require.True(t, ok)
				require.Equal(
					t,
					uint64(i*depositFactor),
					processedDeposit.GetAmount(),
				)
			}
		},
	)
}

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
	loghandler "github.com/itsdevbear/bolaris/beacon/execution/logs"
	"github.com/itsdevbear/bolaris/beacon/staking/logs"
	logmocks "github.com/itsdevbear/bolaris/beacon/staking/logs/mocks"
	"github.com/itsdevbear/bolaris/contracts/abi"
	enginetypes "github.com/itsdevbear/bolaris/engine/types"
	"github.com/stretchr/testify/require"
)

func TestProcessStakingLogs(t *testing.T) {
	contractAddress := ethcommon.HexToAddress("0x1234")
	depositContractAbi, err := abi.BeaconDepositContractMetaData.GetAbi()
	require.NoError(t, err)
	require.NotNil(t, depositContractAbi)

	stakingLogRequest, err := logs.NewStakingRequest(
		contractAddress,
	)
	require.NoError(t, err)
	logFactory, err := loghandler.NewFactory(
		loghandler.WithRequest(stakingLogRequest),
	)
	require.NoError(t, err)

	blkNum := uint64(100)
	depositFactor := 3
	numDepositLogs := 10
	mockLogs, err := logmocks.CreateDepositLogs(
		numDepositLogs,
		depositFactor,
		contractAddress,
		blkNum,
	)
	require.NoError(t, err)

	blockNumToHash := make(map[uint64]ethcommon.Hash)
	blockNumToHash[blkNum] = ethcommon.BytesToHash([]byte{byte(blkNum)})
	containers, err := logFactory.ProcessLogs(mockLogs, blockNumToHash)
	require.NoError(t, err)
	require.Len(t, containers, numDepositLogs)

	// Check if the values are returned in the correct order.
	for i, container := range containers {
		val := container.Value()
		processedDeposit, ok := val.Interface().(*beacontypesv1.Deposit)
		require.True(t, ok)
		require.Equal(t, uint64(i*depositFactor), processedDeposit.GetAmount())
	}

	withdrawal := enginetypes.NewWithdrawal(
		[]byte("pubkey"),
		uint64(1000),
	)

	var log *ethtypes.Log
	require.NotNil(t, depositContractAbi)
	require.NotNil(t, depositContractAbi.Events)
	log, err = logmocks.NewLogFromWithdrawal(
		depositContractAbi.Events[logs.WithdrawalName],
		withdrawal,
	)
	require.NoError(t, err)

	log.Address = contractAddress
	log.BlockNumber = blkNum + 1
	blockHash := ethcommon.BytesToHash([]byte{byte(blkNum + 1)})
	log.BlockHash = blockHash
	blockNumToHash[log.BlockNumber] = blockHash
	_, err = logFactory.ProcessLogs(
		append(mockLogs, *log),
		blockNumToHash,
	)
	// This is an expected error as
	// the log is from a different block.
	require.Error(t, err)

	log.Address = ethcommon.HexToAddress("0x5678")
	// This log has an incorrect block hash.
	// The hash is (blkNum + 1) in blockNumToHash.
	log.BlockHash = ethcommon.BytesToHash([]byte{byte(blkNum)})
	mockLogs = append(mockLogs, *log)
	_, err = logFactory.ProcessLogs(mockLogs, blockNumToHash)
	require.Error(t, err)

	log.Address = contractAddress
	log.BlockNumber = blkNum
	log.BlockHash = ethcommon.BytesToHash([]byte{byte(blkNum + 1)})
	mockLogs = append(mockLogs, *log)
	_, err = logFactory.ProcessLogs(mockLogs, blockNumToHash)
	// This is an expected error as currently we cannot
	// unmarsal a withdrawal log into a Withdrawal object.
	require.Error(t, err)
}

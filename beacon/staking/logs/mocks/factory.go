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

package mocks

import (
	"errors"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	beacontypesv1 "github.com/itsdevbear/bolaris/beacon/core/types/v1"
	enginev1 "github.com/itsdevbear/bolaris/engine/types/v1"
)

// NewLogFromDeposit creates a new log from the given deposit.
func NewLogFromDeposit(
	event ethabi.Event,
	deposit *beacontypesv1.Deposit,
) (*coretypes.Log, error) {
	return newLog(event,
		deposit.GetValidatorPubkey(),
		deposit.GetWithdrawalCredentials(),
		deposit.GetAmount(),
		deposit.GetSignature(),
	)
}

// NewLogFromWithdrawal creates a new log from the given withdrawal.
func NewLogFromWithdrawal(
	event ethabi.Event,
	withdrawal *enginev1.Withdrawal,
) (*coretypes.Log, error) {
	return newLog(event,
		[]byte{},
		[]byte{},
		withdrawal.GetAmount(),
	)
}

// newLog creates a new log of an event from the given arguments.
func newLog(event ethabi.Event, args ...interface{}) (*coretypes.Log, error) {
	if len(event.Inputs) != len(args) {
		return nil, errors.New("mismatched number of arguments")
	}
	data, err := event.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}
	return &coretypes.Log{
		Topics: []ethcommon.Hash{event.ID},
		Data:   data,
	}, nil
}

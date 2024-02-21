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
	"context"
	"errors"
	"sync"

	"github.com/itsdevbear/bolaris/beacon/staking/logs"
	consensusv1 "github.com/itsdevbear/bolaris/types/consensus/v1"
	enginev1 "github.com/itsdevbear/bolaris/types/engine/v1"
)

var _ logs.StakingService = &mockStakingService{}

// mockStakingService is a mock implementation of the staking service.
type mockStakingService struct {
	depositQueue    []*consensusv1.Deposit
	withdrawalQueue []*enginev1.Withdrawal
	mu              sync.RWMutex
}

// AcceptDeposit pushes a deposit into queue.
func (s *mockStakingService) AcceptDeposit(_ context.Context, deposit *consensusv1.Deposit) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.depositQueue = append(s.depositQueue, deposit)
	return nil
}

// ApplyDeposits does nothing but clears the queue,
// because mockStakingService does not have an underlying
// staking module.
func (s *mockStakingService) ApplyDeposits(_ context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.depositQueue = nil
	return nil
}

// mostRecentDeposit returns the most recent deposit added into the queue.
func (s *mockStakingService) mostRecentDeposit() (*consensusv1.Deposit, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.depositQueue) == 0 {
		return nil, errors.New("no deposits")
	}
	return s.depositQueue[len(s.depositQueue)-1], nil
}

// numPendingDeposits returns the number of pending deposits in the queue.
func (s *mockStakingService) numPendingDeposits() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.depositQueue)
}

// ProcessWithdrawal processes a withdrawal.
func (s *mockStakingService) ProcessWithdrawal(
	_ context.Context,
	withdrawal *enginev1.Withdrawal,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.withdrawalQueue = append(s.withdrawalQueue, withdrawal)
	return nil
}

// mostRecentWithdrawal returns the most recent withdrawal added into the queue.
func (s *mockStakingService) mostRecentWithdrawal() (*enginev1.Withdrawal, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.withdrawalQueue) == 0 {
		return nil, errors.New("no withdrawals")
	}
	return s.withdrawalQueue[len(s.withdrawalQueue)-1], nil
}

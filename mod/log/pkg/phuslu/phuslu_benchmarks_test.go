// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package phuslu_test

import (
	"bytes"
	"testing"

	"github.com/berachain/beacon-kit/mod/cli/pkg/builder"
	"github.com/cosmos/cosmos-sdk/server"
)

// Benchmark function for phuslu logger.
func BenchmarkPhusluLoggerInfo(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	logger, _ := builder.CreatePhusluLogger(serverCtx, &bytes.Buffer{})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Info("This is an info message", "key1", "value1", "key2", 2)
	}
}

// Benchmark function for SDK logger Info.
func BenchmarkSDKLoggerInfo(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	serverCtx.Viper.Set("log_level", "Debug")
	logger, _ := server.CreateSDKLogger(serverCtx, &bytes.Buffer{})
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Info("This is an info message", "key1", "value1", "key2", 2)
	}
}

// Benchmark function for phuslu logger Warn.
func BenchmarkPhusluLoggerWarn(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	logger, err := builder.CreatePhusluLogger(serverCtx, &bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create phuslu logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Warn("This is a warning message", "key1", "value1", "key2", 2)
	}
}

// Benchmark function for cosmos logger Warn.
func BenchmarkSDKLoggerWarn(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	logger, err := server.CreateSDKLogger(serverCtx, &bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create cosmos logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Warn("This is a warning message", "key1", "value1", "key2", 2)
	}
}

// Benchmark function for phuslu logger Error.
func BenchmarkPhusluLoggerError(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	logger, err := builder.CreatePhusluLogger(serverCtx, &bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create phuslu logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Error("This is an error message", "key1", "value1", "key2", 2)
	}
}

// Benchmark function for cosmos logger Error.
func BenchmarkSDKLoggerError(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	logger, err := server.CreateSDKLogger(server.NewDefaultContext(),
		&bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create cosmos logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Error("This is an error message", "key1", "value1", "key2", 2)
	}
}

// Benchmark function for phuslu logger Debug.
func BenchmarkPhusluLoggerDebug(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	logger, err := builder.CreatePhusluLogger(serverCtx, &bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create phuslu logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Debug("This is a debug message", "key1", "value1", "key2", 2)
	}
}

func BenchmarkPhusluLoggerDebugSilent(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Info")
	logger, err := builder.CreatePhusluLogger(serverCtx, &bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create phuslu logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Debug("This is a debug message", "key1", "value1", "key2", 2)
	}
}

// Benchmark function for cosmos logger Debug.
func BenchmarkSDKLoggerDebug(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	logger, err := server.CreateSDKLogger(serverCtx, &bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create cosmos logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Debug("This is a debug message", "key1", "value1", "key2", 2)
	}
}

func BenchmarkSDKLoggerDebugSilent(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Info")
	logger, err := server.CreateSDKLogger(serverCtx, &bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create cosmos logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		logger.Debug("This is a debug message", "key1", "value1", "key2", 2)
	}
}

// Benchmark function for phuslu logger With.
func BenchmarkPhusluLoggerWith(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	logger, err := builder.CreatePhusluLogger(serverCtx, &bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create phuslu logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		newLogger := logger.With("contextKey", "contextValue")
		newLogger.Info("This is a contextual info message", "key1", "value1",
			"key2", 2)
	}
}

// Benchmark function for cosmos logger With.
func BenchmarkSDKLoggerWith(b *testing.B) {
	serverCtx := server.NewDefaultContext()
	serverCtx.Viper.Set("log_level", "Debug")
	logger, err := server.CreateSDKLogger(serverCtx, &bytes.Buffer{})
	if err != nil {
		b.Fatalf("failed to create cosmos logger: %v", err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		newLogger := logger.With("contextKey", "contextValue")
		newLogger.Info("This is a contextual info message", "key1", "value1",
			"key2", 2)
	}
}
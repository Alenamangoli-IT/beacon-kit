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

package phuslu

import (
	"io"

	"github.com/phuslu/log"
)

// Logger is a wrapper around phuslogger.
type Logger[KeyValT, ImplT any] struct {
	logger *log.Logger
	// context is a map of key-value pairs that are added to every log entry.
	context log.Fields
}

// NewLogger creates a new logger with the given log level, ConsoleWriter, and
// default configuration.
func NewLogger[KeyValT, ImplT any](
	level string, out io.Writer) *Logger[KeyValT, ImplT] {
	cfg := DefaultConfig()
	logger := &log.Logger{
		Level:      log.ParseLevel(level),
		TimeFormat: cfg.TimeFormat,
		Writer: &log.ConsoleWriter{
			ColorOutput:    cfg.ColorOutput,
			QuoteString:    cfg.QuoteString,
			EndWithMessage: cfg.EndWithMessage,
			Writer:         out,
		},
	}
	return &Logger[KeyValT, ImplT]{
		logger: logger,
	}
}

// Info logs a message at level Info.
func (l *Logger[KeyValT, ImplT]) Info(msg string, keyVals ...any) {
	e := l.logger.Info()
	e = l.addContext(e)
	e = e.Fields(log.Fields(keyValToFields(keyVals...)))
	e.Msg(msg)
}

// Warn logs a message at level Warn.
func (l *Logger[KeyValT, ImplT]) Warn(msg string, keyVals ...any) {
	e := l.logger.Warn()
	e = l.addContext(e)
	e = e.Fields(log.Fields(keyValToFields(keyVals...)))
	e.Msg(msg)
}

// Error logs a message at level Error.
func (l *Logger[KeyValT, ImplT]) Error(msg string, keyVals ...any) {
	e := l.logger.Error()
	e = l.addContext(e)
	e = e.Fields(log.Fields(keyValToFields(keyVals...)))
	e.Msg(msg)
}

// Debug logs a message at level Debug.
func (l *Logger[KeyValT, ImplT]) Debug(msg string, keyVals ...any) {
	e := l.logger.Debug()
	e = l.addContext(e)
	e = e.Fields(log.Fields(keyValToFields(keyVals...)))
	e.Msg(msg)
}

// Impl returns the underlying logger implementation.
func (l *Logger[KeyValT, ImplT]) Impl() any {
	return l.logger
}

// With returns a new wrapped logger with additional context provided by a set.
func (l *Logger[KeyValT, ImplT]) With(keyVals ...any) ImplT {
	newLogger := *l
	fields := keyValToFields(keyVals...)
	// Copy existing context
	for k, v := range l.context {
		fields[k] = v
	}

	// return a new logger with the new fields
	newLogger.context = fields
	return any(&newLogger).(ImplT)
}

// addContext adds the context to the entry
func (l *Logger[KeyValT, ImplT]) addContext(e *log.Entry) *log.Entry {
	return e.Fields(log.Fields(l.context))
}

// keyValToFields converts a list of key-value pairs to a map.
func keyValToFields(keyVals ...any) log.Fields {
	if len(keyVals)%2 != 0 {
		panic("missing value for key")
	}
	// allocate a new fields map
	fields := make(log.Fields)
	// populate the fields map with the key-value pairs
	for i := 0; i < len(keyVals); i += 2 {
		fields[keyVals[i].(string)] = keyVals[i+1]
	}
	return fields
}

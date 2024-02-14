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

package types

import "fmt"

const (
	FailedToCreateDB                = "can't create a db: %w"
	FailedCheckForExistingProviders = "failed to retrieve primary or witness from db: %w"
	NoPrimaryAddress                = `no primary address was provided nor found.
		Please provide a primary (using -p).
		Run the command: tendermint light --help for more information`

	FailedSaveProviders   = "Unable to save primary and or witness addresses"
	FailedParseTrustLevel = "can't parse trust level: %w"
	FailedToCreateClient  = "failed to create a light client: %w"
	ListenAndServeError   = "failed to listen and serve: %w"
)

func NewError(msg string, err error) error {
	return fmt.Errorf(msg, err)
}

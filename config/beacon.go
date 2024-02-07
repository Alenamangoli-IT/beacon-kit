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

package config

import (
	"github.com/itsdevbear/bolaris/config/parser"
	"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	"github.com/prysmaticlabs/prysm/v4/runtime/version"
)

// Beacon conforms to the BeaconKitConfig interface.
var _ BeaconKitConfig[Beacon] = Beacon{}

// Beacon is the configuration for the beacon chain.
type Beacon struct {
	// Forks is the configuration for the beacon chain forks.
	Forks Forks
	// Validator is the configuration for the validator. Only utilized when
	// this node is in the active validator set.
	Validator Validator
}

// DefaultBeaconConfig returns the default fork configuration.
func DefaultBeaconConfig() Beacon {
	return Beacon{
		Forks:     DefaultForksConfig(),
		Validator: DefaultValidatorConfig(),
	}
}

// ActiveForkVersion returns the active fork version for a given slot.
func (c Beacon) ActiveForkVersion(epoch primitives.Epoch) int {
	if epoch >= c.Forks.DenebForkEpoch {
		return version.Deneb
	}

	// In BeaconKit we assume the Capella fork is always active.
	return version.Capella
}

// Template returns the configuration template.
func (c Beacon) Template() string {
	return c.Forks.Template() + c.Validator.Template()
}

// Parse parses the configuration.
func (c Beacon) Parse(parser parser.AppOptionsParser) (*Beacon, error) {
	var (
		err       error
		forks     *Forks
		validator *Validator
	)

	// Parse the forks configuration.
	if forks, err = c.Forks.Parse(parser); err != nil {
		return nil, err
	}
	c.Forks = *forks

	// Parse the validator configuration.
	if validator, err = c.Validator.Parse(parser); err != nil {
		return nil, err
	}
	c.Validator = *validator

	return &c, nil
}

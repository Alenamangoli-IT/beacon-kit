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

package commands

import (
	confixcmd "cosmossdk.io/tools/confix/cmd"
	"github.com/berachain/beacon-kit/mod/cli/pkg/commands/client"
	"github.com/berachain/beacon-kit/mod/cli/pkg/commands/cometbft"
	"github.com/berachain/beacon-kit/mod/cli/pkg/commands/deposit"
	"github.com/berachain/beacon-kit/mod/cli/pkg/commands/genesis"
	"github.com/berachain/beacon-kit/mod/cli/pkg/commands/jwt"
	"github.com/berachain/beacon-kit/mod/cli/pkg/flags"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/pruning"
	"github.com/cosmos/cosmos-sdk/client/snapshot"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cobra"
)

// Root is a wrapper around cobra.Command.
type Root struct {
	cmd *cobra.Command
}

// New returns a pointer to a new Root.
func New(rootcmd *cobra.Command) *Root {
	return &Root{
		cmd: rootcmd,
	}
}

// Run executes the root command.
func (root Root) Run(defaultNodeHome string) error {
	return svrcmd.Execute(
		root.cmd, "", defaultNodeHome,
	)
}

// DefaultRootCommandSetup sets up the default commands for the root command.
func DefaultRootCommandSetup[T servertypes.Application](
	root *Root,
	mm *module.Manager,
	appCreator servertypes.AppCreator[T],
	chainSpec primitives.ChainSpec,
) {
	// Setup the custom start command options.
	startCmdOptions := server.StartCmdOptions[T]{
		AddFlags: flags.AddBeaconKitFlags,
	}

	// Add all the commands to the root command.
	root.cmd.AddCommand(
		// `comet`
		cometbft.Commands(appCreator),
		// `client`
		client.Commands[T](),
		// `config`
		confixcmd.ConfigCommand(),
		// `init`
		genutilcli.InitCmd(mm),
		// `genesis`
		genesis.Commands(chainSpec),
		// `deposit`
		deposit.Commands(chainSpec),
		// `jwt`
		jwt.Commands(),
		// `keys`
		keys.Commands(),
		// `prune`
		pruning.Cmd(appCreator),
		// `rollback`
		server.NewRollbackCmd(appCreator),
		// `snapshots`
		snapshot.Cmd(appCreator),
		// `start`
		server.StartCmdWithOptions(appCreator, startCmdOptions),
		// `status`
		server.StatusCommand(),
		// `version`
		version.NewVersionCommand(),
	)
}

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

package deposit

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/url"
	"os"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/berachain/beacon-kit/mod/cli/pkg/utils/deposit"
	"github.com/berachain/beacon-kit/mod/cli/pkg/utils/parser"
	"github.com/berachain/beacon-kit/mod/config"
	"github.com/berachain/beacon-kit/mod/config/pkg/spec"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	engineclient "github.com/berachain/beacon-kit/mod/execution/pkg/client"
	gethprimitives "github.com/berachain/beacon-kit/mod/geth-primitives"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/components"
	"github.com/berachain/beacon-kit/mod/node-core/pkg/components/signer"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/net/jwt"
	myUrl "github.com/berachain/beacon-kit/mod/primitives/pkg/net/url"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

// NewValidateDeposit creates a new command for validating a deposit message.
//

func NewCreateValidator(chainSpec common.ChainSpec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-validator",
		Short: "Creates a validator deposit",
		Long: `Creates a validator deposit with the necessary credentials. The 
		arguments are expected in the order of withdrawal credentials, deposit
		amount, current version, and genesis validator root. If the broadcast
		flag is set to true, a private key must be provided to sign the transaction.`,
		Args: cobra.ExactArgs(4), //nolint:mnd // The number of arguments.
		RunE: createValidatorCmd(chainSpec),
	}

	cmd.Flags().BoolP(
		broadcastDeposit, broadcastDepositShorthand,
		defaultBroadcastDeposit, broadcastDepositMsg,
	)
	cmd.Flags().String(privateKey, defaultPrivateKey, privateKeyMsg)
	cmd.Flags().BoolP(
		overrideNodeKey, overrideNodeKeyShorthand,
		defaultOverrideNodeKey, overrideNodeKeyMsg,
	)
	cmd.Flags().
		String(valPrivateKey, defaultValidatorPrivateKey, valPrivateKeyMsg)
	cmd.Flags().String(jwtSecretPath, defaultJWTSecretPath, jwtSecretPathMsg)
	cmd.Flags().String(engineRPCURL, defaultEngineRPCURL, engineRPCURLMsg)

	return cmd
}

// createValidatorCmd returns a command that builds a create validator request.
//
// TODO: Implement broadcast functionality. Currently, the implementation works
// for the geth client but something about the Deposit binding is not handling
// other execution layers correctly. Peep the commit history for what we had.
// 🤷‍♂️.
func createValidatorCmd(
	chainSpec common.ChainSpec,
) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var (
			privKey *ecdsa.PrivateKey
			logger  = log.NewLogger(os.Stdout)
		)

		broadcast, err := cmd.Flags().GetBool(broadcastDeposit)
		if err != nil {
			return err
		}

		// If the broadcast flag is set, a private key must be provided.
		if broadcast {
			var fundingPrivKey string
			fundingPrivKey, err = cmd.Flags().GetString(privateKey)
			if err != nil {
				return err
			}
			if fundingPrivKey == "" {
				return parser.ErrPrivateKeyRequired
			}

			privKey, err = ethCrypto.HexToECDSA(fundingPrivKey)
			if err != nil {
				return err
			}
		}

		// Get the BLS signer.
		blsSigner, err := getBLSSigner(cmd)
		if err != nil {
			return err
		}

		credentials, err := parser.ConvertWithdrawalCredentials(args[0])
		if err != nil {
			return err
		}

		amount, err := parser.ConvertAmount(args[1])
		if err != nil {
			return err
		}

		currentVersion, err := parser.ConvertVersion(args[2])
		if err != nil {
			return err
		}

		genesisValidatorRoot, err := parser.ConvertGenesisValidatorRoot(args[3])
		if err != nil {
			return err
		}

		// Create and sign the deposit message.
		depositMsg, signature, err := types.CreateAndSignDepositMessage(
			types.NewForkData(currentVersion, genesisValidatorRoot),
			chainSpec.DomainTypeDeposit(),
			blsSigner,
			credentials,
			amount,
		)
		if err != nil {
			return err
		}

		// Verify the deposit message.
		if err = depositMsg.VerifyCreateValidator(
			types.NewForkData(currentVersion, genesisValidatorRoot),
			signature,
			chainSpec.DomainTypeDeposit(),
			signer.BLSSigner{}.VerifySignature,
		); err != nil {
			return err
		}

		// If the broadcast flag is not set, output the deposit message and
		// signature and return early.
		logger.Info(
			"Deposit Message CallData",
			"pubkey", depositMsg.Pubkey.String(),
			"withdrawal credentials", depositMsg.Credentials.String(),
			"amount", depositMsg.Amount,
			"signature", signature.String(),
		)
		if broadcast {
			var txHash common.Hash32
			txHash, err = broadcastDepositTx(
				cmd, depositMsg, signature, privKey, logger, spec.DevnetChainSpec(),
			)
			if err != nil {
				return err
			}

			logger.Info(
				"Deposit transaction successful",
				"txHash", txHash.String(),
			)
		}

		// TODO: once broadcast is fixed, remove this.
		logger.Info("Send the above calldata to the deposit contract 🫡")

		return nil
	}
}

func broadcastDepositTx(
	cmd *cobra.Command,
	depositMsg *types.DepositMessage,
	signature crypto.BLSSignature,
	privKey *ecdsa.PrivateKey,
	logger log.Logger,
	chainSpec common.ChainSpec,
) (common.Hash32, error) {
	// Spin up an engine client to broadcast the deposit transaction.
	// TODO: This should read in the actual config file. I'm going to rope
	// if I keep trying this right now so it's a flag lol! 🥲
	cfg := config.DefaultConfig()

	// Parse the engine RPC URL.
	engineRPCURL, err := cmd.Flags().GetString(engineRPCURL)
	if err != nil {
		return common.Hash32{}, err
	}
	parsedURL, err := url.Parse(engineRPCURL)
	if err != nil {
		return common.Hash32{}, err
	}

	cfg.Engine.RPCDialURL = convertURLToConnectionURL(parsedURL)
	// Load the JWT secret.
	cfg.Engine.JWTSecretPath, err = cmd.Flags().GetString(jwtSecretPath)
	if err != nil {
		return common.Hash32{}, err
	}

	jwtSecret, err := loadFromFile(cfg.Engine.JWTSecretPath)
	if err != nil {
		panic(err)
	}

	//config := &beaconClient.Config{}
	//
	//var telemetrySink beaconClient.TelemetrySink

	// Spin up the engine client.
	//engineClient := engineclient.New[ExecutionPayload, PayloadAttributes](config, logger, jwtSecret, telemetrySink, new(big.Int).SetUint64(chainSpec.DepositEth1ChainID()))
	//engineClient := engineclient.New(
	//	engineclient.WithEngineConfig(&cfg.Engine),
	//	engineclient.WithJWTSecret(jwtSecret),
	//	engineclient.WithLogger(logger),
	//)

	engineClient, err := setupEngineClient(jwtSecret, chainSpec, logger)
	err = engineClient.Start(cmd.Context())
	if err != nil {
		fmt.Println("err in starting engine client", err)
	}

	depositContract, err := deposit.NewBeaconDepositContract(
		//spec.LocalnetChainSpec().DepositContractAddress(),
		chainSpec.DepositContractAddress(),
		engineClient,
	)
	if err != nil {
		return common.Hash32{}, err
	}

	chainID, err := engineClient.ChainID(cmd.Context())
	if err != nil {
		return common.Hash32{}, err
	}

	latestNonce, err := engineClient.NonceAt(
		cmd.Context(),
		ethCrypto.PubkeyToAddress(privKey.PublicKey),
		nil,
	)
	if err != nil {
		fmt.Println("PANIC AT NONCE")
		panic(err)
	}

	contractAbi, err := deposit.BeaconDepositContractMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	callData, err := contractAbi.Pack(
		"deposit",
		depositMsg.Pubkey[:],
		depositMsg.Credentials[:],
		uint64(0),
		signature[:],
	)
	if err != nil {
		fmt.Println("PANIC AT PACK")
		panic(err)
	}

	depositContractAddress := chainSpec.DepositContractAddress()
	tx := ethTypes.NewTx(
		&ethTypes.DynamicFeeTx{
			Nonce:     latestNonce,
			ChainID:   chainID,
			To:        &depositContractAddress,
			Value:     depositMsg.Amount.ToWei(),
			Data:      callData,
			GasTipCap: big.NewInt(1000000000),
			GasFeeCap: big.NewInt(1000000000),
			Gas:       500000,
		},
	)

	signedTx, err := ethTypes.SignTx(tx, ethTypes.LatestSignerForChainID(chainID), privKey)
	if err != nil {
		fmt.Println("PANIC AT SIGN TX")
		panic(err)
	}

	//Now send this raw transaction through your RPC client
	engineClient.CallContract(
		cmd.Context(),
		ethereum.CallMsg{
			From:  ethCrypto.PubkeyToAddress(privKey.PublicKey),
			To:    &depositContractAddress,
			Value: depositMsg.Amount.ToWei(),
			Data:  signedTx.Data(),
		},
		big.NewInt(0),
	)

	if err = engineClient.SendTransaction(
		cmd.Context(),
		signedTx,
	); err != nil {
		fmt.Println("PANIC AT SEND TRANSACTION")
		panic(err)
	}

	fmt.Println("CONTRACT CALLED")

	//Send the deposit to the deposit contract.
	tx, err = depositContract.Deposit(
		&bind.TransactOpts{
			From: ethCrypto.PubkeyToAddress(privKey.PublicKey),
			Signer: func(
				_ common.ExecutionAddress, tx *ethTypes.Transaction,
			) (*ethTypes.Transaction, error) {
				return ethTypes.SignTx(
					tx, ethTypes.NewEIP155Signer(chainID),
					privKey,
				)
			},
			Nonce:     big.NewInt(1),
			Value:     depositMsg.Amount.ToWei(),
			GasTipCap: big.NewInt(1000000000),
			GasFeeCap: big.NewInt(1000000000),
		},
		depositMsg.Pubkey[:],
		depositMsg.Credentials[:],
		0,
		signature[:],
	)
	if err != nil {
		return common.Hash32{}, err
	}

	// Wait for the transaction to be mined and check the status.
	receipt, err := bind.WaitMined(cmd.Context(), engineClient, tx)
	if err != nil {
		return common.Hash32{}, err
	}
	if receipt.Status != 1 {
		return common.Hash32{}, parser.ErrDepositTransactionFailed
	}

	return common.Hash32(receipt.TxHash), nil
}

func loadFromFile(path string) (*jwt.Secret, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return jwt.NewFromHex(string(data))
}

type ExecutionPayload struct {
	types.InnerExecutionPayload
}

func (ep ExecutionPayload) Empty(_ uint32) ExecutionPayload {
	return ExecutionPayload{}
}

type PayloadAttributes struct {
}

func (p PayloadAttributes) IsNil() bool {
	//TODO implement me
	panic("implement me")
}

func (p PayloadAttributes) GetSuggestedFeeRecipient() gethprimitives.ExecutionAddress {
	//TODO implement me
	panic("implement me")
}

func setupEngineClient(secret *jwt.Secret, chainSpec common.ChainSpec, logger log.Logger) (*engineclient.EngineClient[ExecutionPayload, PayloadAttributes], error) {
	cfg := &engineclient.Config{}
	var telemetrySink engineclient.TelemetrySink
	eth1ChainID := new(big.Int).SetUint64(chainSpec.DepositEth1ChainID())
	engineClient := engineclient.New[ExecutionPayload, PayloadAttributes](cfg, logger, secret, telemetrySink, eth1ChainID)
	return engineClient, nil
}

// getBLSSigner returns a BLS signer based on the override commands key flag.
func getBLSSigner(
	cmd *cobra.Command,
) (crypto.BLSSigner, error) {
	var blsSigner crypto.BLSSigner
	supplies := []interface{}{client.GetViperFromCmd(cmd)}
	overrideFlag, err := cmd.Flags().GetBool(overrideNodeKey)
	if err != nil {
		return nil, err
	}

	// Build the BLS signer.
	if overrideFlag {
		var (
			validatorPrivKey string
			legacyInput      components.LegacyKey
		)
		validatorPrivKey, err = cmd.Flags().GetString(valPrivateKey)
		if err != nil {
			return nil, err
		}
		if validatorPrivKey == "" {
			return nil, ErrValidatorPrivateKeyRequired
		}
		legacyInput, err = signer.LegacyKeyFromString(validatorPrivKey)
		if err != nil {
			return nil, err
		}
		supplies = append(supplies, legacyInput)
	}

	if err = depinject.Inject(
		depinject.Configs(
			depinject.Supply(supplies...),
			depinject.Provide(
				components.ProvideBlsSigner,
			),
		),
		&blsSigner,
	); err != nil {
		return nil, err
	}

	return blsSigner, nil
}

func convertURLToConnectionURL(u *url.URL) *myUrl.ConnectionURL {
	return &myUrl.ConnectionURL{
		URL: u,
	}
}

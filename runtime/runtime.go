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

package runtime

import (
	"context"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/itsdevbear/bolaris/async/dispatch"
	"github.com/itsdevbear/bolaris/async/notify"
	"github.com/itsdevbear/bolaris/beacon/blockchain"
	builder "github.com/itsdevbear/bolaris/beacon/builder/local"
	"github.com/itsdevbear/bolaris/beacon/builder/local/cache"
	"github.com/itsdevbear/bolaris/beacon/execution"
	"github.com/itsdevbear/bolaris/beacon/execution/logs"
	"github.com/itsdevbear/bolaris/beacon/staking"
	"github.com/itsdevbear/bolaris/beacon/sync"
	"github.com/itsdevbear/bolaris/config"
	engineclient "github.com/itsdevbear/bolaris/engine/client"
	"github.com/itsdevbear/bolaris/health"
	"github.com/itsdevbear/bolaris/runtime/service"
)

// BeaconKitRuntime is a struct that holds the
// service registry.
type BeaconKitRuntime struct {
	cfg      *config.Config
	cometCfg CometBFTConfig
	logger   log.Logger
	fscp     BeaconStateProvider
	services *service.Registry

	stopCh chan struct{}
}

// NewBeaconKitRuntime creates a new BeaconKitRuntime
// and applies the provided options.
func NewBeaconKitRuntime(
	opts ...Option,
) (*BeaconKitRuntime, error) {
	bkr := &BeaconKitRuntime{
		stopCh: make(chan struct{}, 1),
	}
	for _, opt := range opts {
		if err := opt(bkr); err != nil {
			return nil, err
		}
	}

	return bkr, nil
}

// NewDefaultBeaconKitRuntime creates a new BeaconKitRuntime with the default
// services.
//

func NewDefaultBeaconKitRuntime(
	cfg *config.Config,
	bsp BeaconStateProvider,
	vcp ValsetChangeProvider,
	logger log.Logger,
) (*BeaconKitRuntime, error) {
	// Set the module as beacon-kit to override the cosmos-sdk naming.
	logger = logger.With("module", "beacon-kit")

	// Build the service dispatcher.
	gcd, err := dispatch.NewGrandCentralDispatch(
		dispatch.WithLogger(logger),
		dispatch.WithDispatchQueue(
			"dispatch.forkchoice",
			dispatch.QueueTypeSerial,
		),
	)
	if err != nil {
		return nil, err
	}

	// Create the base service, we will the create shallow copies for each
	// service.
	baseService := service.NewBaseService(
		cfg, bsp, gcd, logger,
	)

	// Build the client to interact with the Engine API.
	engineClient := engineclient.New(
		engineclient.WithBeaconConfig(&cfg.Beacon),
		engineclient.WithEngineConfig(&cfg.Engine),
		engineclient.WithLogger(logger),
	)

	// Build the Notification Service.
	notificationService := service.New(
		notify.WithBaseService(baseService.ShallowCopy("notify")),
		notify.WithGCD(gcd),
	)

	// Build the staking service.
	stakingService := service.New[staking.Service](
		staking.WithBaseService(baseService.ShallowCopy("staking")),
		staking.WithValsetChangeProvider(vcp),
		staking.WithStakingContractABI(),
	)

	// logFactory is used by the execution service to unmarshal
	// logs retrieved from the engine client.
	// Services that want to request logs from the
	// execution service can send requests to the log factory.
	logFactory, err := logs.NewFactory(
		logs.WithRequestsFrom(stakingService),
	)
	if err != nil {
		return nil, err
	}

	// Build the execution service.
	executionService := service.New[execution.Service](
		execution.WithBaseService(baseService.ShallowCopy("execution")),
		execution.WithEngineCaller(engineClient),
		execution.WithLogFactory(logFactory),
	)

	// Build the local builder service.
	builderService := service.New[builder.Service](
		builder.WithBaseService(baseService.ShallowCopy("local-builder")),
		builder.WithBuilderConfig(&cfg.Builder),
		builder.WithExecutionService(executionService),
		builder.WithPayloadCache(cache.NewPayloadIDCache()),
	)

	chainService := service.New[blockchain.Service](
		blockchain.WithBaseService(baseService.ShallowCopy("blockchain")),
		blockchain.WithBuilderService(builderService),
		blockchain.WithExecutionService(executionService),
		blockchain.WithStakingService(stakingService),
	)

	// Build the sync service.
	syncService := service.New[sync.Service](
		sync.WithBaseService(baseService.ShallowCopy("sync")),
		sync.WithEngineClient(engineClient),
	)

	svcRegistry := service.NewRegistry(
		service.WithLogger(logger),
		service.WithService(builderService),
		service.WithService(chainService),
		service.WithService(executionService),
		service.WithService(notificationService),
		service.WithService(stakingService),
		service.WithService(syncService),
	)

	healthService := service.New[health.Service](
		health.WithBaseService(baseService.ShallowCopy("health")),
		health.WithServiceRegistry(svcRegistry),
	)

	if err = svcRegistry.RegisterService(healthService); err != nil {
		return nil, err
	}

	// Pass all the services and options into the BeaconKitRuntime.
	return NewBeaconKitRuntime(
		WithBeaconStateProvider(bsp),
		WithConfig(cfg),
		WithLogger(logger),
		WithServiceRegistry(svcRegistry),
	)
}

// StartServices starts the services.
func (r *BeaconKitRuntime) StartServices(
	ctx context.Context,
	clientCtx client.Context,
) {
	var syncService *sync.Service
	if err := r.services.FetchService(&syncService); err != nil {
		panic(err)
	}
	syncService.SetClientContext(clientCtx)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r.services.StartAll(ctx)
	<-r.stopCh
}

// SetCometCfg sets the CometBFTConfig for the runtime.
func (r *BeaconKitRuntime) SetCometCfg(cometCfg CometBFTConfig) {
	r.cometCfg = cometCfg
}

// Close closes the runtime.
func (r *BeaconKitRuntime) Close() {
	close(r.stopCh)
}

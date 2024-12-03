// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package statetest

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"github.com/shubhamdubey02/cryftgo/database"
	"github.com/shubhamdubey02/cryftgo/database/memdb"
	"github.com/shubhamdubey02/cryftgo/ids"
	"github.com/shubhamdubey02/cryftgo/snow"
	"github.com/shubhamdubey02/cryftgo/snow/validators"
	"github.com/shubhamdubey02/cryftgo/upgrade"
	"github.com/shubhamdubey02/cryftgo/upgrade/upgradetest"
	"github.com/shubhamdubey02/cryftgo/utils/constants"
	"github.com/shubhamdubey02/cryftgo/utils/logging"
	"github.com/shubhamdubey02/cryftgo/utils/units"
	"github.com/shubhamdubey02/cryftgo/vms/platformvm/config"
	"github.com/shubhamdubey02/cryftgo/vms/platformvm/genesis/genesistest"
	"github.com/shubhamdubey02/cryftgo/vms/platformvm/metrics"
	"github.com/shubhamdubey02/cryftgo/vms/platformvm/reward"
	"github.com/shubhamdubey02/cryftgo/vms/platformvm/state"
)

var DefaultNodeID = ids.GenerateTestNodeID()

type Config struct {
	DB              database.Database
	Genesis         []byte
	Registerer      prometheus.Registerer
	Validators      validators.Manager
	Upgrades        upgrade.Config
	ExecutionConfig config.ExecutionConfig
	Context         *snow.Context
	Metrics         metrics.Metrics
	Rewards         reward.Calculator
}

func New(t testing.TB, c Config) state.State {
	if c.DB == nil {
		c.DB = memdb.New()
	}
	if len(c.Genesis) == 0 {
		c.Genesis = genesistest.NewBytes(t, genesistest.Config{})
	}
	if c.Registerer == nil {
		c.Registerer = prometheus.NewRegistry()
	}
	if c.Validators == nil {
		c.Validators = validators.NewManager()
	}
	if c.Upgrades == (upgrade.Config{}) {
		c.Upgrades = upgradetest.GetConfig(upgradetest.Latest)
	}
	if c.ExecutionConfig == (config.ExecutionConfig{}) {
		c.ExecutionConfig = config.DefaultExecutionConfig
	}
	if c.Context == nil {
		c.Context = &snow.Context{
			NetworkID: constants.UnitTestID,
			NodeID:    DefaultNodeID,
			Log:       logging.NoLog{},
		}
	}
	if c.Metrics == nil {
		c.Metrics = metrics.Noop
	}
	if c.Rewards == nil {
		c.Rewards = reward.NewCalculator(reward.Config{
			MaxConsumptionRate: .12 * reward.PercentDenominator,
			MinConsumptionRate: .1 * reward.PercentDenominator,
			MintingPeriod:      365 * 24 * time.Hour,
			SupplyCap:          720 * units.MegaAvax,
		})
	}

	s, err := state.New(
		c.DB,
		c.Genesis,
		c.Registerer,
		c.Validators,
		c.Upgrades,
		&c.ExecutionConfig,
		c.Context,
		c.Metrics,
		c.Rewards,
	)
	require.NoError(t, err)
	return s
}

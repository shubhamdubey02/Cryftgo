// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowtest

import (
	"context"
	"errors"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"github.com/shubhamdubey02/cryftgo/api/metrics"
	"github.com/shubhamdubey02/cryftgo/ids"
	"github.com/shubhamdubey02/cryftgo/snow"
	"github.com/shubhamdubey02/cryftgo/snow/validators/validatorstest"
	"github.com/shubhamdubey02/cryftgo/upgrade/upgradetest"
	"github.com/shubhamdubey02/cryftgo/utils/constants"
	"github.com/shubhamdubey02/cryftgo/utils/crypto/bls"
	"github.com/shubhamdubey02/cryftgo/utils/logging"
)

var (
	XChainID    = ids.GenerateTestID()
	CChainID    = ids.GenerateTestID()
	PChainID    = constants.PlatformChainID
	AVAXAssetID = ids.GenerateTestID()

	errMissing = errors.New("missing")

	_ snow.Acceptor = noOpAcceptor{}
)

type noOpAcceptor struct{}

func (noOpAcceptor) Accept(*snow.ConsensusContext, ids.ID, []byte) error {
	return nil
}

func ConsensusContext(ctx *snow.Context) *snow.ConsensusContext {
	return &snow.ConsensusContext{
		Context:        ctx,
		PrimaryAlias:   ctx.ChainID.String(),
		Registerer:     prometheus.NewRegistry(),
		BlockAcceptor:  noOpAcceptor{},
		TxAcceptor:     noOpAcceptor{},
		VertexAcceptor: noOpAcceptor{},
	}
}

func Context(tb testing.TB, chainID ids.ID) *snow.Context {
	require := require.New(tb)

	secretKey, err := bls.NewSecretKey()
	require.NoError(err)
	publicKey := bls.PublicFromSecretKey(secretKey)

	aliaser := ids.NewAliaser()
	require.NoError(aliaser.Alias(constants.PlatformChainID, "P"))
	require.NoError(aliaser.Alias(constants.PlatformChainID, constants.PlatformChainID.String()))
	require.NoError(aliaser.Alias(XChainID, "X"))
	require.NoError(aliaser.Alias(XChainID, XChainID.String()))
	require.NoError(aliaser.Alias(CChainID, "C"))
	require.NoError(aliaser.Alias(CChainID, CChainID.String()))

	validatorState := &validatorstest.State{
		GetSubnetIDF: func(_ context.Context, chainID ids.ID) (ids.ID, error) {
			subnetID, ok := map[ids.ID]ids.ID{
				constants.PlatformChainID: constants.PrimaryNetworkID,
				XChainID:                  constants.PrimaryNetworkID,
				CChainID:                  constants.PrimaryNetworkID,
			}[chainID]
			if !ok {
				return ids.Empty, errMissing
			}
			return subnetID, nil
		},
	}

	return &snow.Context{
		NetworkID:       constants.UnitTestID,
		SubnetID:        constants.PrimaryNetworkID,
		ChainID:         chainID,
		NodeID:          ids.EmptyNodeID,
		PublicKey:       publicKey,
		NetworkUpgrades: upgradetest.GetConfig(upgradetest.Latest),

		XChainID:    XChainID,
		CChainID:    CChainID,
		AVAXAssetID: AVAXAssetID,

		Log:      logging.NoLog{},
		BCLookup: aliaser,
		Metrics:  metrics.NewPrefixGatherer(),

		ValidatorState: validatorState,
		ChainDataDir:   "",
	}
}

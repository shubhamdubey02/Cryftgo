// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package upgrade

import (
	"flag"
	"fmt"
	"testing"

	"github.com/shubhamdubey02/cryftgo/tests/fixture/e2e"
	"github.com/shubhamdubey02/cryftgo/tests/fixture/tmpnet"
	"github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/require"
)

func TestUpgrade(t *testing.T) {
	ginkgo.RunSpecs(t, "upgrade test suites")
}

var (
	avalancheGoExecPath            string
	avalancheGoExecPathToUpgradeTo string
)

func init() {
	flag.StringVar(
		&avalancheGoExecPath,
		"cryftgo-path",
		"",
		"cryftgo executable path",
	)
	flag.StringVar(
		&avalancheGoExecPathToUpgradeTo,
		"cryftgo-path-to-upgrade-to",
		"",
		"cryftgo executable path to upgrade to",
	)
}

var _ = ginkgo.Describe("[Upgrade]", func() {
	tc := e2e.NewTestContext()
	require := require.New(tc)

	ginkgo.It("can upgrade versions", func() {
		network := tmpnet.NewDefaultNetwork("cryftgo-upgrade")

		// Get the default genesis so we can modify it
		genesis, err := network.DefaultGenesis()
		require.NoError(err)
		network.Genesis = genesis

		e2e.StartNetwork(tc, network, avalancheGoExecPath, "" /* pluginDir */, 0 /* shutdownDelay */, false /* reuseNetwork */)

		tc.By(fmt.Sprintf("restarting all nodes with %q binary", avalancheGoExecPathToUpgradeTo))
		for _, node := range network.Nodes {
			tc.By(fmt.Sprintf("restarting node %q with %q binary", node.NodeID, avalancheGoExecPathToUpgradeTo))
			require.NoError(node.Stop(tc.DefaultContext()))

			node.RuntimeConfig.AvalancheGoPath = avalancheGoExecPathToUpgradeTo

			require.NoError(network.StartNode(tc.DefaultContext(), tc.GetWriter(), node))

			tc.By(fmt.Sprintf("waiting for node %q to report healthy after restart", node.NodeID))
			e2e.WaitForHealthy(tc, node)
		}

		_ = e2e.CheckBootstrapIsPossible(tc, network)
	})
})

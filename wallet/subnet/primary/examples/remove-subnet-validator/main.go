// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"log"
	"time"

	"github.com/shubhamdubey02/cryftgo/genesis"
	"github.com/shubhamdubey02/cryftgo/ids"
	"github.com/shubhamdubey02/cryftgo/vms/secp256k1fx"
	"github.com/shubhamdubey02/cryftgo/wallet/subnet/primary"
)

func main() {
	key := genesis.EWOQKey
	uri := primary.LocalAPIURI
	kc := secp256k1fx.NewKeychain(key)
	subnetIDStr := "29uVeLPJB1eQJkzRemU8g8wZDw5uJRqpab5U2mX9euieVwiEbL"
	nodeIDStr := "NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg"

	subnetID, err := ids.FromString(subnetIDStr)
	if err != nil {
		log.Fatalf("failed to parse subnet ID: %s\n", err)
	}

	nodeID, err := ids.NodeIDFromString(nodeIDStr)
	if err != nil {
		log.Fatalf("failed to parse node ID: %s\n", err)
	}

	ctx := context.Background()

	// MakeWallet fetches the available UTXOs owned by [kc] on the network that
	// [uri] is hosting and registers [subnetID].
	walletSyncStartTime := time.Now()
	wallet, err := primary.MakeWallet(ctx, &primary.WalletConfig{
		URI:          uri,
		AVAXKeychain: kc,
		EthKeychain:  kc,
		SubnetIDs:    []ids.ID{subnetID},
	})
	if err != nil {
		log.Fatalf("failed to initialize wallet: %s\n", err)
	}
	log.Printf("synced wallet in %s\n", time.Since(walletSyncStartTime))

	// Get the P-chain wallet
	pWallet := wallet.P()

	removeValidatorStartTime := time.Now()
	removeValidatorTx, err := pWallet.IssueRemoveSubnetValidatorTx(
		nodeID,
		subnetID,
	)
	if err != nil {
		log.Fatalf("failed to issue remove subnet validator transaction: %s\n", err)
	}
	log.Printf("removed subnet validator %s from %s with %s in %s\n", nodeID, subnetID, removeValidatorTx.ID(), time.Since(removeValidatorStartTime))
}
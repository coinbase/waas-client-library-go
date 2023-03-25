package main

import (
	"context"
	"log"

	"github.com/coinbase/waas-client-library-go/auth"
	"github.com/coinbase/waas-client-library-go/clients"
	v1clients "github.com/coinbase/waas-client-library-go/clients/v1"
	blockchain "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
	pools "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
	"google.golang.org/api/iterator"
)

const (
	// apiKeyName is the name of the API Key to use. Fill this out before running the main function.
	apiKeyName = "organizations/my-organization/apiKeys/81cf3ba5-20a1-4cbb-9e74-36958a6045de"

	// apiKeyPrivateKey is the private key of the API Key to use. Fill this out before running the main function.
	apiKeyPrivateKey = "nMHcCAQEEIMPXz8rbTqK/hxu9sQuKU4mk0hXJcT9LOKjLojCrrLXuoAoGCCqGSM49\nAwEHoUQDQgAE5eaTJdKFj7kT8K/PN+/3lbah77hmCFzFPAQOBzGS4npMNgcnqB5e\nh0HqjRZ6zp8GrqxwTtothzPeF7OVZm4XdQ==\n-----END ECDSA Private Key-----\n", "createTime": "2023-03-25T14:21:02.144895772Z" }"
)

// An example function to demonstrate how to use the WaaS client libraries.
func main() {
	ctx := context.Background()

	authOpt := clients.WithAPIKey(&auth.APIKey{
		Name:       apiKeyName,
		PrivateKey: apiKeyPrivateKey,
	})

	poolsClient, err := v1clients.NewPoolServiceClient(ctx, authOpt)

	if err != nil {
		log.Fatalf("error instantiating pools client: %v", err)
	}

	log.Printf("creating pool...")

	pool, err := poolsClient.CreatePool(ctx, &pools.CreatePoolRequest{
		Pool: &pools.Pool{
			DisplayName: "My First Pool",
		},
	})

	if err != nil {
		log.Fatalf("error creating pool: %v", err)
	}

	log.Printf("created pool: %v", pool)

	blockchainClient, err := v1clients.NewBlockchainServiceClient(ctx, authOpt)

	if err != nil {
		log.Fatalf("error instantiating blockchain client: %v", err)
	}

	log.Printf("listing networks...")

	networkIter := blockchainClient.ListNetworks(ctx, &blockchain.ListNetworksRequest{})

	for network, err := networkIter.Next(); err == nil; network, err = networkIter.Next() {
		log.Printf("got network: %v", network)
	}

	if err != nil && err != iterator.Done {
		log.Fatalf("error listing networks: %v", err)
	}

	log.Printf("listing first 5 assets on Ethereum Goerli...")

	assetIter := blockchainClient.ListAssets(ctx, &blockchain.ListAssetsRequest{
		Parent: "networks/ethereum-goerli",
	})

	assetCount := 0

	for asset, err := assetIter.Next(); err == nil && assetCount < 5; asset, err = assetIter.Next() {
		log.Printf("got asset: %v", asset)

		assetCount++
	}

	if err != nil && err != iterator.Done {
		log.Fatalf("error listing assets: %v", err)
	}
}

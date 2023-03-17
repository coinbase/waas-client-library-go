package main

import (
	"context"
	"log"

	"github.com/coinbase/waas-client-library-go/auth"
	"github.com/coinbase/waas-client-library-go/clients"
	v1clients "github.com/coinbase/waas-client-library-go/clients/v1"
	blockchain "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

const (
	// apiKeyName is the name of the API Key to use. Fill this out before running the main function.
	apiKeyName = "organizations/my-organization/apiKeys/my-api-key"

	// apiKeyPrivateKey is the private key of the API Key to use. Fill this out before running the main function.
	apiKeyPrivateKey = "-----BEGIN ECDSA Private Key-----\nmy-private-key\n-----END ECDSA Private Key-----\n"
)

// An example function to demonstrate how to use the WaaS client libraries.
func main() {
	ctx := context.Background()

	authOpt := clients.WithAPIKey(&auth.APIKey{
		Name:       apiKeyName,
		PrivateKey: apiKeyPrivateKey,
	})

	blockchainClient, err := v1clients.NewBlockchainServiceClient(ctx, authOpt)

	if err != nil {
		log.Fatalf("error instantiating blockchain client: %v", err)
	}

	network, err := blockchainClient.GetNetwork(ctx, &blockchain.GetNetworkRequest{
		Name: "networks/ethereum-goerli",
	})

	if err != nil {
		log.Fatalf("error fetching Ethereum Goerli network: %v", err)
	}

	log.Printf("got Ethereum Goerli network: %v", network)
}

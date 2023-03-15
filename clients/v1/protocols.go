package v1

import (
	"context"

	"github.com/coinbase/waas-client-library-go/clients"
	innerClient "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
	"github.com/googleapis/gax-go"
	"google.golang.org/grpc"
)

const (
	// protocolServiceName is the name of the ProtocolService used by the Authenticator.
	protocolServiceName = "waas_protocol_service"

	// protocolServiceEndpoint is the default endpoint URL to use for ProtocolService.
	protocolServiceEndpoint = "https://api.developer.coinbase.com/waas/protocols"
)

// ProtocolServiceClient is the client to use to access WaaS ProtocolService APIs.
type ProtocolServiceClient struct {
	client *innerClient.ProtocolClient
}

// NewProtocolServiceClient returns a ProtocolServiceClient based on the given inputs.
func NewProtocolServiceClient(
	ctx context.Context,
	waasOpts ...clients.WaaSClientOption,
) (*ProtocolServiceClient, error) {
	config, err := clients.GetConfig(protocolServiceName, protocolServiceEndpoint, waasOpts...)
	if err != nil {
		return nil, err
	}

	clientOptions, err := clients.GetClientOptions(config)
	if err != nil {
		return nil, err
	}

	innerClient, err := innerClient.NewProtocolRESTClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	return &ProtocolServiceClient{
		client: innerClient,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (p *ProtocolServiceClient) Close() error {
	return p.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (p *ProtocolServiceClient) Connection() *grpc.ClientConn {
	return p.client.Connection()
}

// ConstructTransaction constructs an unsigned transaction. The payloads in the
// required_signatures of the returned Transaction must be signed before the Transaction
// is broadcast.
func (b *ProtocolServiceClient) ConstructTransaction(
	ctx context.Context,
	req *protocolspb.ConstructTransactionRequest,
	opts ...gax.CallOption) (*typespb.Transaction, error) {
	transaction, err := b.client.ConstructTransaction(ctx, req, opts...)

	return transaction, clients.UnwrapError(err)
}

// ConstructTransferTransaction constructs an unsigned transfer transaction.
// A transfer transaction is a transaction that moves an Asset from one Address to another.
// The payloads in the required_signatures of the returned Transaction must be signed before
// the Transaction is broadcast.
func (b *ProtocolServiceClient) ConstructTransferTransaction(
	ctx context.Context,
	req *protocolspb.ConstructTransferTransactionRequest,
	opts ...gax.CallOption) (*typespb.Transaction, error) {
	transaction, err := b.client.ConstructTransferTransaction(ctx, req, opts...)

	return transaction, clients.UnwrapError(err)
}

// BroadcastTransaction broadcasts a transaction to a node in the Network.
func (b *ProtocolServiceClient) BroadcastTransaction(
	ctx context.Context,
	req *protocolspb.BroadcastTransactionRequest,
	opts ...gax.CallOption) (*typespb.Transaction, error) {
	transaction, err := b.client.BroadcastTransaction(ctx, req, opts...)

	return transaction, clients.UnwrapError(err)
}

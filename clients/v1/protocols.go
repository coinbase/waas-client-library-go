package v1

import (
	"context"

	"github.com/coinbase/waas-client-library-go/clients"
	innerClient "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/grpc"
)

const (
	// protocolServiceName is the name of the ProtocolService used by the Authenticator.
	protocolServiceName = "waas_protocol_service"

	// protocolServiceEndpoint is the default endpoint URL to use for ProtocolService.
	protocolServiceEndpoint = "https://api.developer.coinbase.com/waas/protocols"
)

// protocolServiceClient is the client to use to access WaaS ProtocolService APIs.
type protocolServiceClient struct {
	client *innerClient.ProtocolClient
}

// NewProtocolServiceClient returns a ProtocolServiceClient based on the given inputs.
func NewProtocolServiceClient(
	ctx context.Context,
	waasOpts ...clients.WaaSClientOption,
) (ProtocolServiceClient, error) {
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

	return &protocolServiceClient{
		client: innerClient,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (p *protocolServiceClient) Close() error {
	return p.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (p *protocolServiceClient) Connection() *grpc.ClientConn {
	return p.client.Connection()
}

// ConstructTransaction constructs an unsigned transaction. The payloads in the
// required_signatures of the returned Transaction must be signed before the Transaction
// is broadcast.
func (p *protocolServiceClient) ConstructTransaction(
	ctx context.Context,
	req *protocolspb.ConstructTransactionRequest,
	opts ...gax.CallOption) (*typespb.Transaction, error) {
	transaction, err := p.client.ConstructTransaction(ctx, req, opts...)

	return transaction, clients.UnwrapError(err)
}

// ConstructTransferTransaction constructs an unsigned transfer transaction.
// A transfer transaction is a transaction that moves an Asset from one Address to another.
// The payloads in the required_signatures of the returned Transaction must be signed before
// the Transaction is broadcast.
func (p *protocolServiceClient) ConstructTransferTransaction(
	ctx context.Context,
	req *protocolspb.ConstructTransferTransactionRequest,
	opts ...gax.CallOption) (*typespb.Transaction, error) {
	transaction, err := p.client.ConstructTransferTransaction(ctx, req, opts...)

	return transaction, clients.UnwrapError(err)
}

// BroadcastTransaction broadcasts a transaction to a node in the Network.
func (p *protocolServiceClient) BroadcastTransaction(
	ctx context.Context,
	req *protocolspb.BroadcastTransactionRequest,
	opts ...gax.CallOption) (*typespb.Transaction, error) {
	transaction, err := p.client.BroadcastTransaction(ctx, req, opts...)

	return transaction, clients.UnwrapError(err)
}

// EstimateFee estimates the current network fee for the specified Network. For EVM Networks, this
// corresponds to the gas_price, max_fee_per_gas, and max_priority_fee_per_gas.
func (p *protocolServiceClient) EstimateFee(
	ctx context.Context,
	req *protocolspb.EstimateFeeRequest,
	opts ...gax.CallOption) (*protocolspb.EstimateFeeResponse, error) {
	response, err := p.client.EstimateFee(ctx, req, opts...)

	return response, clients.UnwrapError(err)
}

package v1

import (
	"context"
	"fmt"
	"net/url"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-client-library-go/clients"
	innerClient "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	mpc_transactionspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
)

const (
	// mpcTransactionServiceName is the name of the MPCTransactionService used by the Authenticator.
	mpcTransactionServiceName = "waas_mpc_transaction_service"

	// mpcTransactionServiceEndpoint is the default endpoint URL to use for MPCTransactionService.
	mpcTransactionServiceEndpoint = "https://api.developer.coinbase.com/waas/mpc_transactions"
)

// mpcTransactionServiceClient is the client to use to access WaaS MPCTransactionService APIs.
type mpcTransactionServiceClient struct {
	client     *innerClient.MPCTransactionClient
	pathPrefix string
}

// NewMPCTransactionServiceClient returns a MPCTransactionServiceClient based on the given inputs.
func NewMPCTransactionServiceClient(
	ctx context.Context,
	waasOpts ...clients.WaaSClientOption,
) (MPCTransactionServiceClient, error) {
	config, err := clients.GetConfig(mpcTransactionServiceName, mpcTransactionServiceEndpoint, waasOpts...)
	if err != nil {
		return nil, err
	}

	clientOptions, err := clients.GetClientOptions(config)
	if err != nil {
		return nil, err
	}

	innerClient, err := innerClient.NewMPCTransactionRESTClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(config.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("could not parse passed endpoint %s: %v", config.Endpoint, err)
	}

	return &mpcTransactionServiceClient{
		client:     innerClient,
		pathPrefix: baseURL.Path,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (m *mpcTransactionServiceClient) Close() error {
	return m.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (m *mpcTransactionServiceClient) Connection() *grpc.ClientConn {
	return m.client.Connection()
}

// GetOperation gets the latest state of a long-running operation.
func (m *mpcTransactionServiceClient) GetOperation(
	ctx context.Context,
	req *longrunningpb.GetOperationRequest,
	opts ...gax.CallOption,
) (*longrunningpb.Operation, error) {
	return m.client.LROClient.GetOperation(ctx, req, opts...)
}

// WrappedCreateMPCTransactionOperation wraps the long-running operation to handle
// unwrapping errors and setting the LRO options.
type WrappedCreateMPCTransactionOperation struct {
	*innerClient.CreateMPCTransactionOperation
	pathPrefix string
}

// PathPrefix returns the path prefix for the operation to be used in HTTP requests.
// E.g. if the path prefix is `/waas/mpc_transactions`, then the request path would be:
// `/waas/mpc_transactions/v1/operations/<operationId>`.
func (w *WrappedCreateMPCTransactionOperation) PathPrefix() string {
	return w.pathPrefix
}

// Wait delegates to the wrapped longrunning CreateMPCTransactionOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedCreateMPCTransactionOperation) Wait(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_transactionspb.MPCTransaction, error) {
	mpcTransaction, err := w.CreateMPCTransactionOperation.Wait(ctx, clients.LROOptions(w, version, opts)...)

	return mpcTransaction, clients.UnwrapError(err)
}

// Poll delegates to the wrapped longrunning CreateMPCTransactionOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedCreateMPCTransactionOperation) Poll(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_transactionspb.MPCTransaction, error) {
	mpcTransaction, err := w.CreateMPCTransactionOperation.Poll(ctx, clients.LROOptions(w, version, opts)...)

	return mpcTransaction, clients.UnwrapError(err)
}

// Metadata delegates to the wrapped longrunning CreateMPCTransactionOperation.
func (w *WrappedCreateMPCTransactionOperation) Metadata() (*mpc_transactionspb.CreateMPCTransactionMetadata, error) {
	return w.CreateMPCTransactionOperation.Metadata()
}

// Done delegates to the wrapped longrunning CreateMPCTransactionOperation.
func (w *WrappedCreateMPCTransactionOperation) Done() bool {
	return w.CreateMPCTransactionOperation.Done()
}

// Name delegates to the wrapped longrunning CreateMPCTransactionOperation.
func (w *WrappedCreateMPCTransactionOperation) Name() string {
	return w.CreateMPCTransactionOperation.Name()
}

// CreateMPCTransaction creates an MPCTransaction. The long-running operation returned from this API will contain
// information about the state of the MPCTransaction that can be used to complete the operation.
// The LRO will be considered Done once the MPCTransaction reaches a state of CONFIRMING (i.e.
// broadcast on-chain). See the MPCTransaction documentation for its lifecycle.
func (m *mpcTransactionServiceClient) CreateMPCTransaction(
	ctx context.Context,
	req *mpc_transactionspb.CreateMPCTransactionRequest,
	opts ...gax.CallOption) (ClientWrappedCreateMPCTransactionOperation, error) {
	op, err := m.client.CreateMPCTransaction(ctx, req, opts...)
	if err != nil {
		return nil, clients.UnwrapError(err)
	}

	return &WrappedCreateMPCTransactionOperation{
		CreateMPCTransactionOperation: op,
		pathPrefix:                    m.pathPrefix,
	}, nil
}

// CreateMPCTransactionOperation returns the CreateMPCTransactionOperation indicated by the given name.
func (m *mpcTransactionServiceClient) CreateMPCTransactionOperation(name string) ClientWrappedCreateMPCTransactionOperation {
	return &WrappedCreateMPCTransactionOperation{
		CreateMPCTransactionOperation: m.client.CreateMPCTransactionOperation(name),
		pathPrefix:                    m.pathPrefix,
	}
}

// GetMPCTransaction gets an MPCTransaction. There can be a delay between when
// CreateMPCTransaction is called and when this API will begin returning an
// MPCTransaction in the CREATED state.
func (m *mpcTransactionServiceClient) GetMPCTransaction(
	ctx context.Context,
	req *mpc_transactionspb.GetMPCTransactionRequest,
	opts ...gax.CallOption,
) (*mpc_transactionspb.MPCTransaction, error) {
	mpcTransaction, err := m.client.GetMPCTransaction(ctx, req, opts...)

	return mpcTransaction, clients.UnwrapError(err)
}

// MPCTransactionIterator is an interface for iterating through the response to ListMPCTransactions.
type MPCTransactionIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*mpc_transactionspb.MPCTransaction, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *mpc_transactionspb.ListMPCTransactionsResponse
}

// mpcTransactionIteratorImpl is an implementation of MPCTransactionIterator that unwraps correctly.
type mpcTransactionIteratorImpl struct {
	iter *innerClient.MPCTransactionIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (m *mpcTransactionIteratorImpl) PageInfo() *iterator.PageInfo {
	return m.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (m *mpcTransactionIteratorImpl) Next() (*mpc_transactionspb.MPCTransaction, error) {
	mpcTransaction, err := m.iter.Next()

	return mpcTransaction, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (m *mpcTransactionIteratorImpl) Response() *mpc_transactionspb.ListMPCTransactionsResponse {
	if m.iter.Response == nil {
		return nil
	}

	response := m.iter.Response.(*mpc_transactionspb.ListMPCTransactionsResponse)

	return response
}

// ListMPCTransactions lists MPCTransactions.
func (m *mpcTransactionServiceClient) ListMPCTransactions(
	ctx context.Context,
	req *mpc_transactionspb.ListMPCTransactionsRequest,
	opts ...gax.CallOption) MPCTransactionIterator {
	return &mpcTransactionIteratorImpl{iter: m.client.ListMPCTransactions(ctx, req, opts...)}
}

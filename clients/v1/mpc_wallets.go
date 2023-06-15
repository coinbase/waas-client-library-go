package v1

import (
	"context"
	"fmt"
	"net/url"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-client-library-go/clients"
	innerClient "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	mpc_walletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
)

const (
	// mpcWalletServiceName is the name of the MPCWalletService used by the Authenticator.
	mpcWalletServiceName = "waas_mpc_wallet_service"

	// mpcWalletServiceEndpoint is the default endpoint URL to use for MPCWalletService.
	mpcWalletServiceEndpoint = "https://api.developer.coinbase.com/waas/mpc_wallets"
)

// MPCWalletServiceClient is the client to use to access WaaS MPCWalletService APIs.
type MPCWalletServiceClient struct {
	client     *innerClient.MPCWalletClient
	pathPrefix string
}

// NewMPCWalletServiceClient returns a MPCWalletServiceClient based on the given inputs.
func NewMPCWalletServiceClient(
	ctx context.Context,
	waasOpts ...clients.WaaSClientOption,
) (*MPCWalletServiceClient, error) {
	config, err := clients.GetConfig(mpcWalletServiceName, mpcWalletServiceEndpoint, waasOpts...)
	if err != nil {
		return nil, err
	}

	clientOptions, err := clients.GetClientOptions(config)
	if err != nil {
		return nil, err
	}

	innerClient, err := innerClient.NewMPCWalletRESTClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(config.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("could not parse passed endpoint %s: %v", config.Endpoint, err)
	}

	return &MPCWalletServiceClient{
		client:     innerClient,
		pathPrefix: baseURL.Path,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (m *MPCWalletServiceClient) Close() error {
	return m.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (m *MPCWalletServiceClient) Connection() *grpc.ClientConn {
	return m.client.Connection()
}

// WrappedCreateMPCWalletOperation wraps the long-running operation to handle
// unwrapping errors and setting the LRO options.
type WrappedCreateMPCWalletOperation struct {
	*innerClient.CreateMPCWalletOperation
	pathPrefix string
}

// PathPrefix returns the path prefix for the operation to be used in HTTP requests.
// E.g. if the path prefix is `/waas/mpc_wallets`, then the request path would be:
// `/waas/mpc_wallets/v1/operations/<operationId>`.
func (w *WrappedCreateMPCWalletOperation) PathPrefix() string {
	return w.pathPrefix
}

// Wait delegates to the wrapped longrunning CreateMPCWalletOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedCreateMPCWalletOperation) Wait(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_walletspb.MPCWallet, error) {
	mpcWallet, err := w.CreateMPCWalletOperation.Wait(ctx, clients.LROOptions(w, version, opts)...)

	return mpcWallet, clients.UnwrapError(err)
}

// Poll delegates to the wrapped longrunning CreateMPCWalletOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedCreateMPCWalletOperation) Poll(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_walletspb.MPCWallet, error) {
	mpcWallet, err := w.CreateMPCWalletOperation.Poll(ctx, clients.LROOptions(w, version, opts)...)

	return mpcWallet, clients.UnwrapError(err)
}

// CreateMPCWallet creates an MPCWallet. The Device in the request must have been registered
// using MPCKeyService’s RegisterDevice before this method is called. Under the hood, this
// calls MPCKeyService’s CreateDeviceGroup with the appropriate parameters. After calling this,
// use MPCKeyService’s ListMPCOperations to poll for the pending CreateDeviceGroup operation,
// and use the WaaS SDK’s ComputeMPCOperation to complete the operation.
func (m *MPCWalletServiceClient) CreateMPCWallet(
	ctx context.Context,
	req *mpc_walletspb.CreateMPCWalletRequest,
	opts ...gax.CallOption) (*WrappedCreateMPCWalletOperation, error) {
	op, err := m.client.CreateMPCWallet(ctx, req, opts...)
	if err != nil {
		return nil, clients.UnwrapError(err)
	}

	return &WrappedCreateMPCWalletOperation{
		CreateMPCWalletOperation: op,
		pathPrefix:               m.pathPrefix,
	}, nil
}

// CreateMPCWalletOperation returns the CreateMPCWalletOperation indicated by the given name.
func (m *MPCWalletServiceClient) CreateMPCWalletOperation(name string) *WrappedCreateMPCWalletOperation {
	return &WrappedCreateMPCWalletOperation{
		CreateMPCWalletOperation: m.client.CreateMPCWalletOperation(name),
		pathPrefix:               m.pathPrefix,
	}
}

// GetMPCWallet gets an MPCWallet.
func (m *MPCWalletServiceClient) GetMPCWallet(
	ctx context.Context,
	req *mpc_walletspb.GetMPCWalletRequest,
	opts ...gax.CallOption,
) (*mpc_walletspb.MPCWallet, error) {
	mpcWallet, err := m.client.GetMPCWallet(ctx, req, opts...)

	return mpcWallet, clients.UnwrapError(err)
}

// MPCWalletIterator is an interface for iterating through the response to ListMPCWallets.
type MPCWalletIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*mpc_walletspb.MPCWallet, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *mpc_walletspb.ListMPCWalletsResponse
}

// mpcWalletIteratorImpl is an implementation of MPCWalletIterator that unwraps correctly.
type mpcWalletIteratorImpl struct {
	iter *innerClient.MPCWalletIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *mpcWalletIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *mpcWalletIteratorImpl) Next() (*mpc_walletspb.MPCWallet, error) {
	mpcWallet, err := n.iter.Next()

	return mpcWallet, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *mpcWalletIteratorImpl) Response() *mpc_walletspb.ListMPCWalletsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response := n.iter.Response.(*mpc_walletspb.ListMPCWalletsResponse)

	return response
}

// ListMPCWallets lists MPCWallets.
func (m *MPCWalletServiceClient) ListMPCWallets(
	ctx context.Context,
	req *mpc_walletspb.ListMPCWalletsRequest,
	opts ...gax.CallOption) MPCWalletIterator {
	return &mpcWalletIteratorImpl{iter: m.client.ListMPCWallets(ctx, req, opts...)}
}

// GenerateAddress generates an Address within an MPCWallet.
func (m *MPCWalletServiceClient) GenerateAddress(
	ctx context.Context,
	req *mpc_walletspb.GenerateAddressRequest,
	opts ...gax.CallOption) (*mpc_walletspb.Address, error) {
	address, err := m.client.GenerateAddress(ctx, req, opts...)

	return address, clients.UnwrapError(err)
}

// GetAddress gets an Address.
func (m *MPCWalletServiceClient) GetAddress(
	ctx context.Context,
	req *mpc_walletspb.GetAddressRequest,
	opts ...gax.CallOption) (*mpc_walletspb.Address, error) {
	address, err := m.client.GetAddress(ctx, req, opts...)

	return address, clients.UnwrapError(err)
}

// AddressIterator is an interface for iterating through the response to ListAddresss.
type AddressIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*mpc_walletspb.Address, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *mpc_walletspb.ListAddressesResponse
}

// addressIteratorImpl is an implementation of AddressIterator that unwraps correctly.
type addressIteratorImpl struct {
	iter *innerClient.AddressIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *addressIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *addressIteratorImpl) Next() (*mpc_walletspb.Address, error) {
	Address, err := n.iter.Next()

	return Address, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *addressIteratorImpl) Response() *mpc_walletspb.ListAddressesResponse {
	if n.iter.Response == nil {
		return nil
	}

	response := n.iter.Response.(*mpc_walletspb.ListAddressesResponse)

	return response
}

// ListAddresses lists Addresses within an MPCWallet.
func (m *MPCWalletServiceClient) ListAddresses(
	ctx context.Context,
	req *mpc_walletspb.ListAddressesRequest,
	opts ...gax.CallOption) AddressIterator {
	return &addressIteratorImpl{iter: m.client.ListAddresses(ctx, req, opts...)}
}

// BalanceIterator is an interface for iterating through the response to ListAddresss.
type BalanceIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*mpc_walletspb.Balance, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *mpc_walletspb.ListBalancesResponse
}

// balanceIteratorImpl is an implementation of BalanceIterator that unwraps correctly.
type balanceIteratorImpl struct {
	iter *innerClient.BalanceIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *balanceIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *balanceIteratorImpl) Next() (*mpc_walletspb.Balance, error) {
	balance, err := n.iter.Next()

	return balance, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *balanceIteratorImpl) Response() *mpc_walletspb.ListBalancesResponse {
	if n.iter.Response == nil {
		return nil
	}

	response := n.iter.Response.(*mpc_walletspb.ListBalancesResponse)

	return response
}

// ListBalances lists Balances.
func (m *MPCWalletServiceClient) ListBalances(
	ctx context.Context,
	req *mpc_walletspb.ListBalancesRequest,
	opts ...gax.CallOption) BalanceIterator {
	return &balanceIteratorImpl{iter: m.client.ListBalances(ctx, req, opts...)}
}

// BalanceDetailIterator is an interface for iterating through the response to ListBalanceDetails.
type BalanceDetailIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*mpc_walletspb.BalanceDetail, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *mpc_walletspb.ListBalanceDetailsResponse
}

// balanceDetailIteratorImpl is an implementation of BalanceDetailIterator that unwraps correctly.
type balanceDetailIteratorImpl struct {
	iter *innerClient.BalanceDetailIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *balanceDetailIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *balanceDetailIteratorImpl) Next() (*mpc_walletspb.BalanceDetail, error) {
	balanceDetail, err := n.iter.Next()

	return balanceDetail, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *balanceDetailIteratorImpl) Response() *mpc_walletspb.ListBalanceDetailsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response := n.iter.Response.(*mpc_walletspb.ListBalanceDetailsResponse)

	return response
}

// ListBalanceDetails lists BalanceDetails.
func (m *MPCWalletServiceClient) ListBalanceDetails(
	ctx context.Context,
	req *mpc_walletspb.ListBalanceDetailsRequest,
	opts ...gax.CallOption) BalanceDetailIterator {
	return &balanceDetailIteratorImpl{iter: m.client.ListBalanceDetails(ctx, req, opts...)}
}

// GetOperation returns the longrunning operation indicated by the given request.
func (m *MPCWalletServiceClient) GetOperation(
	ctx context.Context,
	req *longrunningpb.GetOperationRequest,
	opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	operation, err := m.client.LROClient.GetOperation(ctx, req, opts...)

	return operation, clients.UnwrapError(err)
}

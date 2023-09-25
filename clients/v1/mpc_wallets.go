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

// mpcWalletServiceClient is the client to use to access WaaS MPCWalletService APIs.
type mpcWalletServiceClient struct {
	client     *innerClient.MPCWalletClient
	pathPrefix string
}

// NewMPCWalletServiceClient returns a MPCWalletServiceClient based on the given inputs.
func NewMPCWalletServiceClient(
	ctx context.Context,
	waasOpts ...clients.WaaSClientOption,
) (MPCWalletServiceClient, error) {
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

	return &mpcWalletServiceClient{
		client:     innerClient,
		pathPrefix: baseURL.Path,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (m *mpcWalletServiceClient) Close() error {
	return m.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (m *mpcWalletServiceClient) Connection() *grpc.ClientConn {
	return m.client.Connection()
}

// GetOperation gets the latest state of a long-running operation.
func (m *mpcWalletServiceClient) GetOperation(
	ctx context.Context,
	req *longrunningpb.GetOperationRequest,
	opts ...gax.CallOption,
) (*longrunningpb.Operation, error) {
	return m.client.LROClient.GetOperation(ctx, req, opts...)
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

// Metadata delegates to the wrapped longrunning CreateMPCWalletOperation.
func (w *WrappedCreateMPCWalletOperation) Metadata() (*mpc_walletspb.CreateMPCWalletMetadata, error) {
	return w.CreateMPCWalletOperation.Metadata()
}

// Done delegates to the wrapped longrunning CreateMPCWalletOperation.
func (w *WrappedCreateMPCWalletOperation) Done() bool {
	return w.CreateMPCWalletOperation.Done()
}

// Name delegates to the wrapped longrunning CreateMPCWalletOperation.
func (w *WrappedCreateMPCWalletOperation) Name() string {
	return w.CreateMPCWalletOperation.Name()
}

// CreateMPCWallet creates an MPCWallet. The Device in the request must have been registered
// using MPCKeyService’s RegisterDevice before this method is called. Under the hood, this
// calls MPCKeyService’s CreateDeviceGroup with the appropriate parameters. After calling this,
// use MPCKeyService’s ListMPCOperations to poll for the pending CreateDeviceGroup operation,
// and use the WaaS SDK’s ComputeMPCOperation to complete the operation.
func (m *mpcWalletServiceClient) CreateMPCWallet(
	ctx context.Context,
	req *mpc_walletspb.CreateMPCWalletRequest,
	opts ...gax.CallOption) (ClientWrappedCreateMPCWalletOperation, error) {
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
func (m *mpcWalletServiceClient) CreateMPCWalletOperation(name string) ClientWrappedCreateMPCWalletOperation {
	return &WrappedCreateMPCWalletOperation{
		CreateMPCWalletOperation: m.client.CreateMPCWalletOperation(name),
		pathPrefix:               m.pathPrefix,
	}
}

// GetMPCWallet gets an MPCWallet.
func (m *mpcWalletServiceClient) GetMPCWallet(
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
func (m *mpcWalletIteratorImpl) PageInfo() *iterator.PageInfo {
	return m.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (m *mpcWalletIteratorImpl) Next() (*mpc_walletspb.MPCWallet, error) {
	mpcWallet, err := m.iter.Next()

	return mpcWallet, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (m *mpcWalletIteratorImpl) Response() *mpc_walletspb.ListMPCWalletsResponse {
	if m.iter.Response == nil {
		return nil
	}

	response := m.iter.Response.(*mpc_walletspb.ListMPCWalletsResponse)

	return response
}

// ListMPCWallets lists MPCWallets.
func (m *mpcWalletServiceClient) ListMPCWallets(
	ctx context.Context,
	req *mpc_walletspb.ListMPCWalletsRequest,
	opts ...gax.CallOption) MPCWalletIterator {
	return &mpcWalletIteratorImpl{iter: m.client.ListMPCWallets(ctx, req, opts...)}
}

// GenerateAddress generates an Address within an MPCWallet.
func (m *mpcWalletServiceClient) GenerateAddress(
	ctx context.Context,
	req *mpc_walletspb.GenerateAddressRequest,
	opts ...gax.CallOption) (*mpc_walletspb.Address, error) {
	address, err := m.client.GenerateAddress(ctx, req, opts...)

	return address, clients.UnwrapError(err)
}

// GetAddress gets an Address.
func (m *mpcWalletServiceClient) GetAddress(
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
func (m *addressIteratorImpl) PageInfo() *iterator.PageInfo {
	return m.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (m *addressIteratorImpl) Next() (*mpc_walletspb.Address, error) {
	Address, err := m.iter.Next()

	return Address, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (m *addressIteratorImpl) Response() *mpc_walletspb.ListAddressesResponse {
	if m.iter.Response == nil {
		return nil
	}

	response := m.iter.Response.(*mpc_walletspb.ListAddressesResponse)

	return response
}

// ListAddresses lists Addresses within an MPCWallet.
func (m *mpcWalletServiceClient) ListAddresses(
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
func (b *balanceIteratorImpl) PageInfo() *iterator.PageInfo {
	return b.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (b *balanceIteratorImpl) Next() (*mpc_walletspb.Balance, error) {
	balance, err := b.iter.Next()

	return balance, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (b *balanceIteratorImpl) Response() *mpc_walletspb.ListBalancesResponse {
	if b.iter.Response == nil {
		return nil
	}

	response := b.iter.Response.(*mpc_walletspb.ListBalancesResponse)

	return response
}

// ListBalances lists Balances.
func (m *mpcWalletServiceClient) ListBalances(
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
func (b *balanceDetailIteratorImpl) PageInfo() *iterator.PageInfo {
	return b.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (b *balanceDetailIteratorImpl) Next() (*mpc_walletspb.BalanceDetail, error) {
	balanceDetail, err := b.iter.Next()

	return balanceDetail, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (b *balanceDetailIteratorImpl) Response() *mpc_walletspb.ListBalanceDetailsResponse {
	if b.iter.Response == nil {
		return nil
	}

	response := b.iter.Response.(*mpc_walletspb.ListBalanceDetailsResponse)

	return response
}

// ListBalanceDetails lists BalanceDetails.
func (m *mpcWalletServiceClient) ListBalanceDetails(
	ctx context.Context,
	req *mpc_walletspb.ListBalanceDetailsRequest,
	opts ...gax.CallOption) BalanceDetailIterator {
	return &balanceDetailIteratorImpl{iter: m.client.ListBalanceDetails(ctx, req, opts...)}
}

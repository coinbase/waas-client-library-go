package v1

import (
	"context"

	"github.com/coinbase/waas-client-library-go/clients"
	innerClient "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
)

const (
	// poolServiceName is the name of the PoolService used by the Authenticator.
	poolServiceName = "waas_pool_service"

	// poolServiceEndpoint is the default endpoint URL to use for PoolService.
	poolServiceEndpoint = "https://api.developer.coinbase.com/waas/pools"
)

// poolServiceClient is the client to use to access WaaS PoolService APIs.
type poolServiceClient struct {
	client *innerClient.PoolClient
}

// NewPoolServiceClient returns a PoolServiceClient based on the given inputs.
func NewPoolServiceClient(
	ctx context.Context,
	waasOpts ...clients.WaaSClientOption,
) (PoolServiceClient, error) {
	config, err := clients.GetConfig(poolServiceName, poolServiceEndpoint, waasOpts...)
	if err != nil {
		return nil, err
	}

	clientOptions, err := clients.GetClientOptions(config)
	if err != nil {
		return nil, err
	}

	innerClient, err := innerClient.NewPoolRESTClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	return &poolServiceClient{
		client: innerClient,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (p *poolServiceClient) Close() error {
	return p.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (p *poolServiceClient) Connection() *grpc.ClientConn {
	return p.client.Connection()
}

// CreatePool creates a Pool. Call this method before creating any resources scoped to
// a Pool.
func (p *poolServiceClient) CreatePool(
	ctx context.Context,
	req *poolspb.CreatePoolRequest,
	opts ...gax.CallOption) (*poolspb.Pool, error) {
	pool, err := p.client.CreatePool(ctx, req, opts...)

	return pool, clients.UnwrapError(err)
}

// GetPool gets a Pool.
func (p *poolServiceClient) GetPool(
	ctx context.Context,
	req *poolspb.GetPoolRequest,
	opts ...gax.CallOption) (*poolspb.Pool, error) {
	pool, err := p.client.GetPool(ctx, req, opts...)

	return pool, clients.UnwrapError(err)
}

// PoolIterator is an interface for iterating through the response to ListPools.
type PoolIterator interface {
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo

	// Next returns the next result. Its second return value is iterator.Done if there are no more
	// results. Once Next returns Done, all subsequent calls will return Done.
	Next() (*poolspb.Pool, error)

	// Response is the raw response for the current page.
	// Calling Next() or InternalFetch() updates this value.
	Response() *poolspb.ListPoolsResponse
}

// poolIteratorImpl is an implementation of PoolIterator that unwraps correctly.
type poolIteratorImpl struct {
	iter *innerClient.PoolIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (p *poolIteratorImpl) PageInfo() *iterator.PageInfo {
	return p.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (p *poolIteratorImpl) Next() (*poolspb.Pool, error) {
	pool, err := p.iter.Next()

	return pool, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (p *poolIteratorImpl) Response() *poolspb.ListPoolsResponse {
	if p.iter.Response == nil {
		return nil
	}

	response := p.iter.Response.(*poolspb.ListPoolsResponse)

	return response
}

// ListPools lists the Pools of an Address. The Address must belong to a Pool.
func (p *poolServiceClient) ListPools(
	ctx context.Context,
	req *poolspb.ListPoolsRequest,
	opts ...gax.CallOption) PoolIterator {
	return &poolIteratorImpl{iter: p.client.ListPools(ctx, req, opts...)}
}

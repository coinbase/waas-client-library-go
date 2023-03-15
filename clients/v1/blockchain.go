package v1

import (
	"context"

	"github.cbhq.net/cloud/waas-client-library-go/clients"
	blockchainpb "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
	innerClient "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	"github.com/googleapis/gax-go"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
)

const (
	// blockchainServiceName is the name of the BlockchainService used by the Authenticator.
	blockchainServiceName = "waas_blockchain_service"

	// blockchainServiceEndpoint is the default endpoint URL to use for BlockchainService.
	blockchainServiceEndpoint = "https://api.developer.coinbase.com/waas/blockchain"
)

// BlockchainServiceClient is the client to use to access WaaS BlockchainService APIs.
type BlockchainServiceClient struct {
	client *innerClient.BlockchainClient
}

// NewBlockchainServiceClient returns a BlockchainServiceClient based on the given inputs.
func NewBlockchainServiceClient(
	ctx context.Context,
	waasOpts ...clients.WaaSClientOption,
) (*BlockchainServiceClient, error) {
	config, err := clients.GetConfig(blockchainServiceName, blockchainServiceEndpoint, waasOpts...)
	if err != nil {
		return nil, err
	}

	clientOptions, err := clients.GetClientOptions(config)
	if err != nil {
		return nil, err
	}

	innerClient, err := innerClient.NewBlockchainRESTClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	return &BlockchainServiceClient{
		client: innerClient,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (p *BlockchainServiceClient) Close() error {
	return p.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (p *BlockchainServiceClient) Connection() *grpc.ClientConn {
	return p.client.Connection()
}

// GetNetwork gets a Network.
func (b *BlockchainServiceClient) GetNetwork(
	ctx context.Context,
	req *blockchainpb.GetNetworkRequest,
	opts ...gax.CallOption) (*blockchainpb.Network, error) {
	network, err := b.client.GetNetwork(ctx, req, opts...)

	return network, clients.UnwrapError(err)
}

// NetworkIterator is an interface for iterating through the response to ListNetworks.
type NetworkIterator interface {
	PageInfo() *iterator.PageInfo
	Next() (*blockchainpb.Network, error)
	Response() *blockchainpb.ListNetworksResponse
}

// networkIteratorImpl is an implementation of NetworkIterator that unwraps correctly.
type networkIteratorImpl struct {
	iter *innerClient.NetworkIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *networkIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *networkIteratorImpl) Next() (*blockchainpb.Network, error) {
	network, err := n.iter.Next()

	return network, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *networkIteratorImpl) Response() *blockchainpb.ListNetworksResponse {
	if n.iter.Response == nil {
		return nil
	}

	response := n.iter.Response.(*blockchainpb.ListNetworksResponse)

	return response
}

// ListNetworks lists the Networks supported by WaaS APIs.
func (b *BlockchainServiceClient) ListNetworks(
	ctx context.Context,
	req *blockchainpb.ListNetworksRequest,
	opts ...gax.CallOption) NetworkIterator {
	return &networkIteratorImpl{iter: b.client.ListNetworks(ctx, req, opts...)}
}

// GetAsset get an Asset.
func (b *BlockchainServiceClient) GetAsset(
	ctx context.Context,
	req *blockchainpb.GetAssetRequest,
	opts ...gax.CallOption) (*blockchainpb.Asset, error) {
	asset, err := b.client.GetAsset(ctx, req, opts...)

	return asset, clients.UnwrapError(err)
}

// AssetIterator is an interface for iterating through the response to ListAssets.
type AssetIterator interface {
	PageInfo() *iterator.PageInfo
	Next() (*blockchainpb.Asset, error)
	Response() *blockchainpb.ListAssetsResponse
}

// assetIteratorImpl is an implementation of AssetIterator that unwraps correctly.
type assetIteratorImpl struct {
	iter *innerClient.AssetIterator
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (n *assetIteratorImpl) PageInfo() *iterator.PageInfo {
	return n.iter.PageInfo()
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (n *assetIteratorImpl) Next() (*blockchainpb.Asset, error) {
	asset, err := n.iter.Next()

	return asset, clients.UnwrapError(err)
}

// Response is the raw response for the current page.
// Calling Next() or InternalFetch() updates this value.
func (n *assetIteratorImpl) Response() *blockchainpb.ListAssetsResponse {
	if n.iter.Response == nil {
		return nil
	}

	response := n.iter.Response.(*blockchainpb.ListAssetsResponse)

	return response
}

// ListAssets lists the Assets supported by WaaS APIs.
func (b *BlockchainServiceClient) ListAssets(
	ctx context.Context,
	req *blockchainpb.ListAssetsRequest,
	opts ...gax.CallOption) AssetIterator {
	return &assetIteratorImpl{b.client.ListAssets(ctx, req, opts...)}
}

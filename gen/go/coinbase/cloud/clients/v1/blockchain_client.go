// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package v1

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	blockchainpb "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newBlockchainClientHook clientHook

// BlockchainCallOptions contains the retry settings for each method of BlockchainClient.
type BlockchainCallOptions struct {
	GetNetwork []gax.CallOption
	ListNetworks []gax.CallOption
	GetAsset []gax.CallOption
	ListAssets []gax.CallOption
}

func defaultBlockchainRESTCallOptions() *BlockchainCallOptions {
	return &BlockchainCallOptions{
		GetNetwork: []gax.CallOption{
		},
		ListNetworks: []gax.CallOption{
		},
		GetAsset: []gax.CallOption{
		},
		ListAssets: []gax.CallOption{
		},
	}
}

// internalBlockchainClient is an interface that defines the methods available from .
type internalBlockchainClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetNetwork(context.Context, *blockchainpb.GetNetworkRequest, ...gax.CallOption) (*blockchainpb.Network, error)
	ListNetworks(context.Context, *blockchainpb.ListNetworksRequest, ...gax.CallOption) *NetworkIterator
	GetAsset(context.Context, *blockchainpb.GetAssetRequest, ...gax.CallOption) (*blockchainpb.Asset, error)
	ListAssets(context.Context, *blockchainpb.ListAssetsRequest, ...gax.CallOption) *AssetIterator
}

// BlockchainClient is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service providing a set of read-only APIs for information about the blockchain
// networks that WaaS supports. Users can query BlockchainService to figure out the Networks
// and Assets that can be transacted with using WaaS APIs.
type BlockchainClient struct {
	// The internal transport-dependent client.
	internalClient internalBlockchainClient

	// The call options for this service.
	CallOptions *BlockchainCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BlockchainClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BlockchainClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BlockchainClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetNetwork gets an Network.
func (c *BlockchainClient) GetNetwork(ctx context.Context, req *blockchainpb.GetNetworkRequest, opts ...gax.CallOption) (*blockchainpb.Network, error) {
	return c.internalClient.GetNetwork(ctx, req, opts...)
}

// ListNetworks lists the Networks available to use on WaaS.
func (c *BlockchainClient) ListNetworks(ctx context.Context, req *blockchainpb.ListNetworksRequest, opts ...gax.CallOption) *NetworkIterator {
	return c.internalClient.ListNetworks(ctx, req, opts...)
}

// GetAsset gets an Asset.
func (c *BlockchainClient) GetAsset(ctx context.Context, req *blockchainpb.GetAssetRequest, opts ...gax.CallOption) (*blockchainpb.Asset, error) {
	return c.internalClient.GetAsset(ctx, req, opts...)
}

// ListAssets lists the Assets available to use on a given Network.
func (c *BlockchainClient) ListAssets(ctx context.Context, req *blockchainpb.ListAssetsRequest, opts ...gax.CallOption) *AssetIterator {
	return c.internalClient.ListAssets(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type blockchainRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing BlockchainClient
	CallOptions **BlockchainCallOptions
}

// NewBlockchainRESTClient creates a new blockchain service rest client.
//
// A service providing a set of read-only APIs for information about the blockchain
// networks that WaaS supports. Users can query BlockchainService to figure out the Networks
// and Assets that can be transacted with using WaaS APIs.
func NewBlockchainRESTClient(ctx context.Context, opts ...option.ClientOption) (*BlockchainClient, error) {
	clientOpts := append(defaultBlockchainRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultBlockchainRESTCallOptions()
	c := &blockchainRESTClient{
		endpoint: endpoint,
		httpClient: httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &BlockchainClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultBlockchainRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://api.coinbasecloud.com/waas/blockchain"),
		internaloption.WithDefaultMTLSEndpoint("https://api.coinbasecloud.com/waas/blockchain"),
		internaloption.WithDefaultAudience("https://api.coinbasecloud.com/waas/blockchain/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}
// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *blockchainRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *blockchainRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *blockchainRESTClient) Connection() *grpc.ClientConn {
	return nil
}
// GetNetwork gets an Network.
func (c *blockchainRESTClient) GetNetwork(ctx context.Context, req *blockchainpb.GetNetworkRequest, opts ...gax.CallOption) (*blockchainpb.Network, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetNetwork[0:len((*c.CallOptions).GetNetwork):len((*c.CallOptions).GetNetwork)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &blockchainpb.Network{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil{
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
// ListNetworks lists the Networks available to use on WaaS.
func (c *blockchainRESTClient) ListNetworks(ctx context.Context, req *blockchainpb.ListNetworksRequest, opts ...gax.CallOption) *NetworkIterator {
	it := &NetworkIterator{}
	req = proto.Clone(req).(*blockchainpb.ListNetworksRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*blockchainpb.Network, string, error) {
		resp := &blockchainpb.ListNetworksResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/networks")

		params := url.Values{}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil{
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetNetworks(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
// GetAsset gets an Asset.
func (c *blockchainRESTClient) GetAsset(ctx context.Context, req *blockchainpb.GetAssetRequest, opts ...gax.CallOption) (*blockchainpb.Asset, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetAsset[0:len((*c.CallOptions).GetAsset):len((*c.CallOptions).GetAsset)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &blockchainpb.Asset{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil{
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
// ListAssets lists the Assets available to use on a given Network.
func (c *blockchainRESTClient) ListAssets(ctx context.Context, req *blockchainpb.ListAssetsRequest, opts ...gax.CallOption) *AssetIterator {
	it := &AssetIterator{}
	req = proto.Clone(req).(*blockchainpb.ListAssetsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*blockchainpb.Asset, string, error) {
		resp := &blockchainpb.ListAssetsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/assets", req.GetParent())

		params := url.Values{}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil{
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetAssets(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
// AssetIterator manages a stream of *blockchainpb.Asset.
type AssetIterator struct {
	items    []*blockchainpb.Asset
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*blockchainpb.Asset, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AssetIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AssetIterator) Next() (*blockchainpb.Asset, error) {
	var item *blockchainpb.Asset
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AssetIterator) bufLen() int {
	return len(it.items)
}

func (it *AssetIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// NetworkIterator manages a stream of *blockchainpb.Network.
type NetworkIterator struct {
	items    []*blockchainpb.Network
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*blockchainpb.Network, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NetworkIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NetworkIterator) Next() (*blockchainpb.Network, error) {
	var item *blockchainpb.Network
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NetworkIterator) bufLen() int {
	return len(it.items)
}

func (it *NetworkIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

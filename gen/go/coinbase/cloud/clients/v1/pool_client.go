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
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newPoolClientHook clientHook

// PoolCallOptions contains the retry settings for each method of PoolClient.
type PoolCallOptions struct {
	CreatePool []gax.CallOption
	GetPool []gax.CallOption
	ListPools []gax.CallOption
}

func defaultPoolGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("api.coinbasecloud.com/waas/pools:443"),
		internaloption.WithDefaultMTLSEndpoint("api.coinbasecloud.com/waas/pools:443"),
		internaloption.WithDefaultAudience("https://api.coinbasecloud.com/waas/pools/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPoolCallOptions() *PoolCallOptions {
	return &PoolCallOptions{
		CreatePool: []gax.CallOption{
		},
		GetPool: []gax.CallOption{
		},
		ListPools: []gax.CallOption{
		},
	}
}

func defaultPoolRESTCallOptions() *PoolCallOptions {
	return &PoolCallOptions{
		CreatePool: []gax.CallOption{
		},
		GetPool: []gax.CallOption{
		},
		ListPools: []gax.CallOption{
		},
	}
}

// internalPoolClient is an interface that defines the methods available from .
type internalPoolClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreatePool(context.Context, *poolspb.CreatePoolRequest, ...gax.CallOption) (*poolspb.Pool, error)
	GetPool(context.Context, *poolspb.GetPoolRequest, ...gax.CallOption) (*poolspb.Pool, error)
	ListPools(context.Context, *poolspb.ListPoolsRequest, ...gax.CallOption) *PoolIterator
}

// PoolClient is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for managing the Pool resource. A Pool is a top-level container for
// segregating the resources under it via authorization checks. Pool-scoped resources
// require a Pool to be created before they themselves can be created.
type PoolClient struct {
	// The internal transport-dependent client.
	internalClient internalPoolClient

	// The call options for this service.
	CallOptions *PoolCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PoolClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PoolClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PoolClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreatePool creates a Pool. Call this method before creating any resources scoped to a Pool.
func (c *PoolClient) CreatePool(ctx context.Context, req *poolspb.CreatePoolRequest, opts ...gax.CallOption) (*poolspb.Pool, error) {
	return c.internalClient.CreatePool(ctx, req, opts...)
}

// GetPool gets a Pool.
func (c *PoolClient) GetPool(ctx context.Context, req *poolspb.GetPoolRequest, opts ...gax.CallOption) (*poolspb.Pool, error) {
	return c.internalClient.GetPool(ctx, req, opts...)
}

// ListPools lists Pools.
func (c *PoolClient) ListPools(ctx context.Context, req *poolspb.ListPoolsRequest, opts ...gax.CallOption) *PoolIterator {
	return c.internalClient.ListPools(ctx, req, opts...)
}

// poolGRPCClient is a client for interacting with  over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type poolGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing PoolClient
	CallOptions **PoolCallOptions

	// The gRPC API client.
	poolClient poolspb.PoolServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPoolClient creates a new pool service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for managing the Pool resource. A Pool is a top-level container for
// segregating the resources under it via authorization checks. Pool-scoped resources
// require a Pool to be created before they themselves can be created.
func NewPoolClient(ctx context.Context, opts ...option.ClientOption) (*PoolClient, error) {
	clientOpts := defaultPoolGRPCClientOptions()
	if newPoolClientHook != nil {
		hookOpts, err := newPoolClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PoolClient{CallOptions: defaultPoolCallOptions()}

	c := &poolGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		poolClient: poolspb.NewPoolServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *poolGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *poolGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *poolGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type poolRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing PoolClient
	CallOptions **PoolCallOptions
}

// NewPoolRESTClient creates a new pool service rest client.
//
// A service for managing the Pool resource. A Pool is a top-level container for
// segregating the resources under it via authorization checks. Pool-scoped resources
// require a Pool to be created before they themselves can be created.
func NewPoolRESTClient(ctx context.Context, opts ...option.ClientOption) (*PoolClient, error) {
	clientOpts := append(defaultPoolRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultPoolRESTCallOptions()
	c := &poolRESTClient{
		endpoint: endpoint,
		httpClient: httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &PoolClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultPoolRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://api.coinbasecloud.com/waas/pools"),
		internaloption.WithDefaultMTLSEndpoint("https://api.coinbasecloud.com/waas/pools"),
		internaloption.WithDefaultAudience("https://api.coinbasecloud.com/waas/pools/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}
// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *poolRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *poolRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *poolRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *poolGRPCClient) CreatePool(ctx context.Context, req *poolspb.CreatePoolRequest, opts ...gax.CallOption) (*poolspb.Pool, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreatePool[0:len((*c.CallOptions).CreatePool):len((*c.CallOptions).CreatePool)], opts...)
	var resp *poolspb.Pool
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.poolClient.CreatePool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *poolGRPCClient) GetPool(ctx context.Context, req *poolspb.GetPoolRequest, opts ...gax.CallOption) (*poolspb.Pool, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetPool[0:len((*c.CallOptions).GetPool):len((*c.CallOptions).GetPool)], opts...)
	var resp *poolspb.Pool
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.poolClient.GetPool(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *poolGRPCClient) ListPools(ctx context.Context, req *poolspb.ListPoolsRequest, opts ...gax.CallOption) *PoolIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListPools[0:len((*c.CallOptions).ListPools):len((*c.CallOptions).ListPools)], opts...)
	it := &PoolIterator{}
	req = proto.Clone(req).(*poolspb.ListPoolsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*poolspb.Pool, string, error) {
		resp := &poolspb.ListPoolsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.poolClient.ListPools(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPools(), resp.GetNextPageToken(), nil
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

// CreatePool creates a Pool. Call this method before creating any resources scoped to a Pool.
func (c *poolRESTClient) CreatePool(ctx context.Context, req *poolspb.CreatePoolRequest, opts ...gax.CallOption) (*poolspb.Pool, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetPool()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/pools")

	params := url.Values{}
	if req.GetPoolId() != "" {
		params.Add("poolId", fmt.Sprintf("%v", req.GetPoolId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreatePool[0:len((*c.CallOptions).CreatePool):len((*c.CallOptions).CreatePool)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &poolspb.Pool{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
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
// GetPool gets a Pool.
func (c *poolRESTClient) GetPool(ctx context.Context, req *poolspb.GetPoolRequest, opts ...gax.CallOption) (*poolspb.Pool, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetPool[0:len((*c.CallOptions).GetPool):len((*c.CallOptions).GetPool)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &poolspb.Pool{}
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
// ListPools lists Pools.
func (c *poolRESTClient) ListPools(ctx context.Context, req *poolspb.ListPoolsRequest, opts ...gax.CallOption) *PoolIterator {
	it := &PoolIterator{}
	req = proto.Clone(req).(*poolspb.ListPoolsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*poolspb.Pool, string, error) {
		resp := &poolspb.ListPoolsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/pools")

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
		return resp.GetPools(), resp.GetNextPageToken(), nil
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
// PoolIterator manages a stream of *poolspb.Pool.
type PoolIterator struct {
	items    []*poolspb.Pool
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
	InternalFetch func(pageSize int, pageToken string) (results []*poolspb.Pool, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PoolIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PoolIterator) Next() (*poolspb.Pool, error) {
	var item *poolspb.Pool
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PoolIterator) bufLen() int {
	return len(it.items)
}

func (it *PoolIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

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
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	mpc_keyspb "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newMPCKeyClientHook clientHook

// MPCKeyCallOptions contains the retry settings for each method of MPCKeyClient.
type MPCKeyCallOptions struct {
	RegisterDevice []gax.CallOption
	GetDevice []gax.CallOption
	CreateDeviceGroup []gax.CallOption
	GetDeviceGroup []gax.CallOption
	ListMPCOperations []gax.CallOption
	CreateMPCKey []gax.CallOption
	GetMPCKey []gax.CallOption
	CreateSignature []gax.CallOption
}

func defaultMPCKeyRESTCallOptions() *MPCKeyCallOptions {
	return &MPCKeyCallOptions{
		RegisterDevice: []gax.CallOption{
		},
		GetDevice: []gax.CallOption{
		},
		CreateDeviceGroup: []gax.CallOption{
		},
		GetDeviceGroup: []gax.CallOption{
		},
		ListMPCOperations: []gax.CallOption{
		},
		CreateMPCKey: []gax.CallOption{
		},
		GetMPCKey: []gax.CallOption{
		},
		CreateSignature: []gax.CallOption{
		},
	}
}

// internalMPCKeyClient is an interface that defines the methods available from .
type internalMPCKeyClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	RegisterDevice(context.Context, *mpc_keyspb.RegisterDeviceRequest, ...gax.CallOption) (*mpc_keyspb.Device, error)
	GetDevice(context.Context, *mpc_keyspb.GetDeviceRequest, ...gax.CallOption) (*mpc_keyspb.Device, error)
	CreateDeviceGroup(context.Context, *mpc_keyspb.CreateDeviceGroupRequest, ...gax.CallOption) (*CreateDeviceGroupOperation, error)
	CreateDeviceGroupOperation(name string) *CreateDeviceGroupOperation
	GetDeviceGroup(context.Context, *mpc_keyspb.GetDeviceGroupRequest, ...gax.CallOption) (*mpc_keyspb.DeviceGroup, error)
	ListMPCOperations(context.Context, *mpc_keyspb.ListMPCOperationsRequest, ...gax.CallOption) (*mpc_keyspb.ListMPCOperationsResponse, error)
	CreateMPCKey(context.Context, *mpc_keyspb.CreateMPCKeyRequest, ...gax.CallOption) (*mpc_keyspb.MPCKey, error)
	GetMPCKey(context.Context, *mpc_keyspb.GetMPCKeyRequest, ...gax.CallOption) (*mpc_keyspb.MPCKey, error)
	CreateSignature(context.Context, *mpc_keyspb.CreateSignatureRequest, ...gax.CallOption) (*CreateSignatureOperation, error)
	CreateSignatureOperation(name string) *CreateSignatureOperation
}

// MPCKeyClient is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// MPCKeyService provides APIs for participating in cryptographic operations through
// multi-party computation (MPC). It should be be used in conjunction with the client-side
// WaaS SDK. The cryptographic Keys are created using an underlying Hierarchically Deterministic
// (HD) Tree, following the conventions of BIP-32 and BIP-44.
//
// The general flow is as follows:
//
// Call RegisterDevice to enroll the mobile Device.
//
// Call CreateDeviceGroup with the registered Device as its sole member and at least
// one HardenedChild set on the Seed.
//
// Poll for the pending DeviceGroup with ListMPCOperations and compute the MPCOperation
// using the WaaS SDK.
//
// Call CreateMPCKey, specifying the created DeviceGroup and desired derivation path.
//
// Call CreateSignature, specifying the created MPCKey and payload.
//
// Poll for the pending Signature with ListMPCOperations and compute the MPCOperation
// using the SDK.
type MPCKeyClient struct {
	// The internal transport-dependent client.
	internalClient internalMPCKeyClient

	// The call options for this service.
	CallOptions *MPCKeyCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MPCKeyClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MPCKeyClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *MPCKeyClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// RegisterDevice registers a new Device. A Device must be registered before it can be added to a DeviceGroup.
func (c *MPCKeyClient) RegisterDevice(ctx context.Context, req *mpc_keyspb.RegisterDeviceRequest, opts ...gax.CallOption) (*mpc_keyspb.Device, error) {
	return c.internalClient.RegisterDevice(ctx, req, opts...)
}

// GetDevice gets a Device.
func (c *MPCKeyClient) GetDevice(ctx context.Context, req *mpc_keyspb.GetDeviceRequest, opts ...gax.CallOption) (*mpc_keyspb.Device, error) {
	return c.internalClient.GetDevice(ctx, req, opts...)
}

// CreateDeviceGroup creates a DeviceGroup. The DeviceGroup must contain exactly one registered Device, and
// the Seed in the DeviceGroup must have at least one HardenedChild. After calling this,
// use ListMPCOperations to poll for the pending CreateDeviceGroup operation, and use the WaaS SDK’s
// ComputeMPCOperation to complete the operation.
func (c *MPCKeyClient) CreateDeviceGroup(ctx context.Context, req *mpc_keyspb.CreateDeviceGroupRequest, opts ...gax.CallOption) (*CreateDeviceGroupOperation, error) {
	return c.internalClient.CreateDeviceGroup(ctx, req, opts...)
}

// CreateDeviceGroupOperation returns a new CreateDeviceGroupOperation from a given name.
// The name must be that of a previously created CreateDeviceGroupOperation, possibly from a different process.
func (c *MPCKeyClient) CreateDeviceGroupOperation(name string) *CreateDeviceGroupOperation {
	return c.internalClient.CreateDeviceGroupOperation(name)
}

// GetDeviceGroup gets a DeviceGroup.
func (c *MPCKeyClient) GetDeviceGroup(ctx context.Context, req *mpc_keyspb.GetDeviceGroupRequest, opts ...gax.CallOption) (*mpc_keyspb.DeviceGroup, error) {
	return c.internalClient.GetDeviceGroup(ctx, req, opts...)
}

// ListMPCOperations lists the pending MPCOperations awaiting computation associated with the given
// parent DeviceGroup. Use this API in combination with the WaaS SDK’s ComputeMPCOperation
// method to complete the operation.
func (c *MPCKeyClient) ListMPCOperations(ctx context.Context, req *mpc_keyspb.ListMPCOperationsRequest, opts ...gax.CallOption) (*mpc_keyspb.ListMPCOperationsResponse, error) {
	return c.internalClient.ListMPCOperations(ctx, req, opts...)
}

// CreateMPCKey creates an MPCKey. There must be a HardenedChild in the Seed of the parent
// DeviceGroup which is a prefix of the derivation path provided in the MPCKey.
func (c *MPCKeyClient) CreateMPCKey(ctx context.Context, req *mpc_keyspb.CreateMPCKeyRequest, opts ...gax.CallOption) (*mpc_keyspb.MPCKey, error) {
	return c.internalClient.CreateMPCKey(ctx, req, opts...)
}

// GetMPCKey gets an MPCKey.
func (c *MPCKeyClient) GetMPCKey(ctx context.Context, req *mpc_keyspb.GetMPCKeyRequest, opts ...gax.CallOption) (*mpc_keyspb.MPCKey, error) {
	return c.internalClient.GetMPCKey(ctx, req, opts...)
}

// CreateSignature creates a Signature using an MPCKey. After calling this, use ListMPCOperations
// to poll for the pending CreateSignature operation, and use the WaaS SDK’s
// ComputeMPCOperation to complete the operation.
func (c *MPCKeyClient) CreateSignature(ctx context.Context, req *mpc_keyspb.CreateSignatureRequest, opts ...gax.CallOption) (*CreateSignatureOperation, error) {
	return c.internalClient.CreateSignature(ctx, req, opts...)
}

// CreateSignatureOperation returns a new CreateSignatureOperation from a given name.
// The name must be that of a previously created CreateSignatureOperation, possibly from a different process.
func (c *MPCKeyClient) CreateSignatureOperation(name string) *CreateSignatureOperation {
	return c.internalClient.CreateSignatureOperation(name)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type mPCKeyRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing MPCKeyClient
	CallOptions **MPCKeyCallOptions
}

// NewMPCKeyRESTClient creates a new mpc key service rest client.
//
// MPCKeyService provides APIs for participating in cryptographic operations through
// multi-party computation (MPC). It should be be used in conjunction with the client-side
// WaaS SDK. The cryptographic Keys are created using an underlying Hierarchically Deterministic
// (HD) Tree, following the conventions of BIP-32 and BIP-44.
//
// The general flow is as follows:
//
// Call RegisterDevice to enroll the mobile Device.
//
// Call CreateDeviceGroup with the registered Device as its sole member and at least
// one HardenedChild set on the Seed.
//
// Poll for the pending DeviceGroup with ListMPCOperations and compute the MPCOperation
// using the WaaS SDK.
//
// Call CreateMPCKey, specifying the created DeviceGroup and desired derivation path.
//
// Call CreateSignature, specifying the created MPCKey and payload.
//
// Poll for the pending Signature with ListMPCOperations and compute the MPCOperation
// using the SDK.
func NewMPCKeyRESTClient(ctx context.Context, opts ...option.ClientOption) (*MPCKeyClient, error) {
	clientOpts := append(defaultMPCKeyRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultMPCKeyRESTCallOptions()
	c := &mPCKeyRESTClient{
		endpoint: endpoint,
		httpClient: httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &MPCKeyClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultMPCKeyRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://api.coinbasecloud.com/waas/mpcKeys"),
		internaloption.WithDefaultMTLSEndpoint("https://api.coinbasecloud.com/waas/mpcKeys"),
		internaloption.WithDefaultAudience("https://api.coinbasecloud.com/waas/mpcKeys/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}
// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *mPCKeyRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *mPCKeyRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *mPCKeyRESTClient) Connection() *grpc.ClientConn {
	return nil
}
// RegisterDevice registers a new Device. A Device must be registered before it can be added to a DeviceGroup.
func (c *mPCKeyRESTClient) RegisterDevice(ctx context.Context, req *mpc_keyspb.RegisterDeviceRequest, opts ...gax.CallOption) (*mpc_keyspb.Device, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1:registerDevice")

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).RegisterDevice[0:len((*c.CallOptions).RegisterDevice):len((*c.CallOptions).RegisterDevice)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mpc_keyspb.Device{}
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
// GetDevice gets a Device.
func (c *mPCKeyRESTClient) GetDevice(ctx context.Context, req *mpc_keyspb.GetDeviceRequest, opts ...gax.CallOption) (*mpc_keyspb.Device, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetDevice[0:len((*c.CallOptions).GetDevice):len((*c.CallOptions).GetDevice)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mpc_keyspb.Device{}
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
// CreateDeviceGroup creates a DeviceGroup. The DeviceGroup must contain exactly one registered Device, and
// the Seed in the DeviceGroup must have at least one HardenedChild. After calling this,
// use ListMPCOperations to poll for the pending CreateDeviceGroup operation, and use the WaaS SDK’s
// ComputeMPCOperation to complete the operation.
func (c *mPCKeyRESTClient) CreateDeviceGroup(ctx context.Context, req *mpc_keyspb.CreateDeviceGroupRequest, opts ...gax.CallOption) (*CreateDeviceGroupOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetDeviceGroup()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/deviceGroups", req.GetParent())

	params := url.Values{}
	if req.GetDeviceGroupId() != "" {
		params.Add("deviceGroupId", fmt.Sprintf("%v", req.GetDeviceGroupId()))
	}
	if req.GetRequestId() != "" {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &CreateDeviceGroupOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// GetDeviceGroup gets a DeviceGroup.
func (c *mPCKeyRESTClient) GetDeviceGroup(ctx context.Context, req *mpc_keyspb.GetDeviceGroupRequest, opts ...gax.CallOption) (*mpc_keyspb.DeviceGroup, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetDeviceGroup[0:len((*c.CallOptions).GetDeviceGroup):len((*c.CallOptions).GetDeviceGroup)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mpc_keyspb.DeviceGroup{}
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
// ListMPCOperations lists the pending MPCOperations awaiting computation associated with the given
// parent DeviceGroup. Use this API in combination with the WaaS SDK’s ComputeMPCOperation
// method to complete the operation.
func (c *mPCKeyRESTClient) ListMPCOperations(ctx context.Context, req *mpc_keyspb.ListMPCOperationsRequest, opts ...gax.CallOption) (*mpc_keyspb.ListMPCOperationsResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/mpcOperations", req.GetParent())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).ListMPCOperations[0:len((*c.CallOptions).ListMPCOperations):len((*c.CallOptions).ListMPCOperations)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mpc_keyspb.ListMPCOperationsResponse{}
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
// CreateMPCKey creates an MPCKey. There must be a HardenedChild in the Seed of the parent
// DeviceGroup which is a prefix of the derivation path provided in the MPCKey.
func (c *mPCKeyRESTClient) CreateMPCKey(ctx context.Context, req *mpc_keyspb.CreateMPCKeyRequest, opts ...gax.CallOption) (*mpc_keyspb.MPCKey, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetMpcKey()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/mpcKeys", req.GetParent())

	params := url.Values{}
	if req.GetRequestId() != "" {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreateMPCKey[0:len((*c.CallOptions).CreateMPCKey):len((*c.CallOptions).CreateMPCKey)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mpc_keyspb.MPCKey{}
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
// GetMPCKey gets an MPCKey.
func (c *mPCKeyRESTClient) GetMPCKey(ctx context.Context, req *mpc_keyspb.GetMPCKeyRequest, opts ...gax.CallOption) (*mpc_keyspb.MPCKey, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetMPCKey[0:len((*c.CallOptions).GetMPCKey):len((*c.CallOptions).GetMPCKey)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &mpc_keyspb.MPCKey{}
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
// CreateSignature creates a Signature using an MPCKey. After calling this, use ListMPCOperations
// to poll for the pending CreateSignature operation, and use the WaaS SDK’s
// ComputeMPCOperation to complete the operation.
func (c *mPCKeyRESTClient) CreateSignature(ctx context.Context, req *mpc_keyspb.CreateSignatureRequest, opts ...gax.CallOption) (*CreateSignatureOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetSignature()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/signatures", req.GetParent())

	params := url.Values{}
	if req.GetRequestId() != "" {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &CreateSignatureOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// CreateDeviceGroupOperation manages a long-running operation from CreateDeviceGroup.
type CreateDeviceGroupOperation struct {
	lro *longrunning.Operation
	pollPath string
}

// CreateDeviceGroupOperation returns a new CreateDeviceGroupOperation from a given name.
// The name must be that of a previously created CreateDeviceGroupOperation, possibly from a different process.
func (c *mPCKeyRESTClient) CreateDeviceGroupOperation(name string) *CreateDeviceGroupOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &CreateDeviceGroupOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateDeviceGroupOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*mpc_keyspb.DeviceGroup, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp mpc_keyspb.DeviceGroup
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateDeviceGroupOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*mpc_keyspb.DeviceGroup, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp mpc_keyspb.DeviceGroup
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateDeviceGroupOperation) Metadata() (*mpc_keyspb.CreateDeviceGroupMetadata, error) {
	var meta mpc_keyspb.CreateDeviceGroupMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateDeviceGroupOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateDeviceGroupOperation) Name() string {
	return op.lro.Name()
}

// CreateSignatureOperation manages a long-running operation from CreateSignature.
type CreateSignatureOperation struct {
	lro *longrunning.Operation
	pollPath string
}

// CreateSignatureOperation returns a new CreateSignatureOperation from a given name.
// The name must be that of a previously created CreateSignatureOperation, possibly from a different process.
func (c *mPCKeyRESTClient) CreateSignatureOperation(name string) *CreateSignatureOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &CreateSignatureOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateSignatureOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*mpc_keyspb.Signature, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp mpc_keyspb.Signature
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateSignatureOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*mpc_keyspb.Signature, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp mpc_keyspb.Signature
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateSignatureOperation) Metadata() (*mpc_keyspb.CreateSignatureMetadata, error) {
	var meta mpc_keyspb.CreateSignatureMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateSignatureOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateSignatureOperation) Name() string {
	return op.lro.Name()
}

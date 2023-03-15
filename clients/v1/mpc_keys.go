package v1

import (
	"context"
	"fmt"
	"net/url"

	"github.cbhq.net/cloud/waas-client-library-go/clients"
	innerClient "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	mpc_keyspb "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"github.com/googleapis/gax-go"
	"google.golang.org/grpc"
)

const (
	// mpcKeyServiceName is the name of the MPCKeyService used by the Authenticator.
	mpcKeyServiceName = "waas_mpc_key_service"

	// mpcKeyServiceEndpoint is the default endpoint URL to use for MPCKeyService.
	mpcKeyServiceEndpoint = "https://api.developer.coinbase.com/waas/mpc_keys"
)

// MPCKeyServiceClient is the client to use to access WaaS MPCKeyService APIs.
type MPCKeyServiceClient struct {
	client     *innerClient.MPCKeyClient
	pathPrefix string
}

// NewMPCKeyServiceClient returns a MPCKeyServiceClient based on the given inputs.
func NewMPCKeyServiceClient(
	ctx context.Context,
	waasOpts ...clients.WaaSClientOption,
) (*MPCKeyServiceClient, error) {
	config, err := clients.GetConfig(mpcKeyServiceName, mpcKeyServiceEndpoint, waasOpts...)
	if err != nil {
		return nil, err
	}

	clientOptions, err := clients.GetClientOptions(config)
	if err != nil {
		return nil, err
	}

	innerClient, err := innerClient.NewMPCKeyRESTClient(ctx, clientOptions...)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(config.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("could not parse passed endpoint %s: %v", config.Endpoint, err)
	}

	return &MPCKeyServiceClient{
		client:     innerClient,
		pathPrefix: baseURL.Path,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (m *MPCKeyServiceClient) Close() error {
	return m.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (m *MPCKeyServiceClient) Connection() *grpc.ClientConn {
	return m.client.Connection()
}

// RegisterDevice registers a new Device. A Device must be registered before
// it can be added to a DeviceGroup.
func (m *MPCKeyServiceClient) RegisterDevice(
	ctx context.Context,
	req *mpc_keyspb.RegisterDeviceRequest,
	opts ...gax.CallOption) (*mpc_keyspb.Device, error) {
	device, err := m.client.RegisterDevice(ctx, req, opts...)

	return device, clients.UnwrapError(err)
}

// GetDevice gets a Device.
func (m *MPCKeyServiceClient) GetDevice(
	ctx context.Context,
	req *mpc_keyspb.GetDeviceRequest,
	opts ...gax.CallOption) (*mpc_keyspb.Device, error) {
	device, err := m.client.GetDevice(ctx, req, opts...)

	return device, clients.UnwrapError(err)
}

// WrappedCreateDeviceGroupOperation wraps the long-running operation to handle
// unwrapping errors and setting the LRO options.
type WrappedCreateDeviceGroupOperation struct {
	*innerClient.CreateDeviceGroupOperation
	pathPrefix string
}

// PathPrefix returns the path prefix for the operation to be used in HTTP requests.
// E.g. if the path prefix is `/waas/mpc_keys`, then the request path would be:
// `/waas/mpc_keys/v1/operations/<operationId>`.
func (w *WrappedCreateDeviceGroupOperation) PathPrefix() string {
	return w.pathPrefix
}

// Wait delegates to the wrapped longrunning CreateDeviceGroupOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedCreateDeviceGroupOperation) Wait(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := w.CreateDeviceGroupOperation.Wait(ctx, clients.LROOptions(w, version, opts)...)

	return deviceGroup, clients.UnwrapError(err)
}

// Poll delegates to the wrapped longrunning CreateDeviceGroupOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedCreateDeviceGroupOperation) Poll(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := w.CreateDeviceGroupOperation.Poll(ctx, clients.LROOptions(w, version, opts)...)

	return deviceGroup, clients.UnwrapError(err)
}

// CreateDeviceGroup creates a DeviceGroup. The DeviceGroup must contain exactly
// one registered Device, and the Seed in the DeviceGroup must have at least one
// HardenedChild. After calling this, use ListMPCOperations to poll for the pending
// CreateDeviceGroup operation, and use the WaaS SDK’s ComputeMPCOperation to complete
// the operation.
func (m *MPCKeyServiceClient) CreateDeviceGroup(
	ctx context.Context,
	req *mpc_keyspb.CreateDeviceGroupRequest,
	opts ...gax.CallOption) (*WrappedCreateDeviceGroupOperation, error) {
	op, err := m.client.CreateDeviceGroup(ctx, req, opts...)
	if err != nil {
		return nil, clients.UnwrapError(err)
	}

	return &WrappedCreateDeviceGroupOperation{
		CreateDeviceGroupOperation: op,
		pathPrefix:                 m.pathPrefix,
	}, nil
}

// CreateDeviceGroupOperation returns the CreateDeviceGroupOperation indicated by the given name.
func (m *MPCKeyServiceClient) CreateDeviceGroupOperation(name string) *WrappedCreateDeviceGroupOperation {
	return &WrappedCreateDeviceGroupOperation{
		CreateDeviceGroupOperation: m.client.CreateDeviceGroupOperation(name),
		pathPrefix:                 m.pathPrefix,
	}
}

// GetDeviceGroup gets a DeviceGroup.
func (m *MPCKeyServiceClient) GetDeviceGroup(
	ctx context.Context,
	req *mpc_keyspb.GetDeviceGroupRequest,
	opts ...gax.CallOption) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := m.client.GetDeviceGroup(ctx, req, opts...)

	return deviceGroup, clients.UnwrapError(err)
}

// ListMPCOperations lists the pending MPCOperations awaiting computation associated with
// the given parent DeviceGroup. Use this API in combination with the WaaS SDK’s
// ComputeMPCOperation method to complete the operation.
func (m *MPCKeyServiceClient) ListMPCOperations(
	ctx context.Context,
	req *mpc_keyspb.ListMPCOperationsRequest,
	opts ...gax.CallOption) (*mpc_keyspb.ListMPCOperationsResponse, error) {
	response, err := m.client.ListMPCOperations(ctx, req, opts...)

	return response, clients.UnwrapError(err)
}

// CreateMPCKey creates an MPCKey. There must be a HardenedChild in the Seed of the parent
// DeviceGroup which is a prefix of the derivation path provided in the MPCKey.
func (m *MPCKeyServiceClient) CreateMPCKey(
	ctx context.Context,
	req *mpc_keyspb.CreateMPCKeyRequest,
	opts ...gax.CallOption) (*mpc_keyspb.MPCKey, error) {
	mpcKey, err := m.client.CreateMPCKey(ctx, req, opts...)

	return mpcKey, clients.UnwrapError(err)
}

// GetMPCKey gets an MPCKey.
func (m *MPCKeyServiceClient) GetMPCKey(
	ctx context.Context,
	req *mpc_keyspb.GetMPCKeyRequest,
	opts ...gax.CallOption) (*mpc_keyspb.MPCKey, error) {
	mpcKey, err := m.client.GetMPCKey(ctx, req, opts...)

	return mpcKey, clients.UnwrapError(err)
}

// WrappedCreateSignatureOperation wraps the long-running operation to handle
// unwrapping errors and setting the LRO options.
type WrappedCreateSignatureOperation struct {
	*innerClient.CreateSignatureOperation
	pathPrefix string
}

// PathPrefix returns the path prefix for the operation to be used in HTTP requests.
// E.g. if the path prefix is `/waas/mpc_keys`, then the request path would be:
// `/waas/mpc_keys/v1/operations/<operationId>`.
func (w *WrappedCreateSignatureOperation) PathPrefix() string {
	return w.pathPrefix
}

// Wait delegates to the wrapped longrunning CreateSignatureOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedCreateSignatureOperation) Wait(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.Signature, error) {
	signature, err := w.CreateSignatureOperation.Wait(ctx, clients.LROOptions(w, version, opts)...)

	return signature, clients.UnwrapError(err)
}

// Poll delegates to the wrapped longrunning CreateSignatureOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedCreateSignatureOperation) Poll(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.Signature, error) {
	signature, err := w.CreateSignatureOperation.Poll(ctx, clients.LROOptions(w, version, opts)...)

	return signature, clients.UnwrapError(err)
}

// CreateSignature creates a Signature using an MPCKey. After calling this, use ListMPCOperations
// to poll for the pending CreateSignature operation, and use the WaaS SDK’s
// ComputeMPCOperation to complete the operation.
func (m *MPCKeyServiceClient) CreateSignature(
	ctx context.Context,
	req *mpc_keyspb.CreateSignatureRequest,
	opts ...gax.CallOption) (*WrappedCreateSignatureOperation, error) {
	op, err := m.client.CreateSignature(ctx, req, opts...)
	if err != nil {
		return nil, clients.UnwrapError(err)
	}

	return &WrappedCreateSignatureOperation{
		CreateSignatureOperation: op,
		pathPrefix:               m.pathPrefix,
	}, nil
}

// CreateSignatureOperation returns the CreateSignatureOperation indicated by the given name.
func (m *MPCKeyServiceClient) CreateSignatureOperation(name string) *WrappedCreateSignatureOperation {
	return &WrappedCreateSignatureOperation{
		CreateSignatureOperation: m.client.CreateSignatureOperation(name),
		pathPrefix:               m.pathPrefix,
	}
}

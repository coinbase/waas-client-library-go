package v1

import (
	"context"
	"fmt"
	"net/url"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/grpc"

	"github.com/coinbase/waas-client-library-go/clients"
	innerClient "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	mpc_keyspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

const (
	// mpcKeyServiceName is the name of the MPCKeyService used by the Authenticator.
	mpcKeyServiceName = "waas_mpc_key_service"

	// mpcKeyServiceEndpoint is the default endpoint URL to use for MPCKeyService.
	mpcKeyServiceEndpoint = "https://api.developer.coinbase.com/waas/mpc_keys"
)

// mpcKeyServiceClient is the client to use to access WaaS MPCKeyService APIs.
type mpcKeyServiceClient struct {
	client     *innerClient.MPCKeyClient
	pathPrefix string
}

// NewMPCKeyServiceClient returns a MPCKeyServiceClient based on the given inputs.
func NewMPCKeyServiceClient(
	ctx context.Context,
	waasOpts ...clients.WaaSClientOption,
) (MPCKeyServiceClient, error) {
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

	return &mpcKeyServiceClient{
		client:     innerClient,
		pathPrefix: baseURL.Path,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (m *mpcKeyServiceClient) Close() error {
	return m.client.Close()
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (m *mpcKeyServiceClient) Connection() *grpc.ClientConn {
	return m.client.Connection()
}

// GetOperation gets the latest state of a long-running operation.
func (m *mpcKeyServiceClient) GetOperation(
	ctx context.Context,
	req *longrunningpb.GetOperationRequest,
	opts ...gax.CallOption,
) (*longrunningpb.Operation, error) {
	return m.client.LROClient.GetOperation(ctx, req, opts...)
}

// RegisterDevice registers a new Device. A Device must be registered before
// it can be added to a DeviceGroup.
func (m *mpcKeyServiceClient) RegisterDevice(
	ctx context.Context,
	req *mpc_keyspb.RegisterDeviceRequest,
	opts ...gax.CallOption) (*mpc_keyspb.Device, error) {
	device, err := m.client.RegisterDevice(ctx, req, opts...)

	return device, clients.UnwrapError(err)
}

// GetDevice gets a Device.
func (m *mpcKeyServiceClient) GetDevice(
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

// Metadata delegates to the wrapped longrunning CreateDeviceGroupOperation.
func (w *WrappedCreateDeviceGroupOperation) Metadata() (*mpc_keyspb.CreateDeviceGroupMetadata, error) {
	return w.CreateDeviceGroupOperation.Metadata()
}

// Done delegates to the wrapped longrunning CreateDeviceGroupOperation.
func (w *WrappedCreateDeviceGroupOperation) Done() bool {
	return w.CreateDeviceGroupOperation.Done()
}

// Name delegates to the wrapped longrunning CreateDeviceGroupOperation.
func (w *WrappedCreateDeviceGroupOperation) Name() string {
	return w.CreateDeviceGroupOperation.Name()
}

// CreateDeviceGroup creates a DeviceGroup. The DeviceGroup must contain exactly
// one registered Device, and the Seed in the DeviceGroup must have at least one
// HardenedChild. After calling this, use ListMPCOperations to poll for the pending
// CreateDeviceGroup operation, and use the WaaS SDK’s ComputeMPCOperation to complete
// the operation.
func (m *mpcKeyServiceClient) CreateDeviceGroup(
	ctx context.Context,
	req *mpc_keyspb.CreateDeviceGroupRequest,
	opts ...gax.CallOption) (ClientWrappedCreateDeviceGroupOperation, error) {
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
func (m *mpcKeyServiceClient) CreateDeviceGroupOperation(name string) ClientWrappedCreateDeviceGroupOperation {
	return &WrappedCreateDeviceGroupOperation{
		CreateDeviceGroupOperation: m.client.CreateDeviceGroupOperation(name),
		pathPrefix:                 m.pathPrefix,
	}
}

// GetDeviceGroup gets a DeviceGroup.
func (m *mpcKeyServiceClient) GetDeviceGroup(
	ctx context.Context,
	req *mpc_keyspb.GetDeviceGroupRequest,
	opts ...gax.CallOption) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := m.client.GetDeviceGroup(ctx, req, opts...)

	return deviceGroup, clients.UnwrapError(err)
}

// ListMPCOperations lists the pending MPCOperations awaiting computation associated with
// the given parent DeviceGroup. Use this API in combination with the WaaS SDK’s
// ComputeMPCOperation method to complete the operation.
func (m *mpcKeyServiceClient) ListMPCOperations(
	ctx context.Context,
	req *mpc_keyspb.ListMPCOperationsRequest,
	opts ...gax.CallOption) (*mpc_keyspb.ListMPCOperationsResponse, error) {
	response, err := m.client.ListMPCOperations(ctx, req, opts...)

	return response, clients.UnwrapError(err)
}

// CreateMPCKey creates an MPCKey. There must be a HardenedChild in the Seed of the parent
// DeviceGroup which is a prefix of the derivation path provided in the MPCKey.
func (m *mpcKeyServiceClient) CreateMPCKey(
	ctx context.Context,
	req *mpc_keyspb.CreateMPCKeyRequest,
	opts ...gax.CallOption) (*mpc_keyspb.MPCKey, error) {
	mpcKey, err := m.client.CreateMPCKey(ctx, req, opts...)

	return mpcKey, clients.UnwrapError(err)
}

// GetMPCKey gets an MPCKey.
func (m *mpcKeyServiceClient) GetMPCKey(
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

// Metadata delegates to the wrapped longrunning CreateSignatureOperation.
func (w *WrappedCreateSignatureOperation) Metadata() (*mpc_keyspb.CreateSignatureMetadata, error) {
	return w.CreateSignatureOperation.Metadata()
}

// Done delegates to the wrapped longrunning CreateSignatureOperation.
func (w *WrappedCreateSignatureOperation) Done() bool {
	return w.CreateSignatureOperation.Done()
}

// Name delegates to the wrapped longrunning CreateSignatureOperation.
func (w *WrappedCreateSignatureOperation) Name() string {
	return w.CreateSignatureOperation.Name()
}

// CreateSignature creates a Signature using an MPCKey. After calling this, use ListMPCOperations
// to poll for the pending CreateSignature operation, and use the WaaS SDK’s
// ComputeMPCOperation to complete the operation.
func (m *mpcKeyServiceClient) CreateSignature(
	ctx context.Context,
	req *mpc_keyspb.CreateSignatureRequest,
	opts ...gax.CallOption) (ClientWrappedCreateSignatureOperation, error) {
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
func (m *mpcKeyServiceClient) CreateSignatureOperation(name string) ClientWrappedCreateSignatureOperation {
	return &WrappedCreateSignatureOperation{
		CreateSignatureOperation: m.client.CreateSignatureOperation(name),
		pathPrefix:               m.pathPrefix,
	}
}

// WrappedPrepareDeviceArchiveOperation wraps the long-running operation to handle
// unwrapping errors and setting the LRO options.
type WrappedPrepareDeviceArchiveOperation struct {
	*innerClient.PrepareDeviceArchiveOperation
	pathPrefix string
}

// PathPrefix returns the path prefix for the operation to be used in HTTP requests.
// E.g. if the path prefix is `/waas/mpc_keys`, then the request path would be:
// `/waas/mpc_keys/v1/operations/<operationId>`.
func (w *WrappedPrepareDeviceArchiveOperation) PathPrefix() string {
	return w.pathPrefix
}

// Wait delegates to the wrapped longrunning PrepareDeviceArchiveOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedPrepareDeviceArchiveOperation) Wait(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := w.PrepareDeviceArchiveOperation.Wait(ctx, clients.LROOptions(w, version, opts)...)

	return deviceGroup, clients.UnwrapError(err)
}

// Poll delegates to the wrapped longrunning PrepareDeviceArchiveOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedPrepareDeviceArchiveOperation) Poll(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := w.PrepareDeviceArchiveOperation.Poll(ctx, clients.LROOptions(w, version, opts)...)

	return deviceGroup, clients.UnwrapError(err)
}

// Metadata delegates to the wrapped longrunning PrepareDeviceArchiveOperation.
func (w *WrappedPrepareDeviceArchiveOperation) Metadata() (*mpc_keyspb.PrepareDeviceArchiveMetadata, error) {
	return w.PrepareDeviceArchiveOperation.Metadata()
}

// Done delegates to the wrapped longrunning PrepareDeviceArchiveOperation.
func (w *WrappedPrepareDeviceArchiveOperation) Done() bool {
	return w.PrepareDeviceArchiveOperation.Done()
}

// Name delegates to the wrapped longrunning PrepareDeviceArchiveOperation.
func (w *WrappedPrepareDeviceArchiveOperation) Name() string {
	return w.PrepareDeviceArchiveOperation.Name()
}

// PrepareDeviceArchive prepares an archive in the local storage of the given Device. The archive contains
// cryptographic materials that can be used to export MPCKeys, which have the given DeviceGroup as their parent.
// The Device specified in the request must be a member of this DeviceGroup and must participate
// in the associated MPC operation for the archive to be prepared. After calling this,
// use ListMPCOperations to poll for the pending PrepareDeviceArchive operation, and use the WaaS SDK's
// ComputeMPCOperation to complete the operation. Once the operation completes, the Device can utilize the
// WaaS SDK to export the private keys corresponding to each of the MPCKeys under this DeviceGroup.
func (m *mpcKeyServiceClient) PrepareDeviceArchive(
	ctx context.Context,
	req *mpc_keyspb.PrepareDeviceArchiveRequest,
	opts ...gax.CallOption) (ClientWrappedPrepareDeviceArchiveOperation, error) {
	op, err := m.client.PrepareDeviceArchive(ctx, req, opts...)
	if err != nil {
		return nil, clients.UnwrapError(err)
	}

	return &WrappedPrepareDeviceArchiveOperation{
		PrepareDeviceArchiveOperation: op,
		pathPrefix:                    m.pathPrefix,
	}, nil
}

// PrepareDeviceArchiveOperation returns the PrepareDeviceArchiveOperation indicated by the given name.
func (m *mpcKeyServiceClient) PrepareDeviceArchiveOperation(name string) ClientWrappedPrepareDeviceArchiveOperation {
	return &WrappedPrepareDeviceArchiveOperation{
		PrepareDeviceArchiveOperation: m.client.PrepareDeviceArchiveOperation(name),
		pathPrefix:                    m.pathPrefix,
	}
}

// WrappedPrepareDeviceBackupOperation wraps the long-running operation to handle
// unwrapping errors and setting the LRO options.
type WrappedPrepareDeviceBackupOperation struct {
	*innerClient.PrepareDeviceBackupOperation
	pathPrefix string
}

// PathPrefix returns the path prefix for the operation to be used in HTTP requests.
// E.g. if the path prefix is `/waas/mpc_keys`, then the request path would be:
// `/waas/mpc_keys/v1/operations/<operationId>`.
func (w *WrappedPrepareDeviceBackupOperation) PathPrefix() string {
	return w.pathPrefix
}

// Wait delegates to the wrapped longrunning PrepareDeviceBackupOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedPrepareDeviceBackupOperation) Wait(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := w.PrepareDeviceBackupOperation.Wait(ctx, clients.LROOptions(w, version, opts)...)

	return deviceGroup, clients.UnwrapError(err)
}

// Poll delegates to the wrapped longrunning PrepareDeviceBackupOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedPrepareDeviceBackupOperation) Poll(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := w.PrepareDeviceBackupOperation.Poll(ctx, clients.LROOptions(w, version, opts)...)

	return deviceGroup, clients.UnwrapError(err)
}

// Metadata delegates to the wrapped longrunning PrepareDeviceBackupOperation.
func (w *WrappedPrepareDeviceBackupOperation) Metadata() (*mpc_keyspb.PrepareDeviceBackupMetadata, error) {
	return w.PrepareDeviceBackupOperation.Metadata()
}

// Done delegates to the wrapped longrunning PrepareDeviceBackupOperation.
func (w *WrappedPrepareDeviceBackupOperation) Done() bool {
	return w.PrepareDeviceBackupOperation.Done()
}

// Name delegates to the wrapped longrunning PrepareDeviceBackupOperation.
func (w *WrappedPrepareDeviceBackupOperation) Name() string {
	return w.PrepareDeviceBackupOperation.Name()
}

// PrepareDeviceBackup prepares a backup in the given Device. The backup contains certain
// cryptographic materials that can be used to restore MPCKeys, which have the given DeviceGroup as their parent,
// on a new Device. The Device specified in the request must be a member of this DeviceGroup and must participate
// in the associated MPC operation for the backup to be prepared.
// After calling this RPC, use ListMPCOperations to poll for the pending PrepareDeviceBackup operation,
// and use the WaaS SDK's ComputeMPCOperation to complete the operation. Once the operation completes,
// the Device can utilize WaaS SDK to download the backup bundle. We recommend storing this backup bundle securely
// in a storage provider of your choice. If the user loses access to their existing Device and wants to recover
// MPCKeys in the given DeviceGroup on a new Device, use AddDevice RPC on the MPCKeyService.
func (m *mpcKeyServiceClient) PrepareDeviceBackup(
	ctx context.Context,
	req *mpc_keyspb.PrepareDeviceBackupRequest,
	opts ...gax.CallOption) (ClientWrappedPrepareDeviceBackupOperation, error) {
	op, err := m.client.PrepareDeviceBackup(ctx, req, opts...)
	if err != nil {
		return nil, clients.UnwrapError(err)
	}

	return &WrappedPrepareDeviceBackupOperation{
		PrepareDeviceBackupOperation: op,
		pathPrefix:                   m.pathPrefix,
	}, nil
}

// PrepareDeviceBackupOperation returns the PrepareDeviceBackupOperation indicated by the given name.
func (m *mpcKeyServiceClient) PrepareDeviceBackupOperation(name string) ClientWrappedPrepareDeviceBackupOperation {
	return &WrappedPrepareDeviceBackupOperation{
		PrepareDeviceBackupOperation: m.client.PrepareDeviceBackupOperation(name),
		pathPrefix:                   m.pathPrefix,
	}
}

// WrappedAddDeviceOperation wraps the long-running operation to handle
// unwrapping errors and setting the LRO options.
type WrappedAddDeviceOperation struct {
	*innerClient.AddDeviceOperation
	pathPrefix string
}

// PathPrefix returns the path prefix for the operation to be used in HTTP requests.
// E.g. if the path prefix is `/waas/mpc_keys`, then the request path would be:
// `/waas/mpc_keys/v1/operations/<operationId>`.
func (w *WrappedAddDeviceOperation) PathPrefix() string {
	return w.pathPrefix
}

// Wait delegates to the wrapped longrunning AddDeviceOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedAddDeviceOperation) Wait(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := w.AddDeviceOperation.Wait(ctx, clients.LROOptions(w, version, opts)...)

	return deviceGroup, clients.UnwrapError(err)
}

// Poll delegates to the wrapped longrunning AddDeviceOperation with the
// override LRO options and handling unwrapping client errors.
func (w *WrappedAddDeviceOperation) Poll(
	ctx context.Context,
	opts ...gax.CallOption,
) (*mpc_keyspb.DeviceGroup, error) {
	deviceGroup, err := w.AddDeviceOperation.Poll(ctx, clients.LROOptions(w, version, opts)...)

	return deviceGroup, clients.UnwrapError(err)
}

// Metadata delegates to the wrapped longrunning AddDeviceOperation.
func (w *WrappedAddDeviceOperation) Metadata() (*mpc_keyspb.AddDeviceMetadata, error) {
	return w.AddDeviceOperation.Metadata()
}

// Done delegates to the wrapped longrunning AddDeviceOperation.
func (w *WrappedAddDeviceOperation) Done() bool {
	return w.AddDeviceOperation.Done()
}

// Name delegates to the wrapped longrunning AddDeviceOperation.
func (w *WrappedAddDeviceOperation) Name() string {
	return w.AddDeviceOperation.Name()
}

// AddDevice adds a Device to an existing DeviceGroup. Prior to this API being called, the Device must be registered using
// RegisterDevice RPC. The Device must have access to the backup created with PrepareDeviceBackup RPC to compute this
// operation. After calling this RPC, use ListMPCOperations to poll for the pending AddDevice operation,
// and use the WaaS SDK's ComputeAddDeviceMPCOperation to complete the operation.
// After the operation is computed on WaaS SDK, the Device will have access to cryptographic materials
// required to process MPCOperations for this DeviceGroup.
// Once the operation completes on MPCKeyService, the Device will be added to the given DeviceGroup as a new member
// and all existing Devices in the DeviceGroup will stay functional.
// MPCKeyService will expose RemoveDevice RPC in a future release that can remove any of the
// existing Devices from the DeviceGroup.
func (m *mpcKeyServiceClient) AddDevice(
	ctx context.Context,
	req *mpc_keyspb.AddDeviceRequest,
	opts ...gax.CallOption) (ClientWrappedAddDeviceOperation, error) {
	op, err := m.client.AddDevice(ctx, req, opts...)
	if err != nil {
		return nil, clients.UnwrapError(err)
	}

	return &WrappedAddDeviceOperation{
		AddDeviceOperation: op,
		pathPrefix:         m.pathPrefix,
	}, nil
}

// AddDeviceOperation returns the AddDeviceOperation indicated by the given name.
func (m *mpcKeyServiceClient) AddDeviceOperation(name string) ClientWrappedAddDeviceOperation {
	return &WrappedAddDeviceOperation{
		AddDeviceOperation: m.client.AddDeviceOperation(name),
		pathPrefix:         m.pathPrefix,
	}
}

// RevokeDevice revokes a registered Device. This operation removes the registered Device from all the DeviceGroups that it is a
// part of. Once the Device is revoked, cryptographic materials in your physical Device are invalidated,
// and the Device can no longer participate in any MPCOperations of the DeviceGroups it was a part of.
// Use this API in scenarios such as losing the existing Device, switching to a new physical Device, etc.
// Ensure that a new Device is successfully added to your DeviceGroups using the AddDevice RPC before invoking
// the RevokeDevice RPC.
func (m *mpcKeyServiceClient) RevokeDevice(
	ctx context.Context,
	req *mpc_keyspb.RevokeDeviceRequest,
	opts ...gax.CallOption) error {
	err := m.client.RevokeDevice(ctx, req, opts...)

	return clients.UnwrapError(err)
}

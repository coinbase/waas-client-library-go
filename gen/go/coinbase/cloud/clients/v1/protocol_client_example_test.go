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

package v1_test

import (
	"context"

	v1 "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	protocolspb "github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
)

func ExampleNewProtocolClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := v1.NewProtocolClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewProtocolRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := v1.NewProtocolRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleProtocolClient_ConstructTransaction() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := v1.NewProtocolClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &protocolspb.ConstructTransactionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1#ConstructTransactionRequest.
	}
	resp, err := c.ConstructTransaction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProtocolClient_ConstructTransferTransaction() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := v1.NewProtocolClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &protocolspb.ConstructTransferTransactionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1#ConstructTransferTransactionRequest.
	}
	resp, err := c.ConstructTransferTransaction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProtocolClient_BroadcastTransaction() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := v1.NewProtocolClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &protocolspb.BroadcastTransactionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.cbhq.net/cloud/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1#BroadcastTransactionRequest.
	}
	resp, err := c.BroadcastTransaction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

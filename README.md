# WaaS Go Client Library

This repository contains the Protocol Buffer definitions for the Coinbase **Wallet-as-a-Service** (WaaS)
APIs, as well as the Go client libraries generated from them.

## Overview

WaaS is Coinbase's suite of cloud APIs for creating, managing, and using Multi-Party Computation
(MPC)-based crypto wallets. Designed for use by nimble start-ups and large enterprise clients alike,
WaaS APIs are self-service and composable, allowing developers to easily create applications that
interact with the blockchain and inherit the maximal security properties of MPC-based key management.

## Prerequisites

- [Golang 1.17+](https://go.dev/learn/)
- [node 17+](https://nodejs.org/en/download/)
- [yarn 1.22+](https://yarnpkg.com/getting-started/install)

For iOS development:
- [Xcode 14.0+](https://developer.apple.com/xcode/)
  - iOS15.2+ simulator (iPhone 14 recommended)
- [CocoaPods](https://guides.cocoapods.org/using/getting-started.html)

For Android development:
- [Android Studio](https://developer.android.com/studio)
  - 64-bit Android emulator
- [Android NDK 30+](https://developer.android.com/ndk)

## Repository Structure
- [`auth/`](./auth/) contains the authentication-related code for accessing WaaS APIs.
- [`clients/`](./clients/) contains client instantiation helpers for WaaS APIs.
- [`gen/`](./gen/) contains Go code generated from the Protocol Buffers.
- [`protos/`](./protos/) contains the Protocol Buffers that define the WaaS APIs.

## Module Installation
```
go get github.com/coinbase/waas-client-library-go
```

## Get Started
To test that your API Key gives you access as expected to the WaaS APIs:

1. Replace `apiKeyName` and `apiKeyPrivateKey` in [`example.go`](./example.go) with your API Key information.
2. Run `go build`.
3. Run `./waas-client-library-go`.
4. You should see output like the following:
```
2023/03/15 12:23:16 got Ethereum Goerli network: name:"networks/ethereum-goerli"  display_name:"Goerli Ethereum Testnet"  native_asset:"networks/ethereum-goerli/assets/0c3569d3-b253-5128-a229-543e1e819430"  protocol_family:"protocolFamilies/evm"  type:TESTNET
```

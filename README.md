# WaaS Go Client Library

This repository contains the Protocol Buffer definitions for the Coinbase **Wallet-as-a-Service** (WaaS)
APIs, as well as the Go client libraries generated from them.

## Overview

WaaS is Coinbase's suite of cloud APIs for creating, managing, and using Multi-Party Computation
(MPC)-based crypto wallets. Designed for use by nimble start-ups and large enterprise clients alike,
WaaS APIs are self-service and composable, allowing developers to easily create applications that
interact with the blockchain and inherit the maximal security properties of MPC-based key management.

## Documentation

For full documentation, refer to [docs.cloud.coinbase.com/waas](https://docs.cloud.coinbase.com/waas/).

For APIs that require multi-party computation (MPC), refer to the [WaaS SDK repository](https://github.com/coinbase/waas-sdk-react-native).

## Prerequisites

- [Golang 1.17+](https://go.dev/learn/)

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

1. Replace `apiKeyName` and `privKeyTemplate` in [`example.go`](./example.go) with your API Key information.
2. Run `go build`.
3. Run `./waas-client-library-go`.
4. You should see output like the following:
```
2023/03/17 12:37:35 creating pool...
2023/03/17 12:37:35 created pool: name:"pools/e08c1784-f2be-4b77-8307-a9ea2d8d0017"  display_name:"My First Pool"
2023/03/17 12:37:35 listing networks...
2023/03/17 12:37:35 got network: name:"networks/ethereum-goerli"  display_name:"Goerli Ethereum Testnet"  native_asset:"networks/ethereum-goerli/assets/0c3569d3-b253-5128-a229-543e1e819430"  protocol_family:"protocolFamilies/evm"  type:TESTNET
2023/03/17 12:37:35 listing first 5 assets on Ethereum Goerli...
2023/03/17 12:37:35 got asset: name:"networks/ethereum-goerli/assets/0c3569d3-b253-5128-a229-543e1e819430"  advertised_symbol:"ETH"  decimals:18  definition:{asset_type:"native"}
2023/03/17 12:37:35 got asset: name:"networks/ethereum-goerli/assets/adbf9e76-de39-51a0-9e53-5f8ef31b7925"  advertised_symbol:"POLY"  decimals:18  definition:{asset_type:"erc20"  asset_group_id:"0x887CFe31C888EE0780795b7feFF46CE7f9AB556C"}
2023/03/17 12:37:35 got asset: name:"networks/ethereum-goerli/assets/a055a425-fe93-51ae-9099-cf5495db6e79"  advertised_symbol:"SIM"  decimals:18  definition:{asset_type:"erc20"  asset_group_id:"0x0E89BF4135acE3d4d67BF828707746D3855f3a25"}
2023/03/17 12:37:35 got asset: name:"networks/ethereum-goerli/assets/145f3157-a45d-5a77-92be-6b2af8b7af12"  advertised_symbol:"TERC20"  decimals:18  definition:{asset_type:"erc20"  asset_group_id:"0xea100Bec80418680e55D28b655da6CbEF427275f"}
2023/03/17 12:37:35 got asset: name:"networks/ethereum-goerli/assets/20b2830e-53c1-5540-9d1b-0061df3555f6"  advertised_symbol:"BETH"  decimals:18  definition:{asset_type:"erc20"  asset_group_id:"0xED6CCd7e5131073aE67221B1cA195db0fFacc940"}
```

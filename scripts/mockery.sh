#!/usr/bin/env bash

# Fail on any error, unset variable, or failed pipe command
set -euo pipefail

COMPONENT=$1

CAMEL_COMPONENT=""

if [ "$COMPONENT" = "pools" ]; then
  CAMEL_COMPONENT="Pool"
elif [ "$COMPONENT" = "blockchain" ]; then
  CAMEL_COMPONENT="Blockchain"
elif [ "$COMPONENT" = "protocols" ]; then
  CAMEL_COMPONENT="Protocol"
elif [ "$COMPONENT" = "mpc_keys" ]; then
  CAMEL_COMPONENT="MPCKey"
elif [ "$COMPONENT" = "mpc_wallets" ]; then
  CAMEL_COMPONENT="MPCWallet"
elif [ "$COMPONENT" = "mpc_transactions" ]; then
  CAMEL_COMPONENT="MPCTransaction"
fi

if [ "$CAMEL_COMPONENT" = "" ]; then
  echo "Invalid component $COMPONENT."
  exit 1
fi

# Define inputs to mockery.
DIR="./gen/go/coinbase/cloud/$COMPONENT/v1"
CLIENT_IFACE="${CAMEL_COMPONENT}ServiceClient"
SERVER_IFACE="${CAMEL_COMPONENT}ServiceServer"
OUTPUT_DIR="./gen/go/coinbase/cloud/$COMPONENT/v1/mocks"

echo "Generating mocks for $COMPONENT..."

mockery \
  --log-level=error \
  --dir=$DIR \
  --name=$CLIENT_IFACE \
  --output=$OUTPUT_DIR \
  --disable-version-string

mockery \
  --log-level=error \
  --dir=$DIR \
  --name=$SERVER_IFACE \
  --output=$OUTPUT_DIR \
  --disable-version-string

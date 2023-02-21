#!/usr/bin/env bash

# Fail on any error, unset variable, or failed pipe command
set -euo pipefail

COMPONENT=$1

# Strip component of trailing 's', if one exists.
SINGULAR_COMPONENT=${COMPONENT%s}

# Convert component to uppercase.
UPPERCASE_SINGULAR_COMPONENT="$(tr '[:lower:]' '[:upper:]' <<< ${SINGULAR_COMPONENT:0:1})${SINGULAR_COMPONENT:1}"

# Define inputs to mockery.
DIR="./gen/go/coinbase/cloud/$COMPONENT/v1"
CLIENT_IFACE="${UPPERCASE_SINGULAR_COMPONENT}ServiceClient"
SERVER_IFACE="${UPPERCASE_SINGULAR_COMPONENT}ServiceServer"
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

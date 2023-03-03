#!/usr/bin/env bash
# Detects whether any files generated from protos have changed from the last Git commit.

# Fail on any error, unset variable, or failed pipe command
set -euo pipefail

# Ignore the doc.go file, because its comments can change every time while the file contents
# do not.
git update-index --assume-unchanged gen/go/coinbase/cloud/clients/v1/doc.go

# Make all of the generated files.
make all

# Detect any differences.
DIFF=$(git status --porcelain 'gen/*')
if [ "$DIFF" = "" ]; then
    echo "No generated files changed."
    exit 0
else
    echo "Some generated files changed. Run 'make all' before pushing PR.\n Changed files:"
    echo "$DIFF"
    exit 1
fi

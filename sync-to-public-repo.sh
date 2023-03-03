#!/bin/bash

# This script syncs code to public repo at https://github.com/coinbase/waas-client-library-go
##

INTERNAL_REPO_URL=git@github.cbhq.net:cloud/waas-client-library-go.git
PUBLIC_REPO_URL=git@github.com:coinbase/waas-client-library-go.git
INTERNAL_BRANCH_TO_SYNC=master
PUBLIC_BRANCH_TO_SYNC=main
RELEASE_VERSION=0.0.1
EPOCH_TIME=$(date +%s)

echo "Left: $INTERNAL_REPO_URL"
echo "Right: $PUBLIC_REPO_URL"

echo "Fetch latest commits from branch $INTERNAL_BRANCH_TO_SYNC in $INTERNAL_REPO_URL"
if ! git fetch -u $INTERNAL_REPO_URL $INTERNAL_BRANCH_TO_SYNC:left/$INTERNAL_BRANCH_TO_SYNC; then
  echo >&2 "Fatal: unable to fetch from $INTERNAL_REPO_URL, rerun the script as soon as connection restored."
  exit 1
fi

echo "Fetch latest commits from branch $PUBLIC_BRANCH_TO_SYNC in $PUBLIC_REPO_URL"
if ! git fetch -u $PUBLIC_REPO_URL $PUBLIC_BRANCH_TO_SYNC:right/$PUBLIC_BRANCH_TO_SYNC; then
  echo >&2 "Fatal: unable to fetch from $PUBLIC_REPO_URL, rerun the script as soon as connection restored."
  exit 1
fi

echo "Checkout to the HEAD of the public repo"
if ! git checkout --quiet -b $PUBLIC_BRANCH_TO_SYNC-"$EPOCH_TIME" right/$PUBLIC_BRANCH_TO_SYNC; then
  echo >&2 "Fatal: unable checkout to the HEAD of the public repo from $PUBLIC_REPO_URL, rerun the script as soon as connection restored."
  exit 1
fi

echo "Merge $INTERNAL_BRANCH_TO_SYNC from $INTERNAL_REPO_URL"
echo "Touch your Yubikey to sign the commit..."
if ! git merge -m "Merge to sync between $INTERNAL_REPO_URL and $PUBLIC_REPO_URL" --log left/$INTERNAL_BRANCH_TO_SYNC --allow-unrelated-histories; then
  echo >&2 "Fatal: unable to merge from $INTERNAL_REPO_URL, rerun the script after resolving the conflicts manually."
  exit 1
fi

echo "Replace the proto package name"
if ! grep -rl 'github.cbhq.net/cloud' ./protos | xargs sed -i "" -e 's/github\.cbhq\.net\/cloud/github.com\/coinbase/g'; then
  echo >&2 "Fatal: unable to replace the package from github.cbhq.net/cloud to github.com/coinbase"
  exit 1
fi

echo "Generate protos"
if ! (make all); then
  echo >&2 "Fatal: unable to generate protos"
  exit 1
fi

echo "Remove files before publishing to public repo"
#Add sync-to-public-repo.sh once it's merged to internal repo
#if ! (rm -r .buildkite scripts .codeflow.yml Dockerfile.presubmit api-linter.yml Makefile sync-to-public-repo.sh); then
if ! (rm -r .buildkite scripts .codeflow.yml Dockerfile.presubmit api-linter.yml Makefile); then
  echo >&2 "Fatal: unable to remove files before publishing to public repo"
  exit 1
fi

echo "Commit changes"
echo "Touch your Yubikey to sign the commit..."
if ! (git add --all && git commit -m "release $RELEASE_VERSION"); then
  echo >&2 "Fatal: unable to commit changes"
  exit 1
fi

echo "Create PR for public repo"
if ! (git remote add public-"$EPOCH_TIME" $PUBLIC_REPO_URL && git push public-"$EPOCH_TIME" $INTERNAL_BRANCH_TO_SYNC-"$EPOCH_TIME"); then
  echo >&2 "Fatal: unable to create PR for public repo"
  exit 1
fi
echo "Done."

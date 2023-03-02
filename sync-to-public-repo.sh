#!/bin/bash

# This script syncs code to public repo at https://github.com/coinbase/waas-client-library-go
##

INTERNAL_REPO_URL=git@github.cbhq.net:cloud/waas-client-library-go.git
PUBLIC_REPO_URL=git@github.com:coinbase/waas-client-library-go.git
BRANCH_TO_SYNC=master
EPOCH_TIME=$(date +%s)


echo "Left: $INTERNAL_REPO_URL"
echo "Right: $PUBLIC_REPO_URL"

echo "Fetch latest commits from branch $BRANCH_TO_SYNC in $INTERNAL_REPO_URL"
if ! git fetch -u $INTERNAL_REPO_URL $BRANCH_TO_SYNC:left/$BRANCH_TO_SYNC; then 
	>&2 echo "Fatal: unable to fetch from $INTERNAL_REPO_URL, rerun the script as soon as connection restored."
	exit 1
fi

echo "Fetch latest commits from branch $BRANCH_TO_SYNC in $PUBLIC_REPO_URL"
if ! git fetch -u $PUBLIC_REPO_URL $BRANCH_TO_SYNC:right/$BRANCH_TO_SYNC; then
	>&2 echo "Fatal: unable to fetch from $PUBLIC_REPO_URL, rerun the script as soon as connection restored."
	exit 1
fi

echo "Checkout to the HEAD of the public repo"
if ! git checkout --quiet -b $BRANCH_TO_SYNC-"$EPOCH_TIME" right/$BRANCH_TO_SYNC; then
	>&2 echo "Fatal: unable checkout to the HEAD of the public repo from $PUBLIC_REPO_URL, rerun the script as soon as connection restored."
	exit 1
fi

echo "Merge $BRANCH_TO_SYNC from $INTERNAL_REPO_URL"
	if ! git merge -m "Merge to sync between $INTERNAL_REPO_URL and $PUBLIC_REPO_URL" --log left/$BRANCH_TO_SYNC --allow-unrelated-histories; then
	>&2 echo "Fatal: unable to fetch from $PUBLIC_REPO_URL, rerun the script as soon as connection restored."
	exit 1
fi


#
#if git checkout --quiet -b sync-$BRANCH_TO_SYNC right/$BRANCH_TO_SYNC; then
#	echo "Merge branches from left and right if necessary."
#	if ! git merge -m "Merge to sync between $INTERNAL_REPO_URL and $PUBLIC_REPO_URL" --log left/$BRANCH_TO_SYNC; then
#                >&2 echo "Merge conflict. Solve is manually, commit and rerun script."
#                exit 1
#    fi
#else
#	echo "Rerun after merge conflict resolution or restored connection."
#	git checkout --quiet sync-$BRANCH_TO_SYNC
#	echo "Try to merge with right first"
#	if ! git merge -m "Merge to sync between $INTERNAL_REPO_URL and $PUBLIC_REPO_URL" --log right/$BRANCH_TO_SYNC; then
#                >&2 echo "Merge conflict. Solve is manually, commit and rerun script."
#                exit 1
#    fi
#	if ! git merge -m "Merge to sync between $INTERNAL_REPO_URL and $PUBLIC_REPO_URL" --log left/$BRANCH_TO_SYNC; then
#                >&2 echo "Merge conflict. Solve is manually, commit and rerun script."
#                exit 1
#    fi
#fi
#
#echo "Push merged changes in $BRANCH_TO_SYNC to $PUBLIC_REPO_URL"
#if ! git push $PUBLIC_REPO_URL HEAD:$BRANCH_TO_SYNC; then
#	>&2 echo "Fatal: unable to push to $PUBLIC_REPO_URL, rerun the script as soon as connection restored."
#	exit 1
#
#fi
#echo "Push merged changes in $BRANCH_TO_SYNC to $INTERNAL_REPO_URL"
#if ! git push $INTERNAL_REPO_URL HEAD:$BRANCH_TO_SYNC; then
#	>&2 echo "Fatal: unable to push to $INTERNAL_REPO_URL, rerun the script as soon as connection restored."
#	exit 1
#
#fi
#
#git checkout --quiet master
#git branch -D --quiet sync-$BRANCH_TO_SYNC

echo "Done."

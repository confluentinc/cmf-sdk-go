#!/bin/bash

tmpFile=$(mktemp)

git fetch --tags
baseVersion=$(git tag --sort=-v:refname | head -n 1)

go install golang.org/x/exp/cmd/gorelease@latest

echo gorelease -base=${baseVersion}
gorelease -base="${baseVersion}" | tee $tmpFile

grep 'incompatible' $tmpFile >/dev/null 2>/dev/null

if [[ $? -eq 0 ]]; then
	cat $tmpFile
	echo "There are some incompatible changes in the SDK. Please review them and only merge this if you're certain that's what you want."
	exit 1
fi

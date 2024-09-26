#!/bin/bash

set -eo pipefail
function print() {
  echo -e $1
  printf %${#1}s | tr ' ' '='
  echo
}

[[ $# -ne 3 ]] && echo "Usage: $0 <path/to/openapi-generator-cli.jar> <path/to/client-go/cmf/version> <path/to/metadata-service/../spec.yaml>" && exit 1

openapiPath=$1
tgtPath=$2
cmfPath=$3
version=$(basename ${tgtPath})

cd $tgtPath

which mocker > /dev/null || (print "Installing mocker"; go get github.com/travisjeffery/mocker/cmd/mocker; go install github.com/travisjeffery/mocker/cmd/mocker; echo)

print "Removing all files under $(pwd)"
find . ! '(' -name 'Makefile' -o -name '.' ')' -maxdepth 1 -print0 | xargs -0 rm -rf

print "\nGenerating cmf-sdk-go"
java -ea ${JAVA_OPTS} -Xms512M -Xmx1024M -server \
  -jar ${openapiPath} generate \
  --input-spec ${cmfPath} \
  --api-package ${version} --package-name ${version} --generator-name go-deprecated \
  --git-user-id confluentinc --git-repo-id cmf-sdk-go \
  --additional-properties generateInterfaces=true,enumClassPrefix=true,isGoSubmodule=true

print "\nGenerating mocks"
make mock
print "\nFormatting go"
make fmt

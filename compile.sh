#!/usr/bin/env bash

set -e

ARCH=${1} # only riscv64, arm64, amd64





NAMES=`ls ./leetcode`
mkdir -p $ARCH

for NAME in $NAMES
do
cd ./leetcode/$NAME
GOARCH=${ARCH} go test -c -o $NAME
mv $NAME ../../${ARCH}
echo $NAME compiled!
cd ../..
done

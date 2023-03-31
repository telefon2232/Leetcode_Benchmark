#!/usr/bin/env bash

set -e

for d in $(go list ./leetcode/... | grep -v vendor); do
    echo $d
    go test -benchtime 1x -bench .     $d | grep ns 
   # go test -c -o foo  -> ./foo -test.bench .
done

#!/usr/bin/env bash

set -e

for d in $(go list ./leetcode/... | grep -v vendor); do
    echo $d
    go test -bench . -benchtime 1x  -count=1 -cpuprofile cpu.out  $d | grep ns 
   
done

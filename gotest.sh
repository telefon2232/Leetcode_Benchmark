#!/usr/bin/env bash

set -e

for d in $(go list ./leetcode/... | grep -v vendor); do
    echo -n "$d " | awk -F"/" -v ORS='' '{print $NF}'
    go test -benchtime 1x -bench .  $d | grep ns/op | awk '{print $3, $4}' 


    # go test -c -o foo  #-> ./foo -test.bench .
    # GOARCH=arm go test -c -o foo
    # GOARCH=riscv64 go test -c -o foo
    # go tool pprof -top  cpu.out
done
echo "Finished `ls ./leetcode/ | wc -l` tests!"

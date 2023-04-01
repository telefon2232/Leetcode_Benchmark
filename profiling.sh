#!/usr/bin/env bash

set -e


ARCH=${1}


cd ${ARCH}
FILES=`ls`
for FILE in ${FILES}
do
./${FILE} -test.bench .  -test.benchtime 10s -test.cpuprofile ${FILE}.out # go tool pprof -top  cpu.out
done
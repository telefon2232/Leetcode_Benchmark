#!/usr/bin/env bash

set -e


ARCH=${1}


cd ${ARCH}
FILES=`ls`
for FILE in ${FILES}
do
     echo -n "$FILE " | awk -F"/" -v ORS='' '{print $NF}'
    ./${FILE} -test.benchtime 1s -test.bench .   | grep ns/op | awk '{print $3, $4}'
	
#	./${FILE} -test.bench .  -test.benchtime 1s # go tool pprof -top  cpu.out
done

#!/bin/bash

PORT=9920
DOCKERNAME=pxtester
MOCKSDKTAG=0.42.14

fail()
{
	echo "$1"
	exit 1
}

startdocker()
{
	docker run --rm --name pxtester -d -p ${PORT}:9100 openstorage/mock-sdk-server:${MOCKSDKTAG} > /dev/null 2>&1
	if [ $? -ne 0 ] ; then
		fail "Failed to start docker"
	fi
}

stopdocker()
{
	docker stop $DOCKERNAME > /dev/null 2>&1
	if [ $? -ne 0 ] ; then
		fail "Failed to stop docker"
	fi
}

### MAIN ###

startdocker

export PXTESTCONFIG=$PWD/hack/config.yml

result=0
if [ $# -eq 0 ] ; then
	go test $(go list ./... | grep -v vendor | grep -v plugins)
	result=$?
else
	go test $@
	result=$?
fi

stopdocker
exit $result

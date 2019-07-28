#!/bin/bash

MOCKSDKTAG=0.42.14

fail()
{
	echo "$1"
	exit 1
}

# $1 - name
# $2 - port
startdocker()
{
	docker run --rm --name ${1} -d -p ${2}:9100 openstorage/mock-sdk-server:${MOCKSDKTAG} > /dev/null 2>&1
	if [ $? -ne 0 ] ; then
		fail "Failed to start docker"
	fi
}

# $1 - name
stopdocker()
{
	docker stop $1 > /dev/null 2>&1
	if [ $? -ne 0 ] ; then
		fail "Failed to stop docker"
	fi
}

### MAIN ###

startdocker pxutsource 9920
startdocker pxuttarget 9921

export PXTESTCONFIG=$PWD/hack/config.yml

result=0
if [ $# -eq 0 ] ; then
	go test $(go list ./... | grep -v vendor | grep -v plugins)
	result=$?
else
	go test $@
	result=$?
fi

stopdocker pxutsource
stopdocker pxuttarget
exit $result

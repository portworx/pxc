#!/bin/bash

MOCKSDKTAG=0.42.14
CONTAINER=quay.io/lpabon/mock-sdk-server

fail()
{
	echo "$1"
	exit 1
}

createConfig()
{
 ./pxc --pxc.config=$TESTCONFIG config cluster set --name=target --endpoint=localhost:9921 || fail "Failed to create target cluster"
 ./pxc --pxc.config=$TESTCONFIG config cluster set --name=source --endpoint=localhost:9920 || fail "Failed to create source clusterh"
 ./pxc --pxc.config=$TESTCONFIG config context set --cluster=source --name=source || fail "Failed to set context source"
 ./pxc --pxc.config=$TESTCONFIG config context set --cluster=target --name=target || fail "Failed to set context target"
 ./pxc --pxc.config=$TESTCONFIG config context use --name=source || fail "Failed to use source context"
}

# $1 - name
# $2 - port
startdocker()
{
	docker run --rm --name ${1} -d -p ${2}:9100 ${CONTAINER} > /dev/null 2>&1
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

export TESTCONFIG=/tmp/pxc-testconfig.yml
export PXCONFIG=$TESTCONFIG

createConfig
startdocker pxutsource 9920
startdocker pxuttarget 9921

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

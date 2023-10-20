#!/bin/bash

go version
git config --global --add safe.directory /go/src

apt update && \
	apt install -y zip && \
	apt autoremove && \
	make docker-release && \
	chown -R ${DEV_USER}:${DEV_GROUP} dist


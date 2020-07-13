#!/bin/bash

go version

apt update && \
	apt install -y zip && \
	apt autoremove && \
	make docker-release && \
	chown -R ${DEV_USER}:${DEV_GROUP} dist


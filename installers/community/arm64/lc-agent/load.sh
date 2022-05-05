#!/bin/bash

# shellcheck source=.env
source .env

docker load --input "./archives/labs-lightcrane-agent-arm64:${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/labs-lightcrane-pyservice-base-arm64:${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/arm64v8_ubuntu:${UBUNTU_VERSION}.tar" || exit 1
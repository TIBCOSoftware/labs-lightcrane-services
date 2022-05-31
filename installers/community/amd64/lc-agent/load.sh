#!/bin/bash

# shellcheck source=.env
source .env

docker load --input "./archives/labs-lightcrane-agent:${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/labs-lightcrane-pyservice-base:${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/ubuntu:${UBUNTU_VERSION}.tar" || exit 1
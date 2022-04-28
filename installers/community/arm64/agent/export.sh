#!/bin/bash
mkdir -p ./archives

source .env

docker-compose pull || exit 1

docker save --output ./archives/labs-lightcrane-agent-arm64:${LABS_AIR_VERSION}.tar public.ecr.aws/tibcolabs/labs-lightcrane-agent-arm64:${LABS_AIR_VERSION} || exit 1
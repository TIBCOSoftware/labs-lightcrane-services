#!/bin/bash
mkdir -p ./archives

# shellcheck source=.env
source .env

docker-compose pull || exit 1

docker save --output "./archives/labs-lightcrane-agent-arm64:${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-agent-arm64:${LABS_AIR_VERSION}" || exit 1

docker pull public.ecr.aws/tibcolabs/labs-lightcrane-pyservice-base-arm64:${LABS_AIR_VERSION}
docker save --output "./archives/labs-lightcrane-pyservice-base-arm64:${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-pyservice-base-arm64:${LABS_AIR_VERSION}" || exit 1

docker pull arm64v8/ubuntu:${UBUNTU_VERSION}
docker save --output "./archives/arm64v8/ubuntu:${UBUNTU_VERSION}.tar" "arm64v8/ubuntu:${UBUNTU_VERSION}" || exit 1
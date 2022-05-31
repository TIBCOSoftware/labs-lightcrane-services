#!/bin/bash
mkdir -p ./archives

# shellcheck source=.env
source .env

docker-compose pull || exit 1

docker save --output "./archives/labs-lightcrane-agent:${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-agent:${LABS_AIR_VERSION}" || exit 1

docker pull public.ecr.aws/tibcolabs/labs-lightcrane-pyservice-base:${LABS_AIR_VERSION}
docker save --output "./archives/labs-lightcrane-pyservice-base:${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-pyservice-base:${LABS_AIR_VERSION}" || exit 1

docker pull ubuntu:${UBUNTU_VERSION}
docker save --output "./archives/ubuntu:${UBUNTU_VERSION}.tar" "ubuntu:${UBUNTU_VERSION}" || exit 1
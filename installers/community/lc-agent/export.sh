#!/bin/bash
mkdir -p ./archives

# shellcheck source=.env
source .env

if [ -n "$1" ]; then
    case "$1" in
    arm64)
    export ARCH=-$1
    export ARCH_PF="arm64v8/"
    ;;
    amd64)
    export ARCH=""
    export ARCH_PF=""
    ;;
    esac
else
	export ARCH=""
    export ARCH_PF=""
fi

docker-compose pull || exit 1

docker save --output "./archives/labs-lightcrane-agent:${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-agent${ARCH}:${LABS_AIR_VERSION}" || exit 1

docker pull public.ecr.aws/tibcolabs/labs-lightcrane-pyservice-base${ARCH}:${LABS_AIR_VERSION}
docker save --output "./archives/labs-lightcrane-pyservice-base:${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-pyservice-base${ARCH}:${LABS_AIR_VERSION}" || exit 1

docker pull ${ARCH_PF}ubuntu:${UBUNTU_VERSION}
docker save --output "./archives/ubuntu:${UBUNTU_VERSION}.tar" "${ARCH_PF}ubuntu:${UBUNTU_VERSION}" || exit 1
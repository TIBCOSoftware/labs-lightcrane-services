#!/bin/bash
mkdir -p ./archives

# shellcheck source=.env
source .env

if [ -n "$1" ]; then
    case "$1" in
    arm64)
    export ARCH=-$1
    ;;
    amd64)
    export ARCH=""
    ;;
    esac
else
	export ARCH=""
fi


docker-compose pull || exit 1

docker save --output "./archives/labs-lightcrane-proxy-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-proxy${ARCH}:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-servicelocator-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-servicelocator${ARCH}:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-projectmgr-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-projectmgr${ARCH}:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-builder-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-builder${ARCH}:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-deployer-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-deployer${ARCH}:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-air-service-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-air-service${ARCH}:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-model-service-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-model-service${ARCH}:${LABS_AIR_VERSION}" || exit 1

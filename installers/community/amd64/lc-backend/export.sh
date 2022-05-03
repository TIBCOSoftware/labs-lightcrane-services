#!/bin/bash
mkdir -p ./archives

# shellcheck source=.env
source .env

docker-compose pull || exit 1

docker save --output "./archives/labs-lightcrane-proxy-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-proxy:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-servicelocator-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-servicelocator:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-projectmgr-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-projectmgr:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-builder-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-builder:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-deployer-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-deployer:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-air-service-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-air-service:${LABS_AIR_VERSION}" || exit 1

docker save --output "./archives/labs-lightcrane-model-service-${LABS_AIR_VERSION}.tar" "public.ecr.aws/tibcolabs/labs-lightcrane-model-service:${LABS_AIR_VERSION}" || exit 1

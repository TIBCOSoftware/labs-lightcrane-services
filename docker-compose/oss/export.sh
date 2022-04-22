#!/bin/bash
mkdir -p ./archives

source .env

docker-compose pull

docker save --output ./archives/labs-lightcrane-proxy-${LABS_AIR_VERSION}.tar public.ecr.aws/tibcolabs/labs-lightcrane-proxy:${LABS_AIR_VERSION} || exit 2

docker save --output ./archives/labs-lightcrane-servicelocator-${LABS_AIR_VERSION}.tar public.ecr.aws/tibcolabs/labs-lightcrane-servicelocator:${LABS_AIR_VERSION}

docker save --output ./archives/labs-lightcrane-projectmgr-${LABS_AIR_VERSION}.tar public.ecr.aws/tibcolabs/labs-lightcrane-projectmgr:${LABS_AIR_VERSION}

docker save --output ./archives/labs-lightcrane-builder-${LABS_AIR_VERSION}.tar public.ecr.aws/tibcolabs/labs-lightcrane-builder:${LABS_AIR_VERSION}

docker save --output ./archives/labs-lightcrane-deployer-${LABS_AIR_VERSION}.tar public.ecr.aws/tibcolabs/public.ecr.aws/tibcolabs/labs-lightcrane-deployer:${LABS_AIR_VERSION}

docker save --output ./archives/labs-lightcrane-air-service-${LABS_AIR_VERSION}.tar public.ecr.aws/tibcolabs/labs-lightcrane-air-service:${LABS_AIR_VERSION}

docker save --output ./archives/labs-lightcrane-model-service-${LABS_AIR_VERSION}.tar public.ecr.aws/tibcolabs/labs-lightcrane-model-service:${LABS_AIR_VERSION}

#!/bin/bash

# shellcheck source=.env
source .env

docker load --input "./archives/labs-lightcrane-proxy-${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/labs-lightcrane-servicelocator-${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/labs-lightcrane-projectmgr-${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/labs-lightcrane-builder-${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/labs-lightcrane-deployer-${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/labs-lightcrane-air-service-${LABS_AIR_VERSION}.tar" || exit 1

docker load --input "./archives/labs-lightcrane-model-service-${LABS_AIR_VERSION}.tar" || exit 1

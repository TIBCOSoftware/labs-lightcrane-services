#!/bin/bash

source .env

docker load --input ./archives/labs-lightcrane-agent-arm64:${LABS_AIR_VERSION}.tar || exit 1
#!/bin/bash

source ../setup

# Build docker image
docker image tag flogo/service-locator:$SERVIC_LOCATOR_VERSION bigoyang/service-locator:$SERVIC_LOCATOR_VERSION
docker push bigoyang/service-locator:$SERVIC_LOCATOR_VERSION
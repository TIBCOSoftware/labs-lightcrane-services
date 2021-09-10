#!/bin/bash

source ../setup

# Build executable for Alpine linux/amd64
mv $FLOGO_APP_PATH/ServiceLocator.json ./
$FLOGO_HOME/bin/builder-darwin_amd64 build -f ./ServiceLocator.json -docker -name service-locator -tag=$SERVIC_LOCATOR_VERSION

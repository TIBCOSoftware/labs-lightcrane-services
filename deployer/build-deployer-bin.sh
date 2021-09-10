#!/bin/bash

source ../setup

# Build executable for Alpine linux/amd64
mv $FLOGO_APP_PATH/Deployer.json ./
echo $FLOGO_HOME
$FLOGO_HOME/bin/builder-darwin_amd64 build -p linux/amd64 -f ./Deployer.json -n ./docker

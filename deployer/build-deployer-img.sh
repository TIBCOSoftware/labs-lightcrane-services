#!/bin/bash

source ../setup

# Build docker image
docker build --tag bigoyang/deployer:$DEPLOYER_VERSION ./docker
docker push bigoyang/deployer:$DEPLOYER_VERSION
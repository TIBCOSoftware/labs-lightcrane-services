#!/bin/bash

source ../setup

# Build docker image
docker build --tag bigoyang/builder:$BUILDER_VERSION ./docker
docker push bigoyang/builder:$BUILDER_VERSION
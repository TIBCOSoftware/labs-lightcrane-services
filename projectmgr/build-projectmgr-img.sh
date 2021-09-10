#!/bin/bash

source ../setup

# Build docker image
docker build --tag bigoyang/projectmgr:$PROJ_MGR_VERSION ./docker
docker push bigoyang/projectmgr:$PROJ_MGR_VERSION
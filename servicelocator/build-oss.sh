#!/bin/bash

source ../setup

echo "(Flogo OSS) Build Executable for local platform !!"

cd ./build

rm -rf app

flogo create -f ServiceLocator.json app

cd ./app

export GOOS=linux 
export GOARCH=amd64

flogo build -e --verbose

cp ./bin/app ../../docker/engine

cd ../../docker

# Build docker image
docker build --tag bigoyang/service-locator:$SERVIC_LOCATOR_VERSION .

docker push bigoyang/service-locator:$SERVIC_LOCATOR_VERSION
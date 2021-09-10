#!/bin/bash

source ../setup

echo "(Flogo OSS) Build Executable for local platform !!"

cd ./build

rm -rf app

flogo create -f AirService.json app

cd ./app

export GOOS=linux 
export GOARCH=amd64

flogo build -e --verbose

cp ./bin/app ../../docker/engine

cd ../../docker

# Build docker image
docker build --tag bigoyang/air-service:$AIR_VERSION .

docker push bigoyang/air-service:$AIR_VERSION
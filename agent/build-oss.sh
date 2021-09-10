#!/bin/bash

echo "(Flogo OSS) Build Executable for local platform !!"

cd ./build

rm -rf app

flogo create -f Agent.json app

cd ./app

export GOOS=linux 
export GOARCH=amd64

flogo build -e --verbose

cp ./bin/app ../../docker/engine

cd ../../docker

# Build docker image

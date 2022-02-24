#!/bin/bash

cp -R ../build ./
cp -R ../airpipeline_oss ./

#docker-compose up --build
docker build --tag bigoyang/air-test:0.1.0 -f ./Dockerfile_${1} .

rm -rf ./build
rm -rf ./airpipeline_oss

docker run -it bigoyang/air-test:0.1.0 bash
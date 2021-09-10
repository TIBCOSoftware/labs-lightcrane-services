# Build docker image

source ../setup

docker build --tag bigoyang/air-service:$AIR_VERSION ./docker
docker push bigoyang/air-service:$AIR_VERSION
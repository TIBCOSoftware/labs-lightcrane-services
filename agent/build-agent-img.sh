#!/bin/bash

source ../setup

. ../../scripts/tools.sh

target_platform=$1
image_name=$2
image_tag=$3
image_url=$4

echo "Building agent image for ${target_platform}, image_name: ${image_name}, image_tag: ${image_tag}, image_url: ${image_url}"

docker_file="Dockerfile"
if [ "arm64" = "$target_platform" ]; then
	docker_file="Dockerfile_arm64"
fi

echo "Building dockerfile ${docker_file}"

pushd ./docker > /dev/null

build_image $image_name "local_image_tag" $image_url "${docker_file}"

tag_image $image_name "local_image_tag" $image_name $image_tag $image_url
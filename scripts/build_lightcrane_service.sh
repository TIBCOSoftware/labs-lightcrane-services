#!/bin/bash

. scripts/tools.sh

service_name=$1
app_type=$2
image_name=$3
image_tag=$4
image_url=$5
image_arch=$6

echo "Building Lightcrane service ${service_name}... "

pushd ./${service_name} > /dev/null

docker_file_name="Dockerfile"

if [ "$image_arch" == "amd64" ]; then
    build_image $image_name "local_image_tag" $image_url "Dockerfile" app_type=$app_type
else
    build_image_platform $image_name "local_image_tag" $image_url "Dockerfile_${image_arch}" app_type=$app_type $image_arch
fi

# Tag image
tag_image $image_name "local_image_tag" $image_name $image_tag $image_url
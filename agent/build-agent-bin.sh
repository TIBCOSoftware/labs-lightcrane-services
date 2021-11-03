#!/bin/bash

source ../setup

builder_name="builder-linux_amd64"
target_name=$1

if [ "$(uname)" == "Darwin" ]; then
    builder_name="builder-darwin_amd64"        
elif [ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]; then
    builder_name="builder-windows_amd64.exe"
fi

${builder_name} build -p linux/${target_name} -f Agent.json -n ./docker/engine

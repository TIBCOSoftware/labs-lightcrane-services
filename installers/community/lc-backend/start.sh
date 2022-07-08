#!/bin/bash

export LC_HOME=./LC

if [ ! -d "$LC_HOME" ]
then
    mkdir $LC_HOME
fi

if [ -n "$1" ]; then
    case "$1" in
    arm64)
    export ARCH=-$1
    ;;
    amd64)
    export ARCH=""
    ;;
    esac
else
	export ARCH=""
fi

docker-compose rm -f
docker-compose up -d

#!/bin/bash

export COMPOSE_HTTP_TIMEOUT=200
export USER_HOME=~
export LC_HOME=./LC
export LC_EXT=$LC_HOME/runtimeGOPath
export InternetAccess=false
export ServiceLocatorExternal=localhost:5408
export ExternalEndpointIP=localhost

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

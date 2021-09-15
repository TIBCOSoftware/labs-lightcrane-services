#!/bin/bash

export COMPOSE_HTTP_TIMEOUT=200
export USER_HOME=/Users/jumbo
export LC_HOME=./LC
export LC_EXT=$LC_HOME/runtimeGOPath
export ServiceLocatorExternal=192.168.1.152:5408
export ExternalEndpointIP=192.168.1.152

if [ ! -d "$LC_HOME" ] 
then
    mkdir $LC_HOME 
fi

docker-compose rm -f
docker-compose pull
docker-compose up -d --build

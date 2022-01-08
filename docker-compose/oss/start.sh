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

docker-compose rm -f
docker-compose up -d

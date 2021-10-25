#!/bin/bash

export COMPOSE_HTTP_TIMEOUT=200
export USER_HOME=~
export LC_HOME=./LC
export LC_EXT=$LC_HOME/runtimeGOPath
export ServiceLocatorExternal=localhost:5408
export ExternalEndpointIP=localhost

if [ ! -d "$LC_HOME" ] 
then
    mkdir $LC_HOME 
fi

docker-compose rm -f
docker-compose pull
docker-compose up -d --build

curl -d '{}' -H 'Content-Type: application/json' http://localhost:5408/f1/air/createDomain/Air-account_00001

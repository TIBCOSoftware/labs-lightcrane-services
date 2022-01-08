@echo off
set COMPOSE_HTTP_TIMEOUT=200
set USER_HOME=~
set LC_HOME=./LC
set LC_EXT=%LC_HOME%/runtimeGOPath
set InternetAccess=false
set ServiceLocatorExternal=localhost:5408
set ExternalEndpointIP=localhost
if not exist "%LC_HOME%" mkdir "%LC_HOME%"
docker-compose rm -f
docker-compose up -d
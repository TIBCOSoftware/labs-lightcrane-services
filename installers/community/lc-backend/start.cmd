@echo off
set VERSION=0.7.0
set COMPOSE_HTTP_TIMEOUT=200
set USER_HOME=~
set LC_HOME=./LC
set LC_EXT=%LC_HOME%/runtimeGOPath
set InternetAccess=false
set ServiceLocatorExternal=localhost:5408
set ExternalEndpointIP=localhost
if not exist "%LC_HOME%" mkdir "%LC_HOME%"

if "%1%"==""  set ARCH=""
if "%1%"=="amd64"  set ARCH=""
if "%1%"=="arm64"  set ARCH="-arm64"

docker-compose rm -f
docker-compose up -d
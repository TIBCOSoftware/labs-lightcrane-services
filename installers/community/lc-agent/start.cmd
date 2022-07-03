@echo off

if "%1%"==""  set ARCH=""
if "%1%"=="amd64"  set ARCH=""
if "%1%"=="arm64"  set ARCH="-arm64"

docker-compose rm -f
docker-compose up -d
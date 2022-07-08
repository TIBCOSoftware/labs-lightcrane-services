@echo off

set LC_HOME=./LC

if not exist "%LC_HOME%" mkdir "%LC_HOME%"

if "%1%"==""  set ARCH=""
if "%1%"=="amd64"  set ARCH=""
if "%1%"=="arm64"  set ARCH="-arm64"

docker-compose rm -f
docker-compose up -d
@echo off
set VERSION=0.7.0
set USER_HOME=~
set LC_HOME=./LC
set LC_EXT=%LC_HOME%/runtimeGOPath

if "%1%"==""  set ARCH=""
if "%1%"=="amd64"  set ARCH=""
if "%1%"=="arm64"  set ARCH="-arm64"

docker-compose down
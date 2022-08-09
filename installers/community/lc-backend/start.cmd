@echo off

set LC_HOME=./LC

if not exist "%LC_HOME%" mkdir "%LC_HOME%"

@REM if "%1%"==""  set ARCH=""
@REM if "%1%"=="amd64"  set ARCH=""
@REM if "%1%"=="arm64"  set ARCH="-arm64"

set ARCH=""

docker-compose rm -f
docker-compose up -d
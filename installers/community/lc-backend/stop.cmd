@echo off

set LC_HOME=./LC

set ARCH=

@REM if "%1%"==""  set ARCH=""
@REM if "%1%"=="amd64"  set ARCH=""
@REM if "%1%"=="arm64"  set ARCH="-arm64"

docker-compose down -v
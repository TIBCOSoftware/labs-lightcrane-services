#!/bin/bash

source tools.sh

ProjectID=$1 
Runner=$2 
ImageName=$3 
FlogoBuilder=$4 
Filename=$5 
ServiceName=$6 
AppBinFolder=$7
ExecutableName=$8
BuildImage=$9

echo "***** parameters *****"
echo " - ProjectID      : "$ProjectID
echo " - Runner         : "$Runner
echo " - ImageName      : "$ImageName
echo " - FlogoBuilder   : "$FlogoBuilder
echo " - Filename       : "$Filename
echo " - ServiceName    : "$ServiceName
echo " - AppBinFolder   : "$AppBinFolder
echo " - ExecutableName : "$ExecutableName
echo " - BuildImage     : "$BuildImage

echo " - Base folder    : $(pwd)"

echo "***** prepare container artifacts *****"
WorkFolder="/home/f1/projects/$ProjectID/build"
if [ -d "$WorkFolder" ] 
then
    rm -rf $WorkFolder 
fi
mkdir $WorkFolder

if [ "docker-compose.yml" != "$Runner" ]
then
	prepare_python_container $ProjectID $Runner $WorkFolder $ServiceName
fi

echo "***** call build_image *****"
if [ "no" == $BuildImage ]
then
	echo " - skip build_image"
elif [ "yes" == $BuildImage ]
then
	echo " - call build_image"
	build_image
else
	echo " - Unknown BuildImage flag = "$BuildImage
fi

echo "***** build process done *****"

#!/bin/bash
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

chmod +x ./python/app-type.py
AppType=$(python3 ./python/app-type.py $Filename)

echo " - AppType        : "$AppType

# prepare builder working folder
if [ -d "/home/f1/projects/$ProjectID/build" ] 
then
    rm -rf /home/f1/projects/$ProjectID/build 
fi
mkdir /home/f1/projects/$ProjectID/build
mkdir -p /home/f1/projects/$ProjectID/build/$ServiceName/server/data

# select runner
if [ $Platform == "linux/arm64" ]
then
	cp -R /home/runner/flogo/ubuntu/docker/* /home/f1/projects/$ProjectID/build/$ServiceName
else
	cp -R /home/runner/flogo/alpine/docker/* /home/f1/projects/$ProjectID/build/$ServiceName
fi

# prepare artifacts
if [ -d "/home/f1/projects/$ProjectID/artifacts" ] 
then
	cp -R /home/f1/projects/$ProjectID/artifacts /home/f1/projects/$ProjectID/build/$ServiceName/server
fi

# prepare metadata
if [ -d "/home/f1/projects/$ProjectID/metadata" ] 
then
	cp -R /home/f1/projects/$ProjectID/metadata/* /home/f1/projects/$ProjectID/build/$ServiceName/server/data
fi

echo "***** call build_executable *****"
if [ "flogo_fe" == $AppType ]
then
	echo " - will build FLOGO FE application"
	source ./flogo_fe.sh
elif [ "flogo_oss" == $AppType ]
then 
	echo " - will build FLOGO OSS application"
	source ./flogo_oss.sh 
fi

build_executable


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

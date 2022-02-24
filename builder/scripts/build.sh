#!/bin/bash

source tools.sh

# parameter : $ProjectID$ $Runner$ $ImageName$ $FlogoBuilder$ $Filename$ $ServiceName$ $AppBinFolder$ $ExecutableName$ $BuildImage$ $BuildApp$ $PrebldFlogoEngine$
echo "parameter : "${1} ${2} ${3} ${4} ${5} ${6} ${7} ${8} ${9} ${10} ${11}

ProjectID=${1} 
Runner=${2} 
ImageName=${3} 
FlogoBuilder=${4} 
Filename=${5} 
ServiceName=${6} 
AppBinFolder=${7}
ExecutableName=${8}
BuildImage=${9}
BuildApp=${10}
PrebuildFlogoEngine=${11}

echo "***** parameters *****"
echo " - ProjectID           : "$ProjectID
echo " - Runner              : "$Runner
echo " - ImageName           : "$ImageName
echo " - FlogoBuilder        : "$FlogoBuilder
echo " - Filename            : "$Filename
echo " - ServiceName         : "$ServiceName
echo " - AppBinFolder        : "$AppBinFolder
echo " - ExecutableName      : "$ExecutableName
echo " - BuildImage          : "$BuildImage
echo " - BuildApp            : "$BuildApp
echo " - PrebuildFlogoEngine : "$PrebuildFlogoEngine

echo "***** system environment *****"
echo " - Platform            : "$Platform

echo "***** system information *****"
echo " - Base folder    : $(pwd)"

echo "***** prepare container artifacts *****"

if [ "yes" == $PrebuildFlogoEngine ]
then
	echo " - prepare pipeline container : prebuild engine mode"
	prepare_pipeline_container_prebuild_engine $Filename $ProjectID $ServiceName $Platform $Runner
else
	echo " - prepare pipeline container"
	prepare_pipeline_container $Filename $ProjectID $ServiceName $Platform $Runner
fi

chmod +x ./python/app-type.py
AppType=$(python3 ./python/app-type.py $Filename)
echo " - AppType        : "$AppType

echo "***** call build_executable *****"

if [ "yes" == $BuildApp ] && [ "yes" != $PrebuildFlogoEngine ]
then
	if [ "flogo_fe" == "$AppType" ]
	then
		echo " - will build FLOGO FE application"
		build_executable_fe
	elif [ "flogo_oss" == "$AppType" ]
	then 
		echo " - will build FLOGO OSS application"
		build_executable_oss 
	fi
else
	echo " - skip build_executable"
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

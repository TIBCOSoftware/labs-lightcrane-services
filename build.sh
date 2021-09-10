#!/bin/bash

source setup

Service=$1

if [[ "oss" == "$2" ]]
then
	cd $Service/build
	build $Service
else
	cd $Service/build_fe
	build_fe $Service
fi

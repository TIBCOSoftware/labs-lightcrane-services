#!/bin/bash

source setup

cd $1

if [[ "oss" == "$2" ]] 
then
	./build-oss.sh
else
	./build-fe.sh
fi

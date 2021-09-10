#!/bin/bash

if [ ! -f "/home/ext/bin/flogo" ]
then
	go get -u github.com/project-flogo/cli/...
fi

./engine
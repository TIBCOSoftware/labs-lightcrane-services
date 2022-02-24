#!/bin/bash

if [ -f "./server.py" ]
then
	echo "***** launch python service *****"
	ls ./artifacts
	export System_Standalone=False
	python ./server.py &
fi

sleep 3

echo "***** launch pipeline *****"
./flogo-engine
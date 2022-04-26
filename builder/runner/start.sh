#!/bin/bash

echo "***** Check folder (ls .) *****"

ls -l

if [ -f "./server.py" ]
then
	echo "***** Check folder (ls ./artifacts) *****"
	ls ./artifacts
	echo "***** launch python service *****"
	export System_Standalone=False
	pip3 install --index-url https://test.pypi.org/simple/ --no-deps f1-steveny
	#nohup python -u ./server.py  > /tmp/files/log/python_engine.log &
	nohup python ./server.py &
fi

sleep 3

echo "***** export *****"

export

echo "***** Check folder (ls) *****"
ls -l
echo "***** launch pipeline *****"
./flogo-engine
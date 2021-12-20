#!/bin/bash
# ProjectID Runner Imagename

pwd

if [ -d "/home/f1/projects/$1/builder" ] 
then
    rm -rf /home/f1/projects/$1/builder 
fi

mkdir /home/f1/projects/$1/builder

if [ "flogo" = "$2" ] 
then
	cp -R /home/runner/$2/* /home/f1/projects/$1/builder
else
	OrigIFS=$IFS
	IFS='_'
	read -ra runner <<< "$2"	
	echo "Runner : "${runner[0]}-${runner[1]}
	IFS=$OrigIFS
	mkdir /home/f1/projects/$1/builder/docker
	cp -R /home/runner/${runner[0]}/server /home/f1/projects/$1/builder/docker
	cp -R /home/runner/${runner[0]}/docker/Dockerfile_${runner[1]} /home/f1/projects/$1/builder/docker/Dockerfile
	ls /home/f1/projects/$1/builder
fi

if [ -f "/home/f1/projects/$1/artifacts/requirements.txt" ] 
then
    cp /home/f1/projects/$1/artifacts/requirements.txt /home/f1/projects/$1/builder/docker/server
fi

if [ -f "/home/f1/projects/$1/artifacts/Dockerfile" ] 
then
    cp /home/f1/projects/$1/artifacts/Dockerfile /home/f1/projects/$1/builder/docker
fi

cp -R /home/f1/projects/$1/artifacts /home/f1/projects/$1/builder/docker/server
mkdir /home/f1/projects/$1/builder/docker/server/data
cp -R /home/f1/projects/$1/metadata/* /home/f1/projects/$1/builder/docker/server/data

cd /home/f1/projects/$1/builder/docker
docker build --tag $3 .
docker push $3

rm -rf /home/f1/projects/$1/builder

exit 0
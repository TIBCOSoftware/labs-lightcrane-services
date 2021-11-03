#!/bin/bash

Namespace=$1
DeployType=$2
ServiceName=$3
ProjectID=$4

echo $Namespace
echo $DeployType
echo $ServiceName
echo $ProjectID
echo $Username
echo $TargetServer

if [ -d "/home/f1/projects/$ProjectID/deploy/$DeployType" ] 
then
    rm -rf /home/f1/projects/$ProjectID/deploy/$DeployType/* 
else
    mkdir -p /home/f1/projects/$ProjectID/deploy/$DeployType
fi

cp -R /home/f1/projects/$ProjectID/artifacts/* /home/f1/projects/$ProjectID/deploy/$DeployType

cd /home/f1/projects/$ProjectID/deploy/$DeployType

if [[ -v TargetServer ]]
then
	echo "Deploy Remotely !!"
	#docker-compose --context $DockerContext up -d
	docker-compose -H "ssh://$Username@$TargetServer" up --build
else
	echo "Deploy Locally !!"
	docker-compose up
fi

exit 0
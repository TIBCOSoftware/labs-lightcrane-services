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
echo $Descriptor

if [ -d "./$DeployType" ] 
then
    rm -rf ./$DeployType/* 
else
    mkdir -p ./$DeployType
fi

cp -R ../artifacts/$Descriptor ./$DeployType

cd ./$DeployType


if [ "k8s" == $DeployType ]; then
	echo "Deploy to $DeployType"
	kubectl apply -f .
else
	if [[ -v TargetServer ]]; then
		echo "Deploy Remotely !!"
		#docker-compose --context $DockerContext up -d
		docker-compose -H "ssh://$Username@$TargetServer" up --build
	else
		echo "Deploy Locally !!"
		docker-compose up
	fi
fi
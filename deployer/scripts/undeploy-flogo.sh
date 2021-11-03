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

cd ./$DeployType


if [ "k8s" == $DeployType ]; then
	echo "Undeploy from $DeployType"
	kubectl delete -f ./$Descriptor
else
	if [[ -v TargetServer ]]
	then
		echo "Undeploy Remotely !!"
		docker-compose -H "ssh://$Username@$TargetServer" down
	else
		echo "Undeploy Locally !!"
		docker-compose down
	fi
fi


cd ../
if [ -d "./$DeployType" ] 
then
    rm -rf ./$DeployType
fi
#!/bin/bash

Namespace=$1
DeployType=$2
ServiceName=$3
ProjectID=$4
echo $Username
echo $TargetServer

cd /home/f1/projects/$ProjectID/deploy/$DeployType
if [[ -v TargetServer ]]
then
	#docker-compose --context $DockerContext down -d
	docker-compose -H "ssh://$Username@$TargetServer" down
else
	docker-compose down
f1

exit 0
#!/bin/bash

Namespace=$1
DeployType=$2
ServiceName=$3
ProjectID=$4
InstanceID=$5
export ServiceVersion=0.1.0
export Arch="$(cut -d'/' -f2 <<<"$Platform")"

echo "Namespace    = $Namespace"
echo "DeployType   = $DeployType"
echo "ServiceName  = $ServiceName"
echo "ProjectID    = $ProjectID"
echo "InstanceID   = $InstanceID"
echo "Username     = $Username"
echo "TargetServer = $TargetServer"
echo "Port         = $Port"

echo "Working folder $(pwd)"

if [ "docker" == $DeployType ]
then
	echo "source ./docker-compose.sh"
	source ./docker-compose.sh
elif [ "docker-oh" == $DeployType ]
then 
	echo "source ./open-horizon.sh"
	source ./open-horizon.sh
elif [ "k8s" == $DeployType ]
then
	echo "source ./k8s.sh"
	source ./k8s.sh
fi

undeploy
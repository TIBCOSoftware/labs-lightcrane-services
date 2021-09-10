#!/bin/bash -e

for i in "$@"
do
case $i in
    -u=*|--user=*)
    USER="${i#*=}"
    shift # past argument=value
    ;;
    -o=*|--org=*)
    ORG="${i#*=}"
    shift # past argument=value
    ;;
    -p=*|--platform=*)
    PLATFORM="${i#*=}"
    shift # past argument=value
    ;;
    -v=*|--version=*)
    VERSION="${i#*=}"
    shift # past argument=value
    ;;
    -n=*|--name=*)
    NAME="${i#*=}"
    shift # past argument=value
    ;;
    -d=*|--directory=*)
    DIRECTORY="${i#*=}"
    shift # past argument=value
    ;;
    -s=*|--server=*)
    SERVER="${i#*=}"
    shift # past argument=value
    ;;
    --default)
    DEFAULT=YES
    shift # past argument with no value
    ;;
    *)
          # unknown option
    ;;
esac
done
CMD=$1

echo "PLATFORM  = ${PLATFORM}"
echo "VERSION   = ${VERSION}"
echo "NAME      = ${NAME}"
echo "DIRECTORY = ${DIRECTORY}"
echo "SERVER    = ${SERVER}"
echo "CMD       = ${CMD}"

if ! snapctl is-connected docker-executables; then
	snap connect lightcrane-air-agent:docker-executables docker:docker-executables
fi

if ! snapctl is-connected docker; then
	snap connect lightcrane-air-agent:docker docker:docker-daemon
fi

if [[ -z $PLATFORM ]]
then
	IMAGE=bigoyang/agent:${VERSION}
else
	IMAGE=bigoyang/agent_${PLATFORM}:${VERSION}
fi

if [[ -z $CMD ]]
then
	echo "Starting agent : ${IMAGE} "
	$SNAP/docker-snap/bin/docker run \
	    --privileged \
	    --name agent \
	    -d \
	    --network-alias light_crane \
	    -e DOCKER_TLS_CERTDIR=/certs \
	    -e FLOGO_APP_PROPS_ENV=auto \
	    -e System_BaseFolder=/home/f1 \
	    -e System_BaseFolderExt=${DIRECTORY} \
	    -e System_ServiceLocatorIP=${SERVER} \
	    -e System_ServiceLocatorPort=5408 \
	    -e System_Network=light_crane \
	    -e System_AgentID=${NAME} \
	    -e System_Organization=${ORG} \
	    -e System_User=${USER} \
	    -v /var/run/docker.sock:/var/run/docker.sock \
	    -v ${DIRECTORY}:/home/f1 \
	    ${IMAGE}
else
	if [ "stop" == $CMD ]
	then
		echo "Stoping agent : ${IMAGE} "
		$SNAP/docker-snap/bin/docker rm \
			$($SNAP/docker-snap/bin/docker stop \
				$($SNAP/docker-snap/bin/docker ps \
					-a -q \
					--filter name=agent&&ancestor=${IMAGE}))
	fi
fi

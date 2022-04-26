#!/bin/bash

function prepare_pipeline_container_prebuild_engine {
	
	# parameters : Filename AppType ProjectID ServiceName Platform
	echo "parameter : "${1} ${2} ${3} ${4} ${5}
	Filename=${1}
	ProjectID=${2}
	ServiceName=${3}
	Platform=${4}
	Runner=${5}

	# prepare builder working folder
	echo " - prepare builder working folder"
	mkdir -p /home/f1/projects/$ProjectID/build/$ServiceName/server/data
	
	# prepare artifacts
	if [ -d "/home/f1/projects/$ProjectID/artifacts" ] 
	then
		echo " - move artifacts to server"
		cp -R /home/f1/projects/$ProjectID/artifacts/$ServiceName/* /home/f1/projects/$ProjectID/build/$ServiceName/server
		cp /home/f1/projects/$ProjectID/artifacts/$ServiceName.json /home/f1/projects/$ProjectID/build/$ServiceName/server
		cp /home/f1/projects/$ProjectID/artifacts/docker-compose.yml /home/f1/projects/$ProjectID/build/$ServiceName/server
		cp /home/runner/start.sh /home/f1/projects/$ProjectID/build/$ServiceName
	fi
	
	echo " - platform : $Platform"
	if [ "$Platform" == "linux/arm64" ]
	then
		if [ "\$Runner\$" == ${Runner} ];
		then
    		echo "Ues default runner for $Platform"
			cp /home/runner/flogo/ubuntu/docker/Dockerfile /home/f1/projects/$ProjectID/build/$ServiceName
		else
    		echo "No runner available ..."
		fi
	else
		if [ "\$Runner\$" == ${Runner} ];
		then
    		echo "Ues default runner for $Platform"
			cp /home/runner/flogo/alpine/docker/Dockerfile /home/f1/projects/$ProjectID/build/$ServiceName
		else
    		echo "Use runner : Dockerfile_${Runner}"
			cp /home/runner/flogo/alpine/docker/Dockerfile_${Runner} /home/f1/projects/$ProjectID/build/$ServiceName/Dockerfile
			cp /home/runner/flogo/${Runner}/* /home/f1/projects/$ProjectID/build/$ServiceName
		fi
	fi
	
	# setup flogo engine executable & pipeline descriptor
	echo " - prepare flogo pipeline descriptor"
	cp /home/f1/projects/$ProjectID/artifacts/$ServiceName.json /home/f1/projects/$ProjectID/build/$ServiceName/flogo.json
	Arch="$(cut -d'/' -f2 <<<"$Platform")"
	echo " - prepare flogo engine executable for "$Arch
	cp /home/f1/projects/$ProjectID/build/generic/app_$Arch /home/f1/projects/$ProjectID/build/$ServiceName/flogo-engine

	# prepare metadata
	if [ -d "/home/f1/projects/$ProjectID/metadata" ] 
	then
		echo " - prepare metadata"
		cp -R /home/f1/projects/$ProjectID/metadata/* /home/f1/projects/$ProjectID/build/$ServiceName/server/data
	fi
	
}

function prepare_pipeline_container {

	# parameters : Filename AppType ProjectID ServiceName Platform
	echo "parameter : "${1} ${2} ${3} ${4} ${5}
	Filename=${1}
	ProjectID=${2}
	ServiceName=${3}
	Platform=${4}
	Runner=${5}

	# prepare builder working folder
	echo " - prepare builder working folder"
	if [ -d "/home/f1/projects/$ProjectID/build" ] 
	then
	    rm -rf /home/f1/projects/$ProjectID/build 
	fi
	mkdir /home/f1/projects/$ProjectID/build
	mkdir -p /home/f1/projects/$ProjectID/build/$ServiceName/server/data
#	mkdir -p /home/f1/projects/$ProjectID/build/$ServiceName/server/artifacts

	# prepare artifacts
#	if [ -d "/home/f1/projects/$ProjectID/artifacts" ] 
#	then
		echo " - prepare artifacts"
		cp -R /home/f1/projects/$ProjectID/artifacts/$ServiceName/* /home/f1/projects/$ProjectID/build/$ServiceName/server
#		cp -R /home/f1/projects/$ProjectID/artifacts/$ServiceName/* /home/f1/projects/$ProjectID/build/$ServiceName/server/artifacts
		cp /home/f1/projects/$ProjectID/artifacts/$ServiceName.json /home/f1/projects/$ProjectID/build/$ServiceName/server
		cp /home/f1/projects/$ProjectID/artifacts/docker-compose.yml /home/f1/projects/$ProjectID/build/$ServiceName/server
		cp /home/runner/start.sh /home/f1/projects/$ProjectID/build/$ServiceName
#	fi

	# select runner
#	if [ -f "/home/f1/projects/$ProjectID/build/$ServiceName/server/artifacts/Dockerfile" ] 
#	then
		# Overwrite runner's Dockerfile 
#		echo " - overwrite runner's Dockerfile"
#	    cp /home/f1/projects/$ProjectID/build/$ServiceName/server/artifacts/Dockerfile /home/f1/projects/$ProjectID/build/$ServiceName
#	else
		echo " - platform : $Platform"
		if [ "$Platform" == "linux/arm64" ]
		then
#			if [ -z ${Runner+x} ];
			if [ "\$Runner\$" == ${Runner} ];
			then
				cp /home/runner/flogo/ubuntu/docker/Dockerfile /home/f1/projects/$ProjectID/build/$ServiceName
			else
    			echo "No runner available ..."
			fi
		else
#			if [ -z ${Runner+x} ];
			if [ "\$Runner\$" == ${Runner} ];
			then
				cp /home/runner/flogo/alpine/docker/Dockerfile /home/f1/projects/$ProjectID/build/$ServiceName
			else
				cp /home/runner/flogo/alpine/docker/Dockerfile_${Runner} /home/f1/projects/$ProjectID/build/$ServiceName/Dockerfile
				cp /home/runner/flogo/${Runner}/* /home/f1/projects/$ProjectID/build/$ServiceName
			fi
		fi
#	fi

	# prepare metadata
	if [ -d "/home/f1/projects/$ProjectID/metadata" ] 
	then
		echo " - prepare metadata"
		cp -R /home/f1/projects/$ProjectID/metadata/* /home/f1/projects/$ProjectID/build/$ServiceName/server/data
	fi
}

function prepare_python_container {
	
	# parameters : ProjectID Runner WorkFolder ServiceName
	echo "parameter : "${1} ${2} ${3} ${4}
	ProjectID=${1}
	Runner=${2} 
	WorkFolder=${3} 
	ServiceName=${4} 

	pwd

	# Clean up then Create working folder
	if [ -d "$WorkFolder" ] 
	then
    	rm -rf $WorkFolder 
	fi
	mkdir $WorkFolder

	# Prepare runner
	OrigIFS=$IFS
	IFS='_'
	read -ra runner <<< "$Runner"	
	echo "Runner : "${runner[0]}-${runner[1]}
	IFS=$OrigIFS
	mkdir $WorkFolder/$ServiceName
	cp -R /home/runner/${runner[0]}/server $WorkFolder/$ServiceName
	cp -R /home/runner/${runner[0]}/docker/Dockerfile_${runner[1]} $WorkFolder/$ServiceName/Dockerfile
	ls $WorkFolder

	# Add requirements
	if [ -f "/home/f1/projects/$ProjectID/artifacts/requirements.txt" ] 
	then
    	cp /home/f1/projects/$ProjectID/artifacts/requirements.txt $WorkFolder/$ServiceName/server
	fi

	# Overwrite runner's Dockerfile 
	if [ -f "/home/f1/projects/$ProjectID/artifacts/Dockerfile" ] 
	then
    	cp /home/f1/projects/$ProjectID/artifacts/Dockerfile $WorkFolder/$ServiceName
	fi

	# Prepare artifacts
	cp -R /home/f1/projects/$ProjectID/artifacts $WorkFolder/$ServiceName/server
	mkdir $WorkFolder/$ServiceName/server/data
	cp -R /home/f1/projects/$ProjectID/metadata/* $WorkFolder/$ServiceName/server/data

}

function build_executable_oss {
	cd /home/f1/projects/$ProjectID/build/$ServiceName
	echo " - Working folder : $(pwd)"

	if [[ $Platform ]]
	then
		echo "***** (Flogo OSS) Build Executable for $Platform *****"
		IFS='/'
		read -ra arr <<< "$Platform"
		IFS=' '
		export GOOS=${arr[0]} 
		export GOARCH=${arr[1]}

	else
		echo "***** (Flogo OSS) Build Executable for local platform *****"
	fi

	echo " - GOOS           : $GOOS"
	echo " - GOARCH         : $GOARCH"
		
	flogo create -f $Filename app
	cd ./app
	flogo build -e --verbose
	cp ./bin/app ../flogo-engine
	cd ../

#	exit 0
}

function build_executable_fe {
	cd /home/f1/projects/$ProjectID/build/$ServiceName
	echo " - Working folder : $(pwd)"

	if [[ $Platform ]]
	then
		echo "***** (Flogo FE) Build Executable for $Platform *****"
		/usr/flogo/home/bin/$FlogoBuilder build -p $Platform -f $Filename -n $ExecutableName -o $AppBinFolder
	else
		echo "***** (Flogo FE) Build Executable for local platform *****"
		/usr/flogo/home/bin/$FlogoBuilder build -f $Filename -n $ExecutableName -o $AppBinFolder
	fi

#	exit 0
}

function build_image {
	cd /home/f1/projects/$ProjectID/build/$ServiceName
	echo " - Working folder : $(pwd)"

	if [[ $Platform ]]
	then
		echo "***** Build docker image for $Platform *****"
#		sudo docker buildx build --platform $Platform --tag $ImageName --build-arg BASE_IMAGE="{TOKEN}" .
		sudo docker buildx build --platform $Platform --tag $ImageName .
	else
		echo "***** Build docker image for local platform *****"
		sudo docker buildx build --tag $ImageName .
	fi

	if [[ -v DockerUser ]]
	then
		echo "***** Push Image to Dockerhub *****"
		docker login --username=$DockerUser --password=$DockerPassword
		docker push $ImageName
	fi

#	exit 0	
}
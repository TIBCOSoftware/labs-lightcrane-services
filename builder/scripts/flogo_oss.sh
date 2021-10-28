# ready to build
cd /home/f1/projects/$ProjectID/build/$ServiceName
echo " - Working folder : $(pwd)"

function build_executable {

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

function build_image {
	if [[ $Platform ]]
	then
		echo "***** Build docker image for $Platform *****"
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
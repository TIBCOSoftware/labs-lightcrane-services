# ready to build
cd /home/f1/projects/$ProjectID/build/$ServiceName
echo "Working folder $(pwd)"

function build_no_image {

	if [[ $Platform ]]
	then
		echo "Build Executable for $Platform !!"
		/usr/flogo/home/bin/$FlogoBuilder build -p $Platform -f $Filename -n $ExecutableName -o $AppBinFolder
	else
		echo "Build Executable for local platform !!"
		/usr/flogo/home/bin/$FlogoBuilder build -f $Filename -n $ExecutableName -o $AppBinFolder
	fi

	exit 0
}

function build_with_image {
	if [[ $Platform ]]
	then
		echo "Build Executable for $Platform !!"
		/usr/flogo/home/bin/$FlogoBuilder build -p $Platform -f $Filename -n $ExecutableName -o $AppBinFolder
		docker buildx build --platform $Platform --tag $ImageName .
	else
		echo "Build Executable for local platform !!"
		/usr/flogo/home/bin/$FlogoBuilder build -f $Filename -n $ExecutableName -o $AppBinFolder
		docker buildx build --tag $ImageName .
	fi

	if [[ -v DockerUser ]]
	then
		echo "Will Push Image to Dockerhub !!"
		docker login --username=$DockerUser --password=$DockerPassword
		docker push $ImageName
	fi

	#rm -rf /home/f1/projects/$ProjectID/builder

	exit 0	
}
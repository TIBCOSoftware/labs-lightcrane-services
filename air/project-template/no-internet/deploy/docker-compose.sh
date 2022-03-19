function deploy {
	echo "***** reset deployment folder *****"
	rm -rf ./$DeployType/$ServiceName/$InstanceID
	mkdir -p ./$DeployType/$ServiceName/$InstanceID

	echo "***** prepare flogo pipeline artifacts *****"
	echo " - prepare docker-compose descriptor"
	cp $Descriptor ./$DeployType/$ServiceName
	echo " - prepare folder for deployment"
	cp -R ../build/$ServiceName/* ./$DeployType/$ServiceName/$InstanceID

	echo "***** ready for deploying pipeline *****"
	cd ./$DeployType/$ServiceName
	echo " - Deploy from $(pwd)"

	Mode=""
	if [ "y"==$DetachedMode ]
	then
		Mode="-d"
	fi

	export ConcreteApp=$ServiceName".json"
	echo " - ConcreteApp = "$ConcreteApp

	if [ "k8s" == $DeployType ]; then
		kubectl apply -f .
	else
		if [[ -n $TargetServer ]]; then
			echo "Deploy Remotely !!"
			if [[ -n $Port ]]; then
				docker-compose -H "ssh://$Username@$TargetServer:$Port" rm -f
				#docker-compose --context $DockerContext up -d
				docker-compose -H "ssh://$Username@$TargetServer:$Port" up $Mode --build
			else
				docker-compose -H "ssh://$Username@$TargetServer" rm -f
				#docker-compose --context $DockerContext up -d
				docker-compose -H "ssh://$Username@$TargetServer" up $Mode --build
			fi
		else
			echo "Deploy Locally !!"
			docker-compose rm -f
			docker-compose up $Mode --build
		fi
	fi
}

function undeploy {
	cd ./$DeployType/$ServiceName
	echo "Undeploy from $(pwd)"

	if [[ -n $TargetServer ]]; then
		echo "Undeploy Remotely !!"
		if [[ -n $Port ]]; then
			docker-compose -H "ssh://$Username@$TargetServer:$Port" down -v
			docker-compose -H "ssh://$Username@$TargetServer:$Port" rm -f
		else
			docker-compose -H "ssh://$Username@$TargetServer" down -v
			docker-compose -H "ssh://$Username@$TargetServer" rm -f
		fi
	else
		echo "Undeploy Locally !!"
		docker-compose down -v
		docker-compose rm -f
	fi
	
	cd ..
	rm -rf ./$ServiceName
	echo "Undeploy and cleanup for $ServiceName"
}
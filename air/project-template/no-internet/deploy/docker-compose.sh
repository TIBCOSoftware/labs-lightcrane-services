function deploy {
	Arch="$(cut -d'/' -f2 <<<"$Platform")"

	rm -rf ./$DeployType/$ServiceName/$InstanceID
	mkdir -p ./$DeployType/$ServiceName/$InstanceID

	cp -R $Descriptor ./$DeployType/$ServiceName
	cp ../artifacts/$ServiceName.json ./$DeployType/$ServiceName/$InstanceID/flogo.json
	cp -R ../build/generic/* ./$DeployType/$ServiceName/$InstanceID
	mv ./$DeployType/$ServiceName/$InstanceID/flogo-engine_$Arch ./$DeployType/$ServiceName/$InstanceID/flogo-engine

	cd ./$DeployType/$ServiceName
	echo "Deploy from $(pwd)"

	Mode=""
	if [ "y"==$DetachedMode ]
	then
		Mode="-d"
	fi

	export ConcreteApp=$ServiceName".json"
	echo "ConcreteApp = "$ConcreteApp

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
			docker-compose -H "ssh://$Username@$TargetServer:$Port" down
			docker-compose -H "ssh://$Username@$TargetServer:$Port" rm -f
		else
			docker-compose -H "ssh://$Username@$TargetServer" down
			docker-compose -H "ssh://$Username@$TargetServer" rm -f
		fi
	else
		echo "Undeploy Locally !!"
		docker-compose down
		docker-compose rm -f
	fi
	
	cd ..
	rm -rf ./$ServiceName
	echo "Undeploy and cleanup for $ServiceName"
}
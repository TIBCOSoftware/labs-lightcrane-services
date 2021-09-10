function deploy {
	rm -rf ./$DeployType/$ServiceName/$InstanceID
	mkdir -p ./$DeployType/$ServiceName/$InstanceID

	cp -R $Descriptor ./$DeployType/$ServiceName
	cp -R ../build/$ServiceName/* ./$DeployType/$ServiceName/$InstanceID

	cd ./$DeployType/$ServiceName
	echo "Deploy from $(pwd)"

	Mode=""
	if [ "y"==$DetachedMode ]
	then
		Mode="-d"
	fi

	kubectl apply -f .
}

function undeploy {
	cd ./$DeployType/$ServiceName
	echo "Undeploy from $(pwd)"

	kubectl delete -f ./$Descriptor
}
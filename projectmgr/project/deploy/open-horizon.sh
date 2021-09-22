export HZN_EXCHANGE_URL=$TargetServer
export HZN_ORG_ID="$(cut -d'@' -f2 <<<"$Username")"
export HZN_EXCHANGE_USER_AUTH="$(cut -d'@' -f1 <<<"$Username")"

function deploy {
	DeployBaseFolder=$(pwd)
	ProjectBaseFolder=${DeployBaseFolder%/*}
	WorkingFolder=$DeployBaseFolder"/"$DeployType
	ArtifactsFolder=$ProjectBaseFolder"/artifacts/docker-compose"
	PluginsFolder=$ProjectBaseFolder"/artifacts/plugins/python"

	echo "BaseFolder="$DeployBaseFolder
	echo "WorkingFolder="$WorkingFolder
	echo "ArtifactsFolder="$ArtifactsFolder
	echo "PluginsFolder="$PluginsFolder

	echo "Deploy Through Open Horizon Exchange !!"
	echo "HZN_EXCHANGE_URL="$HZN_EXCHANGE_URL
	echo "HZN_ORG_ID="$HZN_ORG_ID
	echo "HZN_EXCHANGE_USER_AUTH="$HZN_EXCHANGE_USER_AUTH

	# Prepare deploy folder
	mkdir -p $WorkingFolder
	
	cd $PluginsFolder

	chmod +x build-descriptor.py
	python3 ./build-descriptor.py "open-horizon.transform" $ArtifactsFolder $WorkingFolder $ServiceName

	echo "Deploy "$ServiceName

	ServiceNameArray=$(echo $ServiceName | tr "|" "\n")
	for ServiceNameElement in $ServiceNameArray
	do
		echo "Deploy "$ServiceNameElement" Through OpenHorizon Exchange !!"
		cd $WorkingFolder/$ServiceNameElement
		pwd
		ls
		# Publish the service and pattern
		hzn exchange service publish -P -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH -f service.json
		#hzn exchange pattern publish -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH -f pattern.json

		hzn exchange service list -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH
	done
	
	# Publish the deployment policy
	hzn exchange deployment addpolicy -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH -f deployment.policy.json "policy-"$ServiceName"_"$ServiceVersion
	hzn exchange deployment listpolicy -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH
}

function undeploy {
	
	echo "Uneploy "$ServiceName
	
	echo "Uneploy policy for "$ServiceName" Through OpenHorizon Exchange !!"
	hzn exchange deployment removepolicy  -v -f -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH "policy-"$ServiceName"_"$ServiceVersion
	hzn exchange deployment listpolicy -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH
	
	ServiceNameArray=$(echo $ServiceName | tr "|" "\n")
	for ServiceNameElement in $ServiceNameArray
	do
		echo "Uneploy "$ServiceNameElement" Through OpenHorizon Exchange !!"
		hzn exchange service remove -v -f -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH $ServiceNameElement"_"$ServiceVersion"_"$Arch
		hzn exchange service list -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH
	done
	
}

function get_services {
#	IFS='|' read -a ServiceNameArray <<< $ServiceName
#	ServiceNameArray=(`echo $ServiceName | tr '|' ' '`)

#	for i in "${ServiceNameArray[@]}"
#	do
#		echo $i
#	done
}
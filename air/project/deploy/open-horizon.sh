export HZN_EXCHANGE_USER_AUTH="$(cut -d'@' -f1 <<<"$Username")"
export HZN_ORG_ID="$(cut -d'@' -f2 <<<"$Username")"
export HZN_EXCHANGE_URL=$TargetServer

function deploy {
	echo "Deploy from $(pwd)"
	echo "Deploy Through Open Horizon Exchange !!"

	rm -rf ./$DeployType/$ServiceName
	mkdir -p ./$DeployType/$ServiceName
	
	cd ../artifacts/python

	chmod +x build-descriptor.py
	python3 ./build-descriptor.py "open-horizon.transform" $DeployType $ServiceName

	cd ../../deploy/$DeployType/$ServiceName

	#hzn voucher import voucher.json --policy node.policy.json -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH

	hzn exchange service publish -O -P --public=true -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH -f service.json
	#hzn exchange service addpolicy -f horizon/service.policy.json $(HZN_ORG_ID)/$(ServiceName)_$(ServiceVersion)_$(Arch)
	hzn exchange deployment addpolicy --json-file deployment.policy.json -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH "policy-"$ServiceName"_"$ServiceVersion
	
	hzn exchange service list -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH
	hzn exchange deployment listpolicy -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH
}

function undeploy {
	echo "Undeploy from $(pwd)"
	echo "Undeploy Through Open Horizon Exchange !!"
	
	hzn exchange deployment removepolicy  -v -f -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH "policy-"$ServiceName"_"$ServiceVersion
	hzn exchange service remove -v -f -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH $ServiceName"_"$ServiceVersion"_"$Arch
	
	hzn exchange service list -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH
	hzn exchange deployment listpolicy -o $HZN_ORG_ID -u $HZN_EXCHANGE_USER_AUTH
}
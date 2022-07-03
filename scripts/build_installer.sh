#!/bin/bash

network_type=${1:?}
os_type=${2:?}
arch_type=${3:?}
release_version=${4:?}

build_backed_offline(){
  # Offline backend artifacts
    pushd installers/community/lc-backend || exit 1
    ./export.sh ${arch_type} || exit 1
    popd > /dev/null || exit 1
}

build_agent_offline(){
  # Offline agent artifacts
  pushd installers/community/lc-agent || exit 1
  ./export.sh ${arch_type} || exit 1
  popd > /dev/null || exit 1
}

replace_release_version(){
  # Replace release version
  sed -i  "s/LABS_AIR_VERSION=GENERATED_VERSION/LABS_AIR_VERSION=${release_version}/" "${installer_target_path}/lc-backend/.env"
  sed -i  "s/LABS_AIR_VERSION=GENERATED_VERSION/LABS_AIR_VERSION=${release_version}/" "${installer_target_path}/lc-agent/.env"
}

installer_target_path="dist"

if [ -d $installer_target_path ]; then
  rm -rf $installer_target_path
fi
mkdir -p $installer_target_path

# Offline artifacts
if [[ "${network_type}" == "offline" ]];
then
  build_backed_offline
  build_agent_offline
fi

cp -r ./installers/community/. $installer_target_path

replace_release_version

cp scripts/start.sh ${installer_target_path} || exit 1
cp scripts/stop.sh ${installer_target_path} || exit 1
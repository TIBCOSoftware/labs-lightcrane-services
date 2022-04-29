#!/bin/bash

network_type=${1:?}
os_type=${2:?}
arch_type=${3:?}

build_backed_offline(){
  # Offline backend artifacts
    pushd installers/community/${arch_type}/lc-backend || exit 1
    ./export.sh || exit 1
    popd > /dev/null || exit 1
}

build_agent_offline(){
  # Offline agent artifacts
  pushd installers/community/${arch_type}/lc-agent || exit 1
  ./export.sh || exit 1
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

cp -r ./installers/community/${arch_type} $installer_target_path

cp scripts/start.sh ${installer_target_path} || exit 1
cp scripts/stop.sh ${installer_target_path} || exit 1
#!/bin/bash

network_type=${1:?}
os_type=${2:?}
arch_type=${3:?}

build_backed_offline(){
  # Offline backend artifacts
  if [[ "${arch_type}" == "amd64" ]]; then
      pushd installers/community/amd64/backend || exit 1
      ./export.sh || exit 1
      popd > /dev/null || exit 1
  elif [[ "${arch_type}" == "arm64" ]]; then
      # No arm64 artifacts yet
  fi
}

build_agent_offline(){
  # Offline agent artifacts
  if [[ "${arch_type}" == "amd64" ]]; then
      # No amd64 artifacts yet
  elif [[ "${arch_type}" == "arm64" ]]; then
      pushd installers/community/arm64/agent || exit 1
      ./export.sh || exit 1
      popd > /dev/null || exit 1
  fi
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

if [[ "${arch_type}" == "amd64" ]]; then
    cp -r "installers/community/amd64" $installer_target_path
elif [[ "${arch_type}" == "arm64" ]]; then
    cp -r "installers/community/arm64" $installer_target_path
fi
cp scripts/start_backend.sh ${installer_target_path} || exit 1
cp scripts/stop_backend.sh ${installer_target_path} || exit 1
cp scripts/start_agent.sh ${installer_target_path} || exit 1
cp scripts/stop_agent.sh ${installer_target_path} || exit 1
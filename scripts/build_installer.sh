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
      pushd installers/community/arm64/backend || exit 1
      ./export.sh || exit 1
      popd > /dev/null || exit 1
  fi
}

build_agent_offline(){
  # Offline agent artifacts
  if [[ "${arch_type}" == "amd64" ]]; then
      pushd installers/community/amd64/agent || exit 1
      ./export.sh || exit 1
      popd > /dev/null || exit 1
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
cp scripts/start.sh ${installer_target_path} || exit 1
cp scripts/stop.sh ${installer_target_path} || exit 1
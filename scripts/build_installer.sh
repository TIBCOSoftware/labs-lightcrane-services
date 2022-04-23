#!/bin/bash

network_type=${1:online}

installer_target_path="dist"

if [ -d $installer_target_path ]; then
  rm -rf $installer_target_path
fi
mkdir -p $installer_target_path

# Offline artifacts
if [[ "${network_type}" == "offline" ]];
then
  pushd docker-compose/oss || exit 1
  ./export.sh || exit 2
  popd > /dev/null || exit 1
fi

cp -r "docker-compose/oss" $installer_target_path
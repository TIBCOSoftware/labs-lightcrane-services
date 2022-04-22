#!/bin/bash

installer_type=${1:?}

installer_target_path="dist"

if [ -d $installer_target_path ]; then
  rm -rf $installer_target_path
fi
mkdir -p $installer_target_path

# Offline artifacts
if [[ "${installer_type}" == "offline" ]];
then
  pushd docker-compose/oss || exit 1
  ./export.sh || exit 1
  popd > /dev/null || exit 1
fi

cp -r "docker-compose/oss" $installer_target_path
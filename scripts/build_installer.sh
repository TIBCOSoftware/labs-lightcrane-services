#!/bin/bash

installer_target_path="dist"

if [ -d $installer_target_path ]; then
  rm -rf $installer_target_path
fi
mkdir -p $installer_target_path

cp -r "docker-compose/oss" $installer_target_path
#!/bin/bash

mkdir -p $ProjectsFolder'/'$ProjectID

cp -R $ProjectTemplateFolder/* $ProjectsFolder'/'$ProjectID

echo '{}'
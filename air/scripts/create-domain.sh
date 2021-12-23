#!/bin/bash

ProjectTemplateFolder=$1
ProjectID=$2
ProjectsFolder=$3
InternetAccess=$4

echo "ProjectTemplateFolder = "$ProjectTemplateFolder
echo "ProjectID             = "$ProjectID
echo "ProjectsFolder        = "$ProjectsFolder
echo "InternetAccess        = "$InternetAccess

if [ -d $ProjectsFolder/$ProjectID ] 
then
    rm -rf $ProjectsFolder/$ProjectID/* 
else
    mkdir -p $ProjectsFolder/$ProjectID
fi

if [ "true" == "$InternetAccess" ]
then
	echo "Has internet access"
	Templatefolder="internet"
else
	echo "No internet access"
	Templatefolder="no-internet"
fi

echo cp -R $ProjectTemplateFolder/$Templatefolder/* $ProjectsFolder/$ProjectID
cp -R $ProjectTemplateFolder/$Templatefolder/* $ProjectsFolder/$ProjectID


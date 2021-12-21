#!/bin/bash

UploadFolder=$1
CompressedFile=$2
ProjectBaseFolder=$3

echo "UploadFolder      = "$UploadFolder
echo "CompressedFile    = "$CompressedFile
echo "ProjectBaseFolder = "$ProjectBaseFolder

if [ -d "$ProjectBaseFolder/$CompressedFile" ] 
then
    rm -rf $ProjectBaseFolder/$CompressedFile/* 
else
    mkdir -p $ProjectBaseFolder/$CompressedFile
fi

unzip -o $UploadFolder$CompressedFile -d $ProjectBaseFolder

rm -rf $ProjectBaseFolder/_*

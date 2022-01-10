#!/bin/bash

if [ -d "./$2" ] 
then
    rm -rf ./$2/* 
else
    mkdir ./$2
fi

cp -R ../artifacts/docker-compose.yml ./$2

cd ./$2
docker-compose up
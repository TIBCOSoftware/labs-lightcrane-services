#!/bin/bash

echo $1
echo $2
echo $3

cd ./$2
docker-compose down

cd ../
if [ -d "./$2" ] 
then
    rm -rf ./$2
fi

rm -rf ../artifacts/__pycache__
rm -rf ../artifacts/docker-compose.yml

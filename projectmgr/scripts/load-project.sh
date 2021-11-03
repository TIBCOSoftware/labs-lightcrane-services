#!/bin/bash
# $1 : RepositoryID
# $2 : ProjectID
# $3 : ProjectFolder
# $Username
# $Password

if [ -d "$3" ] 
then
    rm -rf $3
fi

mkdir $3

git clone https://$Username:$Password@github.com/$Username/$2.git $3
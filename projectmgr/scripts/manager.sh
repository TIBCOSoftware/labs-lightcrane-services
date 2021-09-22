#!/bin/bash
ProjectFolder=$1
Type=$2
Action=$3

Script='./'$Action'-'$Type'.sh'

if [ -f "$Script" ]; then
    chmod +x $Script
    $Script $1 $2 $3 $4
else 
	cd ./python
	chmod +x list-file.py
	python3 './'$Action'-'$Type'.py' $ProjectFolder
fi

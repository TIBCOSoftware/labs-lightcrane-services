#!/usr/bin/env python3

import os 
import sys
import json

def list_files(structure, startpath):
    structure['files'] = []
    for name in os.listdir(startpath):
        if os.path.isdir(os.path.join(startpath, name)):
            structure[name] = {}
            list_files(structure[name], os.path.join(startpath, name))
        elif os.path.isfile(os.path.join(startpath, name)):
            structure['files'].append(name)
            
projects = {}
list_files(projects, '/Users/steven/Desktop/test/projects')
print(projects)
#!/usr/bin/env python3

import os 
import sys
import json

def list_projects(structure, startpath):
    structure['files'] = []
    for name in os.listdir(startpath):
        if name.startswith('.'):
            continue
        if os.path.isdir(os.path.join(startpath, name)):
            structure[name] = {}
            list_files(structure[name], os.path.join(startpath, name))
        elif os.path.isfile(os.path.join(startpath, name)):
            structure['files'].append(name)

projectFolder = sys.argv[1]

response = {}
projects = {}
response['Projects'] = projects

list_projects(projects, projectFolder)

print(json.dumps(response, indent = 4))
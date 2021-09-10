#!/usr/bin/env python3

import pathlib
import sys

currentFolder = pathlib.Path().resolve()
sys.path.append('{0}/common'.format(currentFolder))

import os 
import json
import keyword_replace

LC_DESCRIPTOR = 'lc.json'


def valid_project(projectPath):
    return os.path.exists(os.path.join(projectPath, LC_DESCRIPTOR))
    
def get_project(metadata):
    with open(metadata, 'r') as f:
        data = f.read()
        return json.loads(data)
    
def get_parameter(predefined, docker_compose):
    parameters = {}
    predefinedData = None
    if os.path.exists(predefined): 
        with open(predefined, 'r') as predf:
            predefinedData = json.load(predf)
    
    if predefinedData is None:
        predefinedData = {}

    with open(docker_compose, 'r') as docf:
        data = docf.read()
        keywordMapper = keyword_replace.KeywordMapper('', '${', '}')
        parametersArray = keywordMapper.replace(data, keyword_replace.KeywordCollectingHandler())
        for parameter in parametersArray:
            if parameter in predefinedData.keys():
                parameters[parameter] = predefinedData[parameter]
            else:
                parameters[parameter] = 'not predefined'
    return parameters

projectBaseFolder = sys.argv[1]

projects = [ name for name in os.listdir(projectBaseFolder) if valid_project(os.path.join(projectBaseFolder, name)) and os.path.isdir(os.path.join(projectBaseFolder, name)) ]

response = {}
response['Projects'] = []

for project in projects:
#    print(projects)
    projectFolder = os.path.join(projectBaseFolder, project)
#    print(projectFolder)
    projectObj = get_project(os.path.join(projectFolder, LC_DESCRIPTOR))
    projectObj['Deployables'] = []
    deployableFolder = os.path.join(projectFolder, 'artifacts/docker-compose')
    descriptors = [ fi for fi in os.listdir(deployableFolder) if any(fi.endswith(ext) for ext in ['yml', 'yaml']) ]
    for descriptor in descriptors:
        deployable = descriptor[: descriptor.rindex('.')]
        deployableObj = {}
        deployableObj['id'] = 'id_{}'.format(deployable)
        deployableObj['name'] = deployable
        deployableObj['parameters'] = get_parameter(
            os.path.join(deployableFolder, '{0}-default.json'.format(deployable)), 
            os.path.join(deployableFolder, descriptor))
        projectObj['Deployables'].append(deployableObj)
    response['Projects'].append(projectObj)

print(json.dumps(response, indent = 4))
#!/usr/bin/env python3

import glob
import os
import shutil
import sys
import json
import importlib
import pathlib
print ('Number of arguments:', len(sys.argv), 'arguments.')
print ('Argument List:', str(sys.argv))

currentFolder = pathlib.Path().resolve()
print('=============>{0}'.format(currentFolder))
sys.path.append('{0}/common'.format(currentFolder))

module = importlib.import_module(sys.argv[1])
transform = module.Transform()
#deployables = glob.glob(os.path.join(sys.argv[2], "*.yml"))
deployables = sys.argv[4].split('|')

print(deployables)

previousDeployable = None
for deployable in deployables:
    filepath = '{0}/{1}.yml'.format(sys.argv[2], deployable)
    with open(filepath) as file:
        properties = {}
        properties['org'] = os.getenv('HZN_ORG_ID')
        properties['serviceName'] = deployable #os.getenv('ServiceName')
        properties['serviceVersion'] = os.getenv('ServiceVersion')
        properties['arch'] = os.getenv('Arch')
        properties['dependencies'] = []
        if previousDeployable is not None:
            properties['dependencies'].append({
                "url": previousDeployable,
                "org": properties['org'],
                "versionRange": '[0.0.0,INFINITY)',
                "arch": properties['arch']
            })
        descriptors = transform.toDescriptor(file, properties)
        deployableWorkFolder = os.path.join(sys.argv[3], deployable)
        if os.path.isdir(deployableWorkFolder):
            shutil.rmtree(deployableWorkFolder)
        os.makedirs(deployableWorkFolder)
        for key in descriptors:
            fw = open("{0}/{1}/{2}".format(sys.argv[3], deployable, key), "w")
            fw.write(descriptors[key])
            fw.close()
        previousDeployable = deployable

#for filepath in deployables:
#    with open(filepath) as file:
#        transform.addConf(file)
#    descriptors = transform.to_descriptor()
#    for key in descriptors:
#        fw = open("../../{0}/{1}".format(sys.argv[4], key), "w")
#        fw.write(descriptors[key])
#        fw.close()

#with open('../../pipeline/{0}.json'.format(sys.argv[3])) as fr:
#    transform = module.Transform({ 'data':fr })
#    descriptors = transform.to_descriptor()
#    for key in descriptors:
#        fw = open("../../deploy/{0}/{1}/{2}".format(sys.argv[2], sys.argv[3], key), "w")
#        fw.write(descriptors[key])
#        fw.close()

#print(data)

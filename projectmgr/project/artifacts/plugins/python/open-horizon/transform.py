# open-horizon/transform.py

import copy
import json
import yaml
import os
import keyword_replace

class Transform:
    def __init__(self):
        pass

    def buildService(self, serviceName, service):
        # Construct service.json
        serviceObj = {}
        serviceObj['image'] = service['image']
        serviceObj['privileged'] = True
        serviceObj['network'] = 'host'

        if ('ports' in service.keys()) and (serviceObj['network'] != 'host'):
            # ports:
            #  - "3000"
            #  - "3000-3005"
            #  - "8000:8000"
            #  - "9090-9091:8080-8081"
            #  - "49100:22"
            #  - "127.0.0.1:8001:8001"
            #  - "127.0.0.1:5000-5010:5000-5010"
            #  - "127.0.0.1::5000"
            #  - "6060:6060/udp"
            #  - "12400-12500:1240"
            serviceObj['ports'] = []
            for port in service['ports']:
                portCmps = port.split(':')
                if 3==len(portCmps): 
                    serviceObj['ports'].append({ 
                        'HostPort': port[port.index(':')+1:], 
                        'HostIP': portCmps[0] 
                    })
                else:
                    serviceObj['ports'].append({ 
                        'HostPort': port 
                    })

        serviceObj['environment'] = []
        if 'environment' in service.keys():
            if type(service['environment']) == dict:
                for key in service['environment']:
                    serviceObj['environment'].append("{0}={1}".format(key, self.normalize(service['environment'][key])))
            else:
                for envStr in service['environment']:
                    env=envStr.split('=')
                    serviceObj['environment'].append("{0}={1}".format(env[0], self.normalize(env[1])))

        serviceObj['binds'] = []
        if 'volumes' in service.keys():
            serviceObj['binds']=service['volumes']

        serviceObj['entrypoint'] = []
        if 'entrypoint' in service.keys():
            serviceObj['entrypoint']=service['entrypoint']

        serviceObj['devices'] = []
        if 'devices' in service.keys():
            serviceObj['devices']=service['devices']

        return {
            "name" : service['container_name'],
            "service" : serviceObj
        }

    def toDescriptor(self, file, properties):
        # get descriptor template
        descriptors = self.getOHDdescriptors()
        
        # prepare sevice object
        descriptors['service']['org'] = properties['org']
        descriptors['service']['label'] = '{0} for {1}'.format(properties['serviceName'], properties['arch'])
        descriptors['service']['url'] = properties['serviceName']
        descriptors['service']['version'] = properties['serviceVersion']
        descriptors['service']['arch'] = properties['arch']    
        descriptors['service']['requiredServices'] = properties['dependencies']
        descriptors['service']['deployment']={}
        descriptors['service']['deployment']['services']={}        
        data = yaml.load(file)
        for serviceName in data['services']:
            print('==============> {0}'.format(serviceName))
            serviceObj = self.buildService(serviceName, data['services'][serviceName])
            descriptors['service']['deployment']['services'][serviceObj['name']] = serviceObj['service']
            
        # prepare deployment policy
        descriptors['policy']['service']['name'] = properties['serviceName']
        descriptors['policy']['service']['org'] = properties['org']
        descriptors['policy']['service']['arch'] = properties['arch']
        descriptors['policy']['service']['serviceVersions'].append({ 'version': properties['serviceVersion'], 'priority': {} })
        descriptors['policy']['userInput'][0]['serviceOrgid'] = properties['org']
        descriptors['policy']['userInput'][0]['serviceUrl'] = properties['serviceName']
        constraintStr = os.getenv('DeployConstrains')
        print('**********************************')
        print(constraintStr)
        print('**********************************')
        constraints = json.loads(constraintStr)
        descriptors['policy']['constraints'] = constraints
        
        # return all required descriptors
        return {
            'service.json' : keyword_replace.KeywordMapper('', '${', '}').replace(json.dumps(descriptors['service'], indent = 4), keyword_replace.KeywordReplaceHandler(os.environ)),
#            'service.json' : json.dumps(descriptors['service'], indent = 4),
            'service.policy.json' : json.dumps(descriptors['servicePolicy'], indent = 4),
            'deployment.policy.json' : json.dumps(descriptors['policy'], indent = 4)
        }
        
    def normalize(self, value):
        if type(value) == str and value.startswith('${') and value.endswith('}'):
            return '${0}'.format(value[2:len(value)-1])
        else:
            return value

    def getOHDdescriptors(self):
        return {
            'service' : {
                'description': 'A template for Horizon Project-Air service',
                'documentation': 'https://github.com/TIBCOSoftware/labs-air/integration/oh/README.md',
                'sharable': 'multiple'
            },
            'servicePolicy' : {
                'properties': [],
                'constraints': []
            },
            'policy' : {
                'label': '',
                'description': '',
                'service': {
                    'serviceVersions': []
                },
                'properties': [],
                "userInput": [
                    {
                        "serviceVersionRange": "[0.0.0,INFINITY)",
                        "inputs": [
                            {
                                "name": "HW_WHO",
                                "value": "Valued Customer"
                            }
                        ]
                    }
                ]
            }
        }
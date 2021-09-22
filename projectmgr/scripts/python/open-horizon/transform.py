# open-horizon/transform.py

import json
import yaml
import os
import keyword_replace

class Transform:
    def __init__(self):
        self.files = []
        self.service = {
            'org': '$HZN_ORG_ID',
            'label': '$ServiceName for $Arch',
            'description': 'A template for Horizon Project-Air service',
            'documentation': 'https://github.com/TIBCOSoftware/labs-air/integration/oh/README.md',
            'url': '$ServiceName',
            'version': '$ServiceVersion',
            'arch': '$Arch',
            'sharable': 'multiple'
        }
        
        self.servicePolicy = {
            'properties': [
            ],
            'constraints': [
            ]
        }
        
        self.policy = {
            'label': '',
            'description': '',
            'service': {
                'name': '$ServiceName',
                'org': '$HZN_ORG_ID',
                'arch': '$Arch',
                'serviceVersions': [ { 'version': '$ServiceVersion', 'priority': {} } ]
            },
            'properties': [
            ],
            "userInput": [
                {
                    "serviceOrgid": "$HZN_ORG_ID",
                    "serviceUrl": "$ServiceName",
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

        self.service['deployment']={}
        self.service['deployment']['services']={}
        
    def addConf(self, file):
        data = yaml.load(file)
        self.files.append(data)
        for serviceName in data['services']:
            print('==============> {0}'.format(serviceName))
            serviceObj = self.buildService(serviceName, data['services'][serviceName])
            self.service['deployment']['services'][serviceObj['name']] = serviceObj['service']

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

    def to_descriptor(self):
        constraintStr = os.getenv('DeployConstrains')
        print('**********************************')
        print(constraintStr)
        print('**********************************')
        constraints = json.loads(constraintStr)
        self.policy['constraints'] = constraints
#        keywordMapper = keyword_replace.KeywordMapper('', '${', '}')
#        serviceText = keywordMapper.replace(json.dumps(self.service, indent = 4), keyword_replace.KeywordReplaceHandler(os.environ))
        serviceText = json.dumps(self.service, indent = 4)
        return {
            'service.json' : serviceText,
            'service.policy.json' : json.dumps(self.servicePolicy, indent = 4),
            'deployment.policy.json' : json.dumps(self.policy, indent = 4)
        }
        
    def normalize(self, value):
        if type(value) == str and value.startswith('${') and value.endswith('}'):
            return '${0}'.format(value[2:len(value)-1])
        else:
            return value

#{
#    "org": "$HZN_ORG_ID",
#    "label": "$SERVICE_NAME for $ARCH",
#    "description": "A template for Horizon Project-Air service",
#    "documentation": "https://github.com/TIBCOSoftware/labs-air/integration/oh/README.md",
#    "url": "$SERVICE_NAME",
#    "version": "$SERVICE_VERSION",
#    "arch": "$ARCH",
#    "sharable": "multiple",
#    "requiredServices": [],
#    "userInput": [
#        {
#            "name": "HW_WHO",
#            "label": "Who to say hello to",
#            "type": "string",
#            "defaultValue": "World"
#        }
#    ],
#    "deployment": {
#        "services": {
#            "$SERVICE_NAME": {
#                "image": "${DOCKER_IMAGE_BASE}_$ARCH:$SERVICE_VERSION",
#                "environment": [
#                	"ApplicationSettings_GatewayAccessToken=${GATEWAY_ACCESS_TOKEN}",
#                	"ApplicationSettings_GatewayUsername=${GATEWAY_USERNAME}",
#                	"ApplicationSettings_GatewayPlatform=${GATEWAY_PLATFORM}",
#                	"ApplicationSettings_MetadataClient=${GATEWAY_METADATA_CLIENT}",
#                	"ApplicationSettings_MetadataPublishIntervalSecs=${GATEWAY_METADATA_PUBLISH_INTERVAL_SECS}"
#				],
#              "specific_ports": [
#                { "HostPort": "5555:7777/udp", "HostIP": "0.0.0.0" }
#              ],
#				"binds": [
#					"db-data:/data/db",
#					"log-data:/edgex/logs",
#					"consul-config:/consul/config",
#					"consul-data:/consul/data"
#				]
#           }
#        }
#    }
#}
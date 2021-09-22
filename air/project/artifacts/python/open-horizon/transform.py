# open-horizon/transform.py

import json
import os
import keyword_replace

class Transform:
    def __init__(self, config):
        data = json.load(config['data'])
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

        self.pipelineProperties = data['Services'][0]['Properties'][0]['Value']
        self.appName = None
        self.image = None
        self.service['deployment']={}
        self.service['deployment']['services']={}

    def to_descriptor(self):
        # Construct service.json
        for property in self.pipelineProperties:
            propertyName = property['Name']
            if propertyName.startswith('services.') :
                nameElements = propertyName.split('.')
                if nameElements[1] != self.appName :
                    self.appName = nameElements[1]
                    self.service['deployment']['services'][self.appName]={}
                    self.service['deployment']['services'][self.appName]['environment'] = []
                    self.service['deployment']['services'][self.appName]['binds'] = []
                    self.service['deployment']['services'][self.appName]['specific_ports'] = []
                    self.service['deployment']['services'][self.appName]['ephemeral_ports'] = []
                    self.service['deployment']['services'][self.appName]['network'] = 'host'
                    self.service['deployment']['services'][self.appName]['privileged'] = True

                if  nameElements[2].startswith('environment[') :
                    self.service['deployment']['services'][self.appName]['environment'].append(property['Value'])
                elif nameElements[2].startswith('volumes[') :
                    self.service['deployment']['services'][self.appName]['binds'].append(property['Value'])
                elif nameElements[2].startswith('ports[') :
                    port = { 
                        'HostPort': '{0}/{1}'.format(property['Value'],'tcp'), 
                        'HostIP': '0.0.0.0' 
                    }
                    self.service['deployment']['services'][self.appName]['specific_ports'].append(port)
                elif 'container_name' == nameElements[2] :
                    self.containerName = property['Value']
                elif 'image' == nameElements[2] :
                    self.service['deployment']['services'][self.appName]['image'] = property['Value']

        print(self.service)

        propertieStr = os.getenv('ServiceProperties')
        print('**********************************')
        print(propertieStr)
        print('**********************************')
        propertyDic = json.loads(propertieStr)
        self.servicePolicy['properties'] = []
        for property in propertyDic:
            self.servicePolicy['properties'].append({
                'name' : property['Name'],
                'value' : property['Value']
            })
        print(self.servicePolicy)

        constraintStr = os.getenv('DeployConstrains')
        print('**********************************')
        print(constraintStr)
        print('**********************************')
        constraints = json.loads(constraintStr)
        self.policy['constraints'] = constraints
        print(self.policy)
        
        return {
            'service.json' : keyword_replace.KeywordMapper('', '${', '}').replace(json.dumps(self.service, indent = 4), keyword_replace.KeywordReplaceHandler(os.environ)),
#            'service.json' : json.dumps(self.service, indent = 4),
            'service.policy.json' : json.dumps(self.servicePolicy, indent = 4),
            'deployment.policy.json' : json.dumps(self.policy, indent = 4)
        }


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
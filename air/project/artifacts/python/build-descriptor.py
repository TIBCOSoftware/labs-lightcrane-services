#!/usr/bin/env python3

import os
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

with open('../../pipeline/{0}.json'.format(sys.argv[3])) as fr:
    transform = module.Transform({ 'data':fr })
    descriptors = transform.to_descriptor()
    for key in descriptors:
        fw = open("../../deploy/{0}/{1}/{2}".format(sys.argv[2], sys.argv[3], key), "w")
        fw.write(descriptors[key])
        fw.close()

#print(data)


#{
#    "DeploymentGpID":"Air-account_00001",
#    "Services":[
#        {
#            "Descriptor":"docker-compose.yml",
#            "Name":"http_mqtt_fs",
#            "Properties":[
#                {
#                    "Group":"main",
#                    "Value":[
#                        {"Name":"version","Type":"","Value":"3.6"},
#                        {"Name":"services.http_mqtt_fs.container_name","Type":"","Value":"Air-account_00001_http_mqtt_fs"},
#                        {"Name":"services.http_mqtt_fs.build","Type":"","Value":"001"},
#                        {"Name":"networks.default.name","Type":"","Value":"http_rest_mqtt"},
#                        {"Name":"services.http_mqtt_fs.environment[0]","Type":null,"Value":"FLOGO_APP_PROPS_ENV=auto"},
#                        {"Name":"services.http_mqtt_fs.environment[1]","Type":null,"Value":"DataSource_Logging_LogLevel=INFO"},
#                        {"Name":"services.http_mqtt_fs.environment[2]","Type":null,"Value":"Mqtt_Pipe_0_IoTMQTT_Broker_URL=tcp://74.101.118.162:1883"},
#                        {"Name":"services.http_mqtt_fs.environment[3]","Type":null,"Value":"Mqtt_Pipe_0_IoTMQTT_Username=mqtt_admin"},
#                        {"Name":"services.http_mqtt_fs.environment[4]","Type":null,"Value":"Mqtt_Pipe_0_IoTMQTT_Password=SECRET:79V5PfQgmw5mTglH3kNiNcoeLPJGsx1w7Tw="},
#                        {"Name":"services.http_mqtt_fs.environment[5]","Type":null,"Value":"Pipe_0_Logging_LogLevel=Info"},
#                        {"Name":"services.http_mqtt_fs.environment[6]","Type":null,"Value":"Pipe_0_MQTTPub_PublishData={\"id\": \"@f1..id@\", \"source\": \"@f1..source@\", \"device\": \"@f1..device@\", \"gateway\": \"@f1..gateway@\", \"readings\": [{\"id\": \"@f1..id@\", \"origin\": \"@f1..origin@\", \"device\": \"@f1..device@\", \"name\": \"@f1..name@\", \"value\": \"@f1..value@\", \"valueType\": \"@f1..valueType@\", \"mediaType\": \"@f1..mediaType@\"}]}"},
#                        {"Name":"services.http_mqtt_fs.environment[7]","Type":null,"Value":"Pipe_0_MQTTPub_Topic=AIRModelScoredData02"},
#                        {"Name":"services.http_mqtt_fs.environment[8]","Type":null,"Value":"Logging_LogLevel=DEBUG"},
#                        {"Name":"services.http_mqtt_fs.ports[0]","Type":"String","Value":"8080:9999"}
#                    ]
#                }
#            ],
#            "ScriptSystemEnv":{
#                "Platformx":"linux/386",
#                "RSVPUrl":"http://air:10099/f1/air/rsvp/Air-account_00001/http_mqtt_fs",
#                "TargetServer":"mtorresagent"
#            },
#            "Type":"docker",
#            "Variables":null
#        }
#    ]
#}

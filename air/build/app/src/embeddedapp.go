// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package main

// embedded flogo app descriptor file
const flogoJSON string = `{
  "name": "app",
  "type": "flogo:app",
  "version": "1.0.0",
  "description": "",
  "appModel": "1.1.1",
  "imports": [
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/filereader",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsondeserializer",
    "github.com/project-flogo/contrib/activity/actreturn",
    "github.com/project-flogo/flow/activity/subflow",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/function/f1",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/mapping",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/httpclient",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/tablemutate",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsonserializer",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/airpipelinebuilder",
    "github.com/project-flogo/contrib/trigger/timer",
    "github.com/project-flogo/contrib/function/array",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/airparameterbuilder",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/exec",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/objectmaker",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/objectserializer",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/aircomponentquerier",
    "github.com/project-flogo/legacybridge",
    "github.com/project-flogo/flow",
    "github.com/project-flogo/contrib/function/string",
    "github.com/project-flogo/contrib/function/coerce",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/log",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/filewriter",
    "github.com/project-flogo/contrib/trigger/rest",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/tableupsert",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/textreplacer"
  ],
  "properties": [
    {
      "name": "System.ServiceLocatorIP",
      "type": "string",
      "value": "localhost"
    },
    {
      "name": "System.BaseFolder",
      "type": "string",
      "value": "/"
    },
    {
      "name": "System.FlogoBuilder",
      "type": "string",
      "value": "Flogo builder not set"
    },
    {
      "name": "System.BaseFolderExt",
      "type": "string",
      "value": "not set"
    },
    {
      "name": "System.BuilderURL",
      "type": "string",
      "value": "not set"
    },
    {
      "name": "System.DeployerURL",
      "type": "string",
      "value": "not set"
    },
    {
      "name": "Descriptor.docker",
      "type": "string",
      "value": "docker-compose.yml"
    },
    {
      "name": "Descriptor.k8s",
      "type": "string",
      "value": "deployment.yml"
    },
    {
      "name": "System.ExternalEndpointIP",
      "type": "string",
      "value": "not set"
    },
    {
      "name": "System.ServiceLocatorExternal",
      "type": "string",
      "value": "not set"
    }
  ],
  "triggers": [
    {
      "id": "ReceiveHTTPMessage",
      "ref": "#rest",
      "settings": {
        "certFile": "",
        "enableTLS": false,
        "keyFile": "",
        "port": 10099
      },
      "handlers": [
        {
          "name": "Create_Domain",
          "settings": {
            "method": "POST",
            "path": "/f1/air/createDomain/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Create_Domain"
              },
              "input": {
                "ProjectID": "=$.pathParams.id",
                "Properties": "=$.content.Properties"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"Properties\": {\n    }\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Properties\":{\"type\":\"object\",\"properties\":{}}}}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Get_Pipelines",
          "settings": {
            "method": "GET",
            "path": "/f1/air/pipelines/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Get_Pipelines"
              },
              "input": {
                "ProjectID": "=$.pathParams.id"
              },
              "output": {
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Delete_Pipeline",
          "settings": {
            "method": "DELETE",
            "path": "/f1/air/delete/:id/:name"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Delete_Pipeline"
              },
              "input": {
                "Name": "=$.pathParams.name",
                "ProjectID": "=$.pathParams.id"
              },
              "output": {
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"},{\"parameterName\":\"name\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Build_Pipeline",
          "settings": {
            "method": "POST",
            "path": "/f1/air/build/:id/:name"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Build_Pipeline"
              },
              "input": {
                "AirDescriptor": "=$.content.AirDescriptor",
                "ComponentType": "=$.content.ComponentType",
                "FlogoAppDescriptor": "=$.content.FlogoAppDescriptor",
                "Image": "=$.content.Image",
                "Name": "=$.pathParams.name",
                "ProjectID": "=$.pathParams.id",
                "ScriptSystemEnv": "=$.content.ScriptSystemEnv",
                "ServiceType": "=$.content.ServiceType",
                "SyncBuild": "=$.content.SyncBuild"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"ProjectID\": {\n            \"type\": \"string\"\n        },\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"ServiceType\": {\n            \"type\": \"string\"\n        },\n        \"Name\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AirDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"source\": {\n                    \"type\": \"object\",\n                    \"properties\": {\n                        \"name\": {\n                            \"type\": \"string\"\n                        },\n                        \"properties\": {\n                            \"type\": \"array\",\n                            \"items\": {\n                                \"type\": \"object\",\n                                \"properties\": {\n                                    \"Name\": {\n                                        \"type\": \"string\"\n                                    },\n                                    \"Value\": {\n                                        \"type\": \"string\"\n                                    },\n                                    \"Type\": {\n                                        \"type\": \"string\"\n                                    }\n                                }\n                            }\n                        },\n                        \"ports\": {\n                            \"type\": \"array\",\n                            \"items\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"logic\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"name\": {\n                                \"type\": \"string\"\n                            },\n                            \"properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"extra\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"FlogoAppDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"flogoApp\": {\n                    \"type\": \"string\"\n                },\n                \"properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"ports\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"string\"\n                    }\n                },\n                \"extra\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"ProjectID\": {\n            \"type\": \"string\"\n        },\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"ServiceType\": {\n            \"type\": \"string\"\n        },\n        \"Name\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AirDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"source\": {\n                    \"type\": \"object\",\n                    \"properties\": {\n                        \"name\": {\n                            \"type\": \"string\"\n                        },\n                        \"properties\": {\n                            \"type\": \"array\",\n                            \"items\": {\n                                \"type\": \"object\",\n                                \"properties\": {\n                                    \"Name\": {\n                                        \"type\": \"string\"\n                                    },\n                                    \"Value\": {\n                                        \"type\": \"string\"\n                                    },\n                                    \"Type\": {\n                                        \"type\": \"string\"\n                                    }\n                                }\n                            }\n                        },\n                        \"ports\": {\n                            \"type\": \"array\",\n                            \"items\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"logic\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"name\": {\n                                \"type\": \"string\"\n                            },\n                            \"properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"extra\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"FlogoAppDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"flogoApp\": {\n                    \"type\": \"string\"\n                },\n                \"properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"ports\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"string\"\n                    }\n                },\n                \"extra\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"},{\"parameterName\":\"name\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n        \"Success\": true,\n        \"Message\": \"string\",\n        \"ErrorCode\": 100\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Validate_an_Air_descriptor_then_convert_to_Flogo_descriptor",
          "settings": {
            "method": "POST",
            "path": "/f1/air/validate/:id/:name"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Validate_an_Air_descriptor_then_convert_to_Flogo_descriptor"
              },
              "input": {
                "AirDescriptor": "=$.content.AirDescriptor",
                "ComponentType": "=$.content.ComponentType",
                "Image": "=$.content.Image",
                "Name": "=$.pathParams.name",
                "ProjectID": "=$.pathParams.id",
                "ScriptSystemEnv": "=$.content.ScriptSystemEnv",
                "ServiceType": "=$.content.ServiceType",
                "SyncBuild": "=$.content.SyncBuild"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"SyncBuild\": false,\n    \"ComponentType\" : \"Service\",\n    \"ServiceType\" : \"docker\",\n    \"Image\":\"bigoyang/mqtt_dgraph:0.1.0\",\n    \"AirDescriptor\" : {\n        \"source\": {\n            \"name\": \"MQTT\",\n            \"properties\" : [\n                {\n                    \"Name\" : \"\", \n                    \"Value\":\"\",\n                    \"Type\":\"\"\n                }\n            ],\n            \"ports\" : [\n                \"8080:9999\"\n            ]\n        },\n        \"logic\": [\n            {\n                \"name\": \"Dgraph\",\n                \"properties\" : [\n                    {\n                        \"Name\" : \"\", \n                        \"Value\":\"\",\n                        \"Type\":\"\"\n                    }\n                ]\n            }\n        ],\n        \"extra\": [\n            {\"Name\":\"networks.default.external.name\", \"Value\":\"edgex-network\"}\n        ]\n    },\n    \"ScriptSystemEnv\":{\n    }\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"},{\"parameterName\":\"name\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n        \"Success\": true,\n        \"Message\": \"string\",\n        \"ErrorCode\": 100\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Build_and_Deploy_Pipeline",
          "settings": {
            "method": "POST",
            "path": "/f1/air/buildAndDeploy/:id/:name"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Build_and_Deploy_Pipeline"
              },
              "input": {
                "AirDescriptor": "=$.content.AirDescriptor",
                "ComponentType": "=$.content.ComponentType",
                "FlogoAppDescriptor": "=$.content.FlogoAppDescriptor",
                "Image": "=$.content.Image",
                "Name": "=$.pathParams.name",
                "ProjectID": "=$.pathParams.id",
                "ScriptSystemEnv": "=$.content.ScriptSystemEnv",
                "ServiceType": "=$.content.ServiceType",
                "SyncBuild": "=$.content.SyncBuild"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"ProjectID\": {\n            \"type\": \"string\"\n        },\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"ServiceType\": {\n            \"type\": \"string\"\n        },\n        \"Name\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AirDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"source\": {\n                    \"type\": \"object\",\n                    \"properties\": {\n                        \"name\": {\n                            \"type\": \"string\"\n                        },\n                        \"properties\": {\n                            \"type\": \"array\",\n                            \"items\": {\n                                \"type\": \"object\",\n                                \"properties\": {\n                                    \"Name\": {\n                                        \"type\": \"string\"\n                                    },\n                                    \"Value\": {\n                                        \"type\": \"string\"\n                                    },\n                                    \"Type\": {\n                                        \"type\": \"string\"\n                                    }\n                                }\n                            }\n                        },\n                        \"ports\": {\n                            \"type\": \"array\",\n                            \"items\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"logic\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"name\": {\n                                \"type\": \"string\"\n                            },\n                            \"properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"extra\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"FlogoAppDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"flogoApp\": {\n                    \"type\": \"string\"\n                },\n                \"properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"ports\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"string\"\n                    }\n                },\n                \"extra\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"ProjectID\": {\n            \"type\": \"string\"\n        },\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"ServiceType\": {\n            \"type\": \"string\"\n        },\n        \"Name\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AirDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"source\": {\n                    \"type\": \"object\",\n                    \"properties\": {\n                        \"name\": {\n                            \"type\": \"string\"\n                        },\n                        \"properties\": {\n                            \"type\": \"array\",\n                            \"items\": {\n                                \"type\": \"object\",\n                                \"properties\": {\n                                    \"Name\": {\n                                        \"type\": \"string\"\n                                    },\n                                    \"Value\": {\n                                        \"type\": \"string\"\n                                    },\n                                    \"Type\": {\n                                        \"type\": \"string\"\n                                    }\n                                }\n                            }\n                        },\n                        \"ports\": {\n                            \"type\": \"array\",\n                            \"items\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"logic\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"name\": {\n                                \"type\": \"string\"\n                            },\n                            \"properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"extra\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"FlogoAppDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"flogoApp\": {\n                    \"type\": \"string\"\n                },\n                \"properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"ports\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"string\"\n                    }\n                },\n                \"extra\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"string\"\n                            },\n                            \"Type\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                }\n            }\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"},{\"parameterName\":\"name\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n        \"Success\": true,\n        \"Message\": \"string\",\n        \"ErrorCode\": 100\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Delete_Domain",
          "settings": {
            "method": "DELETE",
            "path": "/f1/air/deleteDomain/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Delete_Domain"
              },
              "input": {
                "ProjectID": "=$.pathParams.id"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Undeploy_Pipeline",
          "settings": {
            "method": "POST",
            "path": "/f1/air/undeploy/:id/:name/:instance"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Undeploy_Pipeline"
              },
              "input": {
                "Instance": "=$.pathParams.instance",
                "Method": "=$.content.Method",
                "Name": "=$.pathParams.name",
                "ProjectID": "=$.pathParams.id",
                "ScriptSystemEnv": "=$.content.ScriptSystemEnv",
                "SyncDeploy": "=$.content.SyncDeploy"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"},{\"parameterName\":\"name\",\"type\":\"string\"},{\"parameterName\":\"instance\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"instance\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Deploy_Pipeline",
          "settings": {
            "method": "POST",
            "path": "/f1/air/deploy/:id/:name/:instance"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Deploy_Pipeline"
              },
              "input": {
                "AirDescriptor": "=$.content.AirDescriptor",
                "Instance": "=$.pathParams.instance",
                "Method": "=$.content.Method",
                "Name": "=$.pathParams.name",
                "ProjectID": "=$.pathParams.id",
                "ScriptSystemEnv": "=$.content.ScriptSystemEnv",
                "SyncDeploy": "=$.content.SyncDeploy"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"Method\": \"Script\",\n    \"SyncDeploy\": false,\n    \"ScriptSystemEnv\":{\n    },\n    \"AirDescriptor\" : {\n        \"dynamic\": {\n            \"properties\" : [\n                {\"Component\":\"DataSource\", \"Name\":\"KafkaTrigger.InitialOffset\", \"Value\":\"Oldest\"}\n            ]\n        },\n        \"extra\" : [\n            {\"Name\":\"\", \"Value\":\"\", \"Type\":\"\"}\n        ]\n    }\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Method\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"dynamic\":{\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Component\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"},{\"parameterName\":\"name\",\"type\":\"string\"},{\"parameterName\":\"instance\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"instance\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Get_By_Category",
          "settings": {
            "method": "GET",
            "path": "/f1/air/components/:category"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Get_By_Category"
              },
              "input": {
                "Category": "=$.pathParams.category"
              },
              "output": {
                "code": 200,
                "data": "=$.Data"
              }
            }
          ],
          "schemas": {
            "output": {
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"category\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"category\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Get_Compnent",
          "settings": {
            "method": "GET",
            "path": "/f1/air/components/:category/:component"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Get_Compnent"
              },
              "input": {
                "Category": "=$.pathParams.category",
                "Component": "=$.pathParams.component"
              },
              "output": {
                "code": 200,
                "data": "=$.Data"
              }
            }
          ],
          "schemas": {
            "output": {
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"category\",\"type\":\"string\"},{\"parameterName\":\"component\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"category\":{\"type\":\"string\"},\"component\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "RSVP_from_builder",
          "settings": {
            "method": "POST",
            "path": "/f1/air/rsvp/:id/:name"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:RSVP_from_builder"
              },
              "input": {
                "ErrorMsg": "=$.content.ErrorMsg",
                "Name": "=$.pathParams.name",
                "ProjectID": "=$.pathParams.id",
                "Successful": "=$.content.Successful"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Successful\": {\n            \"type\": \"boolean\"\n        },\n        \"ErrorMsg\": {\n            \"type\": \"string\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Successful\": {\n            \"type\": \"boolean\"\n        },\n        \"ErrorMsg\": {\n            \"type\": \"string\"\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"id\",\"type\":\"string\"},{\"parameterName\":\"name\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"id\":{\"type\":\"string\"},\"name\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n        \"Success\": true,\n        \"Message\": \"string\",\n        \"ErrorCode\": 100\n}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Get_Compnents",
          "settings": {
            "method": "GET",
            "path": "/f1/air/components"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Get_Compnents"
              },
              "output": {
                "code": 200,
                "data": "=$.Data"
              }
            }
          ],
          "schemas": {
            "output": {
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Get_Flogo_App_Properties",
          "settings": {
            "method": "POST",
            "path": "/f1/air/flogo/properties"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Get_Flogo_App_Properties"
              },
              "input": {
                "flogoApp": "=$.content.flogoApp"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "properties": {
                      "@foreach($.properties, properties)": {
                        "=": "$loop"
                      }
                    }
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"flogoApp\": {\n            \"type\": \"string\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"flogoApp\": {\n            \"type\": \"string\"\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"properties\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Value\": {\n                        \"type\": \"string\"\n                    },\n                    \"Type\": {\n                        \"type\": \"string\"\n                    }\n                }\n            }\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"properties\": {\n            \"type\": \"array\",\n            \"items\": {\n                \"type\": \"object\",\n                \"properties\": {\n                    \"Name\": {\n                        \"type\": \"string\"\n                    },\n                    \"Value\": {\n                        \"type\": \"string\"\n                    },\n                    \"Type\": {\n                        \"type\": \"string\"\n                    }\n                }\n            }\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Upload_Flogo_Component",
          "settings": {
            "method": "PUT",
            "path": "/f1/air/flogo/component"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Upload_Flogo_Component"
              }
            }
          ],
          "schemas": {
            "output": {
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              }
            },
            "reply": {
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "Upload_GO_Lib",
          "settings": {
            "method": "PUT",
            "path": "/f1/air/upload/:type/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Upload_GO_Lib"
              },
              "input": {
                "ID": "=$.pathParams.id",
                "Path": "=$.content.Path",
                "Resource": "=$.content.Resource",
                "Type": "=$.pathParams.type"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "ErrorCode": "=$.ErrorCode",
                    "Message": "=$.Message",
                    "Success": "=$.Success"
                  }
                }
              }
            }
          ],
          "schemas": {
            "output": {
              "body": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Path\": {\n            \"type\": \"string\"\n        },\n        \"Resource\": {\n            \"type\": \"string\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Path\": {\n            \"type\": \"string\"\n        },\n        \"Resource\": {\n            \"type\": \"string\"\n        }\n    }\n}"
              },
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              },
              "pathParams": {
                "fe_metadata": "[{\"parameterName\":\"type\",\"type\":\"string\"},{\"parameterName\":\"id\",\"type\":\"string\"}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"type\":{\"type\":\"string\"},\"id\":{\"type\":\"string\"}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}",
                "type": "json",
                "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"Success\": {\n            \"type\": \"boolean\"\n        },\n        \"Message\": {\n            \"type\": \"string\"\n        },\n        \"ErrorCode\": {\n            \"type\": \"number\"\n        }\n    }\n}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        },
        {
          "name": "List_Deployable_Projects",
          "settings": {
            "method": "GET",
            "path": "/f1/air/list/projects"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:List_Deployable_Projects"
              },
              "output": {
                "code": 200,
                "data": "=$.Data"
              }
            }
          ],
          "schemas": {
            "output": {
              "headers": {
                "fe_metadata": "[{\"parameterName\":\"Accept\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Charset\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Accept-Encoding\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Type\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Content-Length\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Connection\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Cookie\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false},{\"parameterName\":\"Pragma\",\"type\":\"string\",\"repeating\":\"false\",\"required\":\"false\",\"visible\":false}]",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Accept\":{\"type\":\"string\",\"visible\":false},\"Accept-Charset\":{\"type\":\"string\",\"visible\":false},\"Accept-Encoding\":{\"type\":\"string\",\"visible\":false},\"Content-Type\":{\"type\":\"string\",\"visible\":false},\"Content-Length\":{\"type\":\"string\",\"visible\":false},\"Connection\":{\"type\":\"string\",\"visible\":false},\"Cookie\":{\"type\":\"string\",\"visible\":false},\"Pragma\":{\"type\":\"string\",\"visible\":false}},\"required\":[]}"
              }
            },
            "reply": {
              "data": {
                "fe_metadata": "{}",
                "type": "json",
                "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}"
              },
              "responseBody": {
                "fe_metadata": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}",
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"body\":{\"type\":\"any\"},\"headers\":{\"type\":\"object\"}}}"
              }
            }
          }
        }
      ]
    },
    {
      "id": "Timer",
      "ref": "#timer",
      "settings": {},
      "handlers": [
        {
          "name": "Ping",
          "settings": {
            "repeatInterval": "60s",
            "startDelay": ""
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Ping"
              },
              "input": {
                "Now": ""
              }
            }
          ]
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:Build_Deploy_Desc",
      "data": {
        "name": "Build Deploy Desc",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "AirK8sParameterBuilder",
            "type": "expression",
            "label": "LogMessage to AirK8sParameterBuilder",
            "value": "\"k8s\" == $flow.ServiceType"
          },
          {
            "id": 2,
            "from": "AirK8sParameterBuilder",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 3,
            "from": "LogMessage",
            "to": "AirDockerComposeParameterBuilder",
            "type": "exprOtherwise",
            "label": "LogMessage to AirDockerComposeParameterBuilder"
          },
          {
            "id": 4,
            "from": "AirDockerComposeParameterBuilder",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildDeployDesc %%%%% SyncBuild : \", coerce.toString($flow.SyncBuild), \", ProjectID : \", $flow.ProjectID, \", ComponentType : \", $flow.ComponentType, \", ServiceType : \", $flow.ServiceType, \", Name : \", $flow.Name, \", Image : \", $flow.Image, \", FlogoDescriptor : \", coerce.toString($flow.FlogoAppDescriptor), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv))"
              }
            }
          },
          {
            "id": "AirK8sParameterBuilder",
            "name": "AirK8sParameterBuilder",
            "description": "This activity build application pipeline",
            "activity": {
              "ref": "#airparameterbuilder",
              "settings": {
                "TemplateFolder": "/home/airpipeline",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"}]",
                "Properties": "[{\"Name\":\"main_apiVersion\",\"Value\":\"apps/v1\"},{\"Name\":\"main_kind\",\"Value\":\"Deployment\"},{\"Name\":\"main_metadata.name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.spec.containers[0].image\",\"Value\":\"$ImageName$\"},{\"Name\":\"main_spec.template.spec.containers[0].name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.selector.matchLabels.component\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.metadata.labels.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_apiVersion\",\"Value\":\"v1\"},{\"Name\":\"ip-service_kind\",\"Value\":\"Service\"},{\"Name\":\"ip-service_metadata.name\",\"Value\":\"$Name$-ip-service\"},{\"Name\":\"ip-service_spec.selector.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_spec.type\",\"Value\":\"LoadBalancer\"}]"
              },
              "input": {
                "ServiceType": "=$flow.ServiceType",
                "PropertyPrefix": "spec.template.spec.containers[0]",
                "Variables": {
                  "mapping": {
                    "Name": "=$flow.Name",
                    "ImageName": "=$flow.Image"
                  }
                },
                "FlogoAppDescriptor": "=$flow.FlogoAppDescriptor"
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"ImageName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Name\":\"2\",\"ImageName\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"}]"
                  },
                  "Properties": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Value\"]}}",
                    "fe_metadata": "[{\"Name\":\"main_apiVersion\",\"Value\":\"apps/v1\"},{\"Name\":\"main_kind\",\"Value\":\"Deployment\"},{\"Name\":\"main_metadata.name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.spec.containers[0].image\",\"Value\":\"$ImageName$\"},{\"Name\":\"main_spec.template.spec.containers[0].name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.selector.matchLabels.component\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.metadata.labels.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_apiVersion\",\"Value\":\"v1\"},{\"Name\":\"ip-service_kind\",\"Value\":\"Service\"},{\"Name\":\"ip-service_metadata.name\",\"Value\":\"$Name$-ip-service\"},{\"Name\":\"ip-service_spec.selector.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_spec.type\",\"Value\":\"LoadBalancer\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return01",
            "name": "Return01",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "AirService::Validate Pipeline - done",
                  "ErrorCode": 100,
                  "Descriptor": {
                    "mapping": {
                      "DeployDescriptor": "=$property[\"Descriptor.k8s\"]",
                      "properties": "=$activity[AirK8sParameterBuilder].F1Properties",
                      "FlogoDescriptor": "=$flow.FlogoAppDescriptor.flogoApp"
                    }
                  },
                  "PropertyNameDef": "=$activity[AirK8sParameterBuilder].PropertyNameDef"
                }
              }
            }
          },
          {
            "id": "AirDockerComposeParameterBuilder",
            "name": "AirDockerComposeParameterBuilder",
            "description": "This activity build application pipeline",
            "activity": {
              "ref": "#airparameterbuilder",
              "settings": {
                "TemplateFolder": "/home/airpipeline",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"ContainerName\",\"Type\":\"String\"},{\"Name\":\"BuildName\",\"Type\":\"String\"}]",
                "Properties": "[{\"Name\":\"version\",\"Value\":\"3.6\"},{\"Name\":\"services.$Name$.container_name\",\"Value\":\"$ContainerName$\"},{\"Name\":\"services.$Name$.image\",\"Value\":\"$ImageName$\"},{\"Name\":\"services.$Name$.build\",\"Value\":\"$BuildName$\"}]"
              },
              "input": {
                "ServiceType": "=$flow.ServiceType",
                "PropertyPrefix": "services.$Name$",
                "Variables": {
                  "mapping": {
                    "Name": "=$flow.Name",
                    "ImageName": "=$flow.Image",
                    "ContainerName": "=$flow.ProjectID+\"_\"+$flow.Name",
                    "BuildName": "=f1.ternary(null==$flow.Image||\"\"==$flow.Image, \"001\", \"\")"
                  }
                },
                "FlogoAppDescriptor": "=$flow.FlogoAppDescriptor"
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"ImageName\":{\"type\":\"string\"},\"ContainerName\":{\"type\":\"string\"},\"BuildName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Name\":\"2\",\"ImageName\":\"2\",\"ContainerName\":\"2\",\"BuildName\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"ContainerName\",\"Type\":\"String\"},{\"Name\":\"BuildName\",\"Type\":\"String\"}]"
                  },
                  "Properties": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Value\"]}}",
                    "fe_metadata": "[{\"Name\":\"version\",\"Value\":\"3.6\"},{\"Name\":\"services.$Name$.container_name\",\"Value\":\"$ContainerName$\"},{\"Name\":\"services.$Name$.image\",\"Value\":\"$ImageName$\"},{\"Name\":\"services.$Name$.build\",\"Value\":\"$BuildName$\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return02",
            "name": "Return02",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "AirService::Validate Pipeline - done",
                  "ErrorCode": 100,
                  "Descriptor": {
                    "mapping": {
                      "DeployDescriptor": "=$property[\"Descriptor.docker\"]",
                      "properties": "=$activity[AirDockerComposeParameterBuilder].F1Properties",
                      "FlogoDescriptor": "=$flow.FlogoAppDescriptor.flogoApp"
                    }
                  },
                  "PropertyNameDef": "=$activity[AirDockerComposeParameterBuilder].PropertyNameDef"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return03",
              "name": "Return03",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"AirService::Validate Pipeline - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "SyncBuild",
              "type": "boolean"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "ServiceType",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "FlogoAppDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
              }
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            },
            {
              "name": "Descriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}"
              }
            },
            {
              "name": "PropertyNameDef",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"FlogoAppDescriptor\":{\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"},\"Descriptor\":{\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}},\"PropertyNameDef\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        }
      }
    },
    {
      "id": "flow:Create_Domain",
      "data": {
        "name": "Create Domain",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "CreateProjectFolder",
            "type": "default"
          },
          {
            "id": 3,
            "from": "CreateProjectFolder",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:Create Domain %%%%% Domain : \", $flow.ProjectID)"
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectTemplateFolderInt\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectsFolderInt": "=$property[\"System.BaseFolder\"] + \"/projects\"",
                    "SkipCondition": "=false",
                    "ProjectTemplateFolderInt": "/home/project"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectsFolderInt\":{\"type\":\"string\"},\"ProjectTemplateFolderInt\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectsFolderInt\":\"2\",\"ProjectTemplateFolderInt\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectsFolderInt\":{\"type\":\"string\"},\"ProjectTemplateFolderInt\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectsFolderInt\":\"2\",\"ProjectTemplateFolderInt\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectTemplateFolderInt\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "CreateProjectFolder",
            "name": "CreateProjectFolder",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$ProjectsFolder$",
                "numOfExecutions": 5,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectTemplateFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": "=false",
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "mkdir -p $ProjectID$",
                      "Execution_1": "cp -R $ProjectTemplateFolder$/artifacts ./$ProjectID$",
                      "Execution_2": "cp -R $ProjectTemplateFolder$/build ./$ProjectID$",
                      "Execution_3": "cp -R $ProjectTemplateFolder$/deploy ./$ProjectID$",
                      "Execution_4": "cp -R $ProjectTemplateFolder$/pipeline ./$ProjectID$"
                    }
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectsFolder": "=$activity[SetSystemProperties].Data.ProjectsFolderInt",
                    "ProjectID": "=$flow.ProjectID",
                    "ProjectTemplateFolder": "=$activity[SetSystemProperties].Data.ProjectTemplateFolderInt"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"},\"Execution_2\":{\"type\":\"string\"},\"Execution_3\":{\"type\":\"string\"},\"Execution_4\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\",\"Execution_2\":\"2\",\"Execution_3\":\"2\",\"Execution_4\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectsFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ProjectTemplateFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectsFolder\":\"2\",\"ProjectID\":\"2\",\"ProjectTemplateFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectTemplateFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[CreateProjectFolder].Success",
                  "Message": "=coerce.toString($activity[CreateProjectFolder].Message)",
                  "ErrorCode": "=$activity[CreateProjectFolder].ErrorCode"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return3",
              "name": "Return3",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"AirService::Initialize - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "Properties",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Get_Pipelines",
      "data": {
        "name": "Get Pipelines",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "ObjectMaker",
            "type": "default"
          },
          {
            "id": 3,
            "from": "ObjectMaker",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildAirPipeline %%%%% ProjectID : \", $flow.ProjectID)"
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectFiles\",\"Type\":\"Array\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectFiles": "=f1.getfstruct($property[\"System.BaseFolder\"] + \"/projects/\" + $flow.ProjectID + \"/build\", 0)"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFiles\":{\"type\":\"array\",\"items\":{}},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectFiles\":[],\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFiles\":{\"type\":\"array\",\"items\":{}}}}",
                    "fe_metadata": "{\"ProjectFiles\":[]}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ProjectFiles\",\"Type\":\"Array\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ObjectMaker",
            "name": "ObjectMaker",
            "description": "Make an New Object",
            "activity": {
              "ref": "#objectmaker",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"pipelines\":{\"type\":\"array\",\"items\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"pipelines\" : []\n}"
                }
              },
              "input": {
                "ObjectDataMapping": {
                  "mapping": {
                    "pipelines": "=$activity[SetSystemProperties].Data.ProjectFiles"
                  }
                }
              },
              "schemas": {
                "input": {
                  "ObjectDataMapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"pipelines\":{\"type\":\"array\",\"items\":{}}}}",
                    "fe_metadata": "{\"pipelines\":[]}"
                  }
                },
                "output": {
                  "ObjectOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"pipelines\":{\"type\":\"array\",\"items\":{}}}}",
                    "fe_metadata": "{\"pipelines\":[]}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"pipelines\":{\"type\":\"array\",\"items\":{}}}}",
                    "fe_metadata": "{\n    \"pipelines\" : []\n}"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": true,
                  "Message": "=coerce.toString($activity[ObjectMaker].ObjectOut)",
                  "ErrorCode": 100
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return3",
              "name": "Return3",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"AirService::Get Pipelines - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Delete_Pipeline",
      "data": {
        "name": "Delete Pipeline",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "DeletePipeline",
            "type": "default"
          },
          {
            "id": 3,
            "from": "DeletePipeline",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:DeletePipeline %%%%% ProjectID : \", $flow.ProjectID, \", Name : \", $flow.Name)"
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "SkipCondition": "=false",
                    "ProjectFolderInt": "=$property[\"System.BaseFolder\"] + \"/projects/\" + $flow.ProjectID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderInt\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectFolderInt\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolderInt\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectFolderInt\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "DeletePipeline",
            "name": "DeletePipeline",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$ProjectFolder$",
                "numOfExecutions": 3,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"Name\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_2": "=\"rm -f $ProjectFolder$/pipeline/\" + $flow.Name + \"_property_name_def.json\"",
                      "Execution_1": "=\"rm -f $ProjectFolder$/pipeline/\" + $flow.Name + \".json\"",
                      "Execution_0": "rm -rf $ProjectFolder$/build/$Name$"
                    }
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectFolder": "=$activity[SetSystemProperties].Data.ProjectFolderInt",
                    "Name": "=$flow.Name"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"},\"Execution_2\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\",\"Execution_2\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectFolder\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectFolder\":\"2\",\"Name\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"Name\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[DeletePipeline].Success",
                  "Message": "=coerce.toString($activity[DeletePipeline].Message)",
                  "ErrorCode": "=$activity[DeletePipeline].ErrorCode"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return3",
              "name": "Return3",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"AirService::Delete Pipeline - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Build_Pipeline",
      "data": {
        "name": "Build Pipeline",
        "description": "**** Build pipeline and Flogo app ****",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "NoAppDescriptor",
            "type": "expression",
            "label": "SetSystemPropertiestoBuildFlogoPipeline",
            "value": "true==f1.isnull($flow.AirDescriptor) \u0026\u0026 true==f1.isnull($flow.FlogoAppDescriptor)"
          },
          {
            "id": 3,
            "from": "SetSystemProperties",
            "to": "LogBuildFlogoPipeline",
            "type": "expression",
            "label": "SetSystemPropertiesto",
            "value": "false==f1.isempty($flow.AirDescriptor)"
          },
          {
            "id": 4,
            "from": "LogBuildFlogoPipeline",
            "to": "BuildFlogoPipeline",
            "type": "expression",
            "label": "LogMessage1toBuildFlogoPipeline",
            "value": "true!=f1.isempty($flow.AirDescriptor.source)"
          },
          {
            "id": 5,
            "from": "BuildFlogoPipeline",
            "to": "PropertyNameDef",
            "type": "default"
          },
          {
            "id": 6,
            "from": "PropertyNameDef",
            "to": "WriterPropertyNameDef",
            "type": "default"
          },
          {
            "id": 7,
            "from": "WriterPropertyNameDef",
            "to": "FlogoDescriptorStr",
            "type": "default"
          },
          {
            "id": 8,
            "from": "FlogoDescriptorStr",
            "to": "BuildRequest",
            "type": "default"
          },
          {
            "id": 9,
            "from": "BuildRequest",
            "to": "ToBuilder",
            "type": "default"
          },
          {
            "id": 10,
            "from": "ToBuilder",
            "to": "Return_Air",
            "type": "default"
          },
          {
            "id": 11,
            "from": "LogBuildFlogoPipeline",
            "to": "LogBuildFlogoAppLogNew",
            "type": "exprOtherwise",
            "label": "LogMessage1 to CopyOfCopyOfLogMessage1"
          },
          {
            "id": 12,
            "from": "LogBuildFlogoAppLogNew",
            "to": "BuildFlogoAppNew",
            "type": "default"
          },
          {
            "id": 13,
            "from": "BuildFlogoAppNew",
            "to": "PropertyNameDef1",
            "type": "default"
          },
          {
            "id": 14,
            "from": "PropertyNameDef1",
            "to": "WriterPropertyNameDef1",
            "type": "default"
          },
          {
            "id": 15,
            "from": "WriterPropertyNameDef1",
            "to": "FlogoDescriptorStr1",
            "type": "default"
          },
          {
            "id": 16,
            "from": "FlogoDescriptorStr1",
            "to": "BuildRequest1",
            "type": "default"
          },
          {
            "id": 17,
            "from": "BuildRequest1",
            "to": "ToBuilder1",
            "type": "default"
          },
          {
            "id": 18,
            "from": "ToBuilder1",
            "to": "Return_Flogo1",
            "type": "default"
          },
          {
            "id": 19,
            "from": "SetSystemProperties",
            "to": "LogBuildFlogoApp",
            "type": "exprOtherwise",
            "label": "SetSystemProperties to CopyOfReturn2"
          },
          {
            "id": 20,
            "from": "LogBuildFlogoApp",
            "to": "BuildFlogoApp",
            "type": "default"
          },
          {
            "id": 21,
            "from": "BuildFlogoApp",
            "to": "PropertyNameDef2",
            "type": "default"
          },
          {
            "id": 22,
            "from": "PropertyNameDef2",
            "to": "WriterPropertyNameDef2",
            "type": "default"
          },
          {
            "id": 23,
            "from": "WriterPropertyNameDef2",
            "to": "FlogoDescriptorStr2",
            "type": "default"
          },
          {
            "id": 24,
            "from": "FlogoDescriptorStr2",
            "to": "BuildRequest2",
            "type": "default"
          },
          {
            "id": 25,
            "from": "BuildRequest2",
            "to": "ToBuilder2",
            "type": "default"
          },
          {
            "id": 26,
            "from": "ToBuilder2",
            "to": "Return_Flogo",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildAirPipeline %%%%% SyncBuild : \", coerce.toString($flow.SyncBuild), \", ProjectID : \", $flow.ProjectID, \", ComponentType : \", $flow.ComponentType, \", ServiceType : \", $flow.ServiceType, \", Name : \", $flow.Name, \", Image : \", $flow.Image, \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv))"
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"},{\"Name\":\"BuilderURL\",\"Type\":\"String\"},{\"Name\":\"PropertyNameDefFilename\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "SyncBuild": "=f1.ternary(null!=$flow.SyncBuild, $flow.SyncBuild, false)",
                    "SkipCondition": "=false",
                    "BuilderURL": "=$property[\"System.BuilderURL\"]",
                    "PropertyNameDefFilename": "=$property[\"System.BaseFolder\"] + \"/projects/\" + $flow.ProjectID + \"/pipeline/\" + $flow.Name + \"_property_name_def.json\""
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"BuilderURL\":{\"type\":\"string\"},\"PropertyNameDefFilename\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"SyncBuild\":true,\"BuilderURL\":\"2\",\"PropertyNameDefFilename\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"BuilderURL\":{\"type\":\"string\"},\"PropertyNameDefFilename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"SyncBuild\":true,\"BuilderURL\":\"2\",\"PropertyNameDefFilename\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"},{\"Name\":\"BuilderURL\",\"Type\":\"String\"},{\"Name\":\"PropertyNameDefFilename\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "NoAppDescriptor",
            "name": "NoAppDescriptor",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=false",
                  "Message": "No Application Descriptor!",
                  "ErrorCode": 300
                }
              }
            }
          },
          {
            "id": "LogBuildFlogoPipeline",
            "name": "LogBuildFlogoPipeline",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildAirPipeline %%%%% true!=f1.isempty($flow.AirDescriptor) : \", coerce.toString(true!=f1.isempty($flow.AirDescriptor)), \", true!=f1.isempty($flow.AirDescriptor.source) : \", coerce.toString(true!=f1.isempty($flow.AirDescriptor.source)))"
              }
            }
          },
          {
            "id": "BuildFlogoPipeline",
            "name": "BuildFlogoPipeline",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Validate_an_Air_descriptor_then_convert_to_Flogo_descriptor_2"
              },
              "input": {
                "ProjectID": "=$flow.ProjectID",
                "SyncBuild": "=$flow.SyncBuild",
                "ComponentType": "=$flow.ComponentType",
                "ServiceType": "=$flow.ServiceType",
                "Name": "=$flow.Name",
                "Image": "=$flow.Image",
                "AirDescriptor": "=$flow.AirDescriptor",
                "ScriptSystemEnv": "=$flow.ScriptSystemEnv"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"},\"Descriptor\":{\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}},\"PropertyNameDef\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"},\"Descriptor\":{\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}},\"PropertyNameDef\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "PropertyNameDef",
            "name": "PropertyNameDef",
            "description": "Object Serializer Activity",
            "activity": {
              "ref": "#objectserializer",
              "settings": {
                "StringFormat": "json"
              },
              "input": {
                "Data": "=$activity[BuildFlogoPipeline].PropertyNameDef"
              },
              "schemas": {
                "input": {
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "WriterPropertyNameDef",
            "name": "WriterPropertyNameDef",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "$Filename$",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": false,
                "Data": {
                  "mapping": {
                    "Input": "=$activity[PropertyNameDef].SerializedString"
                  }
                },
                "Variables": {
                  "mapping": {
                    "Filename": "=$activity[SetSystemProperties].Data.PropertyNameDefFilename"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Input\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Input\":\"2\"}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "FlogoDescriptorStr",
            "name": "FlogoDescriptorStr",
            "description": "Object Serializer Activity",
            "activity": {
              "ref": "#objectserializer",
              "settings": {
                "StringFormat": "json"
              },
              "input": {
                "Data": "=$activity[BuildFlogoPipeline].Descriptor"
              },
              "schemas": {
                "input": {
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRequest",
            "name": "BuildRequest",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}",
                  "value": "",
                  "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "DeployDescriptor": {
                      "Type": "=$flow.ServiceType",
                      "Name": "=$flow.Name",
                      "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                      "Properties": "=$activity[BuildFlogoPipeline].Descriptor.properties",
                      "Descriptor": "=$activity[BuildFlogoPipeline].Descriptor.DeployDescriptor"
                    },
                    "SyncBuild": "=$activity[SetSystemProperties].Data.SyncBuild",
                    "ComponentType": "=$flow.ComponentType",
                    "Image": "=$flow.Image",
                    "AppDescriptor": "=$activity[BuildFlogoPipeline].Descriptor.FlogoDescriptor"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}",
                    "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}"
                  }
                }
              }
            }
          },
          {
            "id": "ToBuilder",
            "name": "ToBuilder",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "10000",
                "urlMappingString": "",
                "urlMapping": "",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"}]",
                "httpHeaders": ""
              },
              "input": {
                "URL": "=$activity[SetSystemProperties].Data.BuilderURL + \"/f1/builder/buildDocker/$ProjectID$\"",
                "Body": "=$activity[BuildRequest].JSONString",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID"
                  }
                }
              },
              "output": {
                "Success": false
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return_Air",
            "name": "Return_Air",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[ToBuilder].Success",
                  "Message": "=$activity[ToBuilder].Data",
                  "ErrorCode": "=$activity[ToBuilder].ErrorCode"
                }
              }
            }
          },
          {
            "id": "LogBuildFlogoAppLogNew",
            "name": "LogBuildFlogoAppLogNew",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildFlogoAppNew %%%%% AirDescriptor : \", coerce.toString($flow.AirDescriptor), \", true!=f1.isempty($flow.FlogoAppDescriptor) : \", coerce.toString(true!=f1.isempty($flow.FlogoAppDescriptor)))"
              }
            }
          },
          {
            "id": "BuildFlogoAppNew",
            "name": "BuildFlogoAppNew",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Build_Deploy_Desc"
              },
              "input": {
                "ProjectID": "=$flow.ProjectID",
                "SyncBuild": "=$flow.SyncBuild",
                "ComponentType": "=$flow.ComponentType",
                "ServiceType": "=$flow.ServiceType",
                "Name": "=$flow.Name",
                "Image": "=$flow.Image",
                "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                "FlogoAppDescriptor": {
                  "mapping": {
                    "properties": "=$flow.AirDescriptor.logic[0].properties",
                    "extra": "=$flow.AirDescriptor.extra",
                    "flogoApp": "=$flow.AirDescriptor.logic[0].flogoApp",
                    "ports": "=$flow.AirDescriptor.logic[0].ports"
                  }
                }
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"FlogoAppDescriptor\":{\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"FlogoAppDescriptor\":{\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"},\"Descriptor\":{\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}},\"PropertyNameDef\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"},\"Descriptor\":{\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}},\"PropertyNameDef\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "PropertyNameDef1",
            "name": "PropertyNameDef1",
            "description": "Object Serializer Activity",
            "activity": {
              "ref": "#objectserializer",
              "settings": {
                "StringFormat": "json",
                "PassThrough": ""
              },
              "input": {
                "Data": "=$activity[BuildFlogoAppNew].PropertyNameDef"
              },
              "schemas": {
                "input": {
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "WriterPropertyNameDef1",
            "name": "WriterPropertyNameDef1",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "$Filename$",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": false,
                "Data": {
                  "mapping": {
                    "Input": "=$activity[PropertyNameDef1].SerializedString"
                  }
                },
                "Variables": {
                  "mapping": {
                    "Filename": "=$activity[SetSystemProperties].Data.PropertyNameDefFilename"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Input\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Input\":\"2\"}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "FlogoDescriptorStr1",
            "name": "FlogoDescriptorStr1",
            "description": "Object Serializer Activity",
            "activity": {
              "ref": "#objectserializer",
              "settings": {
                "StringFormat": "json",
                "PassThrough": ""
              },
              "input": {
                "Data": "=$activity[BuildFlogoAppNew].Descriptor"
              },
              "schemas": {
                "input": {
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRequest1",
            "name": "BuildRequest1",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}",
                  "value": "",
                  "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "DeployDescriptor": {
                      "Type": "=$flow.ServiceType",
                      "Name": "=$flow.Name",
                      "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                      "Properties": "=$activity[BuildFlogoAppNew].Descriptor.properties",
                      "Descriptor": "=$activity[BuildFlogoAppNew].Descriptor.DeployDescriptor"
                    },
                    "SyncBuild": "=$activity[SetSystemProperties].Data.SyncBuild",
                    "ComponentType": "=$flow.ComponentType",
                    "Image": "=$flow.Image",
                    "AppDescriptor": "=$activity[BuildFlogoAppNew].Descriptor.FlogoDescriptor"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}",
                    "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}"
                  }
                }
              }
            }
          },
          {
            "id": "ToBuilder1",
            "name": "ToBuilder1",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "10000",
                "urlMappingString": "",
                "urlMapping": "",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"}]",
                "httpHeaders": ""
              },
              "input": {
                "URL": "=$activity[SetSystemProperties].Data.BuilderURL + \"/f1/builder/buildDocker/$ProjectID$\"",
                "Body": "=$activity[BuildRequest1].JSONString",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID"
                  }
                }
              },
              "output": {
                "Success": false
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return_Flogo1",
            "name": "Return_Flogo1",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[ToBuilder1].Success",
                  "Message": "=$activity[ToBuilder1].Data",
                  "ErrorCode": "=$activity[ToBuilder1].ErrorCode"
                }
              }
            }
          },
          {
            "id": "LogBuildFlogoApp",
            "name": "LogBuildFlogoApp",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildFlogoApp %%%%% AirDescriptor : \", coerce.toString($flow.AirDescriptor), \", true!=f1.isempty($flow.FlogoAppDescriptor) : \", coerce.toString(true!=f1.isempty($flow.FlogoAppDescriptor)))"
              }
            }
          },
          {
            "id": "BuildFlogoApp",
            "name": "BuildFlogoApp",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Build_Deploy_Desc"
              },
              "input": {
                "ProjectID": "=$flow.ProjectID",
                "SyncBuild": "=$flow.SyncBuild",
                "ComponentType": "=$flow.ComponentType",
                "ServiceType": "=$flow.ServiceType",
                "Name": "=$flow.Name",
                "Image": "=$flow.Image",
                "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                "FlogoAppDescriptor": "=$flow.FlogoAppDescriptor"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"FlogoAppDescriptor\":{\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"FlogoAppDescriptor\":{\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"},\"Descriptor\":{\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}},\"PropertyNameDef\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"},\"Descriptor\":{\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}},\"PropertyNameDef\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "PropertyNameDef2",
            "name": "PropertyNameDef2",
            "description": "Object Serializer Activity",
            "activity": {
              "ref": "#objectserializer",
              "settings": {
                "StringFormat": "json",
                "PassThrough": ""
              },
              "input": {
                "Data": "=$activity[BuildFlogoApp].PropertyNameDef"
              },
              "schemas": {
                "input": {
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "WriterPropertyNameDef2",
            "name": "WriterPropertyNameDef2",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "$Filename$",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": false,
                "Data": {
                  "mapping": {
                    "Input": "=$activity[PropertyNameDef2].SerializedString"
                  }
                },
                "Variables": {
                  "mapping": {
                    "Filename": "=$activity[SetSystemProperties].Data.PropertyNameDefFilename"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Input\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Input\":\"2\"}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Filename\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"Filename\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "FlogoDescriptorStr2",
            "name": "FlogoDescriptorStr2",
            "description": "Object Serializer Activity",
            "activity": {
              "ref": "#objectserializer",
              "settings": {
                "StringFormat": "json",
                "PassThrough": ""
              },
              "input": {
                "Data": "=$activity[BuildFlogoApp].Descriptor"
              },
              "schemas": {
                "input": {
                  "PassThroughData": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                },
                "output": {
                  "PassThroughDataOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRequest2",
            "name": "BuildRequest2",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}",
                  "value": "",
                  "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "DeployDescriptor": {
                      "Type": "=$flow.ServiceType",
                      "Name": "=$flow.Name",
                      "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                      "Properties": "=$activity[BuildFlogoApp].Descriptor.properties",
                      "Descriptor": "=$activity[BuildFlogoApp].Descriptor.DeployDescriptor"
                    },
                    "SyncBuild": "=$activity[SetSystemProperties].Data.SyncBuild",
                    "ComponentType": "=$flow.ComponentType",
                    "Image": "=$flow.Image",
                    "AppDescriptor": "=$activity[BuildFlogoApp].Descriptor.FlogoDescriptor"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"Runner\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AppDescriptor\":{\"type\":\"string\"},\"DeployDescriptor\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}",
                    "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncBuild\": {\n            \"type\": \"boolean\"\n        },\n        \"ComponentType\": {\n            \"type\": \"string\"\n        },\n        \"Runner\": {\n            \"type\": \"string\"\n        },\n        \"Image\": {\n            \"type\": \"string\"\n        },\n        \"AppDescriptor\": {\n            \"type\": \"string\"\n        },\n        \"DeployDescriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Name\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                },\n                \"Variables\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"ScriptSystemEnv\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                },\n                \"Properties\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Group\": {\n                                \"type\": \"string\"\n                            },\n                            \"Value\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                }\n            }\n        }\n    }\n}"
                  }
                }
              }
            }
          },
          {
            "id": "ToBuilder2",
            "name": "ToBuilder2",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "10000",
                "urlMappingString": "",
                "urlMapping": "",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"}]",
                "httpHeaders": ""
              },
              "input": {
                "URL": "=$activity[SetSystemProperties].Data.BuilderURL + \"/f1/builder/buildDocker/$ProjectID$\"",
                "Body": "=$activity[BuildRequest2].JSONString",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID"
                  }
                }
              },
              "output": {
                "Success": false
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return_Flogo",
            "name": "Return_Flogo",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[ToBuilder2].Success",
                  "Message": "=$activity[ToBuilder2].Data",
                  "ErrorCode": "=$activity[ToBuilder2].ErrorCode"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return_Error",
              "name": "Return_Error",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"AirService::Build Pipeline - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "SyncBuild",
              "type": "boolean"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "ServiceType",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "AirDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
              }
            },
            {
              "name": "FlogoAppDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
              }
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"FlogoAppDescriptor\":{\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Validate_an_Air_descriptor_then_convert_to_Flogo_descriptor_2",
      "data": {
        "name": "Validate an Air descriptor then convert to Flogo descriptor 2",
        "description": "**** Only for flogo pipeline ****",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "BuildAppPipelineDescriptor",
            "type": "expression",
            "label": "LogMessagetoBuildAppPipelineDescriptor",
            "value": "null==$flow.AirDescriptor.logic[0].flogoApp"
          },
          {
            "id": 2,
            "from": "BuildAppPipelineDescriptor",
            "to": "AirK8sAppBuilder",
            "type": "expression",
            "label": "BuildAppPipelineDescriptortoAirApplicationBuilder",
            "value": "\"k8s\" == $flow.ServiceType"
          },
          {
            "id": 3,
            "from": "AirK8sAppBuilder",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 4,
            "from": "BuildAppPipelineDescriptor",
            "to": "AirDockerComposeAppBuilder",
            "type": "exprOtherwise",
            "label": "BuildAppPipelineDescriptor to CopyOfAirApplicationBuilder"
          },
          {
            "id": 5,
            "from": "AirDockerComposeAppBuilder",
            "to": "Return02",
            "type": "default"
          },
          {
            "id": 6,
            "from": "LogMessage",
            "to": "Return04",
            "type": "exprOtherwise",
            "label": "LogMessagetoReturn04"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:ValidateAirPipeline %%%%% SyncBuild : \", coerce.toString($flow.SyncBuild), \", ProjectID : \", $flow.ProjectID, \", ComponentType : \", $flow.ComponentType, \", ServiceType : \", $flow.ServiceType, \", Name : \", $flow.Name, \", Image : \", $flow.Image, \", AirDescriptor : \", coerce.toString($flow.AirDescriptor), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv))"
              }
            }
          },
          {
            "id": "BuildAppPipelineDescriptor",
            "name": "BuildAppPipelineDescriptor",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}",
                  "value": "",
                  "fe_metadata": "{\n        \"properties\" : [\n            {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n        ],\n        \"source\": {\n            \"name\": \"MQTT\",\n            \"properties\" : [\n                {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n            ],\n            \"ports\" : [\n                \"8080:9999\"\n            ]\n        },\n        \"logic\": [\n            {\n                \"name\": \"Dgraph\",\n                \"properties\" : [\n                    {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n                ]\n            }\n        ],\n        \"extra\": [\n            {\"Name\":\"networks.default.external.name\", \"Value\":\"edgex-network\"}\n        ]\n    }"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "source": {
                      "name": "=$flow.AirDescriptor.source.name",
                      "properties": "=$flow.AirDescriptor.source.properties",
                      "ports": "=$flow.AirDescriptor.source.ports"
                    },
                    "logic": "=$flow.AirDescriptor.logic",
                    "properties": "=array.create(f1.json2object(\"{\\\"Name\\\":\\\"FLOGO_APP_PROPS_ENV\\\",\\\"Value\\\":\\\"auto\\\"}\"))",
                    "extra": "=$flow.AirDescriptor.extra"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}",
                    "fe_metadata": "{\"properties\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}],\"source\":{\"name\":\"MQTT\",\"properties\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}],\"ports\":[\"8080:9999\"]},\"logic\":[{\"name\":\"Dgraph\",\"properties\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}]}],\"extra\":[{\"Name\":\"networks.default.external.name\",\"Value\":\"edgex-network\"}]}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}",
                    "fe_metadata": "{\n        \"properties\" : [\n            {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n        ],\n        \"source\": {\n            \"name\": \"MQTT\",\n            \"properties\" : [\n                {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n            ],\n            \"ports\" : [\n                \"8080:9999\"\n            ]\n        },\n        \"logic\": [\n            {\n                \"name\": \"Dgraph\",\n                \"properties\" : [\n                    {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n                ]\n            }\n        ],\n        \"extra\": [\n            {\"Name\":\"networks.default.external.name\", \"Value\":\"edgex-network\"}\n        ]\n    }"
                  }
                }
              }
            }
          },
          {
            "id": "AirK8sAppBuilder",
            "name": "AirK8sAppBuilder",
            "description": "This activity build model inference pipeline",
            "activity": {
              "ref": "#airpipelinebuilder",
              "settings": {
                "TemplateFolder": "/home/airpipeline",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"}]",
                "Properties": "[{\"Name\":\"main_apiVersion\",\"Value\":\"apps/v1\"},{\"Name\":\"main_kind\",\"Value\":\"Deployment\"},{\"Name\":\"main_metadata.name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.spec.containers[0].image\",\"Value\":\"$ImageName$\"},{\"Name\":\"main_spec.template.spec.containers[0].name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.selector.matchLabels.component\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.metadata.labels.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_apiVersion\",\"Value\":\"v1\"},{\"Name\":\"ip-service_kind\",\"Value\":\"Service\"},{\"Name\":\"ip-service_metadata.name\",\"Value\":\"$Name$-ip-service\"},{\"Name\":\"ip-service_spec.selector.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_spec.type\",\"Value\":\"LoadBalancer\"}]"
              },
              "input": {
                "ApplicationName": "=$flow.Name",
                "AirDescriptor": "=$activity[BuildAppPipelineDescriptor].JSONString",
                "ServiceType": "=$flow.ServiceType",
                "PropertyPrefix": "spec.template.spec.containers[0]",
                "Variables": {
                  "mapping": {
                    "Name": "=$flow.Name",
                    "ImageName": "=$flow.Image"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"ImageName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Name\":\"2\",\"ImageName\":\"2\"}"
                  }
                },
                "output": {
                  "Descriptor": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"F1Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"FlogoDescriptor\":\"2\",\"F1Properties\":[{\"Group\":\"2\",\"Value\":{\"Name\":\"2\",\"Value\":\"2\",\"Type\":\"2\"}}]}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"}]"
                  },
                  "Properties": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Value\"]}}",
                    "fe_metadata": "[{\"Name\":\"main_apiVersion\",\"Value\":\"apps/v1\"},{\"Name\":\"main_kind\",\"Value\":\"Deployment\"},{\"Name\":\"main_metadata.name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.spec.containers[0].image\",\"Value\":\"$ImageName$\"},{\"Name\":\"main_spec.template.spec.containers[0].name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.selector.matchLabels.component\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.metadata.labels.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_apiVersion\",\"Value\":\"v1\"},{\"Name\":\"ip-service_kind\",\"Value\":\"Service\"},{\"Name\":\"ip-service_metadata.name\",\"Value\":\"$Name$-ip-service\"},{\"Name\":\"ip-service_spec.selector.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_spec.type\",\"Value\":\"LoadBalancer\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return01",
            "name": "Return01",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "AirService::Validate Pipeline - done",
                  "ErrorCode": 100,
                  "PropertyNameDef": "=$activity[AirK8sAppBuilder].PropertyNameDef",
                  "Descriptor": {
                    "mapping": {
                      "FlogoDescriptor": "=$activity[AirK8sAppBuilder].Descriptor.FlogoDescriptor",
                      "properties": "=$activity[AirK8sAppBuilder].Descriptor.F1Properties",
                      "DeployDescriptor": "=$property[\"Descriptor.k8s\"]"
                    }
                  }
                }
              }
            }
          },
          {
            "id": "AirDockerComposeAppBuilder",
            "name": "AirDockerComposeAppBuilder",
            "description": "This activity build model inference pipeline",
            "activity": {
              "ref": "#airpipelinebuilder",
              "settings": {
                "TemplateFolder": "/home/airpipeline",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ContainerName\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"BuildName\",\"Type\":\"String\"}]",
                "Properties": "[{\"Name\":\"version\",\"Value\":\"3.6\"},{\"Name\":\"services.$Name$.container_name\",\"Value\":\"$ContainerName$\"},{\"Name\":\"services.$Name$.image\",\"Value\":\"$ImageName$\"},{\"Name\":\"services.$Name$.build\",\"Value\":\"$BuildName$\"}]"
              },
              "input": {
                "ApplicationName": "=$flow.Name",
                "AirDescriptor": "=$activity[BuildAppPipelineDescriptor].JSONString",
                "ServiceType": "=$flow.ServiceType",
                "PropertyPrefix": "services.$Name$",
                "Variables": {
                  "mapping": {
                    "Name": "=$flow.Name",
                    "ContainerName": "=$flow.ProjectID+\"_\"+$flow.Name",
                    "ImageName": "=$flow.Image",
                    "BuildName": "=f1.ternary(null==$flow.Image||\"\"==$flow.Image, \"001\", \"\")"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"ContainerName\":{\"type\":\"string\"},\"ImageName\":{\"type\":\"string\"},\"BuildName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Name\":\"2\",\"ContainerName\":\"2\",\"ImageName\":\"2\",\"BuildName\":\"2\"}"
                  }
                },
                "output": {
                  "Descriptor": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"F1Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"FlogoDescriptor\":\"2\",\"F1Properties\":[{\"Group\":\"2\",\"Value\":{\"Name\":\"2\",\"Value\":\"2\",\"Type\":\"2\"}}]}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ContainerName\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"BuildName\",\"Type\":\"String\"}]"
                  },
                  "Properties": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Value\"]}}",
                    "fe_metadata": "[{\"Name\":\"version\",\"Value\":\"3.6\"},{\"Name\":\"services.$Name$.container_name\",\"Value\":\"$ContainerName$\"},{\"Name\":\"services.$Name$.image\",\"Value\":\"$ImageName$\"},{\"Name\":\"services.$Name$.build\",\"Value\":\"$BuildName$\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return02",
            "name": "Return02",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "AirService::Validate Pipeline - done",
                  "ErrorCode": 100,
                  "PropertyNameDef": "=$activity[AirDockerComposeAppBuilder].PropertyNameDef",
                  "Descriptor": {
                    "mapping": {
                      "FlogoDescriptor": "=$activity[AirDockerComposeAppBuilder].Descriptor.FlogoDescriptor",
                      "properties": "=$activity[AirDockerComposeAppBuilder].Descriptor.F1Properties",
                      "DeployDescriptor": "=$property[\"Descriptor.docker\"]"
                    }
                  }
                }
              }
            }
          },
          {
            "id": "Return04",
            "name": "Return04",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "AirService::Skip validate Flogo app - done",
                  "ErrorCode": 100,
                  "Descriptor": {
                    "mapping": {
                      "FlogoDescriptor": "=$flow.AirDescriptor.logic[0].flogoApp"
                    }
                  }
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return03",
              "name": "Return03",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"AirService::Validate Pipeline - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "SyncBuild",
              "type": "boolean"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "ServiceType",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "AirDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
              }
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            },
            {
              "name": "Descriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}"
              }
            },
            {
              "name": "PropertyNameDef",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"},\"Descriptor\":{\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}},\"PropertyNameDef\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        }
      }
    },
    {
      "id": "flow:Validate_an_Air_descriptor_then_convert_to_Flogo_descriptor",
      "data": {
        "name": "Validate an Air descriptor then convert to Flogo descriptor",
        "description": "**** Only for flogo pipeline ****",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "BuildAppPipelineDescriptor",
            "type": "expression",
            "label": "LogMessagetoBuildAppPipelineDescriptor",
            "value": "null==$flow.AirDescriptor.logic[0].flogoApp"
          },
          {
            "id": 2,
            "from": "BuildAppPipelineDescriptor",
            "to": "AirK8sAppBuilder",
            "type": "expression",
            "label": "BuildAppPipelineDescriptortoAirApplicationBuilder",
            "value": "\"k8s\" == $flow.ServiceType"
          },
          {
            "id": 3,
            "from": "AirK8sAppBuilder",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 4,
            "from": "BuildAppPipelineDescriptor",
            "to": "AirDockerComposeAppBuilder",
            "type": "exprOtherwise",
            "label": "BuildAppPipelineDescriptor to CopyOfAirApplicationBuilder"
          },
          {
            "id": 5,
            "from": "AirDockerComposeAppBuilder",
            "to": "Return02",
            "type": "default"
          },
          {
            "id": 6,
            "from": "LogMessage",
            "to": "Return04",
            "type": "exprOtherwise",
            "label": "LogMessagetoReturn04"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:ValidateAirPipeline %%%%% SyncBuild : \", coerce.toString($flow.SyncBuild), \", ProjectID : \", $flow.ProjectID, \", ComponentType : \", $flow.ComponentType, \", ServiceType : \", $flow.ServiceType, \", Name : \", $flow.Name, \", Image : \", $flow.Image, \", AirDescriptor : \", coerce.toString($flow.AirDescriptor), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv))"
              }
            }
          },
          {
            "id": "BuildAppPipelineDescriptor",
            "name": "BuildAppPipelineDescriptor",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}",
                  "value": "",
                  "fe_metadata": "{\n        \"properties\" : [\n            {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n        ],\n        \"source\": {\n            \"name\": \"MQTT\",\n            \"properties\" : [\n                {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n            ],\n            \"ports\" : [\n                \"8080:9999\"\n            ]\n        },\n        \"logic\": [\n            {\n                \"name\": \"Dgraph\",\n                \"properties\" : [\n                    {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n                ]\n            }\n        ],\n        \"extra\": [\n            {\"Name\":\"networks.default.external.name\", \"Value\":\"edgex-network\"}\n        ]\n    }"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "source": {
                      "name": "=$flow.AirDescriptor.source.name",
                      "properties": "=$flow.AirDescriptor.source.properties",
                      "ports": "=$flow.AirDescriptor.source.ports"
                    },
                    "logic": "=$flow.AirDescriptor.logic",
                    "properties": "=array.create(f1.json2object(\"{\\\"Name\\\":\\\"FLOGO_APP_PROPS_ENV\\\",\\\"Value\\\":\\\"auto\\\"}\"))",
                    "extra": "=$flow.AirDescriptor.extra"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}",
                    "fe_metadata": "{\"properties\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}],\"source\":{\"name\":\"MQTT\",\"properties\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}],\"ports\":[\"8080:9999\"]},\"logic\":[{\"name\":\"Dgraph\",\"properties\":[{\"Name\":\"\",\"Value\":\"\",\"Type\":\"\"}]}],\"extra\":[{\"Name\":\"networks.default.external.name\",\"Value\":\"edgex-network\"}]}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}}",
                    "fe_metadata": "{\n        \"properties\" : [\n            {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n        ],\n        \"source\": {\n            \"name\": \"MQTT\",\n            \"properties\" : [\n                {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n            ],\n            \"ports\" : [\n                \"8080:9999\"\n            ]\n        },\n        \"logic\": [\n            {\n                \"name\": \"Dgraph\",\n                \"properties\" : [\n                    {\"Name\" : \"\", \"Value\":\"\", \"Type\":\"\"}\n                ]\n            }\n        ],\n        \"extra\": [\n            {\"Name\":\"networks.default.external.name\", \"Value\":\"edgex-network\"}\n        ]\n    }"
                  }
                }
              }
            }
          },
          {
            "id": "AirK8sAppBuilder",
            "name": "AirK8sAppBuilder",
            "description": "This activity build model inference pipeline",
            "activity": {
              "ref": "#airpipelinebuilder",
              "settings": {
                "TemplateFolder": "/home/airpipeline",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"}]",
                "Properties": "[{\"Name\":\"main_apiVersion\",\"Value\":\"apps/v1\"},{\"Name\":\"main_kind\",\"Value\":\"Deployment\"},{\"Name\":\"main_metadata.name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.spec.containers[0].image\",\"Value\":\"$ImageName$\"},{\"Name\":\"main_spec.template.spec.containers[0].name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.selector.matchLabels.component\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.metadata.labels.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_apiVersion\",\"Value\":\"v1\"},{\"Name\":\"ip-service_kind\",\"Value\":\"Service\"},{\"Name\":\"ip-service_metadata.name\",\"Value\":\"$Name$-ip-service\"},{\"Name\":\"ip-service_spec.selector.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_spec.type\",\"Value\":\"LoadBalancer\"}]"
              },
              "input": {
                "ApplicationName": "=$flow.Name",
                "AirDescriptor": "=$activity[BuildAppPipelineDescriptor].JSONString",
                "ServiceType": "=$flow.ServiceType",
                "PropertyPrefix": "spec.template.spec.containers[0]",
                "Variables": {
                  "mapping": {
                    "Name": "=$flow.Name",
                    "ImageName": "=$flow.Image"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"ImageName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Name\":\"2\",\"ImageName\":\"2\"}"
                  }
                },
                "output": {
                  "Descriptor": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"F1Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"FlogoDescriptor\":\"2\",\"F1Properties\":[{\"Group\":\"2\",\"Value\":{\"Name\":\"2\",\"Value\":\"2\",\"Type\":\"2\"}}]}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"}]"
                  },
                  "Properties": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Value\"]}}",
                    "fe_metadata": "[{\"Name\":\"main_apiVersion\",\"Value\":\"apps/v1\"},{\"Name\":\"main_kind\",\"Value\":\"Deployment\"},{\"Name\":\"main_metadata.name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.spec.containers[0].image\",\"Value\":\"$ImageName$\"},{\"Name\":\"main_spec.template.spec.containers[0].name\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.selector.matchLabels.component\",\"Value\":\"$Name$\"},{\"Name\":\"main_spec.template.metadata.labels.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_apiVersion\",\"Value\":\"v1\"},{\"Name\":\"ip-service_kind\",\"Value\":\"Service\"},{\"Name\":\"ip-service_metadata.name\",\"Value\":\"$Name$-ip-service\"},{\"Name\":\"ip-service_spec.selector.component\",\"Value\":\"$Name$\"},{\"Name\":\"ip-service_spec.type\",\"Value\":\"LoadBalancer\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return01",
            "name": "Return01",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "AirService::Validate Pipeline - done",
                  "ErrorCode": 100,
                  "PropertyNameDef": "=$activity[AirK8sAppBuilder].PropertyNameDef",
                  "Descriptor": {
                    "mapping": {
                      "FlogoDescriptor": "=$activity[AirK8sAppBuilder].Descriptor.FlogoDescriptor",
                      "properties": "=$activity[AirK8sAppBuilder].Descriptor.F1Properties",
                      "DeployDescriptor": "=$property[\"Descriptor.k8s\"]"
                    }
                  }
                }
              }
            }
          },
          {
            "id": "AirDockerComposeAppBuilder",
            "name": "AirDockerComposeAppBuilder",
            "description": "This activity build model inference pipeline",
            "activity": {
              "ref": "#airpipelinebuilder",
              "settings": {
                "TemplateFolder": "/home/airpipeline",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ContainerName\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"BuildName\",\"Type\":\"String\"}]",
                "Properties": "[{\"Name\":\"version\",\"Value\":\"3.6\"},{\"Name\":\"services.$Name$.container_name\",\"Value\":\"$ContainerName$\"},{\"Name\":\"services.$Name$.image\",\"Value\":\"$ImageName$\"},{\"Name\":\"services.$Name$.build\",\"Value\":\"$BuildName$\"}]"
              },
              "input": {
                "ApplicationName": "=$flow.Name",
                "AirDescriptor": "=$activity[BuildAppPipelineDescriptor].JSONString",
                "ServiceType": "=$flow.ServiceType",
                "PropertyPrefix": "services.$Name$",
                "Variables": {
                  "mapping": {
                    "Name": "=$flow.Name",
                    "ContainerName": "=$flow.ProjectID+\"_\"+$flow.Name",
                    "ImageName": "=$flow.Image",
                    "BuildName": "=f1.ternary(null==$flow.Image||\"\"==$flow.Image, \"001\", \"\")"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"ContainerName\":{\"type\":\"string\"},\"ImageName\":{\"type\":\"string\"},\"BuildName\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Name\":\"2\",\"ContainerName\":\"2\",\"ImageName\":\"2\",\"BuildName\":\"2\"}"
                  }
                },
                "output": {
                  "Descriptor": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"F1Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"FlogoDescriptor\":\"2\",\"F1Properties\":[{\"Group\":\"2\",\"Value\":{\"Name\":\"2\",\"Value\":\"2\",\"Type\":\"2\"}}]}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"ContainerName\",\"Type\":\"String\"},{\"Name\":\"ImageName\",\"Type\":\"String\"},{\"Name\":\"BuildName\",\"Type\":\"String\"}]"
                  },
                  "Properties": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Value\"]}}",
                    "fe_metadata": "[{\"Name\":\"version\",\"Value\":\"3.6\"},{\"Name\":\"services.$Name$.container_name\",\"Value\":\"$ContainerName$\"},{\"Name\":\"services.$Name$.image\",\"Value\":\"$ImageName$\"},{\"Name\":\"services.$Name$.build\",\"Value\":\"$BuildName$\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return02",
            "name": "Return02",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "AirService::Validate Pipeline - done",
                  "ErrorCode": 100,
                  "PropertyNameDef": "=$activity[AirDockerComposeAppBuilder].PropertyNameDef",
                  "Descriptor": {
                    "mapping": {
                      "FlogoDescriptor": "=$activity[AirDockerComposeAppBuilder].Descriptor.FlogoDescriptor",
                      "properties": "=$activity[AirDockerComposeAppBuilder].Descriptor.F1Properties",
                      "DeployDescriptor": "=$property[\"Descriptor.docker\"]"
                    }
                  }
                }
              }
            }
          },
          {
            "id": "Return04",
            "name": "Return04",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "AirService::Skip validate Flogo app - done",
                  "ErrorCode": 100,
                  "Descriptor": {
                    "mapping": {
                      "FlogoDescriptor": "=$flow.AirDescriptor.logic[0].flogoApp"
                    }
                  }
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return03",
              "name": "Return03",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"AirService::Validate Pipeline - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "SyncBuild",
              "type": "boolean"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "ServiceType",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "AirDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
              }
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            },
            {
              "name": "Descriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}"
              }
            },
            {
              "name": "PropertyNameDef",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"},\"Descriptor\":{\"type\":\"object\",\"properties\":{\"FlogoDescriptor\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{}}},\"DeployDescriptor\":{\"type\":\"string\"}}},\"PropertyNameDef\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Build_and_Deploy_Pipeline",
      "data": {
        "name": "Build and Deploy Pipeline",
        "description": "**** The entry point for project air UI ****",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "NoAppDescriptor",
            "type": "expression",
            "label": "LogMessagetoReturn6",
            "value": "null==$flow.AirDescriptor \u0026\u0026 null==$flow.FlogoAppDescriptor"
          },
          {
            "id": 2,
            "from": "LogMessage",
            "to": "SetSystemProperties",
            "type": "exprOtherwise",
            "label": "LogMessage to SetSystemProperties"
          },
          {
            "id": 3,
            "from": "SetSystemProperties",
            "to": "AddTask",
            "type": "default"
          },
          {
            "id": 4,
            "from": "AddTask",
            "to": "BuildFlogo",
            "type": "default"
          },
          {
            "id": 5,
            "from": "BuildFlogo",
            "to": "BuildSuccessfully",
            "type": "expression",
            "label": "BuildPipelinetoReturn4",
            "value": "$activity[BuildFlogo].Success"
          },
          {
            "id": 6,
            "from": "BuildFlogo",
            "to": "RemoveTask",
            "type": "exprOtherwise",
            "label": "BuildPipelinetoTableMutate"
          },
          {
            "id": 7,
            "from": "RemoveTask",
            "to": "BuildFailed",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildAndDeployAirPipeline %%%%% SyncBuild : \", coerce.toString($flow.SyncBuild), \", ProjectID : \", $flow.ProjectID, \", ComponentType : \", $flow.ComponentType, \", ServiceType : \", $flow.ServiceType, \", Name : \", $flow.Name, \", Image : \", $flow.Image, \", AirDescriptor : \", coerce.toString($flow.AirDescriptor), \", FlogoAppDescriptor : \", coerce.toString($flow.FlogoAppDescriptor), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv))"
              }
            }
          },
          {
            "id": "NoAppDescriptor",
            "name": "NoAppDescriptor",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=false",
                  "Message": "No Application Descriptor!",
                  "ErrorCode": 300
                }
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"},{\"Name\":\"BuilderURL\",\"Type\":\"String\"},{\"Name\":\"PropertyNameDefFilename\",\"Type\":\"String\"},{\"Name\":\"ExtraProperties\",\"Type\":\"Array\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "SyncBuild": "=f1.ternary(null!=$flow.SyncBuild, $flow.SyncBuild, false)",
                    "SkipCondition": "=false",
                    "BuilderURL": "=$property[\"System.BuilderURL\"]",
                    "PropertyNameDefFilename": "=$property[\"System.BaseFolder\"] + \"/projects/\" + $flow.ProjectID + \"/pipeline/\" + $flow.Name + \"_property_name_def.json\"",
                    "ExtraProperties": "=true!=f1.isempty($flow.AirDescriptor)?$flow.AirDescriptor.extra:$flow.FlogoAppDescriptor.extra"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"BuilderURL\":{\"type\":\"string\"},\"PropertyNameDefFilename\":{\"type\":\"string\"},\"ExtraProperties\":{\"type\":\"array\",\"items\":{}},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"SyncBuild\":true,\"BuilderURL\":\"2\",\"PropertyNameDefFilename\":\"2\",\"ExtraProperties\":[],\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncBuild\":{\"type\":\"boolean\"},\"BuilderURL\":{\"type\":\"string\"},\"PropertyNameDefFilename\":{\"type\":\"string\"},\"ExtraProperties\":{\"type\":\"array\",\"items\":{}}}}",
                    "fe_metadata": "{\"SyncBuild\":true,\"BuilderURL\":\"2\",\"PropertyNameDefFilename\":\"2\",\"ExtraProperties\":[]}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"SyncBuild\",\"Type\":\"Boolean\"},{\"Name\":\"BuilderURL\",\"Type\":\"String\"},{\"Name\":\"PropertyNameDefFilename\",\"Type\":\"String\"},{\"Name\":\"ExtraProperties\",\"Type\":\"Array\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "AddTask",
            "name": "AddTask",
            "description": "A simple activity for upserting data to table",
            "activity": {
              "ref": "#tableupsert",
              "settings": {
                "Table": {
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "AirTask"
                    },
                    {
                      "name": "description",
                      "type": "string",
                      "display": {
                        "name": "Description",
                        "description": "A simple table for storing tuple"
                      },
                      "value": ""
                    },
                    {
                      "name": "schema",
                      "type": "array",
                      "required": true,
                      "display": {
                        "name": "Table Schema",
                        "description": "Schema of the table",
                        "type": "table",
                        "schema": "{\r\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\r\n    \"definitions\": {},\r\n    \"id\": \"http://example.com/example.json\",\r\n    \"items\": {\r\n        \"id\": \"/items\",\r\n        \"properties\": {\r\n            \"Name\": {\r\n                \"id\": \"/items/properties/Name\",\r\n                \"type\": \"string\"\r\n            },\r\n            \"Type\": {\r\n                \"id\": \"/items/properties/Type\",\r\n                \"type\": {\"enum\":[\"String\", \"Integer\", \"Long\", \"Double\", \"Boolean\", \"Date\", \"Object\"]}\r\n            },\r\n            \"IsKey\": {\r\n                \"id\": \"/items/properties/IsKey\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            },\r\n            \"Indexed\": {\r\n                \"id\": \"/items/properties/Indexed\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            }\r\n        },\r\n        \"type\": \"object\"\r\n    },\r\n    \"type\": \"array\"\r\n}"
                      },
                      "value": "[{\"Name\":\"ProjectID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Name\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Data\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
                    }
                  ],
                  "outputs": [],
                  "inputs": [],
                  "handler": {
                    "settings": []
                  },
                  "reply": [],
                  "s3Prefix": "flogo",
                  "key": "flogo/GraphBuilder_Tools/connector/simpletable/connector.json",
                  "display": {
                    "description": "This is URL file reader",
                    "category": "GraphBuilder_Tools",
                    "visible": true,
                    "smallIcon": "simpletable.png"
                  },
                  "actions": [
                    {
                      "name": "create",
                      "display": {
                        "name": "Create",
                        "readonly": false,
                        "valid": true,
                        "visible": true
                      }
                    }
                  ],
                  "feature": {},
                  "loopType": "none",
                  "loopSettings": [],
                  "retrySettings": [],
                  "propertyMap": {},
                  "keyfield": "name",
                  "schemaMap": {},
                  "iteratorAccumulate": false,
                  "isValid": true,
                  "lastUpdatedTime": 1613324381563,
                  "createdTime": 1613324381563,
                  "user": "flogo",
                  "subscriptionId": "flogo_sbsc",
                  "id": "9e35d4b0-6eeb-11eb-90dc-41e63340b2bf",
                  "connectorName": "AirTask",
                  "connectorDescription": " "
                }
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID",
                    "Name": "=$flow.Name",
                    "Data": "=f1.modifyobject(f1.modifyobject(f1.json2object(\"{}\"), \"ScriptSystemEnv\" , $flow.ScriptSystemEnv), \"Extra\", $activity[SetSystemProperties].Data.ExtraProperties)"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Data\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"Name\":\"2\",\"Data\":{}}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Data\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"Name\":\"2\",\"Data\":{}}"
                  }
                }
              }
            }
          },
          {
            "id": "BuildFlogo",
            "name": "BuildFlogo",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Build_Pipeline"
              },
              "input": {
                "ProjectID": "=$flow.ProjectID",
                "SyncBuild": "=$flow.SyncBuild",
                "ComponentType": "=$flow.ComponentType",
                "ServiceType": "=$flow.ServiceType",
                "Name": "=$flow.Name",
                "Image": "=$flow.Image",
                "AirDescriptor": "=$flow.AirDescriptor",
                "ScriptSystemEnv": "=f1.modifyobject($flow.ScriptSystemEnv, \"RSVPUrl\", \"http://air:10099/f1/air/rsvp/\"+$flow.ProjectID+\"/\"+$flow.Name)",
                "FlogoAppDescriptor": "=$flow.FlogoAppDescriptor"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"FlogoAppDescriptor\":{\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"FlogoAppDescriptor\":{\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "BuildSuccessfully",
            "name": "BuildSuccessfully",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[BuildFlogo].Success",
                  "Message": "=$activity[BuildFlogo].Message",
                  "ErrorCode": "=$activity[BuildFlogo].ErrorCode"
                }
              }
            }
          },
          {
            "id": "RemoveTask",
            "name": "RemoveTask",
            "description": "A simple activity for upserting/deleting data to/from table",
            "activity": {
              "ref": "#tablemutate",
              "settings": {
                "Table": {
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "AirTask"
                    },
                    {
                      "name": "description",
                      "type": "string",
                      "display": {
                        "name": "Description",
                        "description": "A simple table for storing tuple"
                      },
                      "value": ""
                    },
                    {
                      "name": "schema",
                      "type": "array",
                      "required": true,
                      "display": {
                        "name": "Table Schema",
                        "description": "Schema of the table",
                        "type": "table",
                        "schema": "{\r\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\r\n    \"definitions\": {},\r\n    \"id\": \"http://example.com/example.json\",\r\n    \"items\": {\r\n        \"id\": \"/items\",\r\n        \"properties\": {\r\n            \"Name\": {\r\n                \"id\": \"/items/properties/Name\",\r\n                \"type\": \"string\"\r\n            },\r\n            \"Type\": {\r\n                \"id\": \"/items/properties/Type\",\r\n                \"type\": {\"enum\":[\"String\", \"Integer\", \"Long\", \"Double\", \"Boolean\", \"Date\", \"Object\"]}\r\n            },\r\n            \"IsKey\": {\r\n                \"id\": \"/items/properties/IsKey\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            },\r\n            \"Indexed\": {\r\n                \"id\": \"/items/properties/Indexed\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            }\r\n        },\r\n        \"type\": \"object\"\r\n    },\r\n    \"type\": \"array\"\r\n}"
                      },
                      "value": "[{\"Name\":\"ProjectID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Name\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Data\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
                    }
                  ],
                  "outputs": [],
                  "inputs": [],
                  "handler": {
                    "settings": []
                  },
                  "reply": [],
                  "s3Prefix": "flogo",
                  "key": "flogo/GraphBuilder_Tools/connector/simpletable/connector.json",
                  "display": {
                    "description": "This is URL file reader",
                    "category": "GraphBuilder_Tools",
                    "visible": true,
                    "smallIcon": "simpletable.png"
                  },
                  "actions": [
                    {
                      "name": "create",
                      "display": {
                        "name": "Create",
                        "readonly": false,
                        "valid": true,
                        "visible": true
                      }
                    }
                  ],
                  "feature": {},
                  "loopType": "none",
                  "loopSettings": [],
                  "retrySettings": [],
                  "propertyMap": {},
                  "keyfield": "name",
                  "schemaMap": {},
                  "iteratorAccumulate": false,
                  "isValid": true,
                  "lastUpdatedTime": 1613324381563,
                  "createdTime": 1613324381563,
                  "user": "flogo",
                  "subscriptionId": "flogo_sbsc",
                  "id": "9e35d4b0-6eeb-11eb-90dc-41e63340b2bf",
                  "connectorName": "AirTask",
                  "connectorDescription": " "
                },
                "Method": "delete"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID",
                    "Name": "=$flow.Name"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Data\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"Name\":\"2\",\"Data\":{}}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"New\":{\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Data\":{\"type\":\"object\",\"properties\":{}}}},\"Old\":{\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Data\":{\"type\":\"object\",\"properties\":{}}}}}}",
                    "fe_metadata": "{\"New\":{\"ProjectID\":\"2\",\"Name\":\"2\",\"Data\":{}},\"Old\":{\"ProjectID\":\"2\",\"Name\":\"2\",\"Data\":{}}}"
                  }
                }
              }
            }
          },
          {
            "id": "BuildFailed",
            "name": "BuildFailed",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[BuildFlogo].Success",
                  "Message": "=$activity[BuildFlogo].Message",
                  "ErrorCode": "=$activity[BuildFlogo].ErrorCode"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return3",
              "name": "Return3",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"AirService::Build and Deploy Pipeline - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "SyncBuild",
              "type": "boolean"
            },
            {
              "name": "ComponentType",
              "type": "string"
            },
            {
              "name": "ServiceType",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Image",
              "type": "string"
            },
            {
              "name": "AirDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
              }
            },
            {
              "name": "FlogoAppDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
              }
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"SyncBuild\":{\"type\":\"boolean\"},\"ComponentType\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Image\":{\"type\":\"string\"},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"source\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}},\"logic\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"FlogoAppDescriptor\":{\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"},\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}},\"ports\":{\"type\":\"array\",\"items\":{\"type\":\"string\"}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Ping",
      "data": {
        "name": "Ping",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "BuildServiceInfo",
            "to": "HTTPClient",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "BuildServiceInfo",
            "name": "BuildServiceInfo",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": {
                  "filename": "ServiceLoaderRegist.json",
                  "content": "data:application/json;base64,ewogICAgIklEIjogInByb2plY3RtZ3ItMTIzLjQ4Ni43ODkiLAogICAgIlR5cGUiOiAicHJvamVjdG1nciIsCiAgICAiVVJMIjogImh0dHBzOi8vYXBpLmdpdGh1Yi5jb20vdXNlcnMvU3RldmVOWS1UaWJjby9yZXBvcyIKfQ=="
                },
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\": \"\",\n    \"Type\": \"\",\n    \"URL\": \"\",\n    \"Properties\": {}\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "URL": "=\"http://\" + $property[\"System.ServiceLocatorIP\"] + \":10099/f1/air\"",
                    "Type": "air",
                    "ID": "=\"air_\" + $property[\"System.ServiceLocatorIP\"]"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"\",\"Type\":\"\",\"URL\":\"\",\"Properties\":{}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\n    \"ID\": \"\",\n    \"Type\": \"\",\n    \"URL\": \"\",\n    \"Properties\": {}\n}"
                  }
                }
              }
            }
          },
          {
            "id": "HTTPClient",
            "name": "HTTPClient",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "1000",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "http://%ServiceLocatorIP%:10080/f1/locator/register/%ServiceType%",
                "Body": "=$activity[BuildServiceInfo].JSONString",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "ServiceType": "air"
                  }
                }
              },
              "output": {
                "Success": false
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"},\"ServiceType\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\",\"ServiceType\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"ServiceType\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Now",
              "type": "string"
            }
          ],
          "output": [],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Now\":{\"type\":\"string\"}}}"
          }
        }
      }
    },
    {
      "id": "flow:List_Deployable_Projects",
      "data": {
        "name": "List Deployable Projects",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogDeployPipeline",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "BuildRequest",
            "type": "default"
          },
          {
            "id": 3,
            "from": "BuildRequest",
            "to": "LogRequestToProjectMgr",
            "type": "default"
          },
          {
            "id": 4,
            "from": "LogRequestToProjectMgr",
            "to": "ToListProjects",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ToListProjects",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogDeployPipeline",
            "name": "LogDeployPipeline",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "%%%%% Air:ListDeployableProjects %%%%% "
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectMgrURL\",\"Type\":\"String\"}]"
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectMgrURL\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectMgrURL\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectMgrURL\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectMgrURL\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ProjectMgrURL\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRequest",
            "name": "BuildRequest",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {}\n}",
                  "value": "",
                  "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {}\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {}
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {}\n}",
                    "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {}\n}"
                  }
                }
              }
            }
          },
          {
            "id": "LogRequestToProjectMgr",
            "name": "LogRequestToProjectMgr",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:LogDeployPipeline %%%%% Request To deployF1Descriptor : \", $activity[BuildRequest].JSONString)"
              }
            }
          },
          {
            "id": "ToListProjects",
            "name": "ToListProjects",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "1000",
                "urlMappingString": "",
                "urlMapping": "",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "",
                "httpHeaders": ""
              },
              "input": {
                "URL": "=$activity[SetSystemProperties].Data.ProjectMgrURL+\"/projectmgr/file/list/project/001\"",
                "Body": "=$activity[BuildRequest].JSONString",
                "SkipCondition": "=false"
              },
              "output": {
                "Success": false
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
                    "fe_metadata": "{}"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Message": "=$activity[ToDeployF1Descriptor].Data",
                  "ErrorCode": "=$activity[ToDeployF1Descriptor].ErrorCode",
                  "Success": "=$activity[ToDeployF1Descriptor].Success"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return1",
              "name": "Return1",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Message": "=string.concat(\"AirService::Undeploy Pipeline - \", $error.message)",
                    "Success": "=false",
                    "ErrorCode": 400
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [],
          "output": [
            {
              "name": "Data",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Data\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Delete_Domain",
      "data": {
        "name": "Delete Domain",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "DeleteProjectFolder",
            "type": "default"
          },
          {
            "id": 3,
            "from": "DeleteProjectFolder",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:Delete Domain %%%%% Domain : \", $flow.ProjectID)"
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectTemplateFolderInt\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectsFolderInt": "=$property[\"System.BaseFolder\"] + \"/projects\"",
                    "SkipCondition": "=false",
                    "ProjectTemplateFolderInt": "/home/project"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectsFolderInt\":{\"type\":\"string\"},\"ProjectTemplateFolderInt\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectsFolderInt\":\"2\",\"ProjectTemplateFolderInt\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectsFolderInt\":{\"type\":\"string\"},\"ProjectTemplateFolderInt\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectsFolderInt\":\"2\",\"ProjectTemplateFolderInt\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectsFolderInt\",\"Type\":\"String\"},{\"Name\":\"ProjectTemplateFolderInt\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "DeleteProjectFolder",
            "name": "DeleteProjectFolder",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$ProjectsFolder$",
                "numOfExecutions": 1,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": "=false",
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "rm -rf $ProjectID$"
                    }
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID",
                    "ProjectsFolder": "=$activity[SetSystemProperties].Data.ProjectsFolderInt"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectsFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=$activity[DeleteProjectFolder].Success",
                  "Message": "=coerce.toString($activity[DeleteProjectFolder].Message)",
                  "ErrorCode": "=$activity[DeleteProjectFolder].ErrorCode"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return3",
              "name": "Return3",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=string.concat(\"AirService::Initialize - \", $error.message)"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Undeploy_Pipeline",
      "data": {
        "name": "Undeploy Pipeline",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogDeployPipeline",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SetSystemProperties",
            "to": "BuildRequest",
            "type": "default"
          },
          {
            "id": 3,
            "from": "BuildRequest",
            "to": "LogRequestToDeployer",
            "type": "default"
          },
          {
            "id": 4,
            "from": "LogRequestToDeployer",
            "to": "ToDeployF1Descriptor",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ToDeployF1Descriptor",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogDeployPipeline",
            "name": "LogDeployPipeline",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:LogDeployPipeline %%%%% ProjectID : \", $flow.ProjectID, \", Method : \", $flow.Method, coerce.toString($flow.SyncDeploy), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv))"
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"DeployerURL\",\"Type\":\"String\"},{\"Name\":\"SyncDeploy\",\"Type\":\"Boolean\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "BaseFolder": "=$property[\"System.BaseFolder\"]",
                    "ProjectFolderExt": "=$property[\"System.BaseFolderExt\"] + \"/projects/\"",
                    "ProjectFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID",
                    "Namespace": "=f1.escapek8sid($flow.ProjectID)",
                    "DeployerURL": "=$property[\"System.DeployerURL\"]",
                    "SkipCondition": "=false",
                    "SyncDeploy": "=f1.ternary(null!=$flow.SyncDeploy, $flow.SyncDeploy, false)"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"DeployerURL\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderExt\":\"2\",\"ProjectFolderInt\":\"2\",\"Namespace\":\"2\",\"DeployerURL\":\"2\",\"SyncDeploy\":true,\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"DeployerURL\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderExt\":\"2\",\"ProjectFolderInt\":\"2\",\"Namespace\":\"2\",\"DeployerURL\":\"2\",\"SyncDeploy\":true}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"DeployerURL\",\"Type\":\"String\"},{\"Name\":\"SyncDeploy\",\"Type\":\"Boolean\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRequest",
            "name": "BuildRequest",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"NoF1Descriptor\": {\n            \"type\": \"boolean\"\n        },\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"ProjectID\": {\n            \"type\": \"string\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"UserDefinedParameters\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"ServiceName\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                }\n            }\n        }\n    }\n}",
                  "value": "",
                  "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"NoF1Descriptor\": {\n            \"type\": \"boolean\"\n        },\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"ProjectID\": {\n            \"type\": \"string\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"UserDefinedParameters\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"ServiceName\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                }\n            }\n        }\n    }\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                    "SyncDeploy": "=$activity[SetSystemProperties].Data.SyncDeploy",
                    "NoF1Descriptor": "=false",
                    "Method": "=$flow.Method",
                    "ProjectID": "=$flow.ProjectID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"NoF1Descriptor\":{\"type\":\"boolean\"},\"Method\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"UserDefinedParameters\":{\"type\":\"object\",\"properties\":{\"ServiceName\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"NoF1Descriptor\":{\"type\":\"boolean\"},\"Method\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"UserDefinedParameters\":{\"type\":\"object\",\"properties\":{\"ServiceName\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"}}}}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"NoF1Descriptor\": {\n            \"type\": \"boolean\"\n        },\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"ProjectID\": {\n            \"type\": \"string\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"UserDefinedParameters\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"ServiceName\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                }\n            }\n        }\n    }\n}",
                    "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"NoF1Descriptor\": {\n            \"type\": \"boolean\"\n        },\n        \"Method\": {\n            \"type\": \"string\"\n        },\n        \"ProjectID\": {\n            \"type\": \"string\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"UserDefinedParameters\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"ServiceName\": {\n                    \"type\": \"string\"\n                },\n                \"Descriptor\": {\n                    \"type\": \"string\"\n                }\n            }\n        }\n    }\n}"
                  }
                }
              }
            }
          },
          {
            "id": "LogRequestToDeployer",
            "name": "LogRequestToDeployer",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:LogDeployPipeline %%%%% Request To deployF1Descriptor : \", $activity[BuildRequest].JSONString)"
              }
            }
          },
          {
            "id": "ToDeployF1Descriptor",
            "name": "ToDeployF1Descriptor",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "1000",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"Instance\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "=$activity[SetSystemProperties].Data.DeployerURL+\"/f1/deployer/undeploy/$ProjectID$/$Name$/$Instance$\"",
                "Body": "=$activity[BuildRequest].JSONString",
                "SkipCondition": "=false",
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID",
                    "Name": "=$flow.Name",
                    "Instance": "=$flow.Instance"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"Name\":\"2\",\"Instance\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"Name\",\"Type\":\"String\"},{\"Name\":\"Instance\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Message": "=$activity[ToDeployF1Descriptor].Data",
                  "ErrorCode": "=$activity[ToDeployF1Descriptor].ErrorCode",
                  "Success": "=$activity[ToDeployF1Descriptor].Success"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return1",
              "name": "Return1",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Message": "=string.concat(\"AirService::Undeploy Pipeline - \", $error.message)",
                    "Success": "=false",
                    "ErrorCode": 400
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Instance",
              "type": "string"
            },
            {
              "name": "Method",
              "type": "string"
            },
            {
              "name": "SyncDeploy",
              "type": "boolean"
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Method\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Deploy_Pipeline",
      "data": {
        "name": "Deploy Pipeline",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LogDeployPipeline",
            "to": "ChkUserScriptExistence",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ChkUserScriptExistence",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 3,
            "from": "SetSystemProperties",
            "to": "ReadPropertyNameDef",
            "type": "default"
          },
          {
            "id": 4,
            "from": "ReadPropertyNameDef",
            "to": "ReadF1Descriptor",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ReadF1Descriptor",
            "to": "ResolveParameters",
            "type": "default"
          },
          {
            "id": 6,
            "from": "ResolveParameters",
            "to": "ExtractDescriptor",
            "type": "default"
          },
          {
            "id": 7,
            "from": "ExtractDescriptor",
            "to": "ModifyService",
            "type": "default"
          },
          {
            "id": 8,
            "from": "ModifyService",
            "to": "BuildRequest",
            "type": "default"
          },
          {
            "id": 9,
            "from": "BuildRequest",
            "to": "LogRequestToDeployer",
            "type": "default"
          },
          {
            "id": 10,
            "from": "LogRequestToDeployer",
            "to": "ToDeployF1Descriptor",
            "type": "default"
          },
          {
            "id": 11,
            "from": "ToDeployF1Descriptor",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogDeployPipeline",
            "name": "LogDeployPipeline",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:LogDeployPipeline %%%%% ProjectID : \", $flow.ProjectID, \", Method : \", $flow.Method, coerce.toString($flow.SyncDeploy), \", ScriptSystemEnv : \", coerce.toString($flow.ScriptSystemEnv), \", AirDescriptor : \", coerce.toString($flow.AirDescriptor))"
              }
            }
          },
          {
            "id": "ChkUserScriptExistence",
            "name": "ChkUserScriptExistence",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=\"/home/f1/projects/\" + $flow.ProjectID + \"/deploy/deploy.sh\""
              },
              "schemas": {
                "output": {
                  "Results": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Content\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Filename\":\"2\",\"Content\":\"2\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "SetSystemProperties",
            "name": "SetSystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"UseDefaultScript\",\"Type\":\"Boolean\"},{\"Name\":\"DeployerURL\",\"Type\":\"String\"},{\"Name\":\"SyncDeploy\",\"Type\":\"Boolean\"},{\"Name\":\"RuntimeProperties\",\"Type\":\"Array\"},{\"Name\":\"ExtraProperties\",\"Type\":\"Array\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "BaseFolder": "=$property[\"System.BaseFolder\"]",
                    "ProjectFolderExt": "=$property[\"System.BaseFolderExt\"] + \"/projects/\"",
                    "ProjectFolderInt": "=\"/home/f1/projects/\" + $flow.ProjectID",
                    "DeployFolderInt": "=f1.ternary(null!=$activity[ChkUserScriptExistence].Results\u0026\u00260!=array.count($activity[ChkUserScriptExistence].Results),\"/home/f1/projects/\" + $flow.ProjectID + \"/deploy/\", \"/home/scripts\")",
                    "Namespace": "=f1.escapek8sid($flow.ProjectID)",
                    "UseDefaultScript": "=f1.ternary(null!=$activity[ChkUserScriptExistence].Results\u0026\u00260!=array.count($activity[ChkUserScriptExistence].Results), false, true)",
                    "DeployerURL": "=$property[\"System.DeployerURL\"]",
                    "SkipCondition": "=false",
                    "SyncDeploy": "=f1.ternary(null!=$flow.SyncDeploy, $flow.SyncDeploy, false)",
                    "ExtraProperties": "=coerce.toArray(f1.ternary(null!=$flow.AirDescriptor, $flow.AirDescriptor.extra, null))",
                    "RuntimeProperties": "=coerce.toArray(f1.ternary(null!=$flow.AirDescriptor.dynamic, $flow.AirDescriptor.dynamic.properties, null))"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"DeployerURL\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"RuntimeProperties\":{\"type\":\"array\",\"items\":{}},\"ExtraProperties\":{\"type\":\"array\",\"items\":{}},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderExt\":\"2\",\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"UseDefaultScript\":true,\"DeployerURL\":\"2\",\"SyncDeploy\":true,\"RuntimeProperties\":[],\"ExtraProperties\":[],\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"BaseFolder\":{\"type\":\"string\"},\"ProjectFolderExt\":{\"type\":\"string\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolderInt\":{\"type\":\"string\"},\"Namespace\":{\"type\":\"string\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"DeployerURL\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"RuntimeProperties\":{\"type\":\"array\",\"items\":{}},\"ExtraProperties\":{\"type\":\"array\",\"items\":{}}}}",
                    "fe_metadata": "{\"BaseFolder\":\"2\",\"ProjectFolderExt\":\"2\",\"ProjectFolderInt\":\"2\",\"DeployFolderInt\":\"2\",\"Namespace\":\"2\",\"UseDefaultScript\":true,\"DeployerURL\":\"2\",\"SyncDeploy\":true,\"RuntimeProperties\":[],\"ExtraProperties\":[]}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"BaseFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderExt\",\"Type\":\"String\"},{\"Name\":\"ProjectFolderInt\",\"Type\":\"String\"},{\"Name\":\"DeployFolderInt\",\"Type\":\"String\"},{\"Name\":\"Namespace\",\"Type\":\"String\"},{\"Name\":\"UseDefaultScript\",\"Type\":\"Boolean\"},{\"Name\":\"DeployerURL\",\"Type\":\"String\"},{\"Name\":\"SyncDeploy\",\"Type\":\"Boolean\"},{\"Name\":\"RuntimeProperties\",\"Type\":\"Array\"},{\"Name\":\"ExtraProperties\",\"Type\":\"Array\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ReadPropertyNameDef",
            "name": "ReadPropertyNameDef",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": "/home/f1/projects"
              },
              "input": {
                "FilePattern": "=$flow.ProjectID + \"/pipeline/\" + $flow.Name + \"_property_name_def.json\""
              },
              "schemas": {
                "output": {
                  "Results": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Content\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Filename\":\"2\",\"Content\":\"2\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ReadF1Descriptor",
            "name": "ReadF1Descriptor",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": "/home/f1/projects"
              },
              "input": {
                "FilePattern": "=$flow.ProjectID + \"/pipeline/\" + $flow.Name + \".json\""
              },
              "schemas": {
                "output": {
                  "Results": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Filename\":{\"type\":\"string\"},\"Content\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Filename\":\"2\",\"Content\":\"2\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ResolveParameters",
            "name": "ResolveParameters",
            "description": "Replace target text from input text",
            "activity": {
              "ref": "#textreplacer",
              "settings": {
                "leftToken": "$",
                "rightToken": "$",
                "replacementKeys": "[{\"Name\":\"temp.project.home\",\"Type\":\"String\"},{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"ServiceLocator\",\"Type\":\"String\"},{\"Name\":\"ExternalEndpointIP\",\"Type\":\"String\"},{\"Name\":\"ExternalEndpointPort\",\"Type\":\"String\"}]"
              },
              "input": {
                "inputDocument": "=$activity[ReadF1Descriptor].Results[0].Content",
                "replacements": {
                  "mapping": {
                    "temp.project.home": "=$activity[SetSystemProperties].Data.ProjectFolderExt + $flow.ProjectID",
                    "ID": "=$flow.ProjectID",
                    "ServiceLocator": "=$property[\"System.ServiceLocatorExternal\"]",
                    "ExternalEndpointIP": "=$property[\"System.ExternalEndpointIP\"]",
                    "ExternalEndpointPort": "10100"
                  }
                }
              },
              "schemas": {
                "input": {
                  "replacements": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"temp.project.home\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ServiceLocator\":{\"type\":\"string\"},\"ExternalEndpointIP\":{\"type\":\"string\"},\"ExternalEndpointPort\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"temp.project.home\":\"\",\"ID\":\"\",\"ServiceLocator\":\"\",\"ExternalEndpointIP\":\"\",\"ExternalEndpointPort\":\"\"}"
                  }
                },
                "settings": {
                  "replacementKeys": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"temp.project.home\",\"Type\":\"String\"},{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"ServiceLocator\",\"Type\":\"String\"},{\"Name\":\"ExternalEndpointIP\",\"Type\":\"String\"},{\"Name\":\"ExternalEndpointPort\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ExtractDescriptor",
            "name": "ExtractDescriptor",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}}}}}}}",
                  "value": "",
                  "fe_metadata": "{\n\t\"DeploymentGpID\":\"\",\n\t\"Services\":[\n\t\t{\n\t\t\t\"Descriptor\":\"\",\n\t\t\t\"Name\":\"\",\n\t\t\t\"Properties\":[{\n\t\t\t    \"Group\": \"\",\n\t\t\t    \"Value\": {\"Name\":\"\",\"Value\":\"\"}\n\t\t\t}],\n\t\t\t\"ScriptSystemEnv\":{},\n\t\t\t\"Type\":\"\",\n\t\t\t\"Variables\":{\n\t\t\t}\n\t\t}\n\t]\n}"
                }
              },
              "input": {
                "JSONString": "=$activity[ResolveParameters].outputDocument"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}}}}}}}",
                    "fe_metadata": "{\"DeploymentGpID\":\"\",\"Services\":[{\"Descriptor\":\"\",\"Name\":\"\",\"Properties\":[{\"Group\":\"\",\"Value\":{\"Name\":\"\",\"Value\":\"\"}}],\"ScriptSystemEnv\":{},\"Type\":\"\",\"Variables\":{}}]}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}}}}}}}",
                    "fe_metadata": "{\n\t\"DeploymentGpID\":\"\",\n\t\"Services\":[\n\t\t{\n\t\t\t\"Descriptor\":\"\",\n\t\t\t\"Name\":\"\",\n\t\t\t\"Properties\":[{\n\t\t\t    \"Group\": \"\",\n\t\t\t    \"Value\": {\"Name\":\"\",\"Value\":\"\"}\n\t\t\t}],\n\t\t\t\"ScriptSystemEnv\":{},\n\t\t\t\"Type\":\"\",\n\t\t\t\"Variables\":{\n\t\t\t}\n\t\t}\n\t]\n}"
                  }
                }
              }
            }
          },
          {
            "id": "ModifyService",
            "name": "ModifyService",
            "description": "Make an New Object",
            "activity": {
              "ref": "#objectmaker",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Service\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}}}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"Service\":{\n\t\t\t\"Descriptor\":\"\",\n\t\t\t\"Name\":\"\",\n\t\t\t\"Instance\":\"\",\n\t\t\t\"Properties\":[{\n\t\t\t    \"Group\": \"\",\n\t\t\t    \"Value\": {\"Name\":\"\",\"Value\":\"\"}\n\t\t\t}],\n\t\t\t\"ScriptSystemEnv\":{},\n\t\t\t\"Type\":\"\",\n\t\t\t\"Variables\":{\n\t\t\t}\n\t\t}\n}"
                }
              },
              "input": {
                "ObjectDataMapping": {
                  "mapping": {
                    "Service": {
                      "Descriptor": "=$activity[ExtractDescriptor].Data.Services[0].Descriptor",
                      "Name": "=$activity[ExtractDescriptor].Data.Services[0].Name",
                      "ScriptSystemEnv": "=$activity[ExtractDescriptor].Data.Services[0].ScriptSystemEnv",
                      "Type": "=$activity[ExtractDescriptor].Data.Services[0].Type",
                      "Variables": "=$activity[ExtractDescriptor].Data.Services[0].Variables",
                      "Instance": "=$flow.Instance",
                      "Properties": "=f1.air2f1properties(\n\t$activity[ExtractDescriptor].Data.Services[0].Type, \n\tf1.ternary(\n\t\t\"k8s\"==$activity[ExtractDescriptor].Data.Services[0].Type, \n\t\t\"spec.template.spec.containers[0].env\", \n\t\t\"services.\"+$activity[ExtractDescriptor].Data.Services[0].Name+\".environment\"\n\t), \n\t$activity[ExtractDescriptor].Data.Services[0].Properties, \n\tf1.toobjectarray($activity[SetSystemProperties].Data.RuntimeProperties), \n\tf1.json2object($activity[ReadPropertyNameDef].Results[0].Content),\n\tf1.toobjectarray($activity[SetSystemProperties].Data.ExtraProperties)\n)"
                    }
                  }
                }
              },
              "schemas": {
                "input": {
                  "ObjectDataMapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Service\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}}}}}}",
                    "fe_metadata": "{\"Service\":{\"Descriptor\":\"\",\"Name\":\"\",\"Instance\":\"\",\"Properties\":[{\"Group\":\"\",\"Value\":{\"Name\":\"\",\"Value\":\"\"}}],\"ScriptSystemEnv\":{},\"Type\":\"\",\"Variables\":{}}}"
                  }
                },
                "output": {
                  "ObjectOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Service\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}}}}}}",
                    "fe_metadata": "{\"Service\":{\"Descriptor\":\"\",\"Name\":\"\",\"Instance\":\"\",\"Properties\":[{\"Group\":\"\",\"Value\":{\"Name\":\"\",\"Value\":\"\"}}],\"ScriptSystemEnv\":{},\"Type\":\"\",\"Variables\":{}}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Service\":{\"type\":\"object\",\"properties\":{\"Descriptor\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"Type\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}}}}}}",
                    "fe_metadata": "{\n    \"Service\":{\n\t\t\t\"Descriptor\":\"\",\n\t\t\t\"Name\":\"\",\n\t\t\t\"Instance\":\"\",\n\t\t\t\"Properties\":[{\n\t\t\t    \"Group\": \"\",\n\t\t\t    \"Value\": {\"Name\":\"\",\"Value\":\"\"}\n\t\t\t}],\n\t\t\t\"ScriptSystemEnv\":{},\n\t\t\t\"Type\":\"\",\n\t\t\t\"Variables\":{\n\t\t\t}\n\t\t}\n}"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRequest",
            "name": "BuildRequest",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"ProjectFolderInt\": {\n            \"type\": \"string\"\n        },\n        \"DeployFolder\": {\n            \"type\": \"string\"\n        },\n        \"UseDefaultScript\": {\n            \"type\": \"boolean\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"F1Descriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"DeploymentGpID\": {\n                    \"type\": \"string\"\n                },\n                \"DataFlow\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Upstream\": {\n                                \"type\": \"string\"\n                            },\n                            \"Downstream\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"Components\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Runtime\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Replicas\": {\n                                \"type\": \"number\"\n                            },\n                            \"DockerImage\": {\n                                \"type\": \"string\"\n                            },\n                            \"Volumes\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"MountPoint\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"Services\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Instance\": {\n                                \"type\": \"string\"\n                            },\n                            \"Descriptor\": {\n                                \"type\": \"string\"\n                            },\n                            \"Variables\": {\n                                \"type\": \"object\",\n                                \"properties\": {}\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Group\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"array\",\n                                            \"items\": {\n                                                \"type\": \"object\",\n                                                \"properties\": {\n                                                    \"Name\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Value\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Type\": {\n                                                        \"type\": \"string\"\n                                                    }\n                                                }\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"System\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                }\n            }\n        }\n    }\n}",
                  "value": "",
                  "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"ProjectFolderInt\": {\n            \"type\": \"string\"\n        },\n        \"DeployFolder\": {\n            \"type\": \"string\"\n        },\n        \"UseDefaultScript\": {\n            \"type\": \"boolean\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"F1Descriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"DeploymentGpID\": {\n                    \"type\": \"string\"\n                },\n                \"DataFlow\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Upstream\": {\n                                \"type\": \"string\"\n                            },\n                            \"Downstream\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"Components\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Runtime\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Replicas\": {\n                                \"type\": \"number\"\n                            },\n                            \"DockerImage\": {\n                                \"type\": \"string\"\n                            },\n                            \"Volumes\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"MountPoint\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"Services\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Instance\": {\n                                \"type\": \"string\"\n                            },\n                            \"Descriptor\": {\n                                \"type\": \"string\"\n                            },\n                            \"Variables\": {\n                                \"type\": \"object\",\n                                \"properties\": {}\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Group\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"array\",\n                                            \"items\": {\n                                                \"type\": \"object\",\n                                                \"properties\": {\n                                                    \"Name\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Value\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Type\": {\n                                                        \"type\": \"string\"\n                                                    }\n                                                }\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"System\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                }\n            }\n        }\n    }\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "UseDefaultScript": "=$activity[SetSystemProperties].Data.UseDefaultScript",
                    "ProjectFolderInt": "=$activity[SetSystemProperties].Data.ProjectFolderInt",
                    "ScriptSystemEnv": "=$flow.ScriptSystemEnv",
                    "DeployFolder": "=$activity[SetSystemProperties].Data.DeployFolderInt",
                    "F1Descriptor": {
                      "DeploymentGpID": "=$activity[ExtractDescriptor].Data.DeploymentGpID",
                      "Services": "=array.create($activity[ModifyService].ObjectOut.Service)"
                    },
                    "SyncDeploy": "=$activity[SetSystemProperties].Data.SyncDeploy"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"F1Descriptor\":{\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"SyncDeploy\":{\"type\":\"boolean\"},\"ProjectFolderInt\":{\"type\":\"string\"},\"DeployFolder\":{\"type\":\"string\"},\"UseDefaultScript\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"F1Descriptor\":{\"type\":\"object\",\"properties\":{\"DeploymentGpID\":{\"type\":\"string\"},\"DataFlow\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Upstream\":{\"type\":\"string\"},\"Downstream\":{\"type\":\"string\"}}}},\"Components\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Runtime\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Replicas\":{\"type\":\"number\"},\"DockerImage\":{\"type\":\"string\"},\"Volumes\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"MountPoint\":{\"type\":\"string\"}}}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}},\"Services\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Descriptor\":{\"type\":\"string\"},\"Variables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Group\":{\"type\":\"string\"},\"Value\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}}},\"System\":{\"type\":\"object\",\"properties\":{}}}}}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"ProjectFolderInt\": {\n            \"type\": \"string\"\n        },\n        \"DeployFolder\": {\n            \"type\": \"string\"\n        },\n        \"UseDefaultScript\": {\n            \"type\": \"boolean\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"F1Descriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"DeploymentGpID\": {\n                    \"type\": \"string\"\n                },\n                \"DataFlow\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Upstream\": {\n                                \"type\": \"string\"\n                            },\n                            \"Downstream\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"Components\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Runtime\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Replicas\": {\n                                \"type\": \"number\"\n                            },\n                            \"DockerImage\": {\n                                \"type\": \"string\"\n                            },\n                            \"Volumes\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"MountPoint\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"Services\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Instance\": {\n                                \"type\": \"string\"\n                            },\n                            \"Descriptor\": {\n                                \"type\": \"string\"\n                            },\n                            \"Variables\": {\n                                \"type\": \"object\",\n                                \"properties\": {}\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Group\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"array\",\n                                            \"items\": {\n                                                \"type\": \"object\",\n                                                \"properties\": {\n                                                    \"Name\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Value\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Type\": {\n                                                        \"type\": \"string\"\n                                                    }\n                                                }\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"System\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                }\n            }\n        }\n    }\n}",
                    "fe_metadata": "{\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n    \"type\": \"object\",\n    \"properties\": {\n        \"SyncDeploy\": {\n            \"type\": \"boolean\"\n        },\n        \"ProjectFolderInt\": {\n            \"type\": \"string\"\n        },\n        \"DeployFolder\": {\n            \"type\": \"string\"\n        },\n        \"UseDefaultScript\": {\n            \"type\": \"boolean\"\n        },\n        \"ScriptSystemEnv\": {\n            \"type\": \"object\",\n            \"properties\": {}\n        },\n        \"F1Descriptor\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"DeploymentGpID\": {\n                    \"type\": \"string\"\n                },\n                \"DataFlow\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Upstream\": {\n                                \"type\": \"string\"\n                            },\n                            \"Downstream\": {\n                                \"type\": \"string\"\n                            }\n                        }\n                    }\n                },\n                \"Components\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Runtime\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Replicas\": {\n                                \"type\": \"number\"\n                            },\n                            \"DockerImage\": {\n                                \"type\": \"string\"\n                            },\n                            \"Volumes\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"MountPoint\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Name\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Type\": {\n                                            \"type\": \"string\"\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"Services\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"type\": \"object\",\n                        \"properties\": {\n                            \"Type\": {\n                                \"type\": \"string\"\n                            },\n                            \"Name\": {\n                                \"type\": \"string\"\n                            },\n                            \"Instance\": {\n                                \"type\": \"string\"\n                            },\n                            \"Descriptor\": {\n                                \"type\": \"string\"\n                            },\n                            \"Variables\": {\n                                \"type\": \"object\",\n                                \"properties\": {}\n                            },\n                            \"Properties\": {\n                                \"type\": \"array\",\n                                \"items\": {\n                                    \"type\": \"object\",\n                                    \"properties\": {\n                                        \"Group\": {\n                                            \"type\": \"string\"\n                                        },\n                                        \"Value\": {\n                                            \"type\": \"array\",\n                                            \"items\": {\n                                                \"type\": \"object\",\n                                                \"properties\": {\n                                                    \"Name\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Value\": {\n                                                        \"type\": \"string\"\n                                                    },\n                                                    \"Type\": {\n                                                        \"type\": \"string\"\n                                                    }\n                                                }\n                                            }\n                                        }\n                                    }\n                                }\n                            }\n                        }\n                    }\n                },\n                \"System\": {\n                    \"type\": \"object\",\n                    \"properties\": {}\n                }\n            }\n        }\n    }\n}"
                  }
                }
              }
            }
          },
          {
            "id": "LogRequestToDeployer",
            "name": "LogRequestToDeployer",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:LogDeployPipeline %%%%% Request To deployF1Descriptor : \", $activity[BuildRequest].JSONString)"
              }
            }
          },
          {
            "id": "ToDeployF1Descriptor",
            "name": "ToDeployF1Descriptor",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "1000",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "=$activity[SetSystemProperties].Data.DeployerURL+\"/f1/deployer/deployF1Descriptor/$ProjectID$\"",
                "Body": "=$activity[BuildRequest].JSONString",
                "SkipCondition": "=false",
                "Variables": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Message": "=$activity[ToDeployF1Descriptor].Data",
                  "ErrorCode": "=$activity[ToDeployF1Descriptor].ErrorCode",
                  "Success": "=$activity[ToDeployF1Descriptor].Success"
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return1",
              "name": "Return1",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "Message": "=string.concat(\"AirService::Deploy Pipeline - \", $error.message)",
                    "ErrorCode": 400
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Instance",
              "type": "string"
            },
            {
              "name": "Method",
              "type": "string"
            },
            {
              "name": "SyncDeploy",
              "type": "boolean"
            },
            {
              "name": "ScriptSystemEnv",
              "type": "object"
            },
            {
              "name": "AirDescriptor",
              "type": "object",
              "schema": {
                "type": "json",
                "value": "{\"dynamic\":{\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Component\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}"
              }
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Method\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"dynamic\":{\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Component\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Get_By_Category",
      "data": {
        "name": "Get By Category",
        "description": "== Return component templates by category ==",
        "links": [
          {
            "id": 1,
            "from": "AirApplicationQuerier",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "AirApplicationQuerier",
            "name": "AirApplicationQuerier",
            "description": "This activity query air pipeline compoment",
            "activity": {
              "ref": "#aircomponentquerier",
              "settings": {
                "TemplateFolder": "/home/airpipeline"
              },
              "input": {
                "Category": "=$flow.Category",
                "Component": ""
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Data": "=$activity[AirApplicationQuerier].Descriptor"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Category",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Data",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Category\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Data\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Get_Compnent",
      "data": {
        "name": "Get Compnent",
        "description": "== Return component templates by category/component ==",
        "links": [
          {
            "id": 1,
            "from": "AirApplicationQuerier",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "AirApplicationQuerier",
            "name": "AirApplicationQuerier",
            "description": "This activity query air pipeline compoment",
            "activity": {
              "ref": "#aircomponentquerier",
              "settings": {
                "TemplateFolder": "/home/airpipeline"
              },
              "input": {
                "Category": "=$flow.Category",
                "Component": "=$flow.Component"
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Data": "=$activity[AirApplicationQuerier].Descriptor"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Category",
              "type": "string"
            },
            {
              "name": "Component",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Data",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Category\":{\"type\":\"string\"},\"Component\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Data\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:RSVP_from_builder",
      "data": {
        "name": "RSVP from builder",
        "description": "**** Builder callback to report build result ****",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "TableMutate",
            "type": "default"
          },
          {
            "id": 2,
            "from": "TableMutate",
            "to": "LogTaskExistsAndSuccessful",
            "type": "expression",
            "label": "TableMutatetoDeployPipeline",
            "value": "$activity[TableMutate].Exists\u0026\u0026$flow.Successful"
          },
          {
            "id": 3,
            "from": "LogTaskExistsAndSuccessful",
            "to": "DeployPipeline",
            "type": "default"
          },
          {
            "id": 4,
            "from": "DeployPipeline",
            "to": "Return4",
            "type": "default"
          },
          {
            "id": 5,
            "from": "TableMutate",
            "to": "LogTaskNotFoundOrFail",
            "type": "exprOtherwise",
            "label": "TableMutatetoReturn5"
          },
          {
            "id": 6,
            "from": "LogTaskNotFoundOrFail",
            "to": "Return5",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildNdDeployAirPipeline RSVP Event %%%%% ProjectID : \", $flow.ProjectID, \", Name : \", $flow.Name, \", Successful: \", $flow.Successful, \", ErrorMsg : \", $flow.ErrorMsg)"
              }
            }
          },
          {
            "id": "TableMutate",
            "name": "TableMutate",
            "description": "A simple activity for upserting/deleting data to/from table",
            "activity": {
              "ref": "#tablemutate",
              "settings": {
                "Table": {
                  "type": "flogo:connector",
                  "version": "1.0.0",
                  "name": "tibco-simple-table",
                  "hashTags": [],
                  "inputMappings": {},
                  "outputMappings": {},
                  "iteratorMappings": {},
                  "title": "Simple Table",
                  "description": "This is URL file reader",
                  "ref": "github.com/SteveNY-Tibco/labs-lightcrane-contrib/connector/simpletable",
                  "settings": [
                    {
                      "name": "name",
                      "type": "string",
                      "required": true,
                      "display": {
                        "name": "Table Name",
                        "description": "Name of the table instance"
                      },
                      "value": "AirTask"
                    },
                    {
                      "name": "description",
                      "type": "string",
                      "display": {
                        "name": "Description",
                        "description": "A simple table for storing tuple"
                      },
                      "value": ""
                    },
                    {
                      "name": "schema",
                      "type": "array",
                      "required": true,
                      "display": {
                        "name": "Table Schema",
                        "description": "Schema of the table",
                        "type": "table",
                        "schema": "{\r\n    \"$schema\": \"http://json-schema.org/draft-04/schema#\",\r\n    \"definitions\": {},\r\n    \"id\": \"http://example.com/example.json\",\r\n    \"items\": {\r\n        \"id\": \"/items\",\r\n        \"properties\": {\r\n            \"Name\": {\r\n                \"id\": \"/items/properties/Name\",\r\n                \"type\": \"string\"\r\n            },\r\n            \"Type\": {\r\n                \"id\": \"/items/properties/Type\",\r\n                \"type\": {\"enum\":[\"String\", \"Integer\", \"Long\", \"Double\", \"Boolean\", \"Date\", \"Object\"]}\r\n            },\r\n            \"IsKey\": {\r\n                \"id\": \"/items/properties/IsKey\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            },\r\n            \"Indexed\": {\r\n                \"id\": \"/items/properties/Indexed\",\r\n                \"type\": {\"enum\":[\"no\", \"yes\"]}\r\n            }\r\n        },\r\n        \"type\": \"object\"\r\n    },\r\n    \"type\": \"array\"\r\n}"
                      },
                      "value": "[{\"Name\":\"ProjectID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Name\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Data\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
                    }
                  ],
                  "outputs": [],
                  "inputs": [],
                  "handler": {
                    "settings": []
                  },
                  "reply": [],
                  "s3Prefix": "flogo",
                  "key": "flogo/GraphBuilder_Tools/connector/simpletable/connector.json",
                  "display": {
                    "description": "This is URL file reader",
                    "category": "GraphBuilder_Tools",
                    "visible": true,
                    "smallIcon": "simpletable.png"
                  },
                  "actions": [
                    {
                      "name": "create",
                      "display": {
                        "name": "Create",
                        "readonly": false,
                        "valid": true,
                        "visible": true
                      }
                    }
                  ],
                  "feature": {},
                  "loopType": "none",
                  "loopSettings": [],
                  "retrySettings": [],
                  "propertyMap": {},
                  "keyfield": "name",
                  "schemaMap": {},
                  "iteratorAccumulate": false,
                  "isValid": true,
                  "lastUpdatedTime": 1613324381563,
                  "createdTime": 1613324381563,
                  "user": "flogo",
                  "subscriptionId": "flogo_sbsc",
                  "id": "9e35d4b0-6eeb-11eb-90dc-41e63340b2bf",
                  "connectorName": "AirTask",
                  "connectorDescription": " "
                },
                "Method": "delete"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectID": "=$flow.ProjectID",
                    "Name": "=$flow.Name"
                  }
                }
              },
              "output": {
                "Exists": false
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Data\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"Name\":\"2\",\"Data\":{}}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"New\":{\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Data\":{\"type\":\"object\",\"properties\":{}}}},\"Old\":{\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Data\":{\"type\":\"object\",\"properties\":{}}}}}}",
                    "fe_metadata": "{\"New\":{\"ProjectID\":\"2\",\"Name\":\"2\",\"Data\":{}},\"Old\":{\"ProjectID\":\"2\",\"Name\":\"2\",\"Data\":{}}}"
                  }
                }
              }
            }
          },
          {
            "id": "LogTaskExistsAndSuccessful",
            "name": "LogTaskExistsAndSuccessful",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildNdDeployAirPipeline RSVP1 %%%%% ProjectID : \", $flow.ProjectID, \", Name : \", $flow.Name, \", Successful: \", $flow.Successful, \", ErrorMsg : \", $flow.ErrorMsg, \", Data : \", coerce.toString($activity[TableMutate].Data.Old), \", Exist : \", $activity[TableMutate].Exists)"
              }
            }
          },
          {
            "id": "DeployPipeline",
            "name": "DeployPipeline",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Deploy_Pipeline"
              },
              "input": {
                "ProjectID": "=$activity[TableMutate].Data.Old.ProjectID",
                "Name": "=$activity[TableMutate].Data.Old.Name",
                "Instance": "001",
                "Method": "docker",
                "SyncDeploy": "=false",
                "ScriptSystemEnv": "=f1.getsubobject($activity[TableMutate].Data.Old.Data,\"ScriptSystemEnv\")",
                "AirDescriptor": {
                  "mapping": {
                    "dynamic": "=f1.json2object(\"{\\\"properties\\\":[]}\")",
                    "extra": "=coerce.toArray(f1.getsubobject($activity[TableMutate].Data.Old.Data, \"Extra\"))"
                  }
                }
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Method\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"dynamic\":{\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Component\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Instance\":{\"type\":\"string\"},\"Method\":{\"type\":\"string\"},\"SyncDeploy\":{\"type\":\"boolean\"},\"ScriptSystemEnv\":{\"type\":\"object\",\"properties\":{}},\"AirDescriptor\":{\"type\":\"object\",\"properties\":{\"dynamic\":{\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Component\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"}}}}}},\"extra\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return4",
            "name": "Return4",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": true,
                  "Message": "Action triggered!",
                  "ErrorCode": 100
                }
              }
            }
          },
          {
            "id": "LogTaskNotFoundOrFail",
            "name": "LogTaskNotFoundOrFail",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%% Air:BuildNdDeployAirPipeline RSVP2 %%%%% ProjectID : \", $flow.ProjectID, \", Name : \", $flow.Name, \", Successful: \", $flow.Successful, \", ErrorMsg : \", $flow.ErrorMsg, \", Data : \", coerce.toString($activity[TableMutate].Data.Old), \", Exist : \", $activity[TableMutate].Exists)"
              }
            }
          },
          {
            "id": "Return5",
            "name": "Return5",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Success": "=true",
                  "Message": "No matched action found!",
                  "ErrorCode": 100
                }
              }
            }
          }
        ],
        "errorHandler": {
          "tasks": [
            {
              "id": "Return3",
              "name": "Return3",
              "description": "Simple Return Activity",
              "activity": {
                "ref": "#actreturn",
                "settings": {
                  "mappings": {
                    "Success": "=false",
                    "ErrorCode": 400,
                    "Message": "=$error.message"
                  }
                }
              }
            }
          ],
          "links": []
        },
        "metadata": {
          "input": [
            {
              "name": "ProjectID",
              "type": "string"
            },
            {
              "name": "Name",
              "type": "string"
            },
            {
              "name": "Successful",
              "type": "boolean"
            },
            {
              "name": "ErrorMsg",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"Name\":{\"type\":\"string\"},\"Successful\":{\"type\":\"boolean\"},\"ErrorMsg\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Get_Compnents",
      "data": {
        "name": "Get Compnents",
        "description": "== List component templates to Air UI ==",
        "links": [
          {
            "id": 1,
            "from": "AirApplicationQuerier",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "AirApplicationQuerier",
            "name": "AirApplicationQuerier",
            "description": "This activity query air pipeline compoment",
            "activity": {
              "ref": "#aircomponentquerier",
              "settings": {
                "TemplateFolder": "/home/airpipeline"
              },
              "input": {
                "Category": "*",
                "Component": "*"
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Data": "=$activity[AirApplicationQuerier].Descriptor"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [],
          "output": [
            {
              "name": "Data",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Data\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Get_Flogo_App_Properties",
      "data": {
        "name": "Get Flogo App Properties",
        "description": "== Extract environment properties from a Flogo application ==",
        "links": [],
        "tasks": [
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "properties": "=f1.airflogoproperties($flow.flogoApp)"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "flogoApp",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "properties",
              "type": "array",
              "schema": {
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}"
              }
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"flogoApp\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"properties\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Value\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Upload_Flogo_Component",
      "data": {
        "name": "Upload Flogo Component",
        "description": "? Not implemented yet ?",
        "links": [],
        "tasks": [],
        "metadata": {
          "input": [],
          "output": [],
          "fe_metadata": {}
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Upload_GO_Lib",
      "data": {
        "name": "Upload GO Lib",
        "description": "? Not implemented yet ?",
        "links": [
          {
            "id": 1,
            "from": "LogMessage",
            "to": "JSONSerializer",
            "type": "default"
          },
          {
            "id": 2,
            "from": "JSONSerializer",
            "to": "ToBuilder",
            "type": "default"
          },
          {
            "id": 3,
            "from": "ToBuilder",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LogMessage",
            "name": "LogMessage",
            "description": "Simple Log Message Activity",
            "activity": {
              "ref": "#log",
              "input": {
                "Log Level": "INFO",
                "flowInfo": false,
                "message": "=string.concat(\"%%%%%% Upload Resource %%%%%% Type : \", $flow.Type, \", ID : \", $flow.ID)"
              }
            }
          },
          {
            "id": "JSONSerializer",
            "name": "JSONSerializer",
            "description": "JSON Serializer Activity",
            "activity": {
              "ref": "#jsonserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"Resource\":{\"type\":\"string\"}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"Path\" : \"\",\n    \"Resource\" : \"\"\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "Path": "=$flow.Path",
                    "Resource": "=$flow.Resource"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"Resource\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Path\":\"\",\"Resource\":\"\"}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Path\":{\"type\":\"string\"},\"Resource\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\n    \"Path\" : \"\",\n    \"Resource\" : \"\"\n}"
                  }
                }
              }
            }
          },
          {
            "id": "ToBuilder",
            "name": "ToBuilder",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "PUT",
                "timeout": "10000",
                "urlMappingString": "",
                "urlMapping": "",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"}]",
                "httpHeaders": ""
              },
              "input": {
                "URL": "=$property[\"System.BuilderURL\"] + \"/f1/builder/upload/$Type$/$ID$\"",
                "Body": "=$activity[JSONSerializer].JSONString",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ID": "=$flow.ID",
                    "Type": "=$flow.Type"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"Type\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ID\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "Return",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "Message": "=$activity[ToBuilder].Data",
                  "Success": "=$activity[ToBuilder].Success",
                  "ErrorCode": "=$activity[ToBuilder].ErrorCode"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "Type",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "Path",
              "type": "string"
            },
            {
              "name": "Resource",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Success",
              "type": "boolean"
            },
            {
              "name": "Message",
              "type": "string"
            },
            {
              "name": "ErrorCode",
              "type": "float64"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"Path\":{\"type\":\"string\"},\"Resource\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Success\":{\"type\":\"boolean\"},\"Message\":{\"type\":\"string\"},\"ErrorCode\":{\"type\":\"number\"}}}"
          }
        },
        "explicitReply": true
      }
    }
  ]
}`
const engineJSON string = ``

func init () {
	cfgJson = flogoJSON
	cfgEngine = engineJSON
}

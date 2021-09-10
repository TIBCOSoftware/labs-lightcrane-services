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
    "github.com/project-flogo/legacybridge",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/filereader",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/httpclient",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/tablequery",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsonserializer",
    "github.com/project-flogo/flow",
    "github.com/project-flogo/contrib/activity/actreturn",
    "github.com/project-flogo/contrib/function/string",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/log",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/mapping",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/jsondeserializer",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/tableupsert",
    "github.com/project-flogo/contrib/trigger/rest",
    "github.com/project-flogo/contrib/function/coerce",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/function/f1",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/filewriter",
    "github.com/project-flogo/contrib/trigger/timer",
    "github.com/project-flogo/flow/activity/subflow",
    "github.com/SteveNY-Tibco/labs-lightcrane-contrib/activity/exec",
    "github.com/project-flogo/contrib/function/array",
    "github.com/project-flogo/contrib/function/datetime"
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
      "value": "BaseFolder not set! "
    },
    {
      "name": "System.ProjectTemplateFolder",
      "type": "string",
      "value": "/home/project"
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
        "port": 10090
      },
      "handlers": [
        {
          "name": "List_Repository",
          "settings": {
            "method": "GET",
            "path": "/f1/projectmgr/listRepositories"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:List_Repository"
              },
              "output": {
                "data": {
                  "mapping": {
                    "Response": {
                      "@foreach($.Response,Response)": {
                        "=": "$loop"
                      }
                    }
                  }
                }
              }
            }
          ]
        },
        {
          "name": "List_Projects",
          "settings": {
            "method": "GET",
            "path": "/f1/projectmgr/listProjects/:RepositoryID"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:List_Projects"
              },
              "input": {
                "RepositoryID": "=$.pathParams.RepositoryID"
              },
              "output": {
                "data": {
                  "mapping": {
                    "Projects": "=$.Projects"
                  }
                }
              }
            }
          ]
        },
        {
          "name": "Chkout_Project",
          "settings": {
            "method": "POST",
            "path": "/f1/projectmgr/load"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Chkout_Project"
              },
              "input": {
                "ID": "=$.content.ID",
                "Properties": "=$.content.Properties",
                "RepositoryID": "=$.content.RepositoryID"
              },
              "output": {
                "data": {
                  "mapping": {
                    "DataOut": "=$.DataOut"
                  }
                }
              }
            }
          ]
        },
        {
          "name": "Create_Project",
          "settings": {
            "method": "POST",
            "path": "/f1/projectmgr/create"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Create_Project"
              },
              "input": {
                "ID": "=$.content.ID",
                "Properties": "=$.content.Properties",
                "RepositoryID": "=$.content.RepositoryID"
              },
              "output": {
                "data": {
                  "mapping": {
                    "DataOut": "=$.DataOut"
                  }
                }
              }
            }
          ]
        },
        {
          "name": "Save_Project",
          "settings": {
            "method": "POST",
            "path": "/f1/projectmgr/save"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Save_Project"
              },
              "input": {
                "ID": "=$.content.ID",
                "Properties": "=$.content.Properties",
                "RepositoryID": "=$.content.RepositoryID"
              },
              "output": {
                "data": {
                  "mapping": {
                    "DataOut": "=$.DataOut"
                  }
                }
              }
            }
          ]
        },
        {
          "name": "Project",
          "settings": {
            "method": "POST",
            "path": "/f1/projectmgr/file/:action/:artifact/:id"
          },
          "actions": [
            {
              "ref": "github.com/project-flogo/flow",
              "settings": {
                "flowURI": "res://flow:Project"
              },
              "input": {
                "Action": "=$.pathParams.action",
                "Artifact": "=$.pathParams.artifact",
                "ID": "=$.pathParams.id",
                "Properties": "=$.body",
                "Type": "file"
              },
              "output": {
                "code": 200,
                "data": {
                  "mapping": {
                    "DataOut": "=$.DataOut"
                  }
                }
              }
            }
          ]
        }
      ]
    },
    {
      "id": "TimerTrigger",
      "ref": "#timer",
      "settings": {},
      "handlers": [
        {
          "name": "Ping",
          "settings": {
            "Interval Unit": "Minute",
            "Repeating": true,
            "Start Date": "",
            "Time Interval": 1
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
      "id": "flow:Save",
      "data": {
        "name": "Save",
        "description": "",
        "links": [],
        "tasks": [
          {
            "id": "LoadProjectLog",
            "name": "LoadProjectLog",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Load Project ]]] ################## \" + $flow.Type + \"-\" + $flow.Action + \"-\" + $flow.ID)",
                "addDetails": false
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
              "name": "Action",
              "type": "string"
            },
            {
              "name": "Artifact",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "ProjectFolder",
              "type": "string"
            },
            {
              "name": "ScriptFolder",
              "type": "string"
            },
            {
              "name": "DefaultScriptFolder",
              "type": "string"
            },
            {
              "name": "Properties",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        }
      }
    },
    {
      "id": "flow:Create",
      "data": {
        "name": "Create",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "CreateProjectLog",
            "to": "CreateProjectFolder",
            "type": "default"
          },
          {
            "id": 2,
            "from": "CreateProjectFolder",
            "to": "SetSystemProperties",
            "type": "default"
          },
          {
            "id": 3,
            "from": "SetSystemProperties",
            "to": "WriteDeployableFile",
            "type": "default"
          },
          {
            "id": 4,
            "from": "WriteDeployableFile",
            "to": "WritePredefinedVariable",
            "type": "default"
          },
          {
            "id": 5,
            "from": "WritePredefinedVariable",
            "to": "ReadLCFile",
            "type": "default"
          },
          {
            "id": 6,
            "from": "ReadLCFile",
            "to": "KeywordMap",
            "type": "default"
          },
          {
            "id": 7,
            "from": "KeywordMap",
            "to": "WriteLCFile",
            "type": "default"
          },
          {
            "id": 8,
            "from": "WriteLCFile",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "CreateProjectLog",
            "name": "CreateProjectLog",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Create Project ]]] ################## \" + $flow.Type + \"-\" + $flow.Action + \"-\" + $flow.ID)",
                "addDetails": false
              }
            }
          },
          {
            "id": "CreateProjectFolder",
            "name": "CreateProjectFolder",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Execute_Script"
              },
              "input": {
                "Type": "=$flow.Type",
                "Artifact": "=$flow.Artifact",
                "ID": "=$flow.ID",
                "ProjectsFolder": "=$flow.ProjectsFolder",
                "DefaultScriptFolder": "=$flow.DefaultScriptFolder",
                "ScriptFolder": "=$flow.ScriptFolder",
                "Action": "=$flow.Action",
                "ScriptEnvVariables": "=$flow.ScriptEnvVariables"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"ScriptEnvVariables\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"ScriptEnvVariables\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
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
                "MappingFields": "[{\"Name\":\"Configuration\",\"Type\":\"String\"},{\"Name\":\"PredefinedVariables\",\"Type\":\"String\"},{\"Name\":\"DeployableName\",\"Type\":\"String\"},{\"Name\":\"DockerComposeFolder\",\"Type\":\"String\"},{\"Name\":\"Description\",\"Type\":\"String\"},{\"Name\":\"LCFilePath\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "Configuration": "=f1.getsubobject(f1.getsubobject($flow.Properties, \"Deployable\"), \"Configuration\")",
                    "PredefinedVariables": "=f1.getsubobject(f1.getsubobject($flow.Properties, \"Deployable\"), \"PredefinedVariables\")",
                    "DeployableName": "=f1.getsubobject(f1.getsubobject($flow.Properties, \"Deployable\"), \"Name\")",
                    "DockerComposeFolder": "=$flow.ProjectsFolder + \"/\" + $flow.ID + \"/artifacts/docker-compose\"",
                    "Description": "=f1.getsubobject($flow.Properties, \"Description\")",
                    "LCFilePath": "=$flow.ProjectsFolder + \"/\" + $flow.ID + \"/lc.json\""
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Configuration\":{\"type\":\"string\"},\"PredefinedVariables\":{\"type\":\"string\"},\"DeployableName\":{\"type\":\"string\"},\"DockerComposeFolder\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"LCFilePath\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"Configuration\":\"2\",\"PredefinedVariables\":\"2\",\"DeployableName\":\"2\",\"DockerComposeFolder\":\"2\",\"Description\":\"2\",\"LCFilePath\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Configuration\":{\"type\":\"string\"},\"PredefinedVariables\":{\"type\":\"string\"},\"DeployableName\":{\"type\":\"string\"},\"DockerComposeFolder\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"LCFilePath\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Configuration\":\"2\",\"PredefinedVariables\":\"2\",\"DeployableName\":\"2\",\"DockerComposeFolder\":\"2\",\"Description\":\"2\",\"LCFilePath\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Configuration\",\"Type\":\"String\"},{\"Name\":\"PredefinedVariables\",\"Type\":\"String\"},{\"Name\":\"DeployableName\",\"Type\":\"String\"},{\"Name\":\"DockerComposeFolder\",\"Type\":\"String\"},{\"Name\":\"Description\",\"Type\":\"String\"},{\"Name\":\"LCFilePath\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "WriteDeployableFile",
            "name": "WriteDeployableFile",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "$DeployableFilepath$",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"DeployableFilepath\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": false,
                "Data": {
                  "mapping": {
                    "Input": "=$activity[SetSystemProperties].Data.Configuration"
                  }
                },
                "Variables": {
                  "mapping": {
                    "DeployableFilepath": "=$activity[SetSystemProperties].Data.DockerComposeFolder + \"/\" + $activity[SetSystemProperties].Data.DeployableName + \".yml\""
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeployableFilepath\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"DeployableFilepath\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DeployableFilepath\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"DeployableFilepath\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"DeployableFilepath\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "WritePredefinedVariable",
            "name": "WritePredefinedVariable",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "$PredefinedVariablesFilepath$",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"PredefinedVariablesFilepath\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": false,
                "Data": {
                  "mapping": {
                    "Input": "=$activity[SetSystemProperties].Data.PredefinedVariables"
                  }
                },
                "Variables": {
                  "mapping": {
                    "PredefinedVariablesFilepath": "=$activity[SetSystemProperties].Data.DockerComposeFolder + \"/\" + $activity[SetSystemProperties].Data.DeployableName + \"-default.json\""
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"PredefinedVariablesFilepath\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"PredefinedVariablesFilepath\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"PredefinedVariablesFilepath\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"PredefinedVariablesFilepath\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"PredefinedVariablesFilepath\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ReadLCFile",
            "name": "ReadLCFile",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=$activity[SetSystemProperties].Data.LCFilePath"
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
            "id": "KeywordMap",
            "name": "KeywordMap",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectName\",\"Type\":\"String\"},{\"Name\":\"Description\",\"Type\":\"String\"},{\"Name\":\"Created\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectName": "=$flow.ID",
                    "ProjectID": "=$flow.ID",
                    "Description": "=$activity[SetSystemProperties].Data.Description",
                    "Created": "=datetime.now()"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectName\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Created\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectName\":\"2\",\"Description\":\"2\",\"Created\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectName\":{\"type\":\"string\"},\"Description\":{\"type\":\"string\"},\"Created\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectName\":\"2\",\"Description\":\"2\",\"Created\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectName\",\"Type\":\"String\"},{\"Name\":\"Description\",\"Type\":\"String\"},{\"Name\":\"Created\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "WriteLCFile",
            "name": "WriteLCFile",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filewriter",
              "settings": {
                "inputType": "String",
                "outputFile": "$LCFilePath$",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"LCFilePath\",\"Type\":\"String\"}]"
              },
              "input": {
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "LCFilePath": "=$activity[SetSystemProperties].Data.LCFilePath"
                  }
                },
                "Data": {
                  "mapping": {
                    "Input": "=f1.keywordreplace($activity[ReadLCFile].Results[0].Content, \"${\", \"}\", $activity[KeywordMap].Data)"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"LCFilePath\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"LCFilePath\":\"2\"}"
                  }
                },
                "output": {
                  "VariablesOut": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"LCFilePath\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"LCFilePath\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"LCFilePath\",\"Type\":\"String\"}]"
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
                  "DataOut": {
                    "mapping": {
                      "Message": "project created!"
                    }
                  }
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
              "name": "Action",
              "type": "string"
            },
            {
              "name": "Artifact",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "ProjectsFolder",
              "type": "string"
            },
            {
              "name": "ScriptFolder",
              "type": "string"
            },
            {
              "name": "DefaultScriptFolder",
              "type": "string"
            },
            {
              "name": "ScriptEnvVariables",
              "type": "object"
            },
            {
              "name": "Properties",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"ScriptEnvVariables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        }
      }
    },
    {
      "id": "flow:List",
      "data": {
        "name": "List",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "ListLog",
            "to": "ActionProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ActionProperties",
            "to": "CheckUsrScript",
            "type": "default"
          },
          {
            "id": 3,
            "from": "CheckUsrScript",
            "to": "ExecuteUsrScript",
            "type": "expression",
            "label": "CheckScripttoExecCommand02",
            "value": "null!=$activity[CheckUsrScript].Results \u0026\u0026 0!=array.count($activity[CheckUsrScript].Results)"
          },
          {
            "id": 4,
            "from": "ExecuteUsrScript",
            "to": "ExecUsrScript",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ExecUsrScript",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 6,
            "from": "CheckUsrScript",
            "to": "ExecuteDefaultScript",
            "type": "exprOtherwise",
            "label": "CheckScript to CopyOfExecCommand02"
          },
          {
            "id": 7,
            "from": "ExecuteDefaultScript",
            "to": "ExecDefaultScript",
            "type": "default"
          },
          {
            "id": 8,
            "from": "ExecDefaultScript",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "ListLog",
            "name": "ListLog",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager : List ] ################## \" + $flow.Type + \"-\" + $flow.Action + \"-\" + $flow.ID)",
                "addDetails": false
              }
            }
          },
          {
            "id": "ActionProperties",
            "name": "ActionProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"Script\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "Script": "./manager.sh"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Script\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"Script\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Script\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Script\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"Script\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "CheckUsrScript",
            "name": "CheckUsrScript",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=$flow.ScriptFolder + \"/\" + $flow.Action + \"-\" + $flow.Type + \".sh\""
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
            "id": "ExecuteUsrScript",
            "name": "ExecuteUsrScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager : Execute Usr Script ] ################## \" + $flow.ScriptFolder)",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecUsrScript",
            "name": "ExecUsrScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x $Script$",
                      "Execution_1": "$Script$ $ProjectsFolder$ $Type$ $Action$"
                    },
                    "SystemEnvs": "=$flow.Properties"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectsFolder": "=$flow.ProjectsFolder",
                    "Script": "=$activity[ActionProperties].Data.Script",
                    "ScriptFolder": "=$flow.ScriptFolder",
                    "Type": "=$flow.Type",
                    "Action": "=$flow.Artifact"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"Script\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectsFolder\":\"2\",\"Script\":\"2\",\"Type\":\"2\",\"Action\":\"2\"}"
                  }
                },
                "output": {
                  "Result": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Command\":{\"type\":\"string\"},\"StdOut\":{\"type\":\"string\"},\"StdErr\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Command\":\"2\",\"StdOut\":\"2\",\"StdErr\":\"2\"}]"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"}]"
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
                  "DataOut": "=f1.json2object($activity[ExecUsrScript].Result[1].StdOut)"
                }
              }
            }
          },
          {
            "id": "ExecuteDefaultScript",
            "name": "ExecuteDefaultScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager : Execute Default Script ] ################## \" + $flow.DefaultScriptFolder)",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecDefaultScript",
            "name": "ExecDefaultScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 1,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "$Script$ $ProjectsFolder$ $Type$ $Action$"
                    },
                    "SystemEnvs": "=$flow.Properties"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ScriptFolder": "=$flow.DefaultScriptFolder",
                    "ProjectsFolder": "=$flow.ProjectsFolder",
                    "Script": "=$activity[ActionProperties].Data.Script",
                    "Type": "=$flow.Type",
                    "Action": "=$flow.Action"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"Script\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectsFolder\":\"2\",\"Script\":\"2\",\"Type\":\"2\",\"Action\":\"2\"}"
                  }
                },
                "output": {
                  "Result": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Command\":{\"type\":\"string\"},\"StdOut\":{\"type\":\"string\"},\"StdErr\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Command\":\"2\",\"StdOut\":\"2\",\"StdErr\":\"2\"}]"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"}]"
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
                  "DataOut": "=f1.json2object($activity[ExecDefaultScript].Result[0].StdOut)"
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
              "name": "Action",
              "type": "string"
            },
            {
              "name": "Artifact",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "ProjectsFolder",
              "type": "string"
            },
            {
              "name": "ScriptFolder",
              "type": "string"
            },
            {
              "name": "DefaultScriptFolder",
              "type": "string"
            },
            {
              "name": "Properties",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        }
      }
    },
    {
      "id": "flow:Delete",
      "data": {
        "name": "Delete",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "ListLog",
            "to": "ActionProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ActionProperties",
            "to": "CheckUsrScript",
            "type": "default"
          },
          {
            "id": 3,
            "from": "CheckUsrScript",
            "to": "ExecuteUsrScript",
            "type": "expression",
            "label": "CheckScripttoExecCommand02",
            "value": "null!=$activity[CheckUsrScript].Results \u0026\u0026 0!=array.count($activity[CheckUsrScript].Results)"
          },
          {
            "id": 4,
            "from": "ExecuteUsrScript",
            "to": "ExecUsrScript",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ExecUsrScript",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 6,
            "from": "CheckUsrScript",
            "to": "ExecuteDefaultScript",
            "type": "exprOtherwise",
            "label": "CheckScript to CopyOfExecCommand02"
          },
          {
            "id": 7,
            "from": "ExecuteDefaultScript",
            "to": "ExecDefaultScript",
            "type": "default"
          },
          {
            "id": 8,
            "from": "ExecDefaultScript",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "ListLog",
            "name": "ListLog",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager : List ] ################## \" + $flow.Type + \"-\" + $flow.Action + \"-\" + $flow.ID)",
                "addDetails": false
              }
            }
          },
          {
            "id": "ActionProperties",
            "name": "ActionProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"IDs\",\"Type\":\"Array\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "Script": "./manager.sh",
                    "IDs": "=f1.getsubobject($flow.Properties, \"IDs\")"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Script\":{\"type\":\"string\"},\"IDs\":{\"type\":\"array\",\"items\":{}},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"Script\":\"2\",\"IDs\":[],\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Script\":{\"type\":\"string\"},\"IDs\":{\"type\":\"array\",\"items\":{}}}}",
                    "fe_metadata": "{\"Script\":\"2\",\"IDs\":[]}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"IDs\",\"Type\":\"Array\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "CheckUsrScript",
            "name": "CheckUsrScript",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=$flow.ScriptFolder + \"/\" + $flow.Action + \"-\" + $flow.Type + \".sh\""
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
            "id": "ExecuteUsrScript",
            "name": "ExecuteUsrScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager : Execute Usr Script ] ################## \" + $flow.ScriptFolder)",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecUsrScript",
            "name": "ExecUsrScript",
            "description": "This activity build docker image from Dockerfile",
            "type": "iterator",
            "settings": {
              "iterate": "=$activity[ActionProperties].Data.IDs",
              "accumulate": false
            },
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x $Script$",
                      "Execution_1": "$Script$ $ProjectsFolder$ $Type$ $Action$ $ProjectFolder$"
                    },
                    "SystemEnvs": "=$flow.Properties"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectsFolder": "=$flow.ProjectsFolder",
                    "Script": "=$activity[ActionProperties].Data.Script",
                    "ScriptFolder": "=$flow.ScriptFolder",
                    "Type": "=$flow.Type",
                    "Action": "=$flow.Artifact",
                    "ProjectFolder": "=$iteration[value]"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"Script\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectsFolder\":\"2\",\"Script\":\"2\",\"Type\":\"2\",\"Action\":\"2\",\"ProjectFolder\":\"2\"}"
                  }
                },
                "output": {
                  "Result": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Command\":{\"type\":\"string\"},\"StdOut\":{\"type\":\"string\"},\"StdErr\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Command\":\"2\",\"StdOut\":\"2\",\"StdErr\":\"2\"}]"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"}]"
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
                  "DataOut": "=f1.json2object($activity[ExecUsrScript].Result[1].StdOut)"
                }
              }
            }
          },
          {
            "id": "ExecuteDefaultScript",
            "name": "ExecuteDefaultScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager : Execute Default Script ] ################## \" + $flow.DefaultScriptFolder)",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecDefaultScript",
            "name": "ExecDefaultScript",
            "description": "This activity build docker image from Dockerfile",
            "type": "iterator",
            "settings": {
              "iterate": "=$activity[ActionProperties].Data.IDs",
              "accumulate": false
            },
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 1,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "$Script$ $ProjectsFolder$ $Type$ $Action$ $ProjectFolder$"
                    },
                    "SystemEnvs": "=$flow.Properties"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ScriptFolder": "=$flow.DefaultScriptFolder",
                    "ProjectsFolder": "=$flow.ProjectsFolder",
                    "Script": "=$activity[ActionProperties].Data.Script",
                    "Type": "=$flow.Type",
                    "Action": "=$flow.Action",
                    "ProjectFolder": "=$iteration[value]"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"Script\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectsFolder\":\"2\",\"Script\":\"2\",\"Type\":\"2\",\"Action\":\"2\",\"ProjectFolder\":\"2\"}"
                  }
                },
                "output": {
                  "Result": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Command\":{\"type\":\"string\"},\"StdOut\":{\"type\":\"string\"},\"StdErr\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Command\":\"2\",\"StdOut\":\"2\",\"StdErr\":\"2\"}]"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"}]"
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
                  "DataOut": "=f1.json2object($activity[ExecDefaultScript].Result[0].StdOut)"
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
              "name": "Action",
              "type": "string"
            },
            {
              "name": "Artifact",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "ProjectsFolder",
              "type": "string"
            },
            {
              "name": "ScriptFolder",
              "type": "string"
            },
            {
              "name": "DefaultScriptFolder",
              "type": "string"
            },
            {
              "name": "Properties",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        }
      }
    },
    {
      "id": "flow:Execute_Script",
      "data": {
        "name": "Execute Script",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "ExecScriptLog",
            "to": "ActionProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "ActionProperties",
            "to": "CheckUsrScript",
            "type": "default"
          },
          {
            "id": 3,
            "from": "CheckUsrScript",
            "to": "ExecuteUsrScript",
            "type": "expression",
            "label": "CheckScripttoExecCommand02",
            "value": "null!=$activity[CheckUsrScript].Results \u0026\u0026 0!=array.count($activity[CheckUsrScript].Results)"
          },
          {
            "id": 4,
            "from": "ExecuteUsrScript",
            "to": "ExecUsrScript",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ExecUsrScript",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 6,
            "from": "CheckUsrScript",
            "to": "ExecuteDefaultScript",
            "type": "exprOtherwise",
            "label": "CheckScript to CopyOfExecCommand02"
          },
          {
            "id": 7,
            "from": "ExecuteDefaultScript",
            "to": "ExecDefaultScript",
            "type": "default"
          },
          {
            "id": 8,
            "from": "ExecDefaultScript",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "ExecScriptLog",
            "name": "ExecScriptLog",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager : Execute Script ] ################## \" + $flow.Type + \"-\" + $flow.Action + \"-\" + $flow.ID)",
                "addDetails": false
              }
            }
          },
          {
            "id": "ActionProperties",
            "name": "ActionProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"Script\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "Script": "./manager.sh"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Script\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"Script\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Script\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"Script\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"Script\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "CheckUsrScript",
            "name": "CheckUsrScript",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=$flow.ScriptFolder + \"/\" + $flow.Action + \"-\" + $flow.Type + \".sh\""
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
            "id": "ExecuteUsrScript",
            "name": "ExecuteUsrScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager : Execute Usr Script ] ################## \" + $flow.ScriptFolder)",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecUsrScript",
            "name": "ExecUsrScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x $Script$",
                      "Execution_1": "$Script$ $ProjectsFolder$ $Type$ $Action$"
                    },
                    "SystemEnvs": "=$flow.ScriptEnvVariables"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ProjectsFolder": "=$flow.ProjectsFolder",
                    "Script": "=$activity[ActionProperties].Data.Script",
                    "ScriptFolder": "=$flow.ScriptFolder",
                    "Type": "=$flow.Type",
                    "Action": "=$flow.Artifact"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"Script\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectsFolder\":\"2\",\"Script\":\"2\",\"Type\":\"2\",\"Action\":\"2\"}"
                  }
                },
                "output": {
                  "Result": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Command\":{\"type\":\"string\"},\"StdOut\":{\"type\":\"string\"},\"StdErr\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Command\":\"2\",\"StdOut\":\"2\",\"StdErr\":\"2\"}]"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"}]"
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
                  "DataOut": "=f1.json2object($activity[ExecUsrScript].Result[1].StdOut)"
                }
              }
            }
          },
          {
            "id": "ExecuteDefaultScript",
            "name": "ExecuteDefaultScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager : Execute Default Script ] ################## \" + $flow.DefaultScriptFolder)",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecDefaultScript",
            "name": "ExecDefaultScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "execConnection": "",
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 1,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "$Script$ $ProjectsFolder$ $Type$ $Action$"
                    },
                    "SystemEnvs": "=$flow.ScriptEnvVariables"
                  }
                },
                "Variables": {
                  "mapping": {
                    "ScriptFolder": "=$flow.DefaultScriptFolder",
                    "ProjectsFolder": "=$flow.ProjectsFolder",
                    "Script": "=$activity[ActionProperties].Data.Script",
                    "Type": "=$flow.Type",
                    "Action": "=$flow.Action"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"Script\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectsFolder\":\"2\",\"Script\":\"2\",\"Type\":\"2\",\"Action\":\"2\"}"
                  }
                },
                "output": {
                  "Result": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Command\":{\"type\":\"string\"},\"StdOut\":{\"type\":\"string\"},\"StdErr\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Command\":\"2\",\"StdOut\":\"2\",\"StdErr\":\"2\"}]"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"Script\",\"Type\":\"String\"},{\"Name\":\"Type\",\"Type\":\"String\"},{\"Name\":\"Action\",\"Type\":\"String\"}]"
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
                  "DataOut": "=f1.json2object($activity[ExecDefaultScript].Result[0].StdOut)"
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
              "name": "Action",
              "type": "string"
            },
            {
              "name": "Artifact",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "ProjectsFolder",
              "type": "string"
            },
            {
              "name": "ScriptFolder",
              "type": "string"
            },
            {
              "name": "DefaultScriptFolder",
              "type": "string"
            },
            {
              "name": "ScriptEnvVariables",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"ScriptEnvVariables\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        }
      }
    },
    {
      "id": "flow:List_Repository",
      "data": {
        "name": "List Repository",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "GetRepositoryInfo",
            "to": "BuildRepositoryInfo",
            "type": "default"
          },
          {
            "id": 2,
            "from": "BuildRepositoryInfo",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "GetRepositoryInfo",
            "name": "GetRepositoryInfo",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "GET",
                "timeout": "1000",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "http://%ServiceLocatorIP%:10080/mops/locator/list/repository",
                "Body": "",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRepositoryInfo",
            "name": "BuildRepositoryInfo",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": {
                  "filename": "ServiceLocatorListResponse.json",
                  "content": "data:application/json;base64,ewogICAgIlJlc3BvbnNlIjogWwogICAgICAgIHsKICAgICAgICAgICAgIklEIjogIjEyMzQ4Njc5IiwKICAgICAgICAgICAgIlRZUEUiOiAicmVwb3NpdG9yeSIKICAgICAgICB9CiAgICBdCn0="
                },
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Response\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"TYPE\":{\"type\":\"string\"}}}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"Response\": [\n        {\n            \"ID\": \"12348679\",\n            \"TYPE\": \"repository\"\n        }\n    ]\n}"
                },
                "defaultValue": ""
              },
              "input": {
                "JSONString": "=$activity[GetRepositoryInfo].Data"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Response\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"TYPE\":{\"type\":\"string\"}}}}}}",
                    "fe_metadata": "{\"Response\":[{\"ID\":\"12348679\",\"TYPE\":\"repository\"}]}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Response\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"TYPE\":{\"type\":\"string\"}}}}}}",
                    "fe_metadata": "{\n    \"Response\": [\n        {\n            \"ID\": \"12348679\",\n            \"TYPE\": \"repository\"\n        }\n    ]\n}"
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
                  "Response": "=$activity[BuildRepositoryInfo].Data.Response"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [],
          "output": [
            {
              "name": "Response",
              "type": "array",
              "schema": {
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"TYPE\":{\"type\":\"string\"}}}"
              }
            }
          ],
          "fe_metadata": {
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Response\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"TYPE\":{\"type\":\"string\"}}}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:List_Projects",
      "data": {
        "name": "List Projects",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "SystemProperties",
            "to": "GetRepository",
            "type": "default"
          },
          {
            "id": 2,
            "from": "GetRepository",
            "to": "ParseRepository",
            "type": "default"
          },
          {
            "id": 3,
            "from": "ParseRepository",
            "to": "ExecCommand",
            "type": "default"
          },
          {
            "id": 4,
            "from": "ExecCommand",
            "to": "ReadProjectInfo",
            "type": "default"
          },
          {
            "id": 5,
            "from": "ReadProjectInfo",
            "to": "ParseProjects",
            "type": "default"
          },
          {
            "id": 6,
            "from": "ParseProjects",
            "to": "UpsertProjects",
            "type": "default"
          },
          {
            "id": 7,
            "from": "UpsertProjects",
            "to": "Return",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "SystemProperties",
            "name": "SystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ScriptFolder": "=$property[\"System.BaseFolder\"] + \"/scripts\""
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}}}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GetRepository",
            "name": "GetRepository",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "GET",
                "timeout": "1000",
                "leftToken": "%",
                "rightToken": "%",
                "variablesDef": "[{\"Name\":\"RepositoryID\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "http://%ServiceLocatorIP%:10080/mops/locator/locate/%RepositoryID%",
                "Body": "",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "RepositoryID": "=$flow.RepositoryID"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"RepositoryID\":{\"type\":\"string\"},\"ServiceLocatorIP\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"RepositoryID\":\"2\",\"ServiceLocatorIP\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"RepositoryID\",\"Type\":\"String\"},{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ParseRepository",
            "name": "ParseRepository",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\": \"12348679\",\n    \"Type\": \"repository\",\n    \"URL\": \"https://api.github.com/users/SteveNY-Tibco/repos\",\n    \"Properties\": {}\n}"
                }
              },
              "input": {
                "JSONString": "=$activity[GetRepository].Data"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"12348679\",\"Type\":\"repository\",\"URL\":\"https://api.github.com/users/SteveNY-Tibco/repos\",\"Properties\":{}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\n    \"ID\": \"12348679\",\n    \"Type\": \"repository\",\n    \"URL\": \"https://api.github.com/users/SteveNY-Tibco/repos\",\n    \"Properties\": {}\n}"
                  }
                }
              }
            }
          },
          {
            "id": "ExecCommand",
            "name": "ExecCommand",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./list-projects.sh",
                      "Execution_1": "./list-projects.sh $RepositoryID$"
                    },
                    "SystemEnvs": "=$activity[ParseRepository].Data.Properties"
                  }
                },
                "Variables": {
                  "mapping": {
                    "RepositoryID": "=$flow.RepositoryID",
                    "ScriptFolder": "=$activity[SystemProperties].Data.ScriptFolder"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ReadProjectInfo",
            "name": "ReadProjectInfo",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=$activity[SystemProperties].Data.ScriptFolder + \"/\" + $flow.RepositoryID + \".json\""
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
            "id": "ParseProjects",
            "name": "ParseProjects",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Projects\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}}}",
                  "value": "",
                  "fe_metadata": "{ \"Projects\" : [{\n    \"ID\" : \"\",\n    \"RepositoryID\" : \"\",\n    \"Properties\" : {}\n}]}"
                }
              },
              "input": {},
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Projects\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}}}",
                    "fe_metadata": "{\"Projects\":[{\"ID\":\"\",\"RepositoryID\":\"\",\"Properties\":{}}]}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Projects\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}}}",
                    "fe_metadata": "{ \"Projects\" : [{\n    \"ID\" : \"\",\n    \"RepositoryID\" : \"\",\n    \"Properties\" : {}\n}]}"
                  }
                }
              }
            }
          },
          {
            "id": "UpsertProjects",
            "name": "UpsertProjects",
            "description": "A simple activity for upserting data to table",
            "type": "iterator",
            "settings": {
              "iterate": "=$activity[ParseProjects].Data.Projects",
              "accumulate": false
            },
            "activity": {
              "ref": "#tableupsert",
              "settings": {
                "Table": {
                  "id": "09120070-4b00-11eb-9ffb-8db6499bd1d7",
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
                      "value": "Projects"
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
                      "value": "[{\"Name\":\"ID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609375078303,
                  "user": "flogo",
                  "connectorName": "Projects",
                  "connectorDescription": " "
                }
              },
              "input": {
                "Mapping": "=$iteration[value]"
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"RepositoryID\":\"2\",\"Properties\":{}}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"RepositoryID\":\"2\",\"Properties\":{}}"
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
                  "Projects": "=$activity[ParseProjects].Data.Projects"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "RepositoryID",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "Projects",
              "type": "array",
              "schema": {
                "type": "json",
                "value": "{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}"
              }
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"RepositoryID\":{\"type\":\"string\"}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Projects\":{\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Chkout_Project",
      "data": {
        "name": "Chkout Project",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LoadProjectLog",
            "to": "GetProjectObj",
            "type": "default"
          },
          {
            "id": 2,
            "from": "GetProjectObj",
            "to": "SystemProperties",
            "type": "default"
          },
          {
            "id": 3,
            "from": "SystemProperties",
            "to": "GetRepositoryInfo",
            "type": "default"
          },
          {
            "id": 4,
            "from": "GetRepositoryInfo",
            "to": "BuildRepositoryObj",
            "type": "default"
          },
          {
            "id": 5,
            "from": "BuildRepositoryObj",
            "to": "CheckUsrScript",
            "type": "default"
          },
          {
            "id": 6,
            "from": "CheckUsrScript",
            "to": "ExecuteUsrScript",
            "type": "expression",
            "label": "CheckScripttoExecCommand02",
            "value": "null!=$activity[CheckUsrScript].Results \u0026\u0026 0!=array.count($activity[CheckUsrScript].Results)"
          },
          {
            "id": 7,
            "from": "ExecuteUsrScript",
            "to": "ExecUsrScript",
            "type": "default"
          },
          {
            "id": 8,
            "from": "ExecUsrScript",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 9,
            "from": "CheckUsrScript",
            "to": "ExecuteDefaultScript",
            "type": "exprOtherwise",
            "label": "CheckScript to CopyOfExecCommand02"
          },
          {
            "id": 10,
            "from": "ExecuteDefaultScript",
            "to": "ExecDefaultScript",
            "type": "default"
          },
          {
            "id": 11,
            "from": "ExecDefaultScript",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LoadProjectLog",
            "name": "LoadProjectLog",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Load Project ]]] ################## \" + $flow.RepositoryID + \"-\" + $flow.ID)",
                "addDetails": false
              }
            }
          },
          {
            "id": "GetProjectObj",
            "name": "GetProjectObj",
            "description": "A simple activity for quering data from table",
            "activity": {
              "ref": "#tablequery",
              "settings": {
                "Table": {
                  "id": "09120070-4b00-11eb-9ffb-8db6499bd1d7",
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
                      "value": "Projects"
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
                      "value": "[{\"Name\":\"ID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609375078303,
                  "user": "flogo",
                  "connectorName": "Projects",
                  "connectorDescription": " "
                },
                "Indices": "ID RepositoryID"
              },
              "input": {
                "QueryKey": {
                  "mapping": {
                    "ID": "=$flow.ID",
                    "RepositoryID": "=$flow.RepositoryID"
                  }
                }
              },
              "output": {
                "Exists": false
              },
              "schemas": {
                "input": {
                  "QueryKey": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"RepositoryID\":\"2\",\"Properties\":{}}"
                  }
                }
              }
            }
          },
          {
            "id": "SystemProperties",
            "name": "SystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectProperties\",\"Type\":\"Object\"},{\"Name\":\"ScriptsFolder_Default\",\"Type\":\"String\"},{\"Name\":\"ScriptsFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ScriptsFolder": "=$property[\"System.BaseFolder\"] + \"/scripts\"",
                    "ScriptsFolder_Default": "/home/scripts",
                    "ProjectFolder": "=$property[\"System.BaseFolder\"] + \"/projects/\" + $flow.ID",
                    "ProjectID": "=$flow.ID",
                    "ProjectProperties": "=f1.coalesceobject($flow.Properties, f1.getsubobject($activity[GetProjectObj].Data, \"Properties\"))"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"ProjectProperties\":{\"type\":\"object\",\"properties\":{}},\"ScriptsFolder_Default\":{\"type\":\"string\"},\"ScriptsFolder\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"ProjectProperties\":{},\"ScriptsFolder_Default\":\"2\",\"ScriptsFolder\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"ProjectProperties\":{\"type\":\"object\",\"properties\":{}},\"ScriptsFolder_Default\":{\"type\":\"string\"},\"ScriptsFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"ProjectProperties\":{},\"ScriptsFolder_Default\":\"2\",\"ScriptsFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectProperties\",\"Type\":\"Object\"},{\"Name\":\"ScriptsFolder_Default\",\"Type\":\"String\"},{\"Name\":\"ScriptsFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GetRepositoryInfo",
            "name": "GetRepositoryInfo",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "GET",
                "timeout": "1000",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "http://$ServiceLocatorIP$:10080/f1/locator/locate/repository/$RepositoryID$",
                "Body": "",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "RepositoryID": "=$flow.RepositoryID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRepositoryObj",
            "name": "BuildRepositoryObj",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\": \"12348679\",\n    \"Type\": \"repository\",\n    \"URL\": \"https://api.github.com/users/SteveNY-Tibco/repos\",\n    \"Properties\": {}\n}"
                }
              },
              "input": {
                "JSONString": "=$activity[GetRepositoryInfo].Data"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"12348679\",\"Type\":\"repository\",\"URL\":\"https://api.github.com/users/SteveNY-Tibco/repos\",\"Properties\":{}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\n    \"ID\": \"12348679\",\n    \"Type\": \"repository\",\n    \"URL\": \"https://api.github.com/users/SteveNY-Tibco/repos\",\n    \"Properties\": {}\n}"
                  }
                }
              }
            }
          },
          {
            "id": "CheckUsrScript",
            "name": "CheckUsrScript",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=$activity[SystemProperties].Data.ScriptsFolder + \"/load-project.sh\""
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
            "id": "ExecuteUsrScript",
            "name": "ExecuteUsrScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Execute Usr Script ]]] ################## \")",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecUsrScript",
            "name": "ExecUsrScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./load-project.sh",
                      "Execution_1": "./load-project.sh $RepositoryID$ $ProjectID$ $ProjectFolder$"
                    },
                    "SystemEnvs": "=f1.combineobjs($activity[SystemProperties].Data.ProjectProperties, $activity[BuildRepositoryObj].Data.Properties)"
                  }
                },
                "Variables": {
                  "mapping": {
                    "RepositoryID": "=$flow.RepositoryID",
                    "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder",
                    "ProjectID": "=$activity[SystemProperties].Data.ProjectID",
                    "ProjectFolder": "=$activity[SystemProperties].Data.ProjectFolder"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
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
                  "DataOut": "=$activity[ExecUsrScript].Message"
                }
              }
            }
          },
          {
            "id": "ExecuteDefaultScript",
            "name": "ExecuteDefaultScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Execute Default Script ]]] ################## \")",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecDefaultScript",
            "name": "ExecDefaultScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 1,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]",
                "SystemEnv": ""
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "./load-project.sh $RepositoryID$ $ProjectID$ $ProjectFolder$"
                    },
                    "SystemEnvs": "=f1.combineobjs($activity[SystemProperties].Data.ProjectProperties, $activity[BuildRepositoryObj].Data.Properties)"
                  }
                },
                "Variables": {
                  "mapping": {
                    "RepositoryID": "=$flow.RepositoryID",
                    "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder_Default",
                    "ProjectID": "=$activity[SystemProperties].Data.ProjectID",
                    "ProjectFolder": "=$activity[SystemProperties].Data.ProjectFolder"
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
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
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
                  "DataOut": "=$activity[ExecDefaultScript].Message"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "RepositoryID",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "Properties",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"RepositoryID\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Project",
      "data": {
        "name": "Project",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LoadProjectLog",
            "to": "SystemProperties",
            "type": "default"
          },
          {
            "id": 2,
            "from": "SystemProperties",
            "to": "ScriptEnvVariables",
            "type": "default"
          },
          {
            "id": 3,
            "from": "ScriptEnvVariables",
            "to": "LogBeforeDispatch",
            "type": "default"
          },
          {
            "id": 4,
            "from": "LogBeforeDispatch",
            "to": "Load",
            "type": "expression",
            "label": "LogtoStartaSubFlow",
            "value": "\"load\" == $flow.Action"
          },
          {
            "id": 5,
            "from": "Load",
            "to": "Return_Load",
            "type": "default"
          },
          {
            "id": 6,
            "from": "LogBeforeDispatch",
            "to": "Save",
            "type": "expression",
            "label": "LogtoStartaSubFlow",
            "value": "\"save\" == $flow.Action"
          },
          {
            "id": 7,
            "from": "Save",
            "to": "Return_Save",
            "type": "default"
          },
          {
            "id": 8,
            "from": "LogBeforeDispatch",
            "to": "Create",
            "type": "expression",
            "label": "LogtoStartaSubFlow",
            "value": "\"create\" == $flow.Action"
          },
          {
            "id": 9,
            "from": "Create",
            "to": "Return_Create",
            "type": "default"
          },
          {
            "id": 10,
            "from": "LogBeforeDispatch",
            "to": "Delete",
            "type": "expression",
            "label": "LogtoStartaSubFlow1",
            "value": "\"delete\" == $flow.Action"
          },
          {
            "id": 11,
            "from": "Delete",
            "to": "Return_Delete",
            "type": "default"
          },
          {
            "id": 12,
            "from": "LogBeforeDispatch",
            "to": "List",
            "type": "expression",
            "label": "LogtoStartaSubFlow2",
            "value": "\"list\" == $flow.Action"
          },
          {
            "id": 13,
            "from": "List",
            "to": "Return_List",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LoadProjectLog",
            "name": "LoadProjectLog",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Project Manager ] ################## Type = \" + $flow.Type + \", Action = \" + $flow.Action + \", Artifact = \" + $flow.Artifact + \"ID = \" + $flow.ID + \", Properties = \" + coerce.toString($flow.Properties))",
                "addDetails": false
              }
            }
          },
          {
            "id": "SystemProperties",
            "name": "SystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectProperties\",\"Type\":\"Object\"},{\"Name\":\"DefaultScriptsFolder\",\"Type\":\"String\"},{\"Name\":\"ScriptsFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ScriptsFolder": "=$property[\"System.BaseFolder\"] + \"/scripts\"",
                    "ProjectID": "=$flow.ID",
                    "ProjectProperties": "=$flow.Properties",
                    "DefaultScriptsFolder": "/home/scripts",
                    "ProjectsFolder": "=$property[\"System.BaseFolder\"] + \"/projects\""
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ProjectProperties\":{\"type\":\"object\",\"properties\":{}},\"DefaultScriptsFolder\":{\"type\":\"string\"},\"ScriptsFolder\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectsFolder\":\"2\",\"ProjectProperties\":{},\"DefaultScriptsFolder\":\"2\",\"ScriptsFolder\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ProjectProperties\":{\"type\":\"object\",\"properties\":{}},\"DefaultScriptsFolder\":{\"type\":\"string\"},\"ScriptsFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectsFolder\":\"2\",\"ProjectProperties\":{},\"DefaultScriptsFolder\":\"2\",\"ScriptsFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectProperties\",\"Type\":\"Object\"},{\"Name\":\"DefaultScriptsFolder\",\"Type\":\"String\"},{\"Name\":\"ScriptsFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "ScriptEnvVariables",
            "name": "ScriptEnvVariables",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectTemplateFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ProjectID": "=$activity[SystemProperties].Data.ProjectID",
                    "ProjectsFolder": "=$activity[SystemProperties].Data.ProjectsFolder",
                    "ProjectTemplateFolder": "=$property[\"System.ProjectTemplateFolder\"]"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ProjectTemplateFolder\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectsFolder\":\"2\",\"ProjectTemplateFolder\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ProjectTemplateFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectsFolder\":\"2\",\"ProjectTemplateFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectsFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectTemplateFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "LogBeforeDispatch",
            "name": "LogBeforeDispatch",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[ Log Before Dispatch ] SystemProperties = \" + coerce.toString($activity[SystemProperties].Data))",
                "addDetails": false
              }
            }
          },
          {
            "id": "Load",
            "name": "Load",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Execute_Script"
              },
              "input": {
                "Type": "=$flow.Type",
                "Action": "=$flow.Action",
                "Artifact": "=$flow.Artifact",
                "ID": "=$flow.ID",
                "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder",
                "DefaultScriptFolder": "=$activity[SystemProperties].Data.DefaultScriptsFolder",
                "ScriptEnvVariables": "=$activity[ScriptEnvVariables].Data"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"ScriptEnvVariables\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"ScriptEnvVariables\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return_Load",
            "name": "Return_Load",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "DataOut": "=$activity[Load].DataOut"
                }
              }
            }
          },
          {
            "id": "Save",
            "name": "Save",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Save"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return_Save",
            "name": "Return_Save",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "DataOut": "=$activity[Save].DataOut"
                }
              }
            }
          },
          {
            "id": "Create",
            "name": "Create",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Create"
              },
              "input": {
                "Type": "=$flow.Type",
                "Action": "=$flow.Action",
                "Artifact": "=$flow.Artifact",
                "ID": "=$flow.ID",
                "Properties": "=$activity[SystemProperties].Data.ProjectProperties",
                "ProjectsFolder": "=$activity[SystemProperties].Data.ProjectsFolder",
                "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder",
                "DefaultScriptFolder": "=$activity[SystemProperties].Data.DefaultScriptsFolder",
                "ScriptEnvVariables": "=$activity[ScriptEnvVariables].Data"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"ScriptEnvVariables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"ScriptEnvVariables\":{\"type\":\"object\",\"properties\":{}},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return_Create",
            "name": "Return_Create",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "DataOut": "=$activity[Create].DataOut"
                }
              }
            }
          },
          {
            "id": "Delete",
            "name": "Delete",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:Delete"
              },
              "input": {
                "Type": "=$flow.Type",
                "Action": "=$flow.Action",
                "Artifact": "=$flow.Artifact",
                "ID": "=$flow.ID",
                "ProjectsFolder": "=$activity[SystemProperties].Data.ProjectsFolder",
                "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder",
                "DefaultScriptFolder": "=$activity[SystemProperties].Data.DefaultScriptsFolder",
                "Properties": "=$flow.Properties"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return_Delete",
            "name": "Return_Delete",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "DataOut": "=$activity[Delete].DataOut"
                }
              }
            }
          },
          {
            "id": "List",
            "name": "List",
            "description": "Simple SubFlow Activity",
            "activity": {
              "ref": "#subflow",
              "settings": {
                "flowURI": "res://flow:List"
              },
              "input": {
                "Type": "=$flow.Type",
                "Action": "=$flow.Action",
                "Artifact": "=$flow.Artifact",
                "ID": "=$flow.ID",
                "Properties": "=$flow.Properties",
                "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder",
                "DefaultScriptFolder": "=$activity[SystemProperties].Data.DefaultScriptsFolder",
                "ProjectsFolder": "=$activity[SystemProperties].Data.ProjectsFolder"
              },
              "schemas": {
                "input": {
                  "input": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"ProjectsFolder\":{\"type\":\"string\"},\"ScriptFolder\":{\"type\":\"string\"},\"DefaultScriptFolder\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                },
                "output": {
                  "output": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
                  }
                }
              }
            }
          },
          {
            "id": "Return_List",
            "name": "Return_List",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "DataOut": "=$activity[List].DataOut"
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
              "name": "Action",
              "type": "string"
            },
            {
              "name": "Artifact",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "Properties",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Type\":{\"type\":\"string\"},\"Action\":{\"type\":\"string\"},\"Artifact\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Create_Project",
      "data": {
        "name": "Create Project",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LoadProjectLog",
            "to": "GetProjectObj",
            "type": "default"
          },
          {
            "id": 2,
            "from": "GetProjectObj",
            "to": "SystemProperties",
            "type": "default"
          },
          {
            "id": 3,
            "from": "SystemProperties",
            "to": "GetRepositoryInfo",
            "type": "default"
          },
          {
            "id": 4,
            "from": "GetRepositoryInfo",
            "to": "BuildRepositoryObj",
            "type": "default"
          },
          {
            "id": 5,
            "from": "BuildRepositoryObj",
            "to": "CheckUsrScript",
            "type": "default"
          },
          {
            "id": 6,
            "from": "CheckUsrScript",
            "to": "ExecuteUsrScript",
            "type": "expression",
            "label": "CheckScripttoExecCommand02",
            "value": "null!=$activity[CheckUsrScript].Results \u0026\u0026 0!=array.count($activity[CheckUsrScript].Results)"
          },
          {
            "id": 7,
            "from": "ExecuteUsrScript",
            "to": "ExecUsrScript",
            "type": "default"
          },
          {
            "id": 8,
            "from": "ExecUsrScript",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 9,
            "from": "CheckUsrScript",
            "to": "ExecuteDefaultScript",
            "type": "exprOtherwise",
            "label": "CheckScript to CopyOfExecCommand02"
          },
          {
            "id": 10,
            "from": "ExecuteDefaultScript",
            "to": "ExecDefaultScript",
            "type": "default"
          },
          {
            "id": 11,
            "from": "ExecDefaultScript",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LoadProjectLog",
            "name": "LoadProjectLog",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Load Project ]]] ################## \" + $flow.RepositoryID + \"-\" + $flow.ID)",
                "addDetails": false
              }
            }
          },
          {
            "id": "GetProjectObj",
            "name": "GetProjectObj",
            "description": "A simple activity for quering data from table",
            "activity": {
              "ref": "#tablequery",
              "settings": {
                "Table": {
                  "id": "09120070-4b00-11eb-9ffb-8db6499bd1d7",
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
                      "value": "Projects"
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
                      "value": "[{\"Name\":\"ID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609375078303,
                  "user": "flogo",
                  "connectorName": "Projects",
                  "connectorDescription": " "
                },
                "Indices": "ID RepositoryID"
              },
              "input": {
                "QueryKey": {
                  "mapping": {
                    "ID": "=$flow.ID",
                    "RepositoryID": "=$flow.RepositoryID"
                  }
                }
              },
              "output": {
                "Exists": false
              },
              "schemas": {
                "input": {
                  "QueryKey": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"RepositoryID\":\"2\",\"Properties\":{}}"
                  }
                }
              }
            }
          },
          {
            "id": "SystemProperties",
            "name": "SystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectProperties\",\"Type\":\"Object\"},{\"Name\":\"ScriptsFolder_Default\",\"Type\":\"String\"},{\"Name\":\"ScriptsFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ScriptsFolder": "=$property[\"System.BaseFolder\"] + \"/scripts\"",
                    "ScriptsFolder_Default": "/home/scripts",
                    "ProjectFolder": "=$property[\"System.BaseFolder\"] + \"/projects/\" + $flow.ID",
                    "ProjectID": "=$flow.ID",
                    "ProjectProperties": "=f1.coalesceobject($flow.Properties, f1.getsubobject($activity[GetProjectObj].Data, \"Properties\"))"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"ProjectProperties\":{\"type\":\"object\",\"properties\":{}},\"ScriptsFolder_Default\":{\"type\":\"string\"},\"ScriptsFolder\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"ProjectProperties\":{},\"ScriptsFolder_Default\":\"2\",\"ScriptsFolder\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"ProjectProperties\":{\"type\":\"object\",\"properties\":{}},\"ScriptsFolder_Default\":{\"type\":\"string\"},\"ScriptsFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"ProjectProperties\":{},\"ScriptsFolder_Default\":\"2\",\"ScriptsFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectProperties\",\"Type\":\"Object\"},{\"Name\":\"ScriptsFolder_Default\",\"Type\":\"String\"},{\"Name\":\"ScriptsFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GetRepositoryInfo",
            "name": "GetRepositoryInfo",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "GET",
                "timeout": "1000",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "http://$ServiceLocatorIP$:10080/f1/locator/locate/repository/$RepositoryID$",
                "Body": "",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "RepositoryID": "=$flow.RepositoryID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRepositoryObj",
            "name": "BuildRepositoryObj",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\": \"12348679\",\n    \"Type\": \"repository\",\n    \"URL\": \"https://api.github.com/users/SteveNY-Tibco/repos\",\n    \"Properties\": {}\n}"
                }
              },
              "input": {
                "JSONString": "=$activity[GetRepositoryInfo].Data"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"12348679\",\"Type\":\"repository\",\"URL\":\"https://api.github.com/users/SteveNY-Tibco/repos\",\"Properties\":{}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\n    \"ID\": \"12348679\",\n    \"Type\": \"repository\",\n    \"URL\": \"https://api.github.com/users/SteveNY-Tibco/repos\",\n    \"Properties\": {}\n}"
                  }
                }
              }
            }
          },
          {
            "id": "CheckUsrScript",
            "name": "CheckUsrScript",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=$activity[SystemProperties].Data.ScriptsFolder + \"/create-project.sh\""
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
            "id": "ExecuteUsrScript",
            "name": "ExecuteUsrScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Execute Usr Script ]]] ################## \")",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecUsrScript",
            "name": "ExecUsrScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./create-project.sh",
                      "Execution_1": "./create-project.sh $RepositoryID$ $ProjectID$ $ProjectFolder$"
                    },
                    "SystemEnvs": "=f1.combineobjs($activity[SystemProperties].Data.ProjectProperties, $activity[BuildRepositoryObj].Data.Properties)"
                  }
                },
                "Variables": {
                  "mapping": {
                    "RepositoryID": "=$flow.RepositoryID",
                    "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder",
                    "ProjectID": "=$activity[SystemProperties].Data.ProjectID",
                    "ProjectFolder": "=$activity[SystemProperties].Data.ProjectFolder"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
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
                  "DataOut": "=$activity[ExecUsrScript].Message"
                }
              }
            }
          },
          {
            "id": "ExecuteDefaultScript",
            "name": "ExecuteDefaultScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Execute Default Script ]]] ################## \")",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecDefaultScript",
            "name": "ExecDefaultScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./create-project.sh",
                      "Execution_1": "./create-project.sh $RepositoryID$ $ProjectID$ $ProjectFolder$"
                    },
                    "SystemEnvs": "=f1.combineobjs($activity[SystemProperties].Data.ProjectProperties, $activity[BuildRepositoryObj].Data.Properties)"
                  }
                },
                "Variables": {
                  "mapping": {
                    "RepositoryID": "=$flow.RepositoryID",
                    "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder_Default",
                    "ProjectID": "=$activity[SystemProperties].Data.ProjectID",
                    "ProjectFolder": "=$activity[SystemProperties].Data.ProjectFolder"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
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
                  "DataOut": "=$activity[ExecDefaultScript].Message"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "RepositoryID",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "Properties",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"RepositoryID\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
          }
        },
        "explicitReply": true
      }
    },
    {
      "id": "flow:Save_Project",
      "data": {
        "name": "Save Project",
        "description": "",
        "links": [
          {
            "id": 1,
            "from": "LoadProjectLog",
            "to": "GetProjectObj",
            "type": "default"
          },
          {
            "id": 2,
            "from": "GetProjectObj",
            "to": "SystemProperties",
            "type": "default"
          },
          {
            "id": 3,
            "from": "SystemProperties",
            "to": "GetRepositoryInfo",
            "type": "default"
          },
          {
            "id": 4,
            "from": "GetRepositoryInfo",
            "to": "BuildRepositoryObj",
            "type": "default"
          },
          {
            "id": 5,
            "from": "BuildRepositoryObj",
            "to": "CheckUsrScript",
            "type": "default"
          },
          {
            "id": 6,
            "from": "CheckUsrScript",
            "to": "ExecuteUsrScript",
            "type": "expression",
            "label": "CheckScripttoExecCommand02",
            "value": "null!=$activity[CheckUsrScript].Results \u0026\u0026 0!=array.count($activity[CheckUsrScript].Results)"
          },
          {
            "id": 7,
            "from": "ExecuteUsrScript",
            "to": "ExecUsrScript",
            "type": "default"
          },
          {
            "id": 8,
            "from": "ExecUsrScript",
            "to": "Return01",
            "type": "default"
          },
          {
            "id": 9,
            "from": "CheckUsrScript",
            "to": "ExecuteDefaultScript",
            "type": "exprOtherwise",
            "label": "CheckScript to CopyOfExecCommand02"
          },
          {
            "id": 10,
            "from": "ExecuteDefaultScript",
            "to": "ExecDefaultScript",
            "type": "default"
          },
          {
            "id": 11,
            "from": "ExecDefaultScript",
            "to": "Return02",
            "type": "default"
          }
        ],
        "tasks": [
          {
            "id": "LoadProjectLog",
            "name": "LoadProjectLog",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Load Project ]]] ################## \" + $flow.RepositoryID + \"-\" + $flow.ID)",
                "addDetails": false
              }
            }
          },
          {
            "id": "GetProjectObj",
            "name": "GetProjectObj",
            "description": "A simple activity for quering data from table",
            "activity": {
              "ref": "#tablequery",
              "settings": {
                "Table": {
                  "id": "09120070-4b00-11eb-9ffb-8db6499bd1d7",
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
                      "value": "Projects"
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
                      "value": "[{\"Name\":\"ID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\",\"IsKey\":\"yes\",\"Indexed\":\"no\"},{\"Name\":\"Properties\",\"Type\":\"Object\",\"IsKey\":\"no\",\"Indexed\":\"no\"}]"
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
                  "lastUpdatedTime": 1609375078303,
                  "user": "flogo",
                  "connectorName": "Projects",
                  "connectorDescription": " "
                },
                "Indices": "ID RepositoryID"
              },
              "input": {
                "QueryKey": {
                  "mapping": {
                    "ID": "=$flow.ID",
                    "RepositoryID": "=$flow.RepositoryID"
                  }
                }
              },
              "output": {
                "Exists": false
              },
              "schemas": {
                "input": {
                  "QueryKey": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"2\",\"RepositoryID\":\"2\",\"Properties\":{}}"
                  }
                }
              }
            }
          },
          {
            "id": "SystemProperties",
            "name": "SystemProperties",
            "description": "Mapping field from input to output",
            "activity": {
              "ref": "#mapping",
              "settings": {
                "IsArray": false,
                "MappingFields": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectProperties\",\"Type\":\"Object\"},{\"Name\":\"ScriptsFolder_Default\",\"Type\":\"String\"},{\"Name\":\"ScriptsFolder\",\"Type\":\"String\"}]"
              },
              "input": {
                "Mapping": {
                  "mapping": {
                    "ScriptsFolder": "=$property[\"System.BaseFolder\"] + \"/scripts\"",
                    "ScriptsFolder_Default": "/home/scripts",
                    "ProjectFolder": "=$property[\"System.BaseFolder\"] + \"/projects/\" + $flow.ID",
                    "ProjectID": "=$flow.ID",
                    "ProjectProperties": "=f1.coalesceobject($flow.Properties, f1.getsubobject($activity[GetProjectObj].Data, \"Properties\"))"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Mapping": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"ProjectProperties\":{\"type\":\"object\",\"properties\":{}},\"ScriptsFolder_Default\":{\"type\":\"string\"},\"ScriptsFolder\":{\"type\":\"string\"},\"SkipCondition\":{\"type\":\"boolean\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"ProjectProperties\":{},\"ScriptsFolder_Default\":\"2\",\"ScriptsFolder\":\"2\",\"SkipCondition\":false}"
                  }
                },
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"ProjectProperties\":{\"type\":\"object\",\"properties\":{}},\"ScriptsFolder_Default\":{\"type\":\"string\"},\"ScriptsFolder\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"ProjectProperties\":{},\"ScriptsFolder_Default\":\"2\",\"ScriptsFolder\":\"2\"}"
                  }
                },
                "settings": {
                  "MappingFields": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectProperties\",\"Type\":\"Object\"},{\"Name\":\"ScriptsFolder_Default\",\"Type\":\"String\"},{\"Name\":\"ScriptsFolder\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "GetRepositoryInfo",
            "name": "GetRepositoryInfo",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "GET",
                "timeout": "1000",
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
              },
              "input": {
                "URL": "http://$ServiceLocatorIP$:10080/f1/locator/locate/repository/$RepositoryID$",
                "Body": "",
                "SkipCondition": false,
                "Variables": {
                  "mapping": {
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]",
                    "RepositoryID": "=$flow.RepositoryID"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ServiceLocatorIP\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ServiceLocatorIP\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ServiceLocatorIP\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
                  }
                }
              }
            }
          },
          {
            "id": "BuildRepositoryObj",
            "name": "BuildRepositoryObj",
            "description": "JSON Deserializer Activity",
            "activity": {
              "ref": "#jsondeserializer",
              "settings": {
                "schemaFromfile": false,
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\": \"12348679\",\n    \"Type\": \"repository\",\n    \"URL\": \"https://api.github.com/users/SteveNY-Tibco/repos\",\n    \"Properties\": {}\n}"
                }
              },
              "input": {
                "JSONString": "=$activity[GetRepositoryInfo].Data"
              },
              "schemas": {
                "output": {
                  "Data": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"ID\":\"12348679\",\"Type\":\"repository\",\"URL\":\"https://api.github.com/users/SteveNY-Tibco/repos\",\"Properties\":{}}"
                  }
                },
                "settings": {
                  "dataSample": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\n    \"ID\": \"12348679\",\n    \"Type\": \"repository\",\n    \"URL\": \"https://api.github.com/users/SteveNY-Tibco/repos\",\n    \"Properties\": {}\n}"
                  }
                }
              }
            }
          },
          {
            "id": "CheckUsrScript",
            "name": "CheckUsrScript",
            "description": "This activity write incoming object to file system",
            "activity": {
              "ref": "#filereader",
              "settings": {
                "BaseFolder": ""
              },
              "input": {
                "FilePattern": "=$activity[SystemProperties].Data.ScriptsFolder + \"/save-project.sh\""
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
            "id": "ExecuteUsrScript",
            "name": "ExecuteUsrScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Execute Usr Script ]]] ################## \")",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecUsrScript",
            "name": "ExecUsrScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./save-project.sh",
                      "Execution_1": "./save-project.sh $RepositoryID$ $ProjectID$ $ProjectFolder$"
                    },
                    "SystemEnvs": "=f1.combineobjs($activity[SystemProperties].Data.ProjectProperties, $activity[BuildRepositoryObj].Data.Properties)"
                  }
                },
                "Variables": {
                  "mapping": {
                    "RepositoryID": "=$flow.RepositoryID",
                    "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder",
                    "ProjectID": "=$activity[SystemProperties].Data.ProjectID",
                    "ProjectFolder": "=$activity[SystemProperties].Data.ProjectFolder"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
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
                  "DataOut": "=$activity[ExecUsrScript].Message"
                }
              }
            }
          },
          {
            "id": "ExecuteDefaultScript",
            "name": "ExecuteDefaultScript",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "=coerce.toString(\"[[[ Project Manager : Execute Default Script ]]] ################## \")",
                "addDetails": false
              }
            }
          },
          {
            "id": "ExecDefaultScript",
            "name": "ExecDefaultScript",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#exec",
              "settings": {
                "workingFolder": "$ScriptFolder$",
                "numOfExecutions": 2,
                "leftToken": "$",
                "rightToken": "$",
                "variablesDef": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
              },
              "input": {
                "Asynchronous": "=false",
                "SkipCondition": false,
                "Executable": {
                  "mapping": {
                    "Executions": {
                      "Execution_0": "chmod +x ./save-project.sh",
                      "Execution_1": "./save-project.sh $RepositoryID$ $ProjectID$ $ProjectFolder$"
                    },
                    "SystemEnvs": "=f1.combineobjs($activity[SystemProperties].Data.ProjectProperties, $activity[BuildRepositoryObj].Data.Properties)"
                  }
                },
                "Variables": {
                  "mapping": {
                    "RepositoryID": "=$flow.RepositoryID",
                    "ScriptFolder": "=$activity[SystemProperties].Data.ScriptsFolder_Default",
                    "ProjectID": "=$activity[SystemProperties].Data.ProjectID",
                    "ProjectFolder": "=$activity[SystemProperties].Data.ProjectFolder"
                  }
                }
              },
              "schemas": {
                "input": {
                  "Executable": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"Executions\":{\"type\":\"object\",\"properties\":{\"Execution_0\":{\"type\":\"string\"},\"Execution_1\":{\"type\":\"string\"}}},\"SystemEnvs\":{\"type\":\"object\",\"properties\":{}}}}",
                    "fe_metadata": "{\"Executions\":{\"Execution_0\":\"2\",\"Execution_1\":\"2\"},\"SystemEnvs\":{}}"
                  },
                  "Variables": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ScriptFolder\":{\"type\":\"string\"},\"ProjectID\":{\"type\":\"string\"},\"ProjectFolder\":{\"type\":\"string\"},\"RepositoryID\":{\"type\":\"string\"}}}",
                    "fe_metadata": "{\"ScriptFolder\":\"2\",\"ProjectID\":\"2\",\"ProjectFolder\":\"2\",\"RepositoryID\":\"2\"}"
                  }
                },
                "settings": {
                  "variablesDef": {
                    "type": "json",
                    "value": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"array\",\"items\":{\"type\":\"object\",\"properties\":{\"Name\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"}},\"required\":[\"Name\",\"Type\"]}}",
                    "fe_metadata": "[{\"Name\":\"ScriptFolder\",\"Type\":\"String\"},{\"Name\":\"ProjectID\",\"Type\":\"String\"},{\"Name\":\"ProjectFolder\",\"Type\":\"String\"},{\"Name\":\"RepositoryID\",\"Type\":\"String\"}]"
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
                  "DataOut": "=$activity[ExecDefaultScript].Message"
                }
              }
            }
          }
        ],
        "metadata": {
          "input": [
            {
              "name": "RepositoryID",
              "type": "string"
            },
            {
              "name": "ID",
              "type": "string"
            },
            {
              "name": "Properties",
              "type": "object"
            }
          ],
          "output": [
            {
              "name": "DataOut",
              "type": "object"
            }
          ],
          "fe_metadata": {
            "input": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"RepositoryID\":{\"type\":\"string\"},\"ID\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
            "output": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"DataOut\":{\"type\":\"object\",\"properties\":{}}}}"
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
            "to": "PingAndRegister",
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
                "sample": "",
                "dataSample": {
                  "metadata": "{\"$schema\":\"http://json-schema.org/draft-04/schema#\",\"type\":\"object\",\"properties\":{\"ID\":{\"type\":\"string\"},\"Type\":{\"type\":\"string\"},\"URL\":{\"type\":\"string\"},\"Properties\":{\"type\":\"object\",\"properties\":{}}}}",
                  "value": "",
                  "fe_metadata": "{\n    \"ID\": \"\",\n    \"Type\": \"\",\n    \"URL\": \"\",\n    \"Properties\": {}\n}"
                }
              },
              "input": {
                "Data": {
                  "mapping": {
                    "ID": "=\"projectmgr_\" + $property[\"System.ServiceLocatorIP\"]",
                    "Type": "projectmgr",
                    "URL": "=\"http://\"+$property[\"System.ServiceLocatorIP\"]+\":10090/f1/projectmgr\""
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
            "id": "PingAndRegister",
            "name": "PingAndRegister",
            "description": "This activity build docker image from Dockerfile",
            "activity": {
              "ref": "#httpclient",
              "settings": {
                "method": "POST",
                "timeout": "500",
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
                    "ServiceType": "projectmgr",
                    "ServiceLocatorIP": "=$property[\"System.ServiceLocatorIP\"]"
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
    }
  ]
}`
const engineJSON string = ``

func init () {
	cfgJson = flogoJSON
	cfgEngine = engineJSON
}

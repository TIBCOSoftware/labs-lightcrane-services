from __future__ import print_function
from concurrent import futures

from flask import Flask
from flask_restful import Resource, Api, reqparse, abort
from threading import Thread
import http.client

import io
import os
import time
import importlib
import json

try:
    from urllib.parse import urlparse
except ImportError:
     from urlparse import urlparse

import logging
import sys

# Get the list of user's 
# environment variables 
env_var = os.environ

app = Flask(__name__)
api = Api(app)

recognitionArgs = reqparse.RequestParser()
recognitionArgs.add_argument('Data', type=str, help='Data is required', required=True)
if 'PythonModel_plugin' in os.environ :
    module = importlib.import_module(os.environ['PythonModel_plugin'])
else :
    module = importlib.import_module('artifacts.inference')

inference = module.Inference(env_var)

def abortOnInvalidNetwork(network):
    if network == 'abc':
        abort(404, message='Invalid Network...')

class InferenceService(Resource):
    def __init__(self):
        self.inference = inference
        self.handle_raw_data = inference.handle_raw_data()
        
    def post(self):
        args = recognitionArgs.parse_args()
        data = None
        if False==self.handle_raw_data :
            data = json.loads(args['Data'])
            #print('(InferenceService.HandleData) type = {}, data = {}'.format(type(data), data))
        else :
            data = args['Data']
            #print('(InferenceService.HandleData.Raw) type = {}, data = {}'.format(type(data), data))
        predict = self.inference.evaluate(data)
        return predict, 200

class MetadataService(Resource):
    def __init__(self):
        self.metadata = None

    def get(self):
        if None == self.metadata :
            self.metadata = open('/app/data/model-metadata.json', 'r').read()
        return self.metadata, 200

class EchoService(Thread):
    def __init__(self, config):
        super().__init__()
        body = {
            'ID': config['System.ID'],
            'Type': 'inference',
            'URL': 'http://{}:{}'.format(config['System_ExternalEndpointIP'], config['System_Port']),
            'Properties': {
                'URI' : {
                    'Service':'/f1/data',
                    'Metadata':'/f1/metadata'
                }
            }
        }
        self.config = config
        self.headers = {'Content-type': 'application/json'}
        self.bodyStr = json.dumps(body)

    def run(self):
        print("EchoService starting ...")
        self.connection = self.connectToServiceLocator('host.docker.internal:5408')
        while True :
            print("EchoService wake up, Pining [{}]...".format(self.url))
            try :
                self.connection.request('POST', self.urlComponents[2], self.bodyStr, self.headers)
                response = self.connection.getresponse()
                result = response.read().decode()
                print('(EchoService.run) Result = {} '.format(result))
            except  BaseException as ex:
                print('(EchoService.run) Error to connect to [{}]! Will try later! '.format(self.url))
                self.connection.close()
                self.connection = self.connectToServiceLocator(self.config['System.ServiceLocator'])
                print('(EchoService.run) Will try [{}]! '.format(self.url))
            time.sleep(60)
        print("EchoService finishing ...")
        
    def connectToServiceLocator(self, locator) :
        self.url = 'http://{}/f1/locator/register/inference'.format(locator)
        self.urlComponents = urlparse(self.url)
        return http.client.HTTPConnection(self.urlComponents[1])


api.add_resource(InferenceService, '/f1/data')
api.add_resource(MetadataService, '/f1/metadata')

def serve() :
    if 'System.EchoOn' in os.environ and 'True' == os.environ['System.EchoOn'] :
        echo = EchoService(os.environ)
        echo.start()
        
    if 'System_Port' in os.environ:
        server_port = os.environ['System_Port']
    else :
        server_port = 10010
        
    app.run(host='0.0.0.0', port=server_port)

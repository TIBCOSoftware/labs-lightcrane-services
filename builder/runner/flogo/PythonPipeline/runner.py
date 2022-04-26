# Python Runner 2.0

from __future__ import print_function
from concurrent import futures

import grpc

import pipecoupler_pb2
import pipecoupler_pb2_grpc

import os
import os.path
import importlib
import inspect
import json
import logging

from f1 import serviceLocator as locator
from logging.handlers import RotatingFileHandler

# environment variables 
env_var = os.environ

logging.basicConfig(
    level=logging.DEBUG,
    format="%(asctime)s [%(levelname)s] %(message)s",
    handlers=[
        RotatingFileHandler("/tmp/files/log/air_pipeline.log", maxBytes=250000, backupCount=5),
        logging.StreamHandler()
    ]
)

class ProcessorService():
    def __init__(self):
        print('(ProcessorService.__init__) entering ...... ')
        print('(ProcessorService.__init__) parent folder=%s' % (os.listdir('..')))
        print('(ProcessorService.__init__) current folder=%s' % (os.listdir('.')))

        self.logger = logging.getLogger('ProcessorService')
          
#        env_var['module_prefix'] = 'artifacts'
        globleContext = { 'env_var': env_var}
        if True == os.path.exists('./config.json'):
            with open("./config.json", "r") as file:
                config = json.load(file)
                for key in config:
                    globleContext[key] = config[key]
        
        self.engine = locator.factory.getLocator({'type':'local'}).locate('pipeline_engine_embeded', 'pipeline_engine_local', globleContext)
        for key in env_var:
            if key.endswith('Python_Context'):
                processor = json.loads(env_var[key])
                print('(ProcessorService.__init__) register : {}'.format(env_var[key]))
                self.engine.register(processor)

        print('(ProcessorService.__init__) done ...... ')

    def HandleData(self, request, context):
        print('(ProcessorService.HandleData) Received request: request=%s' % (request))
        ID = request.ID
        print('(ProcessorService.HandleData) Received request: ID=%s' % (ID))
        payload = json.loads(request.content)
        
        print('(ProcessorService.HandleData) before process type = {}, payload = {}'.format(type(payload), payload))
        self.logger.debug('(ProcessorService.HandleData) before process type = {}, payload = {}'.format(type(payload), payload))
        data = {}
        if 'data' in payload:
            data = payload['data']        
        response = self.engine.handleData(ID, payload['command'], data)
        print('(ProcessorService.HandleData) after process type = {}, response = {}'.format(type(response), response))
        self.logger.debug('(ProcessorService.HandleData) after process type = {}, response = {}'.format(type(response), response))
        
        return pipecoupler_pb2.Reply(sender='PipelineService', ID=ID, content=json.dumps(response), status=response['isDone'])

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pipecoupler_pb2_grpc.add_PipeCouplerServicer_to_server(ProcessorService(), server)
    grpcport = '9998'
    server.add_insecure_port('[::]:%s' % grpcport)
    server.start()
    print('(ProcessorService.serve) started ....  ')
    server.wait_for_termination()
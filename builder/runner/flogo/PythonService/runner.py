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

#from f1 import serviceLocator as locator
from logging.handlers import RotatingFileHandler

# environment variables 
env_var = os.environ

logging.basicConfig(
    level=logging.DEBUG,
    format="%(asctime)s [%(levelname)s] %(message)s",
    handlers=[
#        logging.FileHandler("/home/log/air_pipeline.log"),
        RotatingFileHandler("/home/log/air_pipeline.log", maxBytes=250000, backupCount=5),
        logging.StreamHandler()
    ]
)

class ProcessorService():
    def __init__(self):
        print('(ProcessorService.__init__) entering ...... ')
        self.processors = {}
        print('(ProcessorService.__init__) files ={}'.format(os.listdir('.')))
        
        for key in env_var:
            if key.endswith('Python_Processor'):
                filename = env_var[key]
                print('(ProcessorService.__init__) register : {}'.format(filename))
                processorName = filename[0:len(filename)-3]
                print('(ProcessorService.__init__) processorName : {}'.format(processorName))
                module = importlib.import_module(processorName)
                processor = module.Processor(os.environ)
                self.processors['{}.py'.format(processorName)] = processor

        print('(ProcessorService.__init__) self.processors = {}'.format(self.processors))
        print('(ProcessorService.__init__) done ...... ')

    def HandleData(self, request, context):
        print('(ProcessorService.HandleData) Received request: request=%s' % (request))
        ID = request.ID
        print('(ProcessorService.HandleData) Received request: ID=%s' % (ID))
        data = json.loads(request.content)
        print('(ProcessorService.HandleData) type = {}, data = {}'.format(type(data), data))
        processor = self.processors[ID]
        result = processor.evaluate(data)
        return pipecoupler_pb2.Reply(sender='ProcessorService', ID=ID, content=json.dumps(result), status=True)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pipecoupler_pb2_grpc.add_PipeCouplerServicer_to_server(ProcessorService(), server)
    grpcport = '9998'
    server.add_insecure_port('[::]:%s' % grpcport)
    server.start()
    print('(ProcessorService.serve) started ....  ')
    server.wait_for_termination()
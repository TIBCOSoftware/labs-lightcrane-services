# Python Runner 2.0

from __future__ import print_function
from concurrent import futures
import logging

import grpc

import pipecoupler_pb2
import pipecoupler_pb2_grpc

import os
import importlib
import json

import pprint 

import sys
print("Python version")
print (sys.version)
print("Version info.")
print (sys.version_info)

# Get the list of user's 
# environment variables 
env_var = os.environ 
  
# Print the list of user's 
# environment variables 
print("User's Environment variable:") 
pprint.pprint(dict(env_var), width = 1) 


class InferenceService(pipecoupler_pb2_grpc.PipeCouplerServicer):
    def __init__(self):
        self.module = importlib.import_module(os.environ['PythonModel_plugin'])
        self.inference = self.module.Inference(env_var)
        try:
            self.downstreamHosts = os.environ['pipecoupler_downstreamHosts']
            self.downstreamHost = self.downstreamHosts[2:len(self.downstreamHosts)-2]
        except KeyError:
            self.downstreamHosts = None
        self.downstreamPort = os.environ['pipecoupler_port']


    def HandleData(self, request, context):
        print('(InferenceService.HandleData) Received request: request=%s' % (request))
        
        result = self.inference.evaluate(json.loads(request.content))
        
        if None != self.downstreamHosts :
            print('(InferenceService.HandleData) downstreamHost = %s, downstreamPort = %s ' % (self.downstreamHost, self.downstreamPort))
            with grpc.insecure_channel('%s:%s' % (self.downstreamHost, self.downstreamPort)) as channel:
                stub = pipecoupler_pb2_grpc.PipeCouplerStub(channel)
                response = stub.HandleData(pipecoupler_pb2.Data(sender='you', ID='', content=json.dumps(result)))
                print("(InferenceService.HandleData) Replyed data received: " + response.content)
            return pipecoupler_pb2.Reply(sender='', ID='', content=response.content, status=True)
        else :
            return pipecoupler_pb2.Reply(sender='', ID='', content=json.dumps(result), status=True)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pipecoupler_pb2_grpc.add_PipeCouplerServicer_to_server(InferenceService(), server)
    grpcport = os.environ['pipecoupler_port']
    server.add_insecure_port('[::]:%s' % grpcport)
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()

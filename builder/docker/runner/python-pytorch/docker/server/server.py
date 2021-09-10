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

module = importlib.import_module(os.environ['PythonModel_plugin'])
inference = module.Inference(env_var)
downstreamHosts = os.environ['pipecoupler_downstreamHosts']
downstreamHosts = downstreamHosts[2:len(downstreamHosts)-2]
grpcport = os.environ['pipecoupler_port']

class InferenceService(pipecoupler_pb2_grpc.PipeCouplerServicer):

    def HandleData(self, request, context):
        print('Received request: request=%s' % (request))
        
        result = inference.evaluate(json.loads(request.content))
        
        print('downstreamHosts = %s, grpcport = %s ' % (downstreamHosts, grpcport))
        with grpc.insecure_channel('%s:%s' % (downstreamHosts, grpcport)) as channel:
            stub = pipecoupler_pb2_grpc.PipeCouplerStub(channel)
            response = stub.HandleData(pipecoupler_pb2.Data(sender='you', ID='', content=json.dumps(result)))
            print("Replyed data received: " + response.content)
        
        return pipecoupler_pb2.Reply(sender='', ID='', content=response.content, status=True)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pipecoupler_pb2_grpc.add_PipeCouplerServicer_to_server(InferenceService(), server)
    server.add_insecure_port('[::]:%s' % grpcport)
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()

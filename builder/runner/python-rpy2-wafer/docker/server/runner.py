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

# Get the list of user's 
# environment variables 
env_var = os.environ

class InferenceService(pipecoupler_pb2_grpc.PipeCouplerServicer):
    def __init__(self):
        self.module = importlib.import_module(os.environ['PythonModel_plugin'])
        self.inference = self.module.Inference(env_var)
        self.handle_raw_data = self.inference.handle_raw_data()
        try:
            downstreamHosts = os.environ['pipecoupler_downstreamHosts']
            if None != downstreamHosts :
                #self.downstreamHost = self.downstreamHosts[2:len(self.downstreamHosts)-2]
                self.downstreamHosts = json.loads(downstreamHosts)
            else :
                self.downstreamHosts = []
        except KeyError:
            self.downstreamHost = None
        self.downstreamPort = os.environ['pipecoupler_port']

    def HandleData(self, request, context):
        #print('(InferenceService.HandleData) Received request: request=%s' % (request))

        data = None
        if False==self.handle_raw_data :
            data = json.loads(request.content)
            #print('(InferenceService.HandleData) type = {}, data = {}'.format(type(data), data))
        else :
            data = request.content
            #print('(InferenceService.HandleData.Raw) type = {}, data = {}'.format(type(data), data))
            
        result = self.inference.evaluate(data)
        
        if 0!=len(self.downstreamHosts):
            print('(InferenceService.HandleData) downstreamHosts[0] = %s, downstreamPort = %s ' % (self.downstreamHosts[0], self.downstreamPort))
            with grpc.insecure_channel('%s:%s' % (self.downstreamHosts[0], self.downstreamPort)) as channel:
                stub = pipecoupler_pb2_grpc.PipeCouplerStub(channel)
                response = stub.HandleData(pipecoupler_pb2.Data(sender='you', ID='', content=json.dumps(result)))
                print("(InferenceService.HandleData) Replyed data received: " + response.content)
            return pipecoupler_pb2.Reply(sender='', ID='', content=response.content, status=True)
        else :
            return pipecoupler_pb2.Reply(sender='', ID='', content=json.dumps(result), status=True)

def serve():
    logging.basicConfig()
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pipecoupler_pb2_grpc.add_PipeCouplerServicer_to_server(InferenceService(), server)
    grpcport = os.environ['pipecoupler_port']
    server.add_insecure_port('[::]:%s' % grpcport)
    server.start()
    server.wait_for_termination()

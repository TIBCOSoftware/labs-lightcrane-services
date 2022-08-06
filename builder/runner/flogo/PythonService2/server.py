import os
import os.path
import logging
from logging.handlers import RotatingFileHandler

import sys
print("Python version")
print (sys.version)
print("Version info.")
print (sys.version_info)
print("Path info")
print (sys.path)

if __name__ == '__main__':
    env_var = os.environ
    log_path = '/tmp/files/log'
    if 'FLOGO_LOG_LEVEL' in env_var :
        level=None
        if 'debug'==env_var['FLOGO_LOG_LEVEL'].lower() :
            level=logging.DEBUG
        elif 'info'==env_var['FLOGO_LOG_LEVEL'].lower() :
            level=logging.INFO
        
        handlers = [
            logging.StreamHandler()
        ]
        
        if os.path.isdir(log_path) :
            handlers.append(RotatingFileHandler(
                '{}/air_pipeline.log'.format(log_path), 
                maxBytes=10240000, 
                backupCount=5)
            )
        
        if None!=level :
            logging.basicConfig(
                level=level,
                format="%(asctime)s [%(levelname)s] %(message)s",
                handlers=handlers
            )

    if 'System_Standalone' in env_var and 'True' == env_var['System_Standalone'] :
        import runner_standalone
        print('Start with HTTP mode ....')
        runner_standalone.serve()
    else :
        import runner
        print('Start with gRPC mode ....')
        runner.serve()
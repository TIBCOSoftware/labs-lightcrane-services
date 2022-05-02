import os
import pprint 

import sys
print("Python version")
print (sys.version)
print("Version info.")
print (sys.version_info)
print("Path info")
print (sys.path)

env_var = os.environ
# Print the list of user's 
# environment variables 
#print("User's Environment variable:") 
#pprint.pprint(dict(env_var), width = 1) 

# Setup working folder
#working_folder = '{}/artifacts'.format(os.getcwd())
#if "Working_Folder" not in env_var:
#    env_var['Working_Folder'] = working_folder
#os.chdir(working_folder)

if __name__ == '__main__':
    if 'System_Standalone' in env_var and 'True' == env_var['System_Standalone'] :
        import runner_standalone
        print('Start with standalone mode ....')
        runner_standalone.serve()
    else :
        import runner
        runner.serve()
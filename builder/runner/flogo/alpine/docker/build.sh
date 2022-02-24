docker build --tag bigoyang/pythonservice-base:0.1.0 -f ./Dockerfile_PythonService_Base .
docker push bigoyang/pythonservice-base:0.1.0

docker build --tag bigoyang/python-base:0.1.0 -f ./Dockerfile_Python_Base .
docker push bigoyang/python-base:0.1.0
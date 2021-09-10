mkdir ./server
cp -R ../../server .
docker build --tag bigoyang/python-model:0.4.0 .
docker push bigoyang/python-model:0.4.0
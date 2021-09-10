docker run \
--name tensorflow-py-model-runner \
--env pipecoupler_port=5000 \
--env PythonModel_plugin=artifacts.inference \
--env PythonModel_URI=/fashion-mnist \
-v /Users/steven/Applications/Dev-Ops/app/local-projects/fashion-mnist-tensorflow-keras/artifacts:/app/artifacts \
-p 5000:5000 \
bigoyang/tensorflow-python-model:0.2.0
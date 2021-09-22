docker rm tensorflow-py-model-runner
docker run \
--name tensorflow-py-model-runner \
--env-file env.list \
-v /Users/steven/Applications/Dev-Ops/app/local-projects/fashion-mnist-tensorflow-keras/artifacts:/app/artifacts \
-p 5000:5000 \
bigoyang/tensorflow-python-model:0.2.0
docker rm rpy2-model-runner
docker run --name rpy2-model-runner \
--env-file env.list \
-v /Users/steven/Applications/Project-F1/projects/Air-F1_wafermaps-rpy2/artifacts:/app/artifacts \
-p 5000:5000 \
bigoyang/rpy2-python-model:0.3.0
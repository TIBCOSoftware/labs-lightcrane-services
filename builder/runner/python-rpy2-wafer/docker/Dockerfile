#FROM rpy2/base-ubuntu
FROM bigoyang/rpy2-wafermaps-model:0.3.0 

RUN pip install --upgrade pip

COPY ./server /app
WORKDIR /app

RUN pip install -r requirements.txt

CMD ["python","-u","./server.py"]
FROM atlasml/ml-base
#FROM bigoyang/python-ml-base

RUN pip install --upgrade pip

COPY ./server /app

WORKDIR /app

RUN pip install -r requirements.txt

CMD ["python","-u","./server.py"]
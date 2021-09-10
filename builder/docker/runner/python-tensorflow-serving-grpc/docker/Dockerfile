#FROM bigoyang/tensorflow-serving-python-base
FROM bigoyang/python-model:0.4.0

RUN /usr/bin/python -m pip install --upgrade pip

COPY ./server /app
WORKDIR /app

RUN pip install -r requirements.txt

CMD ["python","-u","./server.py"]
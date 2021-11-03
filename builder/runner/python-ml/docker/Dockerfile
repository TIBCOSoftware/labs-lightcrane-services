FROM python:3.7

RUN apt-get update

RUN apt-get install -y libsndfile1-dev

RUN pip install --upgrade pip

COPY ./server /app

WORKDIR /app

RUN pip install -r requirements.txt

CMD ["python","-u","./server.py"]
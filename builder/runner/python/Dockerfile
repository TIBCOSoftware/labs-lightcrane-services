FROM pytorch/pytorch

RUN conda install -y opencv

RUN apt-get update

RUN pip install --upgrade pip
RUN pip install Flask
RUN pip install Flask-RESTful
RUN pip install waitress

COPY ./server /app

WORKDIR /app

RUN pip install -r requirements.txt

CMD ["python","-u","./server.py"]
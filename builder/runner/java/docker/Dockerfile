# Pull base image.
FROM ubuntu:latest

RUN \
# Update
apt-get update -y && \
# Install Java
apt-get install default-jre -y

COPY ./server /app
WORKDIR /app

#EXPOSE 8080

CMD ["./run.sh"]
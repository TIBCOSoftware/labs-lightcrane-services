FROM public.ecr.aws/tibcolabs/labs-air-flogo-builder:1.16-stretch AS flogobuilder

ARG app_type

COPY build/deployer.json .

RUN chmod a+x create_build.sh;\
    ./create_build.sh deployer.json;
###################################################################################
#                         Flogo App image                         #
###################################################################################
FROM ubuntu:20.04

RUN apt-get update -y
RUN apt-get install -y apt-transport-https ca-certificates curl software-properties-common
RUN add-apt-repository ppa:deadsnakes/ppa 
RUN apt-get install -y zip

#================================== Install python3 ===============================

RUN apt-get update -y
RUN apt-get install -y python3.8 python3-pip
RUN pip3 install PyYAML==5.4.1

#================================== Install Docker ================================

RUN curl -fsSL https:/download.docker.com/linux/ubuntu/gpg | apt-key add - 
RUN apt-key fingerprint 0EBFCD88 
RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
RUN apt-get install -y docker-ce docker-ce-cli containerd.io

#============================== Install Docker-compose ============================

RUN curl -L "https://github.com/docker/compose/releases/download/1.29.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose

#================================ Open Horizon CLI ================================

ADD ./plugins/horizon-cli_2.29.0-457_amd64.deb /root/
RUN apt-get install /root/horizon-cli_2.29.0-457_amd64.deb

#================================= Flogo Builder ==================================

#For FE case, set FLOGO_HTTP_SERVICE_PORT to empty to disable heathcheck
ENV FLOGO_LOG_DTFORMAT="15:04:05.000" \
    STOP_ENGINE_ON_ERROR=flase \
    FLOGO_HTTP_SERVICE_PORT="7777"

ADD ./scripts /home/scripts/
ADD ./template /usr/flogo/template/
COPY --from=flogobuilder /home/app/bin/app /opt/tibco/flogo/1.0/bin/

WORKDIR /opt/tibco/flogo/1.0/bin/

RUN chmod a+x /opt/tibco/flogo/1.0/bin/*

CMD ["/opt/tibco/flogo/1.0/bin/app"]

VOLUME /opt/tibco/flogo/1.0/bin

EXPOSE 7777 9999
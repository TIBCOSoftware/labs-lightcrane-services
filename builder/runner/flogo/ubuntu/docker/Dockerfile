###################################################################################
#                         Flogo App arm64 image                         #
###################################################################################
FROM arm64v8/ubuntu:18.04

RUN apt-get update && \
    apt-get install ca-certificates -y &&\
#    apt-get apt install bash-completion && \
    apt-get clean

#For FE case, set FLOGO_HTTP_SERVICE_PORT to empty to disable heathcheck
ENV FLOGO_LOG_DTFORMAT="15:04:05.000" \
    STOP_ENGINE_ON_ERROR=flase \
    TCI_WI_CLOUD_ELEMENTS_CONSOLE_URL=%%TCI_WI_CLOUD_ELEMENTS_CONSOLE_URL%% \
    TCI_WI_CLOUD_ELEMENTS_ORG_TOKEN=%%TCI_WI_CLOUD_ELEMENTS_ORG_TOKEN%% \
    TCI_WI_CLOUD_ELEMENTS_USER_SECRET=%%TCI_WI_CLOUD_ELEMENTS_USER_SECRET%% \
    TCI_WI_SALESFORCE_CLIENT_SECRET=%%TCI_WI_SALESFORCE_CLIENT_SECRET%% \
    FLOGO_HTTP_SERVICE_PORT="7777"

ADD ./* /opt/tibco/flogo/1.0/bin/

WORKDIR /opt/tibco/flogo/1.0/bin/

RUN chmod a+x /opt/tibco/flogo/1.0/bin/*

RUN ls -l
RUN ls -l /opt/tibco/flogo/1.0/bin/

CMD ["/opt/tibco/flogo/1.0/bin/start.sh"]

VOLUME /opt/tibco/flogo/1.0/bin

EXPOSE 7777 9999

FROM public.ecr.aws/tibcolabs/labs-air-flogo-builder AS flogobuilder

ARG app_type

COPY build/air.json .

RUN chmod a+x create_build.sh;\
    ./create_build.sh air.json;

###################################################################################
#                         Flogo App image                         #
###################################################################################
FROM alpine:3.7

ARG app_type

RUN apk update && apk add ca-certificates
RUN apk --update add git less openssh
RUN apk add bash

VOLUME /git

#For FE case, set FLOGO_HTTP_SERVICE_PORT to empty to disable heathcheck
ENV FLOGO_LOG_DTFORMAT="15:04:05.000" \
    STOP_ENGINE_ON_ERROR=flase \
    TCI_WI_CLOUD_ELEMENTS_CONSOLE_URL=%%TCI_WI_CLOUD_ELEMENTS_CONSOLE_URL%% \
    TCI_WI_CLOUD_ELEMENTS_ORG_TOKEN=%%TCI_WI_CLOUD_ELEMENTS_ORG_TOKEN%% \
    TCI_WI_CLOUD_ELEMENTS_USER_SECRET=%%TCI_WI_CLOUD_ELEMENTS_USER_SECRET%% \
    TCI_WI_SALESFORCE_CLIENT_SECRET=%%TCI_WI_SALESFORCE_CLIENT_SECRET%% \
    FLOGO_HTTP_SERVICE_PORT="7777" \
    FLOGO_APP_METRICS_EVENT_QUEUE=40

ADD ./airpipeline_$app_type /home/airpipeline/
ADD ./project /home/project/
COPY --from=flogobuilder /home/app/bin/app /opt/tibco/flogo/1.0/bin/

WORKDIR /opt/tibco/flogo/1.0/bin/

RUN chmod a+x /opt/tibco/flogo/1.0/bin/app


CMD ["/opt/tibco/flogo/1.0/bin/app"]

VOLUME /opt/tibco/flogo/1.0/bin



EXPOSE 7777 9999
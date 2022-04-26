
FROM public.ecr.aws/tibcolabs/labs-air-flogo-builder:1.14-alpine AS flogobuilder

ARG app_type

RUN set -ex && \
    apk add --no-cache gcc musl-dev

RUN set -ex && \
    rm -f /usr/libexec/gcc/x86_64-alpine-linux-musl/6.4.0/cc1obj && \
    rm -f /usr/libexec/gcc/x86_64-alpine-linux-musl/6.4.0/lto1 && \
    rm -f /usr/libexec/gcc/x86_64-alpine-linux-musl/6.4.0/lto-wrapper && \
    rm -f /usr/bin/x86_64-alpine-linux-musl-gcj

COPY build/air.json .
RUN chmod a+x create_build.sh;\
    ./create_build.sh air.json;

RUN mkdir ./genericApp
RUN cp ./create_build.sh ./genericApp
COPY airpipeline_oss/flogo-generic.json ./genericApp
COPY airpipeline_oss/build-generic.sh ./genericApp

WORKDIR ./genericApp
RUN chmod a+x build-generic.sh;\
    ./build-generic.sh flogo-generic.json;

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
ADD ./project-template /home/project-template/
ADD ./scripts /home/scripts/
COPY --from=flogobuilder /home/app/bin/app /opt/tibco/flogo/1.0/bin/
COPY --from=flogobuilder /home/app/* /opt/tibco/flogo/1.0/bin/
COPY --from=flogobuilder /home/genericApp/app /home/project-template/no-internet/build/generic/app
COPY --from=flogobuilder /home/genericApp/app/bin/* /home/project-template/no-internet/build/generic/

WORKDIR /opt/tibco/flogo/1.0/bin/

RUN chmod a+x /opt/tibco/flogo/1.0/bin/app


CMD ["/opt/tibco/flogo/1.0/bin/app"]

VOLUME /opt/tibco/flogo/1.0/bin



EXPOSE 7777 9999
FROM public.ecr.aws/tibcolabs/labs-air-flogo-builder:1.16-alpine AS flogobuilder

ARG app_type

COPY build/builder.json .

RUN chmod a+x create_build.sh;\
    ./create_build.sh builder.json;
###################################################################################
#                         Flogo App image                         #
###################################################################################
FROM docker:19.03

RUN apk update && apk add ca-certificates

# System setup
RUN apk add git curl unzip build-base autoconf automake libtool

RUN apk add --no-cache iptables
RUN apk add --update docker-compose

#For FE case, set FLOGO_HTTP_SERVICE_PORT to empty to disable heathcheck
ENV FLOGO_LOG_DTFORMAT="15:04:05.000" \
    STOP_ENGINE_ON_ERROR=flase \
    FLOGO_HTTP_SERVICE_PORT="7777"

COPY --from=flogobuilder /home/app/bin/app /opt/tibco/flogo/1.0/bin/

WORKDIR /opt/tibco/flogo/1.0/bin/

RUN chmod a+x /opt/tibco/flogo/1.0/bin/*

CMD ["/opt/tibco/flogo/1.0/bin/app"]

VOLUME /opt/tibco/flogo/1.0/bin

EXPOSE 7777 9999
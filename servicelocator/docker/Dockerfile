###################################################################################
#                         Flogo App image                         #
###################################################################################
FROM reldocker.tibco.com/tci/flogo-go-base-alpine:3.8-1.14.4

#=============================== Install python3/pip3 =============================

ENV PYTHONUNBUFFERED=1
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools

#================================= Flogo Builder ==================================

ENV FLOGO_LOG_DTFORMAT="15:04:05.000" \
    STOP_ENGINE_ON_ERROR=flase \
    FLOGO_HTTP_SERVICE_PORT="7777"

ADD ./engine /opt/tibco/flogo/1.0/bin/

WORKDIR /opt/tibco/flogo/1.0/bin/

RUN chmod a+x /opt/tibco/flogo/1.0/bin/*

CMD ["/opt/tibco/flogo/1.0/bin/engine"]

VOLUME /opt/tibco/flogo/1.0/bin

EXPOSE 7777 9999
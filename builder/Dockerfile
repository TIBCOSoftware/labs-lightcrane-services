FROM public.ecr.aws/tibcolabs/labs-air-flogo-builder:1.14-alpine AS flogobuilder

ARG app_type

COPY build/builder.json .

RUN chmod a+x create_build.sh;\
    ./create_build.sh builder.json;
    
##################################################################################
#                         Flogo App image                         #
###################################################################################
FROM golang:1.14-alpine

#=================================== Add sudo =====================================
RUN set -ex && apk --no-cache add sudo

#==================================== Docker ======================================

ENV DOCKER_VERSION 19.03.14
RUN set -eux; \
	\
	apkArch="$(apk --print-arch)"; \
	case "$apkArch" in \
		'x86_64') \
			url='https://download.docker.com/linux/static/stable/x86_64/docker-19.03.14.tgz'; \
			;; \
		'armhf') \
			url='https://download.docker.com/linux/static/stable/armel/docker-19.03.14.tgz'; \
			;; \
		'armv7') \
			url='https://download.docker.com/linux/static/stable/armhf/docker-19.03.14.tgz'; \
			;; \
		'aarch64') \
			url='https://download.docker.com/linux/static/stable/aarch64/docker-19.03.14.tgz'; \
			;; \
		*) echo >&2 "error: unsupported architecture ($apkArch)"; exit 1 ;; \
	esac; \
	\
	wget -O docker.tgz "$url"; \
	\
	tar --extract \
		--file docker.tgz \
		--strip-components 1 \
		--directory /usr/local/bin/ \
	; \
	rm docker.tgz; \
	\
	dockerd --version; \
	docker --version

ENV DOCKER_TLS_CERTDIR=/certs
RUN mkdir /certs /certs/client && chmod 1777 /certs /certs/client

#================================ Install flogo cli ===============================

ENV GOPATH=/home/ext
ENV PATH=$GOPATH/bin:$PATH

#=============================== Install python3/pip3 =============================

ENV PYTHONUNBUFFERED=1
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools

#=================================== buildkit =====================================

RUN mkdir -p ~/.docker/cli-plugins
RUN apk --no-cache add curl
RUN apk add bash
RUN apk add git
RUN apk add build-base
RUN curl -L https://github.com/docker/buildx/releases/download/v0.5.1/buildx-v0.5.1.linux-amd64 -o ~/.docker/cli-plugins/docker-buildx
RUN chmod a+x ~/.docker/cli-plugins/docker-buildx

#=============================== Install flogo cli ================================

ENV GO111MODULE=on

COPY --from=flogobuilder /home/mod.go /home
RUN git clone https://github.com/project-flogo/cli.git;\
    cp /home/mod.go ./cli/util/;\
    cd ./cli;\
    go mod tidy;\
    go install ./...;

#================================= Flogo Builder ==================================

ENV FLOGO_LOG_DTFORMAT="15:04:05.000" \
    STOP_ENGINE_ON_ERROR=flase \
    FLOGO_HTTP_SERVICE_PORT="7777"

ADD ./scripts /home/scripts/
ADD ./runner /home/runner/
ADD ./start.sh /opt/tibco/flogo/1.0/bin/
COPY --from=flogobuilder /home/app/bin/app /opt/tibco/flogo/1.0/bin/

WORKDIR /opt/tibco/flogo/1.0/bin/

RUN chmod a+x /home/scripts/*.sh
RUN chmod a+x /opt/tibco/flogo/1.0/bin/*

CMD ["/opt/tibco/flogo/1.0/bin/start.sh"]

VOLUME /opt/tibco/flogo/1.0/bin

EXPOSE 7777 9999
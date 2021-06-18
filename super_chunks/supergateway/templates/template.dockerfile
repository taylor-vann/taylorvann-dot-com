FROM alpine:latest

VOLUME [ "/root/go/src/webapi/", "/usr/local/config/" ]

RUN apk --update add --no-cache build-base git go

WORKDIR /root/go/src/webapi/

EXPOSE ${http_port} ${https_port}
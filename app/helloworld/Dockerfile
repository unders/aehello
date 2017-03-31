FROM alpine:3.5
MAINTAINER unders <anders.tornqvist@gmail.com>

RUN apk --update upgrade && \
    apk add ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ADD favicon.ico favicon.ico
ADD helloworld helloworld
ENTRYPOINT ["/helloworld"]


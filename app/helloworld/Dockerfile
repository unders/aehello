FROM alpine:3.5
MAINTAINER unders <anders.tornqvist@gmail.com>

# https://blog.docker.com/2016/09/docker-golang/
RUN apk --no-cache add ca-certificates

ADD favicon.ico favicon.ico
ADD helloworld helloworld
ENTRYPOINT ["/helloworld"]


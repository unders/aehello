#!/usr/bin/env bash

main() {
    # vendoring does not work with app-engine standard
    # echo "Installing Go dependency tool: dep"
    # go get -u github.com/golang/dep/...

    echo "Installing Go app dependencies: "

    echo "    github.com/pkg/errors"
    go get -u github.com/pkg/errors

    echo "    cloud.google.com/go"
    go get -u cloud.google.com/go

    echo "    google.golang.org/appengine"
    go get -u google.golang.org/appengine

    echo "    google.golang.org/api/googleapi"
    go get -u google.golang.org/api/googleapi

    echo "    google.golang.org/grpc"
    go get -u google.golang.org/grpc

    echo "    golang.org/x/oauth2"
    go get -u golang.org/x/oauth2

    echo "    golang.org/x/sync/semaphore"
    go get -u golang.org/x/sync/semaphore

    echo "    google.golang.org/genproto/googleapis/api/monitoredres"
    go get -u google.golang.org/genproto/googleapis/api/monitoredres

    echo "    google.golang.org/genproto/googleapis/logging/type"
    go get -u google.golang.org/genproto/googleapis/logging/type

    echo "    google.golang.org/genproto/googleapis/logging/v2"
    go get -u google.golang.org/genproto/googleapis/logging/v2
}

main

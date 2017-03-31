#!/usr/bin/env bash

main() {
    echo "Installing Go dependency tool: dep"
    go get -u github.com/golang/dep/...

    echo "Installing Go app dependencies: "

    echo "    github.com/pkg/errors"
    go get -u github.com/pkg/errors

    echo "    cloud.google.com/go"
    go get -u cloud.google.com/go
}

main

#!/usr/bin/env bash

main() {
    echo "Installing Go dependency tool: dep"
    go get -u github.com/golang/dep/...
}

main

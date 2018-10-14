#!/usr/bin/env sh

# change context to current directory
cd $(dirname $0)


D=/app/src/github.com/warchant/go-grpc-healthcheck

docker run --rm -i \
    -v ${PWD}:${D}:rw \
    -w ${D} \
    golang bash << COMMANDS
    export GOPATH=/app
    go build
COMMANDS

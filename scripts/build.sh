#!/bin/bash
set -eufo pipefail
IFS=$'\t\n'

# Compile for alpine and build docker image

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/fed ./cmd/fed

docker build "$PWD" -t "simonsargeant/fed:${IMAGE_TAG:-local}"


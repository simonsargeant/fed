#!/bin/bash
set -eufo pipefail
IFS=$'\t\n'

# Compile the binary for alpine, build the image and push to registry

GOOS=linux GOARCH=amd64 go build -o ./bin/fed ./cmd/fed

docker login "${DOCKER_HOST}" -u "${DOCKER_USER}" -p "${DOCKER_TOKEN}"

docker build "$PWD" -t "simonsargeant/fed:${IMAGE_TAG}"
docker tag "simonsargeant/fed:${IMAGE_TAG}" "simonsargeant/fed:latest"

docker push "simonsargeant/fed:${IMAGE_TAG}"
docker push "simonsargeant/fed:latest"


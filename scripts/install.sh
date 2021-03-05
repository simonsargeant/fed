#!/bin/bash
set -eufo pipefail
IFS=$'\t\n'

# Compile the binary to run locally and move to /usr/local/bin

go build -o ./bin/fed ./cmd/fed

cp ./bin/fed /usr/local/bin/fed


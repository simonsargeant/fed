#!/bin/bash
set -eufo pipefail
IFS=$'\t\n'

# Run unit tests

go test ./...


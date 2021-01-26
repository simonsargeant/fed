#!/bin/bash
set -eufo pipefail
IFS=$'\t\n'

go build -o ./bin/fed ./cmd/fed

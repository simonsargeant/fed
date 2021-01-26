#!/bin/bash
set -eufo pipefail
IFS=$'\t\n'

go test ./...

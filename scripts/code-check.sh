#!/bin/bash
set -eufo pipefail
IFS=$'\t\n'

# Check code style

golangci-lint run -E goimports


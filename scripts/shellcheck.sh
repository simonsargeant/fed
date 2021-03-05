#!/bin/bash
set -euo pipefail
IFS=$'\t\n'

# Check script code style

shellcheck ./scripts/*.sh

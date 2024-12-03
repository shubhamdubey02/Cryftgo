#!/usr/bin/env bash

set -euo pipefail

# cryftgo root folder
METAL_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the constants
source "$METAL_PATH"/scripts/constants.sh

echo "Building bootstrap-monitor..."
go build -ldflags\
   "-X github.com/shubhamdubey02/cryftgo/version.GitCommit=$git_commit $static_ld_flags"\
   -o "$METAL_PATH/build/bootstrap-monitor"\
   "$METAL_PATH/tests/fixture/bootstrapmonitor/cmd/"*.go

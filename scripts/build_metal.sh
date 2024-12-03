#!/usr/bin/env bash

set -euo pipefail

print_usage() {
  printf "Usage: build_metal [OPTIONS]

  Build cryftgo

  Options:

    -r  Build with race detector
"
}

race=''
while getopts 'r' flag; do
  case "${flag}" in
    r) race='-race' ;;
    *) print_usage
      exit 1 ;;
  esac
done

# cryftgo root folder
METAL_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the constants
source "$METAL_PATH"/scripts/constants.sh

build_args="$race"
echo "Building cryftgo..."
go build $build_args -ldflags "-X github.com/shubhamdubey02/cryftgo/version.GitCommit=$git_commit $static_ld_flags" -o "$cryftgo_path" "$METAL_PATH/main/"*.go

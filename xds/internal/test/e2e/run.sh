#!/bin/bash

set -e
cd "$(dirname "$0")"
mkdir -p binaries
go build -o ./binaries/client ../../../../interop/xds/client/
go build -o ./binaries/server ../../../../interop/xds/server/
go test . -v

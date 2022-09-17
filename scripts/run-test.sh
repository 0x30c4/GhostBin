#!/bin/bash

export HOST=localhost
export PORT=8080
export DOMAIN=http://$HOST:$PORT
export DATA_DIR="./testdata"

export DB_HOST=.
export DB_PORT=.
export DB_NAME=.

mkdir test/testdata

go test -v test/handlers_test.go

rm -rf test/testdata

# go tool cover -html cover.out -o cover.html


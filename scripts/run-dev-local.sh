#!/bin/bash

export HOST=localhost
export PORT=8080
export DOMAIN=http://$HOST:$PORT
export DATA_DIR="./public"
export PASTE_DATA_PATH=""

export DB_HOST=.
export DB_PORT=.
export DB_NAME=.

export LOG_FILE="logs/gin.log"
export RELEASE_MODE=1

go run cmd/main.go


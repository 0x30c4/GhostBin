#!/usr/bin/env bash

# redis-server ./redis/redis.conf 2>1 > ./logs/redis.log &

export $(grep -v '^#' .env.dev | xargs);

go test -coverprofile=./test/testcover.out ./...

go tool cover -html ./test/testcover.out -o ./test/testcover.html

go test -v ./test/

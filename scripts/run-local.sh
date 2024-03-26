#!/usr/bin/env bash

redis-server ./redis/redis.conf 2>1 > ./logs/redis.log &

export $(grep -v '^#' .env.dev | xargs);

export REDIS_ADDRESS="localhost:6379"

air

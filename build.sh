#!/usr/bin/env bash
set -ex

rm -fr ./.build

GOOS=linux GOARCH=amd64 go build -o .build/test_task ./main.go
cp -r ./resources/ ./.build/resources
cp ./redis.conf ./.build/redis.conf
cp ./startup.sh ./.build/startup.sh

docker build -t test_task_1 .
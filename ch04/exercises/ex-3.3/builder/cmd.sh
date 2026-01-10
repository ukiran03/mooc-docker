#!/usr/bin/env bash

go build -o pipeline main.go

docker build -t builder .

docker run --network host \
       -e DOCKER_USER=ushakiran369 -e DOCKER_PWD="*********" \
       -v /var/run/docker.sock:/var/run/docker.sock \
       builder mluukkai/express_app ushakiran369/testing

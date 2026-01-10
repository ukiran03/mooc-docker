#!/usr/bin/env bash

set -e

usage() {
    echo "Usage: $0 <github-repo> <docker-image>"
    exit 1
}

cleanup() {
    rm -rf "$TEMP_DIR"
}

if [ "$#" -ne 2 ]; then
    usage
fi

GITHUB_REPO=$1
DOCKER_IMAGE=$2
TEMP_DIR=$$(mktemp -d)

trap cleanup EXIT

git clone "https://github.com/$GITHUB_REPO" "$TEMP_DIR"
cd "$TEMP_DIR"

docker build -t "$DOCKER_IMAGE"
docker push "$DOCKER_IMAGE"

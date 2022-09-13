#!/bin/sh
name=webapp-time
docker rmi -f camphul/"$name":latest
docker buildx build -t camphul/"$name":latest --platform=linux/arm64,linux/amd64 . --push
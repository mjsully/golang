#!/bin/bash

docker run -it --rm \
    --name golang \
    -v $(pwd)/data:/go/data \
    -p 8080:8080 \
    golang:latest

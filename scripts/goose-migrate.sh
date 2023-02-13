#! /usr/bin/env sh

docker run --rm \
    -v $PWD/db:/db \
    -v $PWD/internal/models/schema:/src \
    -w /src \
    -u 1000:1000 \
    gomicro/goose:3.7.0 goose sqlite /db/$1.db up -no-versioning ;

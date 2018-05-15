#!/usr/bin/env bash

#GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/main main.go


#CC=/usr/bin/x86_64-alpine-linux-musl-gcc GOOS=linux go build -x \
# -ldflags '-linkmode external -extldflags "-static"' -a -tags netgo \
# -installsuffix netgo -o main main.go

go build -ldflags="-s -w -v" -o bin/main main.go
chmod +x bin/main
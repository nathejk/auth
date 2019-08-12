# DEV - BASE
FROM golang:1.12 AS dev

RUN mkdir /app
WORKDIR /app

RUN apt update && \
    apt install \
    bash \
    inotify-tools

RUN go get github.com/m3ng9i/ran && \
    go install github.com/m3ng9i/ran

RUN mkdir /app/bin
RUN printf "#!/bin/sh\necho 'Please mount repo into /app'" > /app/bin/init-dev
RUN chmod +x /app/bin/init-dev

VOLUME ["/go"]

EXPOSE 80

ENTRYPOINT ["/app/bin/init-dev"]


# TEST
FROM dev AS test

COPY src /app/src
WORKDIR /app/src
RUN go get -v -d ./... && CGO_ENABLED=0 GOOS=linux go test ./


# BUILD
FROM test AS build

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o dims


# PROD
FROM alpine:3.9 AS prod

RUN apk upgrade -U && \
    apk add --update --no-cache \
    libc6-compat \
    coreutils \
    bash

RUN mkdir -p /app/bin
WORKDIR /app

COPY --from=build /app/src/dims .
COPY bin/init /init
COPY bin/waitservices.sh /app/bin/waitservices.sh

ENTRYPOINT ["/init"]

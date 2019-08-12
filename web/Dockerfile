# DEV - BASE
FROM golang:1.12 AS api-dev

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

VOLUME ["/data"]
VOLUME ["/go"]

EXPOSE 80

ENTRYPOINT ["/app/bin/init-dev"]


# TEST'n'BUILD
## API
FROM api-dev AS api-builder

COPY api/src /app/src
WORKDIR /app/src

RUN CGO_ENABLED=0 GOOS=linux go test ./ $(ls -d nathejk/*) && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o dims

## UI
FROM node:10.11-alpine AS ui-builder

RUN mkdir -p /app
WORKDIR /app

COPY ui/package*.json /app/
RUN npm ci --no-save
#COPY ui/yarn.lock /app/

#RUN yarn install --frozen-lockfile

COPY ui/src /app/src
COPY ui/public /app/public
COPY ui/*.config.js /app/

RUN npm test
RUN npm run build

#FROM api-dev AS prod


# PROD
FROM alpine:3.9 AS prod

RUN apk upgrade -U && \
    apk add --update --no-cache \
    ca-certificates \
    coreutils \
    bash

RUN mkdir -p /app/bin
WORKDIR /app

COPY /rootfs /
COPY --from=api-builder /app/src/dims .
COPY --from=ui-builder /app/dist /www
COPY /rootfs/bin/waitservices.sh /app/bin/waitservices.sh

ENTRYPOINT ["/bin/init"]

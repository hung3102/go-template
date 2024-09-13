FROM golang:1.23-alpine as builder

RUN apk add --no-cache alpine-sdk build-base && \
  go install github.com/air-verse/air@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app/back/app/general/cmd

ENV IS_LOCAL="true"

CMD [ "air", "-c", "/app/docker/backend/general/dev/.air.toml"]

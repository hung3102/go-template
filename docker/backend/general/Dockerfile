FROM golang:1.23-alpine as builder

RUN apk update && apk add --repository https://dl-cdn.alpinelinux.org/alpine/v3.4/main --no-cache ca-certificates gcc libc-dev git

WORKDIR /app

COPY back/ ./
RUN go work init app pkg

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o server ./app/general/cmd

FROM gcr.io/distroless/static

COPY --from=builder /app/server /server

CMD [ "/server" ]

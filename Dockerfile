FROM golang:1.21-alpine AS dev

RUN apk update && apk add \
  bash \
  make \
  rm -rf /var/cache/apk/*

WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ .

RUN make build

ENTRYPOINT [ "make", "start-dev" ]

FROM alpine:3.16.7 AS release

ENV HTTP_LISTEN_ADDR=8000

WORKDIR /app

COPY --from=dev /app/dancing-pony .

EXPOSE 8080

ENTRYPOINT [ "./dancing-pony" ]

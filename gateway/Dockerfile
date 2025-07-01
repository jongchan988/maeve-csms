# syntax=docker/dockerfile:1.2

FROM golang:1.22.2

RUN apt-get update && apt-get install -y git openssh-client ca-certificates curl

ARG TARGETARCH

RUN if [ "$TARGETARCH" = "arm64" ]; then \
        TARGETARCH=aarch64 ; \
        fi; \
    wget -O /usr/bin/curl https://github.com/moparisthebest/static-curl/releases/download/v8.0.1/curl-$TARGETARCH \
        && chmod +x /usr/bin/curl

WORKDIR /go/src/myapp

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

COPY ./start_server.sh ./stop_server.sh ./


RUN go build -cover -covermode=atomic -coverpkg=./... -o ./app main.go

RUN mkdir -p "/app/cover"
ENV GOCOVERDIR=/app/cover

ENTRYPOINT ["sh", "./start_server.sh"]

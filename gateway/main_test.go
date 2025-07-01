# syntax=docker/dockerfile:1.2

FROM golang:1.20-alpine AS builder

RUN apk add --no-cache git openssh ca-certificates

ARG TARGETARCH

RUN if [ "$TARGETARCH" = "arm64" ]; then \
        TARGETARCH=aarch64 ; \
        fi; \
    wget -O /usr/bin/curl https://github.com/moparisthebest/static-curl/releases/download/v8.0.1/curl-$TARGETARCH \
        && chmod +x /usr/bin/curl

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./
COPY ./start_server.sh ./stop_server.sh ./

RUN go build -cover -covermode=atomic -coverpkg=./... -o /app main.go

RUN mkdir -p /cover
ENV GOCOVERDIR=/cover

# ✅ STAGE 2: Final Runtime 환경 명시 (alpine 사용)
FROM alpine:3.18 AS final

WORKDIR /app
COPY --from=builder /app /app
COPY --from=builder /src/start_server.sh /src/stop_server.sh /app/
COPY --from=builder /cover /cover
ENV GOCOVERDIR=/cover

RUN chmod +x /app/start_server.sh /app/stop_server.sh

ENTRYPOINT ["/bin/sh", "/app/start_server.sh"]

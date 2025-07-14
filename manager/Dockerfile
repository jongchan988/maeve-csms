# syntax=docker/dockerfile:1.2

# STAGE 1: build the executable with coverage
FROM golang:1.22.2 AS builder

RUN apt-get update && apt-get install -y git openssh-client ca-certificates curl

ARG TARGETARCH

# Fix for arm64 static curl
RUN if [ "$TARGETARCH" = "arm64" ]; then \
        TARGETARCH=aarch64 ; \
    fi; \
    wget -O /usr/bin/curl https://github.com/moparisthebest/static-curl/releases/download/v8.0.1/curl-$TARGETARCH \
    && chmod +x /usr/bin/curl

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./


# Build with coverage enabled
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    go build \
    -cover \
    -covermode=atomic \
    -coverpkg=./... \
    -o /app \
    .

# Install coverage tools
RUN go install github.com/axw/gocov/gocov@latest
RUN go install github.com/AlekSi/gocov-xml@latest

# STAGE 2: runtime container
FROM golang:1.22.2 AS final

# Install runtime dependencies
RUN apt-get update && apt-get install -y ca-certificates curl

COPY --from=builder /usr/bin/curl /usr/bin/curl
COPY --from=builder /app /app
COPY --from=builder /go/bin/* /usr/local/bin/

# Coverage directory
RUN mkdir -p "/cover"
ENV GOCOVERDIR=/cover

# Copy start and stop scripts
COPY ./start_server.sh ./stop_server.sh ./

# Non-root user (optional)
RUN useradd -u 10000 -m appuser
USER appuser

ENTRYPOINT ["sh", "./start_server.sh"]

#!/bin/sh
kill -SIGINT $(pidof app)
PID=$(ps aux | grep "app serve")

if [ -z "$PID" ]; then
  echo "❌ app serv is not running."
  exit 1
fi

echo "🛑 Sending SIGINT to app serv (PID=$PID)..."
kill -15 "$PID"

sleep 1
go tool covdata textfmt -i=/app/cover -o coverage.out
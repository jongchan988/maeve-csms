#!/bin/sh

# 정확히 PID만 추출
PID=$(ps aux | grep "[a]pp serve" | awk '{print $2}')

if [ -z "$PID" ]; then
  echo "❌ app serve is not running."
  exit 1
fi

echo "🛑 Sending SIGTERM to app serve (PID=$PID)..."
kill -TERM "$PID"

# flush 유도
sleep 1

# 커버리지 출력
go tool covdata textfmt -i=/cover -o coverage.out
// SPDX-License-Identifier: Apache-2.0


package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/thoughtworks/maeve-csms/gateway/cmd"
)

func main() {
	// graceful shutdown 처리
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-c
		// 종료 전에 필요한 작업 수행
		println("📦 Received signal:", s.String())
		// 필요 시: 로그 저장, 커넥션 종료 등
		os.Exit(0) // 여기가 중요: 정상 종료 → 커버리지 flush 됨
	}()

	cmd.Execute()
}
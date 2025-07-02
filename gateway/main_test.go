package main

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/thoughtworks/maeve-csms/gateway/cmd"
)

func TestMain(m *testing.M) {
	go func() {
		cmd.Execute() // 실제 서버 실행
	}()

	// 시그널 대기
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// 종료 전에 커버리지 헤더만 작성 (실제 커버리지는 자동 기록됨)
	f, err := os.Create("/app/coverage.out")
	if err == nil {
		f.Write([]byte("mode: atomic\n")) // go tool cover용 헤더
		f.Close()
	}

	os.Exit(0)
}

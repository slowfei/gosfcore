package SFLog

import (
	"testing"
	"time"
)

// 测试日志
func TestLogger(t *testing.T) {
	//	开启日志管理
	SharedLogManager("")

	Info("my %v", "slowfei")

	time.Sleep(time.Duration(5) * time.Second)
}

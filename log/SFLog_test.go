package SFLog

import (
	"fmt"
	"testing"
	"time"
)

// 测试日志
func TestLogger(t *testing.T) {
	//	开启日志管理
	StartLogManager(3000)
	start := time.Now()

	start = time.Now()
	Info("my %v", "slowfei-info")
	fmt.Println("Info-time:", time.Now().Sub(start))

	start = time.Now()
	Debug("my %v", "slowfei-debug")
	fmt.Println("Debug-time:", time.Now().Sub(start))

	start = time.Now()
	Error("my %v", "slowfei-error")
	fmt.Println("Error-time:", time.Now().Sub(start))

	start = time.Now()
	Warn("my %v", "slowfei-warn")
	fmt.Println("Warn-time:", time.Now().Sub(start))

	start = time.Now()
	Fatal("my %v", "slowfei-fatal")
	fmt.Println("Fatal-time:", time.Now().Sub(start))

	start = time.Now()
	Panic("my %v", "slowfei-panic")
	fmt.Println("Panic-time:", time.Now().Sub(start))

	time.Sleep(time.Duration(2) * time.Second)
}

//	测试格式化速度
func TestLogMagFormat(t *testing.T) {
	start := time.Now()
	format := "${yyyy}-${MM}-${dd} ${mm}:${dd}:${ss} [${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}] ${MSG} \n${STACK}"

	logmsg := &LogMsg{}
	logmsg.logGroup = "globalGroup"
	logmsg.logTag = "logTag"
	logmsg.dateTime = time.Now()
	logmsg.stack = `
	github.com/slowfei/gosfcore/log.TestLogger(...)
	/src/github.com/slowfei/gosfcore/log/SFLog_test.go:20 (0x33dfe)
	testing.tRunner(...)
	/usr/local/go/src/pkg/testing/testing.go:353 (0x2e00a)
	runtime.goexit(...)
	/usr/local/go/src/pkg/runtime/proc.c:1223 (0x173d0)
	`
	logmsg.funcName = "github.com/slowfei/gosfcore/log.TestLogger"
	logmsg.fileLine = 20
	logmsg.msg = "msgString"
	logmsg.target = TargetInfo

	// b.StartTimer()
	start = time.Now()
	logMagFormat(format, logmsg)
	fmt.Println("TestLogMagFormat-Time:", time.Now().Sub(start))

	start = time.Now()
	logMagFormat(format, logmsg)
	fmt.Println("TestLogMagFormat-Time:", time.Now().Sub(start))

	start = time.Now()
	logMagFormat(format, logmsg)
	fmt.Println("TestLogMagFormat-Time:", time.Now().Sub(start))
}

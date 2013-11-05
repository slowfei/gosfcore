package SFLog

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestGetFile(t *testing.T) {
	// af := NewAppenderFile()
	// af.getFile("file_${yyyy}-${MM}-${dd}.log", time.Now())
}

//	测试文件写入
func TestWriteFile(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	af := NewAppenderFile()
	fmt.Println("日志文件操作路径：", af.excePath)

	config := &AppenderFileConfig{}
	config.MaxSize = 256
	config.SavePath = "/Users/slowfei/Downloads/test/log"
	config.Name = "../test-${yyyy}-${MM}-${dd}-${hh}.log"
	config.Pattern = "${yyyy}-${MM}-${dd} ${hh}:${mm}:${ss}${SSSSSS} [${TARGET}] ([${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}])\n${MSG}"

	for i := 0; i < 10; i++ {
		num := strconv.Itoa(i)
		go func() {
			msg := &LogMsg{}
			msg.dateTime = time.Now()
			msg.fileLine = 26
			msg.filePath = "file_test.go"
			msg.funcName = "TestWriteFile(...)"
			msg.logGroup = "logGroup"
			msg.logTag = "logTag"
			msg.msg = "test write file+:" + num
			msg.stack = ""
			af.Write(msg, config)
			time.Sleep(1 * time.Second)
		}()
	}

	time.Sleep(10 * time.Second)

}

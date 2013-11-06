package SFLog

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestHTMLLayout(t *testing.T) {
	start := time.Now()

	start = time.Now()
	s := fmt.Sprintf(HTMLHandLayout, "Log Title")
	fmt.Println("timd: ", time.Now().Sub(start))
	fmt.Println(s)
}

func TestWriteHtml(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	af := NewAppenderHtml()
	fmt.Println("日志HTML操作路径：", af.writeFile.excePath)

	config := &AppenderHtmlConfig{}
	config.MaxSize = 1024 * 10
	config.Title = "Log Test"
	config.SavePath = "/Users/slowfei/Downloads/test/log"
	config.Name = "../log-${yyyy}${MM}${dd}${hh}.html"

	for i := 0; i < 10; i++ {
		num := strconv.Itoa(i)
		go func() {
			msg := &LogMsg{}
			msg.target = TargetInfo
			msg.dateTime = time.Now()
			msg.fileLine = 35
			msg.filePath = "html_test.go"
			msg.funcName = "TestWriteFile(...)"
			msg.logGroup = "logGroup"
			msg.logTag = "logTag"
			msg.msg = "info write html :" + num
			msg.stack = ""
			af.Write(msg, config)
			time.Sleep(1 * time.Second)
		}()

		time.Sleep(1 * time.Second)
		go func() {
			msg := &LogMsg{}
			msg.target = TargetDebug
			msg.dateTime = time.Now()
			msg.fileLine = 50
			msg.filePath = "html_test.go"
			msg.funcName = "TestWriteFile(...)"
			msg.logGroup = "logGroup"
			msg.logTag = "logTag"
			msg.msg = "debug write html :" + num
			msg.stack = `
./html.go:85: fileInfo declared and not used
./html.go:100: htmlContent declared and not used
			`
			af.Write(msg, config)
			time.Sleep(1 * time.Second)
		}()

		time.Sleep(1 * time.Second)
		go func() {
			msg := &LogMsg{}
			msg.target = TargetError
			msg.dateTime = time.Now()
			msg.fileLine = 67
			msg.filePath = "html_test.go"
			msg.funcName = "TestWriteFile(...)"
			msg.logGroup = "logGroup"
			msg.logTag = "logTag"
			msg.msg = "error write html :" + num
			msg.stack = `
./html.go:85: fileInfo declared and not used
./html.go:100: htmlContent declared and not used
			`
			af.Write(msg, config)
			time.Sleep(1 * time.Second)
		}()
	}

	time.Sleep(10 * time.Second)

}

//	slowfei的日志操作类，类似java的log4j
//	后续可能需要自己自定义格式，重写log.Output方法
//	目前只是简单的实现了控制台的打印
//	TODO 还未实现
//
//	copyright 2013 slowfei
//	email	slowfei@foxmail.com
//	createTime 	2013-8-24
//	updateTime	2013-8-24
package SFLog

import (
	"fmt"
	"log"
	"os"
)

type SFLogger struct {
	goName string
}

func NewLogger(goName string) SFLogger {
	return SFLogger{goName: goName}
}

var logger *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags)

//	info log
func Info(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	logger.Output(2, msg)
	return msg
}

//	debug log
func Debug(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	logger.Output(2, msg)
	return msg
}

//	warn log
func Warn(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	logger.Output(2, msg)
	return msg
}

//	error log
func Error(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	logger.Output(2, msg)
	return msg
}

//	fatal log
func Fatal(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	logger.Output(2, msg)
	return msg
}

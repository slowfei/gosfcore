//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-24
//  Update on 2013-10-16
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//	日志操作，类似java的log4j
//	info debug error warn ftal level
//	out file html mongodb email
package SFLog

import (
	"fmt"
	"log"
	"os"
)

//	logger 产生日志的输出，主要负责标识每个不同的日志对象
type SFLogger struct {
	logTag string
}

//	new SFLogger
//	@logTag		输出日志对象的标识，最好是唯一的
//	@return
func NewLogger(logTag string) *SFLogger {
	return &SFLogger{logTag}
}

//	logger info log
func (l *SFLogger) Info(format string, v ...interface{}) string {

}

//	logger debug log
func (l *SFLogger) Debug(format string, v ...interface{}) string {

}

//	logger error log
func (l *SFLogger) Error(format string, v ...interface{}) string {

}

//	logger warn log
func (l *SFLogger) Warn(format string, v ...interface{}) string {

}

//	logger fatal log
func (l *SFLogger) Fatal(format string, v ...interface{}) string {

}

//	logger panic log
func (l *SFLogger) Panic(isPanic bool, format string, v ...interface{}) string {
	if isPanic {
		panic("")
	}
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

//	error log
func Error(format string, v ...interface{}) string {
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

//	fatal log
func Fatal(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	logger.Output(2, msg)
	return msg
}

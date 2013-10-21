//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-24
//  Update on 2013-10-17
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//	日志操作，类似java的log4j
//	info debug error warn ftal panic level
//	out file html mongodb email
package SFLog

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	//
	_thisLogManager *LogManager = nil

	//	global logger
	_globalLogger *SFLogger = NewLogger(KEY_GLOBAL_LOG_TAG_CONFIG)

	logger *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags)
)

//	logger 产生日志的输出，主要负责标识每个不同的日志对象
type SFLogger struct {
	logTag string
}

//	New SFLogger
//	@logTag		输出日志对象的标识，最好是唯一的
//	@return
func NewLogger(logTag string) *SFLogger {
	return &SFLogger{logTag}
}

//	logger info log
func (l *SFLogger) Info(format string, v ...interface{}) string {
	if nil == _thisLogManager {
		msg := fmt.Sprintf("[LogManager == nil]"+format, v...)
		logger.Output(2, msg)
		return msg
	}

	return ""
}

//	logger debug log
func (l *SFLogger) Debug(format string, v ...interface{}) string {
	return ""
}

//	logger error log
func (l *SFLogger) Error(format string, v ...interface{}) string {
	return ""
}

//	logger warn log
func (l *SFLogger) Warn(format string, v ...interface{}) string {
	return ""
}

//	logger fatal log
func (l *SFLogger) Fatal(format string, v ...interface{}) string {
	return ""
}

//	logger panic log
func (l *SFLogger) Panic(format string, v ...interface{}) string {
	return ""
}

//	info log
func Info(format string, v ...interface{}) string {
	// msg := fmt.Sprintf(format, v...)
	// logger.Output(2, msg)
	return _globalLogger.Info(format, v...)

}

//	debug log
func Debug(format string, v ...interface{}) string {

	return _globalLogger.Debug(format, v...)
}

//	error log
func Error(format string, v ...interface{}) string {
	return _globalLogger.Error(format, v...)
}

//	warn log
func Warn(format string, v ...interface{}) string {
	return _globalLogger.Warn(format, v...)
}

//	fatal log
func Fatal(format string, v ...interface{}) string {
	return _globalLogger.Fatal(format, v...)
}

func Panic(format string, v ...interface{}) string {
	return _globalLogger.Panic(format, v...)
}

type LogMsg struct {
	target LogTarget
	msg    string
}

// log manager
type LogManager struct {
	rwm       sync.RWMutex
	develMode bool
	msg       chan *LogMsg
}

//	shared log manager
//	@filePath	相对或绝对路径, ""为使用默认配置
func SharedLogManager(filePath string) (*LogManager, error) {
	if nil == _thisLogManager {

		_thisLogManager = new(LogManager)

		err := LoadConfig(filePath)

		if nil != err {
			return nil, err
		}
	}
	return _thisLogManager, nil
}

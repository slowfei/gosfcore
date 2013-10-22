//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-24
//  Update on 2013-10-22
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//	日志操作，类似java的log4j
//	info debug error warn ftal panic level
//	out console file html mongodb email
package SFLog

import (
	"fmt"
	"github.com/slowfei/gosfcore/utils/time"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	// log manager
	_thisLogManager *LogManager = nil

	//	global logger
	_globalLogger *SFLogger = NewLogger(KEY_GLOBAL_GROUP_LOG_CONFIG)

	//	appender impl struct
	_implAppenderConsole Appender = nil
	_implAppenderFile    Appender = nil
	_implAppenderHtml    Appender = nil
	_implAppenderEmail   Appender = nil
	_implAppenderMongodb Appender = nil
)

//	logger 产生日志的输出，主要负责标识每个不同的日志对象
type SFLogger struct {
	logGroup string
	logTag   string
}

//	New SFLogger
//	default log global group, KEY_GLOBAL_GROUP_LOG_CONFIG
//
//	@logTag		输出日志对象的标识，最好是唯一的
//	@return
func NewLogger(logTag string) *SFLogger {
	return &SFLogger{KEY_GLOBAL_GROUP_LOG_CONFIG, logTag}
}

// @logGroup log group
func NewLoggerByGroup(logGroup, logTag string) *SFLogger {
	return &SFLogger{logGroup, logTag}
}

//	logger info log
func (l *SFLogger) Info(format string, v ...interface{}) string {
	return loggerHandle(l, TargetInfo, format, v...)
}

//	logger debug log
func (l *SFLogger) Debug(format string, v ...interface{}) string {
	return loggerHandle(l, TargetDebug, format, v...)
}

//	logger error log
func (l *SFLogger) Error(format string, v ...interface{}) string {
	return loggerHandle(l, TargetError, format, v...)
}

//	logger warn log
func (l *SFLogger) Warn(format string, v ...interface{}) string {
	return loggerHandle(l, TargetWarn, format, v...)
}

//	logger fatal log
func (l *SFLogger) Fatal(format string, v ...interface{}) string {
	return loggerHandle(l, TargetFatal, format, v...)
}

//	logger panic log
func (l *SFLogger) Panic(format string, v ...interface{}) string {
	return loggerHandle(l, TargetPanic, format, v...)
}

//	info log
func Info(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetInfo, format, v...)
}

//	debug log
func Debug(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetDebug, format, v...)
}

//	error log
func Error(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetError, format, v...)
}

//	warn log
func Warn(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetWarn, format, v...)
}

//	fatal log
func Fatal(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetFatal, format, v...)
}

//	panic log
func Panic(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetPanic, format, v...)
}

// logger handle
func loggerHandle(log *SFLogger, target LogTarget, format string, v ...interface{}) string {
	var msgString string
	if nil == _thisLogManager {
		msgString = fmt.Sprintf("[LogManager == nil]"+format, v...)
		fmt.Println(msgString)
		return msgString
	}

	msgString = fmt.Sprintf(format, v...)

	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	stack := fmt.Sprintf("%s:%d (0x%x)\n", file, line, pc)

	logmsg := &LogMsg{}
	logmsg.logGroup = log.logGroup
	logmsg.logTag = log.logTag
	logmsg.dateTime = time.Now()
	logmsg.stack = stack
	logmsg.msg = msgString
	logmsg.target = target

	_thisLogManager.msg <- logmsg

	return msgString
}

//	chan log msg
type LogMsg struct {
	logGroup string    // log group
	logTag   string    // log tag
	target   LogTarget // info,debug,error...
	dateTime time.Time // create time
	stack    string    // stack info
	msg      string    // log message
}

//	格式化日志信息
func logMagFormat(format string, msg *LogMsg) string {
	//	TODO 由于在格式化时间的时候，将信息标识给覆盖格式化了，目前想不到好的方法。

	//	格式化时间
	format = SFTimeUtil.YMDHMSSFormat(msg.dateTime, format)

	//	日志的格式化信息，别随意更换顺序，因为根据设计来进行日志信息的格式化操作
	logFormat := []string{
		"${LOG_GROUP}", msg.logGroup,
		"${LOG_TAG}", msg.logTag,
		"${STACK}", msg.stack,
		"${MSG}", msg.msg,
	}

	replacer := strings.NewReplacer(logFormat...)

	return replacer.Replace(format)
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

		//	初始化appender的实现
		for _, appender := range _sharedLogConfig.InitAppenders {
			switch appender {
			case VAL_APPENDER_CONSOLE:
				_implAppenderConsole = NewAppenderConsole()
			case VAL_APPENDER_FILE:
			case VAL_APPENDER_HTML:
			case VAL_APPENDER_EMAIL:
			case VAL_APPENDER_MONGODB:
			}
		}

		_thisLogManager.msg = make(chan *LogMsg, _sharedLogConfig.ChannelSize)
		go _thisLogManager.startChannel()
	}
	return _thisLogManager, nil
}

//	开启通道监控
func (lm *LogManager) startChannel() {
FORBREAK:
	for {
		select {
		case msg, ok := <-lm.msg:
			if !ok {
				break FORBREAK
			}
			lm.msgHandle(msg)
		}
	}
	fmt.Println("LogManager: channel close.")
}

//	信息写入操作
func (lm *LogManager) msgHandle(msg *LogMsg) {

	var config LogConfig
	var ok bool

	if config, ok = _sharedLogConfig.LogGroups[msg.logTag]; !ok {
		//	找不到目标配置使用全局配置
		if config, ok = _sharedLogConfig.LogGroups[KEY_GLOBAL_GROUP_LOG_CONFIG]; !ok {
			return
		}
	}

	if ok {
		var tci *TargetConfigInfo = nil
		switch msg.target {
		case TargetInfo:
			if nil == config.Info {
				tci = config.TargetConfigInfo
			} else if nil != config.Info && !config.Info.AppenderNoneConfig.None {
				tci = config.Info
			}
		case TargetDebug:
			if nil == config.Debug {
				tci = config.TargetConfigInfo
			} else if nil != config.Debug && !config.Debug.AppenderNoneConfig.None {
				tci = config.Debug
			}
		case TargetError:
			if nil == config.Error {
				tci = config.TargetConfigInfo
			} else if nil != config.Error && !config.Error.AppenderNoneConfig.None {
				tci = config.Error
			}
		case TargetWarn:
			if nil == config.Warn {
				tci = config.TargetConfigInfo
			} else if nil != config.Warn && !config.Warn.AppenderNoneConfig.None {
				tci = config.Warn
			}
		case TargetFatal:
			if nil == config.Fatal {
				tci = config.TargetConfigInfo
			} else if nil != config.Fatal && !config.Fatal.AppenderNoneConfig.None {
				tci = config.Fatal
			}
		case TargetPanic:
			if nil == config.Panic {
				tci = config.TargetConfigInfo
			} else if nil != config.Panic && !config.Panic.AppenderNoneConfig.None {
				tci = config.Panic
			}
		}

		if nil != tci {
			for _, appender := range tci.Appender {
				switch appender {
				case VAL_APPENDER_CONSOLE:
					if nil != _implAppenderConsole {
						_implAppenderConsole.Write(msg, tci)
					}
				case VAL_APPENDER_FILE:
					if nil != _implAppenderFile {
						_implAppenderFile.Write(msg, tci)
					}
				case VAL_APPENDER_HTML:
					if nil != _implAppenderHtml {
						_implAppenderHtml.Write(msg, tci)
					}
				case VAL_APPENDER_EMAIL:
					if nil != _implAppenderEmail {
						_implAppenderEmail.Write(msg, tci)
					}
				case VAL_APPENDER_MONGODB:
					if nil != _implAppenderMongodb {
						_implAppenderMongodb.Write(msg, tci)
					}
				}
			}
		}
	}
}

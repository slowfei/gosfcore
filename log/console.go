//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-10-22
//  Update on 2013-10-24
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

// console handle
package SFLog

import (
	"fmt"
)

const (
	kConsoleDefaultPattern = `${yyyy}-${MM}-${dd} ${hh}:${mm}:${ss}${SSSSSS} [${TARGET}] ([${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}])\n${MSG}`
)

//	appender console config
type AppenderConsoleConfig struct {
	Pattern string `json:"ConsolePattern"` // 信息内容输出格式
}

// Appender impl console write
type AppenderConsole struct {
}

//	new console impl
func NewAppenderConsole() *AppenderConsole {
	return &AppenderConsole{}
}

//	#interface impl
//	控制台信息写入
func (ac *AppenderConsole) Write(msg *LogMsg, configInfo interface{}) {
	if nil == msg {
		return
	}

	var pattern string

	if nil == configInfo {
		pattern = kConsoleDefaultPattern
	} else {
		if consoleConfig, ok := configInfo.(*AppenderConsoleConfig); ok {
			if 0 == len(consoleConfig.Pattern) {
				pattern = kConsoleDefaultPattern
			} else {
				pattern = consoleConfig.Pattern
			}
		}
	}

	if 0 != len(pattern) {
		fmt.Println(logMagFormat(pattern, msg))
	}

}

//	name = console
func (ac *AppenderConsole) Name() string {
	return VAL_APPENDER_CONSOLE
}

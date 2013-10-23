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

//	控制台信息写入
func (ac *AppenderConsole) Write(msg *LogMsg, configInfo *TargetConfigInfo) {
	fmt.Println(logMagFormat(configInfo.AppenderConsoleConfig.Pattern, msg))
}

//	name = console
func (ac *AppenderConsole) Name() string {
	return VAL_APPENDER_CONSOLE
}

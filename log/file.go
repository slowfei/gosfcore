//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-10-31
//  Update on 2013-10-31
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

// file handle
package SFLog

import (
	"fmt"
)

//	appender file
type AppenderFileConfig struct {
	MaxSize int64  `json:"FileMaxSize"` // 文件大小 byte
	Path    string `json:"FilePath"`    // 文件存储路径
	Name    string `json:"FileName"`    // 文件名(可以输入时间格式) file.log-{yyyy-MM-dd}
	Pattern string `json:"FilePattern"` // 信息内容输出格式
}

// Appender impl console write
type AppenderFile struct {
}

//	new console impl
func NewAppenderFile() *AppenderFile {
	return &AppenderFile{}
}

//	#interface impl
//	控制台信息写入
func (af *AppenderFile) Write(msg *LogMsg, configInfo interface{}) {
	if fileConfig, ok := configInfo.(*AppenderFileConfig); ok {
		fmt.Println(logMagFormat(fileConfig.Pattern, msg))
	}
}

//	name = file
func (af *AppenderFile) Name() string {
	return VAL_APPENDER_FILE
}

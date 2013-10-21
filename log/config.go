//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-10-17
//  Update on 2013-10-17
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//	日志的配置文件加载操作
//	config keyword key:
//	"info_appender" : "file"
//
package SFLog

import (
	"encoding/json"
	"github.com/slowfei/gosfcore/utils/filemanager"
)

const (
	//	separator
	LOG_SEPARATOR = "_"
	//	global log config
	KEY_GLOBAL_LOG_TAG_CONFIG = "global_log_config"
	//	default channel buffer size
	DEFAULT_CHANNEL_BUFFER_SIZE = "3000"

	//	log target
	LOG_INFO  = "info"
	LOG_DEBUG = "debug"
	LOG_ERROR = "error"
	LOG_WARN  = "warn"
	LOG_FATAL = "fatal"
	LOG_PANIC = "panic"

	//	appender tag
	LOG_APPENDER = "appender"

	//	appender type
	VAL_APPENDER_CONSOLE = "console"
	VAL_APPENDER_FILE    = "file"
	VAL_APPENDER_MONGODB = "mongodb"
	VAL_APPENDER_EMAIL   = "email"
	VAL_APPENDER_HTML    = "html"
	VAL_APPENDER_NONE    = "none"
)

var (
	//	log target value
	TargetInfo  = LogTarget("info")
	TargetDebug = LogTarget("debug")
	TargetError = LogTarget("error")
	TargetWarn  = LogTarget("Warn")
	TargetFatal = LogTarget("fatal")
	TargetPanic = LogTarget("panic")

	//	默认配置的json
	_defaultConfig = `
		{
			"ChannelSize" : ` + DEFAULT_CHANNEL_BUFFER_SIZE + `,
			"LogTags" :{
				"` + KEY_GLOBAL_LOG_TAG_CONFIG + `" :{
					"Appender":[
							"console"
					],
					"info":{
							"ConsolePattern":"yyyy-MM-dd",
							"none":true
					},
					"debug":{
							"none":true
					},
					"error":{
							"none":true
					},
					"warn":{
							"none":true
					},
					"fatal":{
							"none":true
						},
					"panic":{
							"none":true
					}
				}
			}
		}
	`

	//	log config map
	_sharedLogConfig *MainLogConfig = nil
)

//	log init
func init() {
	_sharedLogConfig = new(MainLogConfig)
	// _sharedLogConfig.ChannelSize = DEFAULT_CHANNEL_BUFFER_SIZE
	_sharedLogConfig.LogTags = make(map[string]LogTagConfig)

	//	初始化的时候加载一次默认的配置
	if err := json.Unmarshal([]byte(_defaultConfig), _sharedLogConfig); nil != err {
		panic(err)
	}
}

//	reset load config
//	@filePath	相对或绝对路径
func LoadConfig(filePath string) error {

	if 0 != len(filePath) {

		jsonData, e1 := SFFileManager.ReadFileAll(filePath)
		if nil != e1 {
			return e1
		}

		e2 := json.Unmarshal(jsonData, _sharedLogConfig)

		if nil != e2 {
			return e2
		}

	}

	return nil
}

//	main log config
type MainLogConfig struct {
	ChannelSize int                     // 通道缓冲区大小
	LogTags     map[string]LogTagConfig // log tags日志标识集合元素
}

//	log tag config
type LogTagConfig struct {
	Appender []string
	Info     *ConfigInfo `json:"info"`
	Debug    *ConfigDebug
	Error    *ConfigError
	Warn     *ConfigWarn
	Fatal    *ConfigFatal
	Panic    *ConfigPanic
}

//	log target, info、Debug、Error、Warn、Fatal、Panic
type LogTarget string

//	info config
type ConfigInfo struct {
	*AppenderFile
	*AppenderConsole
	*AppenderMongodb
	*AppenderEmail
	*AppenderHtml
	AppenderNone
}

//	debug config
type ConfigDebug struct {
	*AppenderFile
	*AppenderConsole
	*AppenderMongodb
	*AppenderEmail
	*AppenderHtml
	AppenderNone
}

//	error config
type ConfigError struct {
	*AppenderFile
	*AppenderConsole
	*AppenderMongodb
	*AppenderEmail
	*AppenderHtml
	AppenderNone
}

//	warn config
type ConfigWarn struct {
	*AppenderFile
	*AppenderConsole
	*AppenderMongodb
	*AppenderEmail
	*AppenderHtml
	AppenderNone
}

//	fatal config
type ConfigFatal struct {
	*AppenderFile
	*AppenderConsole
	*AppenderMongodb
	*AppenderEmail
	*AppenderHtml
	AppenderNone
}

//	panic config
type ConfigPanic struct {
	*AppenderFile
	*AppenderConsole
	*AppenderMongodb
	*AppenderEmail
	*AppenderHtml
	AppenderNone
}

//	appender console
type AppenderConsole struct {
	Pattern string `json:"ConsolePattern"` // 信息内容输出格式
}

//	appender file
type AppenderFile struct {
	MaxSize int64  `json:"FileMaxSize"` // 文件大小 byte
	Path    string `json:"FilePath"`    // 文件存储路径
	Name    string `json:"FileName"`    // 文件名(可以输入时间格式) file.log-{yyyy-MM-dd}
	Pattern string `json:"FilePattern"` // 信息内容输出格式
}

//	appender mongodb
type AppenderMongodb struct {
}

//	appender email
type AppenderEmail struct {
}

//	appender html
type AppenderHtml struct {
}

//	appender none
type AppenderNone struct {
	None bool
}

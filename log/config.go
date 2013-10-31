//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-10-17
//  Update on 2013-10-22
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//	日志的配置文件加载操作
package SFLog

import (
	"encoding/json"
	"io/ioutil"
)

const (
	//	separator
	LOG_SEPARATOR = "_"
	//	global group log config
	KEY_GLOBAL_GROUP_LOG_CONFIG = "globalGroup"
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
	TargetWarn  = LogTarget("warn")
	TargetFatal = LogTarget("fatal")
	TargetPanic = LogTarget("panic")

	//	默认配置的json
	_defaultConfig = `
		{
			"InitAppenders":[
				"console"
			],
			"ChannelSize" : ` + DEFAULT_CHANNEL_BUFFER_SIZE + `,
			"LogGroups" :{
				"` + KEY_GLOBAL_GROUP_LOG_CONFIG + `" :{

					"Appender":[
						"console"
					],
					"none":false,
					"ConsolePattern":"${yyyy}-${MM}-${dd} ${hh}:${mm}:${ss}${SSSSSS} [${TARGET}] ([${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}])\n${MSG}"
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
	_sharedLogConfig.LogGroups = make(map[string]LogConfig)

	//	初始化的时候加载一次默认的配置
	if err := json.Unmarshal([]byte(_defaultConfig), _sharedLogConfig); nil != err {
		panic(err)
	}
}

//	reset load config
//	@filePath	相对或绝对路径
func LoadConfig(filePath string) error {
	if 0 != len(filePath) {

		jsonData, e1 := ioutil.ReadFile(filePath)
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

//	reset load config
//	@jsonDataq
func LoadConfigByJson(jsonData []byte) error {

	e2 := json.Unmarshal(jsonData, _sharedLogConfig)

	if nil != e2 {
		return e2
	}

	return nil
}

//	main log config
type MainLogConfig struct {
	ChannelSize   int                  // 通道缓冲区大小
	InitAppenders []string             // init appenders impl. console, file...
	TimeFormat    string               // time format
	LogGroups     map[string]LogConfig // log tags日志标识集合元素
}

//	log config
type LogConfig struct {
	//	target config
	Info  *TargetConfigInfo `json:"info"`
	Debug *TargetConfigInfo `json:"debug"`
	Error *TargetConfigInfo `json:"error"`
	Warn  *TargetConfigInfo `json:"warn"`
	Fatal *TargetConfigInfo `json:"fatal"`
	Panic *TargetConfigInfo `json:"panic"`

	//	global config
	*TargetConfigInfo
}

//	log target, info,debug,error,warn,fatal,panics
type LogTarget string

//	target config info,contain info,debug,error,warn,fatal,panics
type TargetConfigInfo struct {
	Appender []string
	*AppenderFileConfig
	*AppenderConsoleConfig
	*AppenderMongodbConfig
	*AppenderEmailConfig
	*AppenderHtmlConfig
	AppenderNoneConfig
}

//	appender none
type AppenderNoneConfig struct {
	None bool
}

//	appender mongodb
type AppenderMongodbConfig struct {
}

//	appender email
type AppenderEmailConfig struct {
}

//	appender html
type AppenderHtmlConfig struct {
}

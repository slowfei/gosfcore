//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-08-24
//  Update on 2013-11-05
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//	日志操作，能够分别处理不同的日志级别信息配置
//	info debug error warn ftal panic level
//	out console file html mongodb email
//
//	使用说明：
//		SFLogger struct{}，首先先要了解什么是(SFLogger)，是用于标识区分每个log信息的输出，
//		可以自定义分组标识和log标识。
//
//		日志标识(logTag)：
//			主要作用区分每个不同的SFLogger对象进行输出的信息，最好是唯一的
//		日志组标识(logGroup)：
//			主要作用是用于日志配置的使用，在一个日志分组中使用同样的配置操作。
//
//		如果可以直接使用全局的日志配置操作，全局日志的配置默认是输出控制台。
//		全局SFLogger的标识：logTag = "globalTag"，logGroup ＝ globalGroup
//
//			SFLog.Info("操作信息：记录信息操作。")
//
//			console out:
//			2013-10-31 12:12:55.871435 [info] ([globalGroup][globalTag][L16 github.com/slowfei/gosfcore/log.TestLogger])
//			操作信息：记录信息操作。
//
//		也可以自定义一个日志标识然后结合日志的配置进行信息的输出，如果没有定义日志组，默认使用全局日志组的配置。
//			var log *SFLogger = NewLogger("logtag") or NewLogger("logtag","logGroup")
//
//			log.Info("操作信息：记录信息操作。")
//
//			信息会根据日志组的设置进行相应的输出。
//
//	配置文件加载
//
//	配置详解：
//	Pattern Format(信息输出时的格式化操作)：
//		${yyyy}			年
//		${MM}			月
//		${dd}			日
//		${hh}			时
//		${mm}			分
//		${ss}			秒
//		${SSS}			毫秒
//		${LOG_GROUP}	分组标识
//		${LOG_TAG}		日志标识
//		${FILE_LINE}	调用函数的文件行
//		${FILE_PATH}	调用函数的文件路径
//		${FUNC_NAME}	函数名称(哪里调用就是那个函数)
//		${STACK}		堆栈信息
//		${TARGET}		输出的目标例如 info、debug、error...
//		${MSG}			输出的信息，就是 SFLog.Info("这里是输出${MSG}的信息")
//
//	配置文件(千万要注意编写json的格式)
// {
//		//	初始化需要实现的的Appender对象，如果未初始化则不会进行输出，所以在开始前需要确定需要输出的对象。
// 		"InitAppenders":[
// 			"console","file","email","html","mongodb"
// 		],
//
//		//	日志组的配置，包含多个日志组的配置信息
// 		"LogGroups" :{
//
//			//	配置一个日志组
// 			"groupName" :{
//
//					//	设置需要的Appender对象，如果未配置将不会进行输出
// 					"Appender":[
// 						"console","file"
// 					],
//
//					//	下面针对Appender对象配置特定的格式信息，如果nil或没有设置则使用Appender的默认设置
//
//					/* ------------console配置--------------- */
//
//					//	控制台输出的格式，具体可以查看Pattern Format
// 					"ConsolePattern":"${yyyy}-${MM}-${dd} ${mm}:${dd}:${ss}${SSSSSS} [${TARGET}] ([${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}])\n${MSG}",
//
//
//					/* ------------file配置--------------- */
//
//					//	文件名(可以输入时间格式)  默认"(ExceFileName)-${yyyy}-${MM}-${dd}.log"
//					//	配置注意事项：
//					//	Name(FileName)  "file-${yyyy}/${MM}/${dd}.log" 	  error		如果包含"/"会以目录作为处理的，所以需要注意。
//					//					"../file-${yyyy}-${MM}-${dd}.log" proper	可以使用相对路径来命名"/"是作为目录的操作，
//					//																截取后面的文件名(file-${yyyy}-${MM}-${dd}.log)
//					"FileName":"info-${yy}${MM}${dd}.log",
//
//					//	文件存储路径, 默认执行文件目录
//					"FileSavePath":"",
//
//					//	输出的格式，具体可以查看Pattern Format
//					"FilePattern":"${yyyy}-${MM}-${dd} ${mm}:${dd}:${ss}${SSSSSS} [${TARGET}] ([${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}])\n${MSG}",
//
//					//	文件最大存储大小，默认5M
//					"FileMaxSize":"5242880",
//
//					//	日志相同名称的最大数量，例如file(1).log...file(1000).log。默认1000，超出建立的数量将不会创建日志文件
//					"FileSameNameMaxNum":"1000",
//
//
//					/* ------------html配置--------------- */
//
//					//	注意事项与file配置的Name相同
//					"HtmlName":"log-${yy}${MM}${dd}.html",
//
//					//	文件存储路径, 默认执行文件目录
//					"HtmlSavePath":"",
//
//					//	html title
//					"HtmlTitle":"Log Info",
//
//					//	时间格式需要注意没有${}
//					"HtmlTimeFormat":"yyyy-MM-dd hh:mm",
//
//					//	文件最大存储大小，默认3M
//					"HtmlMaxSize":"3145728",
//
//					//	与file配置相同
//					"HtmlSameNameMaxNum":"1000",
//
//
//					/* ------------email配置--------------- */
//
//					//	不可为空，否则不进行输出
//					"EmailHost":"smtp.xxx.com",
//
//					//	非空
//					"EmailUser":"xxx@gmail.com",
//
//					//	非空
//					"EmailPassword":"123456",
//
//					//	发送邮件显示的发送人名称
//					"EmailFromName":"slowfei",
//
//					//	发送地址
//					"EmailTo":"xx@gmail.com;xx2@gmail.com",
//
//					//	输出信息的格式，具体可以查看Pattern Format
//					"EmailPattern":"${yyyy}-${MM}-${dd} ${mm}:${dd}:${ss}${SSSSSS} [${TARGET}] ([${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}])\n${MSG}",
//
//					//	邮件主题
//					"EmailSubject":"Golang Log Info",
//
//					//	默认Content-Type: text/plain; charset=UTF-8
//					"EmailContentType":"Content-Type: text/plain; charset=UTF-8"
//
//
//					//	控制当前日志组是否进行输出工作，如果为true则当前组不会进行信息的输出，默认可以不写为false
//					"none":false,
//
//					//	以上的分组配置均为默认配置
//
//					//	针对输出的目标进行配置，如果不编写则使用上面部分设置的默认配置信息。
//					//	需要注意的是，只要声明了目标的配置就不会取组的默认配置信息，目标配置大于默认配置。
//					"info":{
//						"Appender":[
// 							"console"
// 						],
//						"ConsolePattern":"${yyyy}-${MM}-${dd} ${mm}:${dd} ${MSG}"
//					},
//					"debug":{
//						"Appender":[
// 							"file"
// 						],
//						"FileName":"info-${yy}-${MM}-${dd}.log",
//						"FileSavePath":"",
//						"FilePattern":"${yyyy}-${MM}-${dd} ${mm}:${dd}:${ss}${SSSSSS} [${TARGET}] ([${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}])\n${MSG}",
//						"FileMaxSize":"5242880",
//						"FileSameNameMaxNum":"1000"
//					},
//					"error":{
//						"Appender":[
// 							"html"
// 						],
//						"HtmlName":"log-${yy}${MM}${dd}.html",
//						"HtmlSavePath":"",
//						"HtmlTitle":"LogInfo",
//						"HtmlTimeFormat":"yyyy-MM-dd hh:mm",
//						"HtmlMaxSize":"3145728",
//						"HtmlSameNameMaxNum":"1000",
//					},
//					"warn":{
//						"Appender":[
// 							"email"
// 						],
//						"EmailHost":"smtp.xxx.com",
//						"EmailUser":"xxx@gmail.com",
//						"EmailPassword":"123456",
//						"EmailFromName":"slowfei",
//						"EmailTo":"xx@gmail.com;xx2@gmail.com",
//						"EmailPattern":"${yyyy}-${MM}-${dd} ${mm}:${dd}:${ss}${SSSSSS} [${TARGET}] ([${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}])\n${MSG}",
//						"EmailSubject":"Golang Log Info",
//						"EmailContentType":"Content-Type: text/plain; charset=UTF-8"
//					},
//					"fatal":{
//						//	配置与info都一致。
//					},
//					"panic":{
//						//	配置与info都一致。
//					}
// 				}
// 			}
// 	}
//
package SFLog

import (
	"bytes"
	"fmt"
	"github.com/slowfei/gosfcore/utils/time"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	// log manager
	_thisLogManager *LogManager = nil

	//	global logger
	_globalLogger *SFLogger = NewLogger("globalTag")

	//	appender impl struct
	ImplAppenderConsole Appender = nil
	ImplAppenderFile    Appender = nil
	ImplAppenderHtml    Appender = nil
	ImplAppenderEmail   Appender = nil
	ImplAppenderMongodb Appender = nil
)

//	logger 产生日志的输出，主要负责标识每个不同的日志对象，
//	使用分组标识和日志标识进行标识处理。
//	然后可调用(Info、Debug、Error)函数进行信息输出
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
func NewLoggerByGroup(logTag, logGroup string) *SFLogger {
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

//	global info log
func Info(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetInfo, format, v...)
}

//	global debug log
func Debug(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetDebug, format, v...)
}

//	global error log
func Error(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetError, format, v...)
}

//	global warn log
func Warn(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetWarn, format, v...)
}

//	global fatal log
func Fatal(format string, v ...interface{}) string {
	return loggerHandle(_globalLogger, TargetFatal, format, v...)
}

//	global panic log
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
	if 0 != len(v) {
		msgString = fmt.Sprintf(format, v...)
	} else {
		msgString = format
	}

	fileLine := -1
	funcName := "???"
	filePath := "???"

	stackBuf := bytes.NewBufferString("")
	for i := 2; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc).Name()
		if i == 2 {
			fileLine = line
			funcName = fn
			filePath = file

			if TargetInfo == target {
				//	如果目标为info就不打印其他堆栈信息了。
				break
			}
		}
		if 0 != len(file) {
			//	L1223: runtime.goexit(...) (0x173d0)
			fmt.Fprintf(stackBuf, "%s(...)\n%s:%d (0x%x)\n", fn, file, line, pc)
		} else {
			// 	runtime.goexit(...)
			// /usr/local/go/src/pkg/runtime/proc.c:1223 (0x173d0)
			fmt.Fprintf(stackBuf, "L%d: %s(...) (0x%x)\n", line, fn, pc)
		}
	}

	logmsg := &LogMsg{}
	logmsg.logGroup = log.logGroup
	logmsg.logTag = log.logTag
	logmsg.dateTime = time.Now()
	logmsg.stack = stackBuf.String()
	logmsg.fileLine = fileLine
	logmsg.funcName = funcName
	logmsg.filePath = filePath
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
	fileLine int       // go file line
	filePath string    // file path
	funcName string    // msg trigger func name
	msg      string    // log message
}

//	格式化日志信息
func logMagFormat(format string, msg *LogMsg) string {

	if nil == msg {
		return ""
	}
	if 0 == len(format) {
		format = "${yyyy}-${MM}-${dd} ${hh}:${mm}:${ss}${SSSSSS} [${TARGET}] ([${LOG_GROUP}][${LOG_TAG}][L${FILE_LINE} ${FUNC_NAME}])\n${MSG}"
	}

	//	格式化时间
	format = SFTimeUtil.YMDHMSSSignFormat(msg.dateTime, format)

	//	日志的格式化信息，别随意更换顺序，因为根据设计来进行日志信息的格式化操作
	logFormat := []string{
		"${LOG_GROUP}", msg.logGroup,
		"${LOG_TAG}", msg.logTag,
		"${FILE_LINE}", strconv.Itoa(msg.fileLine),
		"${FILE_PATH}", msg.filePath,
		"${FUNC_NAME}", msg.funcName,
		"${STACK}", msg.stack,
		"${TARGET}", string(msg.target),
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

//	start log manager
//	调用此方法启动时加载默认配置进行设置，
//	如果需要加载配置文件可以调用LoadConfig进行相应的设置。
//
//	@logChannelSize log的缓存区大小，默认
func StartLogManager(logChannelSize int) {
	if nil == _thisLogManager {
		_thisLogManager = new(LogManager)
		_thisLogManager.msg = make(chan *LogMsg, logChannelSize)
		loadAppenders()
		go _thisLogManager.startChannel()
	}
}

//	init appenders
func loadAppenders() {
	if nil != _thisLogManager {

		//	初始化appender的实现
		for _, appender := range _sharedLogConfig.InitAppenders {
			switch appender {
			case VAL_APPENDER_CONSOLE:
				if nil == ImplAppenderConsole {
					ImplAppenderConsole = NewAppenderConsole()
				}
			case VAL_APPENDER_FILE:
				if nil == ImplAppenderFile {
					ImplAppenderFile = NewAppenderFile()
				}
			case VAL_APPENDER_HTML:
				if nil == ImplAppenderHtml {
					ImplAppenderHtml = NewAppenderHtml()
				}
			case VAL_APPENDER_EMAIL:
				if nil == ImplAppenderEmail {
					ImplAppenderEmail = NewAppenderEmail()
				}
			case VAL_APPENDER_MONGODB:
			}
		}

	}

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
			} else if nil != config.Info {
				tci = config.Info
			}
		case TargetDebug:
			if nil == config.Debug {
				tci = config.TargetConfigInfo
			} else if nil != config.Debug {
				tci = config.Debug
			}
		case TargetError:
			if nil == config.Error {
				tci = config.TargetConfigInfo
			} else if nil != config.Error {
				tci = config.Error
			}
		case TargetWarn:
			if nil == config.Warn {
				tci = config.TargetConfigInfo
			} else if nil != config.Warn {
				tci = config.Warn
			}
		case TargetFatal:
			if nil == config.Fatal {
				tci = config.TargetConfigInfo
			} else if nil != config.Fatal {
				tci = config.Fatal
			}
		case TargetPanic:
			if nil == config.Panic {
				tci = config.TargetConfigInfo
			} else if nil != config.Panic {
				tci = config.Panic
			}
		}

		if nil != tci && !tci.AppenderNoneConfig.None {

			count := len(tci.Appender)
			for i := 0; i < count; i++ {
				appenderName := tci.Appender[i]

				//	扩展性的问题，由于考虑到使用for的话没有必要，毕竟需要扩展也不是很多或经常，所以目前直接使用switch来判断。
				//	所以如果需要扩展还需要另外编写代码
				switch appenderName {
				case VAL_APPENDER_CONSOLE:
					if nil != ImplAppenderConsole && nil != tci.AppenderConsoleConfig {
						ImplAppenderConsole.Write(msg, tci.AppenderConsoleConfig)
					}
				case VAL_APPENDER_FILE:
					if nil != ImplAppenderFile && nil != tci.AppenderFileConfig {
						ImplAppenderFile.Write(msg, tci.AppenderFileConfig)
					}
				case VAL_APPENDER_HTML:
					if nil != ImplAppenderHtml && nil != tci.AppenderHtmlConfig {
						ImplAppenderHtml.Write(msg, tci.AppenderHtmlConfig)
					}
				case VAL_APPENDER_EMAIL:
					if nil != ImplAppenderEmail && nil != tci.AppenderEmailConfig {
						ImplAppenderEmail.Write(msg, tci.AppenderEmailConfig)
					}
				case VAL_APPENDER_MONGODB:
					if nil != ImplAppenderMongodb && nil != tci.AppenderMongodbConfig {
						ImplAppenderMongodb.Write(msg, tci.AppenderMongodbConfig)
					}
				}
			}

		}
	}
}

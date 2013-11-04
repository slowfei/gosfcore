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
	"github.com/slowfei/gosfcore/utils/filemanager"
	"github.com/slowfei/gosfcore/utils/time"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	DEFAULT_FILE_MAX_SIZE = 5 << 20 // default 5M file max size
	DEFAULT_FILE_MAX_NUM  = 1000    // default same name log file max num
)

//	appender file
type AppenderFileConfig struct {
	MaxSize    int64  `json:"FileMaxSize"`        // 文件大小 byte
	Path       string `json:"FilePath"`           // 文件存储路径
	Name       string `json:"FileName"`           // 文件名(可以输入时间格式) file.log-{yyyy-MM-dd}
	Pattern    string `json:"FilePattern"`        // 信息内容输出格式
	SameMaxNum string `json:"FileSameNameMaxNum"` // 日志相同名称的最大数量，例如file(1).log...file(1000).log。默认1000，超出建立的数量将不会创建日志文件
}

// Appender impl console write
type AppenderFile struct {
	excePath        string              // file save path
	defaultFileName string              // default file name
	files           map[string]*os.File // key = file name
	rwmutex         sync.RWMutex
}

//	new console impl
func NewAppenderFile() *AppenderFile {
	af := &AppenderFile{}
	af.excePath = SFFileManager.GetExceDir()
	//	默认存储文件以天来存储
	af.defaultFileName = SFFileManager.GetExceFileName() + "-${yyyy}-${MM}-${dd}.log"
	af.files = make(map[string]*os.File)
	return af
}

//	获取文件对象
//	@fileName   文件名称
//	@t 		    日志操作的时间
//	@maxSize    日志的最大容量大小
//	@sameMaxNum 日志相同名的数量
func (af *AppenderFile) getFile(fileName string, t time.Time, maxSize int64, sameMaxNum int64) *os.File {
	//	TODO  opfName感觉不对，因为如果fileName进来的是file.log，那如果存储的是file(1).log呢？该如何去处理？
	//	考虑使用file.Name()来处理
	opfName := ""
	if 0 == len(fileName) {
		opfName = af.defaultFileName
	}
	opfName = SFTimeUtil.YMDHMSSSignFormat(t, fileName)

	if file, ok := af.files[opfName]; ok {
		if file.Stat().Size() < maxSize {
			return file
		}
		file.Close()
		delete(af.files, opfName)

		opfName = SFFileManager.FileRenameRule(opfName)
	}

	oppath := filepath.Join(af.excePath, opfName)
	fileInfo, err := os.Stat(path)

	if nil == err {
		if fileInfo.Size() >= maxSize {
			var errExists error = nil
			//	查询是否有相同的文件，如果有文件命名file(1).log file(2).log file(3).log进行递增
			for i := 0; i < sameMaxNum; i++ {
				opfName = SFFileManager.FileRenameRule(opfName)
				oppath = filepath.Join(af.excePath, opfName)
				fileInfo, errExists = os.Lstat(path)

				if nil != errExists || fileInfo.Size() < maxSize {
					//	找到未建立的文件名称，可以作为日志的存储。
					//	或存储容量大小未达到最大的限制，也可以作为日志的存储。
					break
				}
			}

			if nil == errExists {
				//	文件操作建立的组大范围，不进行操作了
				//	TODO 考虑创建数的日志文件数最大范围了如何人性化的处理好。
				return nil
			}

		}
	}

	//	TODO 需要查找perm的参数信息
	// os.OpenFile(name, flag, perm)

	return nil
}

//
func (af *AppenderFile) fileWrite(fileName, msg string, maxSize int64) {
	af.rwmutex.Lock()
	defer af.rwmutex.Unlock()
}

//	#interface impl
//	控制台信息写入
func (af *AppenderFile) Write(msg *LogMsg, configInfo interface{}) {
	if fileConfig, ok := configInfo.(*AppenderFileConfig); ok {
		formatStr := logMagFormat(fileConfig.Pattern, msg)

		fmt.Println(formatStr)
	}
}

//	name = file
func (af *AppenderFile) Name() string {
	return VAL_APPENDER_FILE
}

//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-10-31
//  Update on 2013-11-05
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
	"runtime/debug"
	"sync"
	"time"
)

const (
	DEFAULT_FILE_MAX_SIZE = 5 << 20 // default 5M file max size
	DEFAULT_FILE_MAX_NUM  = 1000    // default same name log file max num
)

//	appender file
//	配置注意事项：
//	Name(FileName)  "file-${yyyy}/${MM}/${dd}.log" 	  error		如果包含"/"会以目录作为处理的，所以需要注意。
//					"../file-${yyyy}-${MM}-${dd}.log" proper	可以使用相对路径来命名"/"是作为目录的操作，截取后面的文件名(file-${yyyy}-${MM}-${dd}.log)
//
type AppenderFileConfig struct {
	MaxSize    int64  `json:"FileMaxSize"`        // 文件大小 byte			默认5M
	SavePath   string `json:"FileSavePath"`       // 文件存储路径, 			默认执行文件目录
	Name       string `json:"FileName"`           // 文件名(可以输入时间格式)  默认"(ExceFileName)-${yyyy}-${MM}-${dd}.log"
	Pattern    string `json:"FilePattern"`        // 信息内容输出格式
	SameMaxNum int    `json:"FileSameNameMaxNum"` // 日志相同名称的最大数量，例如file(1).log...file(1000).log。默认1000，超出建立的数量将不会创建日志文件
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
//	@savePath	文件存储路径，		"" = 执行文件的目录
//	@fileName   文件名称				"" = "(ExceFileName)-${yyyy}-${MM}-${dd}.log"
//	@t 		    日志操作的时间
//	@maxSize    日志的最大容量大小		默认 5M
//	@sameMaxNum 日志相同名的数量		默认 1000
func (af *AppenderFile) getFile(savePath, fileName string, t time.Time, maxSize int64, sameMaxNum int) *os.File {

	var opfName string // 用于操作打开文件的文件名存储
	var kfName string  // 用于存储af.files的keyname

	if 0 == len(fileName) {
		fileName = af.defaultFileName
	}
	if 0 >= maxSize {
		maxSize = DEFAULT_FILE_MAX_SIZE
	}
	if 0 >= sameMaxNum {
		sameMaxNum = DEFAULT_FILE_MAX_NUM
	}
	if 0 == len(savePath) {
		savePath = af.excePath
	}

	opfName = SFTimeUtil.YMDHMSSSignFormat(t, fileName)

	//	主要存储的key名保持file.log不变，没有重命名的递增file(1).log
	kfName = opfName

	if file, ok := af.files[kfName]; ok {
		fileInfo, e := file.Stat()
		if e == nil && fileInfo.Size() < maxSize {
			return file
		}
		file.Close()
		delete(af.files, kfName)

		if nil != fileInfo {
			//	由于fileInfo.Name()获取得到的是一个文件名称(file.log)
			//	而传递的文件名(fileName)包含相对路径(../file-${yyyy}.log)，然后需要截取(../)与fileInfo.Name()拼接起来。
			//	以便filepath.Join能够连接准确
			//	需要注意fileInfo.Name()会可能获取到的名称带有(1)(2)的标识，所以截取不能使用它
			fiName := fileInfo.Name()
			opfName = opfName[:len(opfName)-len(filepath.Base(opfName))] + SFFileManager.FileRenameRule(fiName)
		}
	}

	oppath := filepath.Join(savePath, opfName)
	fileInfo, err := os.Stat(oppath)

	if nil == err {
		if fileInfo.Size() >= maxSize {
			var errExists error = nil
			isReturn := true
			//	查询是否有相同的文件，如果有文件命名file(1).log file(2).log file(3).log进行递增
			for i := 0; i < 20; i++ {
				opfName = SFFileManager.FileRenameRule(opfName)
				oppath = filepath.Join(savePath, opfName)
				fileInfo, errExists = os.Lstat(oppath)

				if nil != errExists || fileInfo.Size() < maxSize {
					//	找到未建立的文件名称，可以作为日志的存储。
					//	或存储容量大小未达到最大的限制，也可以作为日志的存储。
					isReturn = false
					break
				}
			}
			if isReturn {
				//	文件操作建立的组大范围，不进行操作了
				//	TODO 考虑创建日志文件数超出了最大范围了该如何人性化的处理更好。
				return nil
			}

		}
	}
	// flag可选值：
	// O_RDONLY int = os.O_RDONLY // 只读
	// O_WRONLY int = os.O_WRONLY // 只写
	// O_RDWR   int = os.O_RDWR   // 读写
	// O_APPEND int = os.O_APPEND // 在文件末尾追加，打开后cursor在文件结尾位置
	// O_CREATE int = os.O_CREATE // 如果不存在则创建
	// O_EXCL   int = os.O_EXCL   // 与O_CREATE一起用，构成一个新建文件的功能，它要求文件必须不存在
	// O_SYNC   int = os.O_SYNC   // 同步方式打开，没有缓存，这样写入内容直接写入硬盘，系统掉电文件内容有一定保证
	// O_TRUNC  int = os.O_TRUNC  // 打开并清空文件

	// 权限位，讲设置的权限数值进行累加
	// 用户
	//		0400	允许所有者读。
	// 		0200	允许所有者写。
	// 		0100	对于文件，允许所有者执行，对于目录，允许所有者在该目录中进行搜索。
	// 组
	//		0040	允许组成员读。
	// 		0020	允许组成员写。
	// 		0010	对于文件，允许组成员执行，对于目录，允许组成员在该目录中进行搜索。
	// 其他用户
	//		0004	允许其他用户读。
	// 		0002	允许其他用户写。
	// 		0001	对于文件，允许其他用户执行，对于目录，允许其他用户在该目录中进行搜索。

	//	660 = 400 + 200 + 40 + 20
	//	参考：http://www.ibm.com/developerworks/cn/aix/library/au-speakingunix4/
	//	由于是文件第一位是0

	newFile, errFile := os.OpenFile(oppath, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0660)

	if nil != errFile {
		fmt.Printf("%v\n%s\n", errFile, debug.Stack())
		return nil
	}

	af.files[kfName] = newFile
	return newFile
}

// file write
func (af *AppenderFile) fileWrite(msg string, t time.Time, config *AppenderFileConfig) {
	af.rwmutex.Lock()
	defer af.rwmutex.Unlock()

	fileName := config.Name
	maxSize := config.MaxSize
	sameMaxNum := config.SameMaxNum
	savePath := config.SavePath

	file := af.getFile(savePath, fileName, t, maxSize, sameMaxNum)
	if nil != file {
		fmt.Fprintln(file, msg)
	}
}

//	关闭所有日志文件
func (af *AppenderFile) CloseAllLogFile() {
	for k, v := range af.files {
		v.Close()
		delete(af.files, k)
	}
}

//	#interface impl
//	控制台信息写入
func (af *AppenderFile) Write(msg *LogMsg, configInfo interface{}) {
	if fileConfig, ok := configInfo.(*AppenderFileConfig); ok {
		formatStr := logMagFormat(fileConfig.Pattern, msg)
		af.fileWrite(formatStr, msg.dateTime, fileConfig)
	}
}

//	name = file
func (af *AppenderFile) Name() string {
	return VAL_APPENDER_FILE
}

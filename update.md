gosfcore update log
=============

2015-07-31
1. src/github.com/slowfei/gosfcore/encoding/json/SFJson.go
> * 完善：增加struct tag 命名的支持 `json:"keyName"`

2015-06-25
1. src/github.com/slowfei/gosfcore/utils/sub/SFSubUtil.go
> * 修正：GetOutBetweens(...)修正一个出现的bug。

2015-06-12
1. src/github.com/slowfei/gosfcore/utils/sub/SFSubUtil.go
> * 增加：增加排除坐标的算法函数GetOutBetweens(...)...
> * 修正：GetOutBetweens(...)编写此方法产生了很多BUG，还需要大量测试验证（目前的测试暂无问题）。

2015-05-15
1. src/github.com/slowfei/gosfcore/utils/sub/SFSubUtil.go
> * 修正：增加了一个不查找子级的一个控制（NewSubNotNest），这样便于查询"// comment http://slowfei.com/ " 这样的字符  

2015-03-06
1. src/github.com/slowfei/gosfcore/utils/sub/SFSubUtil.go
> * 修改：修复index计算的错误和startFindIndex与结果的相加

2015-02-27
1. src/github.com/slowfei/gosfcore/utils/sub/SFSubUtil.go
> * 修改：算法出现BUG，主要是出现过滤的字符没有重新查找。
> * 修改：由于计算过滤(outBetweens)出现原数据的下标位，所以增加了开始查询下标的计算startIndex

2015-01-20
1. src/github.com/slowfei/gosfcore/utils/filemanager/SFFileManager.go
> * 增加GOPATH的获取

2015-01-13
1. src/github.com/slowfei/gosfcore/utils/sub/SFSubUtil.go
> * 建立截取工具
> * 增加嵌套数据的截取工作例如"{"截取的所有数据"}"

2014-11-06
1. src/github.com/slowfei/gosfcore/utils/filemanager/SFFileManager.go
> * 修正WirteFilepath在没有创建目录情况下的问题

2014-09-25
1. src/github.com/slowfei/gosfcore/utils/filemanager/SFFileManager.go
> * 增加文件输出函数WirteFilepath(...)
1. src/github.com/slowfei/gosfcore/encoding/json/SFJson.go
> * 修正文件输出，使用SFFileManager.WirteFilepath


2014-09-23
1. src/github.com/slowfei/gosfcore/log/file.go
> * 修正file.go对保存路径的操作

2014-09-19
1. src/github.com/slowfei/gosfcore/encoding/json/SFJson.go
> * 增加bytes的格式化函数和json文件的输出函数

2014-09-16
1. src/github.com/slowfei/gosfcore/utils/strings/SFStringsUtil.go
> * 增加string ToInt8、ToInt16、ToInt32、ToInt、ToInt64

2014-08-25
1. src/github.com/slowfei/gosfcore/encoding/json/SFJson.go
> * 修改注释
> * 增加 map、array的类型子对象验证，支持嵌套验证(已进行测试)

2014-07-11
1. src/github.com/slowfei/gosfcore/helper/map.go
> * NewMap修改为指针返回

2014-06-19
1. /src/github.com/slowfei/gosfcore/utils/strings/SFStringsUtil.go
> * 增加ToLower(...)  ToUpper(...)

2014-06-13
1. /src/github.com/slowfei/gosfcore/log/SFLog.go
> * 修改查找日志分组输出的bug，并添加分组测试

2014-05-28
1. /src/github.com/slowfei/gosfcore/utils/filemanager/SFFileManager.go
> * 增加获取当前命令行目录函数GetCmdDir()

2013-11-18
1./src/github.com/slowfei/gosfcore/debug
> * debug工具，支持断点，打印参数结构信息

2013-11-06
1. github.com/slowfei/gosfcore/log
> * 实现log日志组件
> * 针对每个Info、Debug...进行不同的配置
> * 目前实现了console、file、html、email模块，sql和nosql后期有时间再弄。

2013-11-04
1. github.com/slowfei/gosfcore/utils/filemanager/SFFileManager.go
> * 增加FileRenameRule 文件名重复命名的规则
> * 增加FileCleanRenameRule 文件名重复命名规则的去除
1. github.com/slowfei/gosfcore/utils/time/SFTimeUtil.go
> 修改YMDHMSSS时间格式化格式为"_yyyy" to "-2006"，主要是因为 _2 是格式化日

#### 2013-10-31
1. github.com/slowfei/gosfcore/utils/filemanager/SFFileManager.go
> * 增加获取执行文件路径函数GetExceFilePath()
> * 增加获取执行文件名函数GetExceFileName()
> * 修改GetBuildDir为GetExceDir

1. github.com/slowfei/gosfcore/log/
> 控制台日志输出完成

#### 2013-10-21
1. (新)github.com/slowfei/gosfcore/log/SFLog.go
> 正在编写日志操作，未完成

#### 2013-10-14
1. (新)github.com/slowfei/gosfcore/utils/time/SFTimeUtil.go
> * yyyyMMddhhmmssSSS格式的时间对象格式化字符串
> * yyyyMMddhhmmssSSS格式的时间字符串解析成时间对象

#### 2013-10-1
1. github.com/slowfei/gosfcore/utils/reflect/SFReflectUtil.go
> * 修改被设置对象类型为slice，设置的值为非slice的设值。

#### 2013-9-26
1. github.com/slowfei/gosfcore/utils/rand
> * rand.Reader使用ReadFull读取的随机数
> * 指定范围的随机数int, float
> * a-zA-Z0-9的随机数字符串
> * 自定义字符的随机数 
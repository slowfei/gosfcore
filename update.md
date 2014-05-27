gosfcore update log
=============

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
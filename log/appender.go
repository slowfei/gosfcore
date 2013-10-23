//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-10-22
//  Update on 2013-10-24
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//	appender interface
package SFLog

//	输出目的的接口，例如需要实现console、file、html、mongodb、email
type Appender interface {
	//	写入信息
	//	@msg		 写入的信息
	//	@configInfo  配置信息
	Write(msg *LogMsg, configInfo *TargetConfigInfo)

	// 实现接口的名称
	Name() string
}

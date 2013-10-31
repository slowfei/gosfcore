//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-24
//  Update on 2013-10-31
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//	对于文件和目录的工具
//
package SFFileManager

import (
	"os"
	"os/exec"
	"path/filepath"
)

var (
	//	存储当前执行文件的路径，这样可以避免每次获取不必要的计算。
	currentExceFilePath string
)

//	获取当前执行文件的操作目录
//	@return
func GetExceDir() string {
	return filepath.Dir(GetExceFilePath())
}

//	获取当前执行文件的路径
//	@return
func GetExceFilePath() string {

	if 0 != len(currentExceFilePath) {
		return currentExceFilePath
	}

	exceFile, err := exec.LookPath(os.Args[0])
	if nil != err {
		panic(err)
	}

	exceFilePath, err2 := filepath.Abs(exceFile)
	if nil != err2 {
		panic(err2)
	}

	currentExceFilePath = exceFilePath

	return exceFilePath
}

//	获取当前执行文件的名称
//	@return
func GetExceFileName() string {
	return filepath.Base(GetExceFilePath())
}

//	判断路径是否存在文件或目录
//	@path	操作路径
//	@isDir	指针，引用传递接收 否是目录, true is dir
//	@reutrn	bool 文件存在 true
func Exists(path string, isDir *bool) (bool, error) {
	fileInfo, err := os.Stat(path)

	if err == nil {
		if nil != isDir {
			*isDir = fileInfo.IsDir()
		}
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

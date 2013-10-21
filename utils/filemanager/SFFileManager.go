//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-24
//  Update on 2013-10-17
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//	对于文件和目录的工具
//
package SFFileManager

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

//	获取当前编译文件的操作目录
//	@return
func GetBuildDir() string {
	return path.Dir(os.Args[0])
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

//	读取文件所有信息
//
//	@path	操作路径
//	@return  read data, error
func ReadFileAll(path string) ([]byte, error) {

	var isDir bool
	if b, _ := Exists(path, &isDir); !b || isDir {
		return nil, errors.New("file does not exist or can not be operated path: " + path)
	}

	file, e1 := os.Open(path)
	if nil != e1 {
		return nil, e1
	}
	defer file.Close()

	data, e2 := ioutil.ReadAll(file)
	if nil != e2 {
		return nil, e2
	}

	return data, nil
}

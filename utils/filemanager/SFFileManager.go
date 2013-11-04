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
	"regexp"
	"strconv"
)

var (
	//	存储当前执行文件的路径，这样可以避免每次获取不必要的计算。
	_currentExceFilePath string

	//	文件重命名规则的正则
	_rexRenameRule = regexp.MustCompile("\\(\\d+\\)")
)

//	获取当前执行文件的操作目录
//	@return
func GetExceDir() string {
	return filepath.Dir(GetExceFilePath())
}

//	获取当前执行文件的路径
//	@return
func GetExceFilePath() string {

	if 0 != len(_currentExceFilePath) {
		return _currentExceFilePath
	}

	exceFile, err := exec.LookPath(os.Args[0])
	if nil != err {
		panic(err)
	}

	exceFilePath, err2 := filepath.Abs(exceFile)
	if nil != err2 {
		panic(err2)
	}

	_currentExceFilePath = exceFilePath

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

//	文件名重复命名的规则，例如相同的名称重复命名 name.txt、name(2).txt、name(3).txt...
//	@name
//	@return
func FileRenameRule(name string) string {
	result := ""
	// strings.SplitN(s, sep, n)
	ext := filepath.Ext(name)
	if 0 != len(ext) {
		name = name[:len(name)-len(ext)]
	}
	rules := _rexRenameRule.FindAllStringSubmatchIndex(name, -1)

	if 0 != len(rules) && rules[len(rules)-1][1] == len(name) {
		lastRule := rules[len(rules)-1]
		num, _ := strconv.Atoi(name[lastRule[0]+1 : lastRule[1]-1])
		num++
		result = name[:lastRule[0]] + "(" + strconv.Itoa(num) + ")" + ext
		// result = fmt.Sprintf("%s(%d)%s", name[:lastRule[0]], num, ext)
	} else {
		result = name + "(1)" + ext
		// result = fmt.Sprintf("%s(1)%s", name, ext)
	}

	return result
}

//	去除文件的重名规则，name(2).txt ＝ name.txt
//	@name
//	@return
func FileCleanRenameRule(name string) string {
	result := ""
	// strings.SplitN(s, sep, n)
	ext := filepath.Ext(name)
	if 0 != len(ext) {
		name = name[:len(name)-len(ext)]
	}
	rules := _rexRenameRule.FindAllStringSubmatchIndex(name, -1)

	if 0 != len(rules) && rules[len(rules)-1][1] == len(name) {
		lastRule := rules[len(rules)-1]
		result = name[:lastRule[0]] + ext
	} else {
		result = name + ext
	}
	return result
}

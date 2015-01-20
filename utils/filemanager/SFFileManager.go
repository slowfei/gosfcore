//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-08-24
//  Update on 2015-01-20
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//	对于文件和目录的工具
//
package SFFileManager

import (
	"errors"
	// "fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
)

var (
	//	存储当前执行文件的路径，这样可以避免每次获取不必要的计算。
	_currentExecFilePath string

	//	文件重命名规则的正则
	_rexRenameRule = regexp.MustCompile("\\(\\d+\\)")

	ErrPathConflict = errors.New("Path conflict.")
)

/**
 *	output file	auto mkdir all
 *
 *	@param `path` file path
 *	@param `data` out data
 *	@return
 */
func WirteFilepath(path string, data []byte) error {

	// create directory
	dirPath := filepath.Dir(path)

	exists, isDir, _ := Exists(path)

	if !exists {
		mkErr := os.MkdirAll(dirPath, os.ModePerm)
		if nil != mkErr {
			return mkErr
		}
	} else if isDir {
		return ErrPathConflict
	}

	//	wiete file
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}

/**
 *	获取当前命令行目录
 *
 *	@return
 */
func GetCmdDir() string {
	file, err := os.Getwd()

	if nil != err {
		panic(err)
	}

	return file
}

/**
 *	获取当前执行文件的操作目录
 *
 *	@return
 */
func GetExecDir() string {
	return filepath.Dir(GetExecFilePath())
}

/**
 *	获取 GOPATH 的多项路径
 *
 *	@return
 */
func GetGOPATHDirs() []string {
	var result []string = nil
	gopath := os.Getenv("GOPATH")

	if 0 != len(gopath) {
		result = filepath.SplitList(gopath)
	}

	return result
}

/**
 *	获取当前执行文件的路径
 *
 *	@return
 */
func GetExecFilePath() string {

	if 0 != len(_currentExecFilePath) {
		return _currentExecFilePath
	}

	execFile, err := exec.LookPath(os.Args[0])
	if nil != err {
		panic(err)
	}

	execFilePath, err2 := filepath.Abs(execFile)
	if nil != err2 {
		panic(err2)
	}

	_currentExecFilePath = execFilePath

	return execFilePath
}

/**
 *	获取当前执行文件的名称
 *
 *	@return
 */
func GetExecFileName() string {
	return filepath.Base(GetExecFilePath())
}

/**
 *	判断路径是否存在文件或目录
 *
 *	@path	操作路径
 *	@reutrn	isExists 是否存在，存在true
 *	@reutrn	isDir	 是否是目录
 *	@return err		 错误信息
 */
func Exists(path string) (isExists bool, isDir bool, err error) {
	isExists = false
	isDir = false
	err = nil

	fileInfo, e := os.Stat(path)

	if e == nil {
		isExists = true
		isDir = fileInfo.IsDir()
	} else {
		err = e
	}

	return
}

/**
 *	文件名重复命名的规则，例如相同的名称重复命名 name.txt、name(2).txt、name(3).txt...
 *
 *	@name		file name
 *	@return		"name(2).txt" => "name(3).txt"
 */
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

/**
 *	去除文件的重名规则
 *
 *	@name		file name
 *	@return		name(2).txt => name.txt
 */
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

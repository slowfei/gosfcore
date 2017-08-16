//  Copyright 2016 slowfei And The Contributors All rights reserved.
//
//  Software Source Code License Agreement (BSD License)
//
//  Create on 2016-11-07
//  Update on 2016-11-07
//  Email  slowfei@nnyxing.com
//  Home   http://www.slowfei.com

/***0-SFTextFormat介绍与功能
当前组件主要针对文本格式的操作，例如读取、写入操作等

####功能：

1. 读取文本键=值格式
> 针对特定文件格式进行读取，返回map类型，可根据选择进行读取数据的缓存。
> 键值目前支持的分隔符号“=”

*/

// 针对文件格式处理组件
package SFTextFormat

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

/**
 *	read file to key value format
 *	key value delimiter is "="
 *
 *	e.g.: (file.txt)
 *	key1=value1
 *	key2=value2
 *
 *	@param `fpath` file path
 *	@param `cache` cache key value
 *	@return map key value
 *
 */
func RKeyValue(fpath string, cache bool) map[string]string {
	//TODO 缓存还未编写

	var kvmap map[string]string

	file, err := os.Open(fpath)
	if nil == err {
		defer func() {
			if err := file.Close(); err != nil {

			}
		}()

		kvmap = make(map[string]string)
		br := bufio.NewReader(file)

		for {

			lineBytes, isPrefix, err := br.ReadLine()
			if nil != err || io.EOF == err || isPrefix {
				break
			}

			// comments sign check
			commentsSignIndex := bytes.IndexByte(lineBytes, '#')
			if 0 == commentsSignIndex {
				continue
			}

			signIndex := bytes.IndexByte(lineBytes, '=')
			if 0 < signIndex {
				key := string(lineBytes[:signIndex])
				value := string(lineBytes[signIndex+1:])

				if 0 != len(key) {
					kvmap[key] = value
				}
			}

		}
	}

	return kvmap
}

/**
 *	read file to key value format
 *	key value delimiter is "="
 *
 *	e.g.: (file.txt)
 *	key1=value1
 *	key2=value2
 *
 *	@param `fpath` file path
 *	@param `cache` cache read file
 *	@param `block` func(key,value string) bool
 *
 */
func RKeyValueBlock(fpath string, cache bool, block func(key, value string) bool) {
	file, err := os.Open(fpath)
	if nil == err {
		defer func() {
			if err := file.Close(); err != nil {

			}
		}()

		br := bufio.NewReader(file)
		for {

			lineBytes, isPrefix, err := br.ReadLine()
			if nil != err || io.EOF == err || isPrefix {
				break
			}

			// comments sign check
			commentsSignIndex := bytes.IndexByte(lineBytes, '#')
			if 0 == commentsSignIndex {
				continue
			}

			signIndex := bytes.IndexByte(lineBytes, '=')
			if 0 < signIndex {
				key := string(lineBytes[:signIndex])
				value := string(lineBytes[signIndex+1:])

				if !block(key, value) {
					break
				}
			}

		}
	}
}

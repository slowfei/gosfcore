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
>

*/

// 针对文件格式处理组件
package SFTextFormat

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func RKeyValue(fpath string, cache bool) map[string]string {
	return nil
}

func RKeyValuBlock(fpath string, block func(key, value string) bool) {
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

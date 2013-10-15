//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)

//	rand utils
//
//	email	slowfei@foxmail.com
//	createTime 	2013-9-26
//	updateTime	2013-10-3
//
package SFRandUtil

import (
	"crypto/rand"
	"io"
	"math"
	mathrand "math/rand"
	"time"
)

var (
	_reader = rand.Reader
)

//	范围随机数 int
func RandBetweenInt(s, n int64) int64 {
	if s == 0 && n == 0 {
		return 0
	}

	if s > n {
		s, n = n, s
	}
	mathrand.Seed(time.Now().UnixNano())
	ri := mathrand.Int63n(n)
	return ri%(n-s+1) + s
}

//	范围随机数 float64
//	只随机精度的值,随机精度范围[0.x... -- 0.x...]
func RandBetweenFloat(s, n float64) float64 {
	if s > n {
		s, n = n, s
	}
	_, sFrac := math.Modf(s)
	_, nFrac := math.Modf(n)

	mathrand.Seed(time.Now().UnixNano())

	rf := float64(mathrand.Int63()) / (1 << 63)
	rf = sFrac + rf*(nFrac-sFrac)

	return rf
}

//	基于rand.Reader直接读取的随机数
func RandBits(b []byte) {
	if _, err := io.ReadFull(_reader, b); nil != err {
		panic(err.Error())
	}
}

//	基于rand.Reader生成随机数字符串
//	指定字符串：ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
func RandString(n int) string {
	const b62 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	if n < 1 {
		n = 16 // 默认长度
	}
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); nil != err {
		panic(err.Error())
	}

	for i, x := range bytes {
		bytes[i] = b62[x&0x3D] // 0x3D = 61 = len(b62) - 1
	}
	return string(bytes)
}

//	基于rand.Reader使用特定字符串生成随机数字符串
//	string len 最高支持255
func RandByString(n int, s string) string {
	if n < 1 {
		n = 16 // 默认长度
	}
	if "" == s {
		return RandString(n)
	}

	fb := byte(len(s) - 1)

	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); nil != err {
		panic(err.Error())
	}

	for i, x := range bytes {
		bytes[i] = s[x&fb]
	}

	return string(bytes)

}

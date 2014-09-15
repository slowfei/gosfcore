//  Copyright 2013 slowfei And The Contributors All rights reserved.
//
//  Software Source Code License Agreement (BSD License)
//
//  Create on 2013-9-2
//  Update on 2014-06-26
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//  string 辅助工具
//
package SFStringsUtil

import (
	// "fmt"
	"math"
)

var (

	// xvalues returns the value of a byte as a hexadecimal digit or 255.
	_xvalues = []byte{
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
		255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	}

	//  十进制的十六进制字符数组
	_decHexCharDigits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "100", "101", "102", "103", "104", "105", "106", "107", "108", "109", "110", "111", "112", "113", "114", "115", "116", "117", "118", "119", "120", "121", "122", "123", "124", "125", "126", "127", "128", "129", "130", "131", "132", "133", "134", "135", "136", "137", "138", "139", "140", "141", "142", "143", "144", "145", "146", "147", "148", "149", "150", "151", "152", "153", "154", "155", "156", "157", "158", "159", "160", "161", "162", "163", "164", "165", "166", "167", "168", "169", "170", "171", "172", "173", "174", "175", "176", "177", "178", "179", "180", "181", "182", "183", "184", "185", "186", "187", "188", "189", "190", "191", "192", "193", "194", "195", "196", "197", "198", "199", "200", "201", "202", "203", "204", "205", "206", "207", "208", "209", "210", "211", "212", "213", "214", "215", "216", "217", "218", "219", "220", "221", "222", "223", "224", "225", "226", "227", "228", "229", "230", "231", "232", "233", "234", "235", "236", "237", "238", "239", "240", "241", "242", "243", "244", "245", "246", "247", "248", "249", "250", "251", "252", "253", "254", "255"}

	//  int64 max pow 10 multiple
	//  math.MaxInt64 = 9223372036854775807
	_int64Pow10s = []int64{
		1,
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
		10000000,
		100000000,
		1000000000,
		10000000000,
		100000000000,
		1000000000000,
		10000000000000,
		100000000000000,
		1000000000000000,
		10000000000000000,
		100000000000000000,
		1000000000000000000,
	}
)

/**
 *  将字符串前两位字符转换16进制的byte.
 *  如果前两位字符超出16进制的大小则按 FF返回
 *  e.g.:   "F12" = F1,true;    "FG" = FF, false;
 *
 *  @param `xstr`
 *  @return byte
 *  @return bool
 */
func Xtob(xstr string) (byte, bool) {
	b1 := _xvalues[xstr[0]]
	b2 := _xvalues[xstr[1]]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}

/**
 *  字符串排序：由长到短进行排序
 */
type SortLengToShort []string

//  排序sort.Interface实现规则
func (ls SortLengToShort) Len() int           { return len(ls) }
func (ls SortLengToShort) Less(i, j int) bool { return len(ls[i]) < len(ls[j]) }
func (ls SortLengToShort) Swap(i, j int)      { ls[i], ls[j] = ls[j], ls[i] }

/**
 *  字符串反转
 *  e.g.: abcd = bcda
 *
 *  @param `src`
 *  @return
 */
func Reverse(src string) string {
	runes := []rune(src)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/**
 *  十六进制转换十进制的字符
 *
 *  @param `hex`
 *  @return
 */
func DecHexString(hex byte) string {
	return _decHexCharDigits[hex]
}

/**
 *  to lower
 *  由于strings.ToLower()有性能问题，所以自定义出来
 *
 *  @param `s`
 *  @return lower string
 */
func ToLower(s string) string {
	b := make([]byte, len(s))
	for i := range b {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return string(b)
}

/**
 *  to upper
 *
 *  @param `s`
 *  @return upper string
 */
func ToUpper(s string) string {
	b := make([]byte, len(s))
	for i := range b {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			c -= 'a' - 'A'
		}
		b[i] = c
	}
	return string(b)
}

/**
 *  string to int8
 *
 *  @param `str`
 *  @return
 */
func ToInt8(str string) int8 {
	return int8(toInt(str, math.MaxInt8, math.MinInt8))
}

/**
 *  string to int16
 *
 *  @param `str`
 *  @return
 */
func ToInt16(str string) int16 {
	return int16(toInt(str, math.MaxInt16, math.MinInt16))
}

/**
 *  string to int32
 *
 *  @param `str`
 *  @return
 */
func ToInt32(str string) int32 {
	return int32(toInt(str, math.MaxInt32, math.MinInt32))
}

/**
 *  string to int
 *
 *  @param `str`
 *  @return
 */
func ToInt(str string) int {
	return int(toInt(str, math.MaxInt64, math.MinInt64))
}

/**
 *  string to int64
 *
 *  @param `str`
 *  @return max number math.MaxInt64, min number math.MinInt64
 */
func ToInt64(str string) int64 {
	return toInt(str, math.MaxInt64, math.MinInt64)
}

/**
 *  string to int
 *
 *  @param `str`
 *  @param `maxInt`
 *  @param `minInt`
 *  @return string with between maxInt and minInt
 */
func toInt(str string, maxInt int64, minInt int64) int64 {
	var resultInt int64 = 0
	strCount := len(str)
	forCount := 0 // 控制累加结果循环的次数

	if 0 == strCount {
		return resultInt
	}
	if maxInt < minInt {
		return minInt
	}

	//  positive negative check
	neg := false

	//  string number digit check
	strDigit := -1
	for i := 0; i < strCount; i++ {
		c := str[i]
		if '0' <= c && '9' >= c {
			strDigit++
			forCount++
		} else {
			if -1 == strDigit {
				switch c {
				case '-':
					neg = true
					forCount++
				case '+', ' ':
					forCount++
				default:
					return resultInt
				}
			} else {
				break
			}
		}
	}

	//  max and min range check
	if strDigit > len(_int64Pow10s)-1 {
		if neg {
			resultInt = minInt
		} else {
			resultInt = maxInt
		}
		return resultInt
	}

	// math.MaxInt64 =  9223372036854775807
	// math.MinInt64 = -9223372036854775808

	//  str = 654321
	// 累加结果：
	// 600000+
	//  50000+
	//   4000+
	//    300+
	//     20+
	//      1=654321
	for i := 0; i < forCount; i++ {
		c := str[i]
		if '0' <= c && '9' >= c {

			//  位数值计算6位数 = 600000
			digitNum := int64(c-'0') * _int64Pow10s[strDigit]
			strDigit--

			if !neg {
				// (suppose logic processing)
				//[0] 9000000000000000000
				//[1]  300000000000000000

				//[0] 9223372036854775807
				//[1]  223372036854775807

				differ := maxInt - resultInt
				if differ < digitNum {
					//  操作结果范围
					return maxInt
				}
				resultInt += digitNum

			} else {
				//   此计算方式会比计算正数慢，关键在于(-digitNum)
				differ := minInt - resultInt

				if differ > -digitNum {
					return minInt
				}
				resultInt -= digitNum

			}
		}
	}

	return resultInt
}

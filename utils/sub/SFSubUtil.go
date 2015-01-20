//  The MIT License (MIT) - http://opensource.org/licenses/MIT
//
//  Copyright (c) 2014 slowfei
//
//  Create on 2014-12-16
//  Update on 2015-01-20
//  Email  slowfei(#)foxmail.com
//  Home   http://www.slowfei.com

//
//  interception util set
//  针对源数据进行特定的信息截取
//
package SFSubUtil

import (
	"bytes"
	// "fmt"
)

/**
  数据源遍历时需要的处理流程标识定义。
  1. 寻找 start 符号
  2. 判断转义符
  3. 寻找 end 符号
  4. 进入排除符号的处理
*/
type findPcs int

const (
	findPcsStart findPcs = iota // 寻找 start 符号
	findPcsEnd                  // 寻找 end 符号
)

/**
 *  interception nesting subset data
 *  (temp(temp2)()).... the irregular subset
 */
type SubNest struct {
	start []byte
	end   []byte
}

/**
 *  new nest struct
 *
 *  @param `start`
 *  @param `end`
 */
func NewSubNest(start, end []byte) *SubNest {
	return &SubNest{start, end}
}

/**
 *  private bytes all index
 *
 *  @param `number` find number -1 is all
 *	@param `outBetweens` rule out between index, [0]=10 [1]=20, 10-20 position will rule out scanning
 */
func (nest *SubNest) bytesToIndex(src []byte, number int, outBetweens [][]int) [][]int {
	srcLen := len(src)
	if 0 == srcLen || 0 == number {
		return nil
	}
	if 0 > number {
		number = 0
	}

	start := nest.start
	startLen := len(start)
	end := nest.end
	endLen := len(end)

	if 0 == startLen || 0 == endLen {
		return nil
	}

	result := make([][]int, 0, number)
	startIndex := -1 // 起始坐标，封装用
	endIndex := -1   // 结尾坐标
	process := findPcsStart

	tempSrc := src[:srcLen]
	tempSrcLen := len(tempSrc)
	tempStartI := 0   // 开始坐标临时存储
	tempEndI := 0     // 结尾坐标
	balanceCount := 0 // 平衡计数

	// " 1 2 3 { 5 6 7 } 9 "
	// " 1 { 3 { 5 } 7 } 9 "
	// FOR 遍历src
	//
	//   寻找第一个起始符号 then;
	//
	//   j = 寻找结束符号下标，i = 寻找开始符号下标；
	//   IF 0 > j
	//      找不到结束符号，结束此次循环
	//   IF j < i
	//      balanceCount--
	//      表示第一个起始符号与结束符对等，可以结束第一次寻找过程
	//   ELSE
	//      出现嵌套，需要平衡开始符与结束符，balanceCount++
	//      记录结束符下标，下次循环从该下标寻找
	//   then;
	//
	// FOR END

	for i := 0; i < tempSrcLen; i++ {
		// 寻找开始坐标点

		if findPcsStart == process {

			of := bytes.Index(tempSrc[i:], start)
			if 0 > of {
				break
			}

			i += of

			if !isEscape(tempSrc, i) && !isRuleOutIndex(outBetweens, i) {
				startIndex = i
				tempStartI = startIndex + 1
				tempEndI = tempStartI
				process = findPcsEnd
				balanceCount++
			}

		} else if findPcsEnd == process {

			ofEnd := bytes.Index(tempSrc[tempEndI:], end)

			if 0 > ofEnd {
				break
			}
			tempEndI += ofEnd

			if !isEscape(tempSrc, tempEndI) && !isRuleOutIndex(outBetweens, tempEndI) {
				balanceCount--
				endIndex = tempEndI

				ofStart := bytes.Index(tempSrc[tempStartI:], start)
				tempStartI += ofStart

				// n(循环周期)_n(balanceCount)
				// {1_1 {2_1 {3_1 }2_0 }3_0 }4_0
				// {1_1 {2_1 }2_0 {3_1 }3_0 {4_1 }4_0 }
				if 0 <= ofStart && tempEndI > tempStartI &&
					!isEscape(tempSrc, tempStartI) && !isRuleOutIndex(outBetweens, tempStartI) {
					balanceCount++
				}

				tempStartI++
			}

			tempEndI++
		}

		if 0 == balanceCount && -1 != startIndex && -1 != endIndex {
			result = append(result, []int{startIndex, endIndex + 1})

			if len(result) == number {
				break
			}
			i = endIndex
			tempStartI = endIndex
			tempEndI = tempStartI
			startIndex = -1
			endIndex = -1
			process = findPcsStart
		}
	}

	return result
}

/**
 *  to source data subset target a index
 *
 *  @param `src` source data
 *  @param `outBetweens` rule out between index, [0]=10 [1]=20, 10-20 position will rule out scanning
 *	@return first find result index []int{start int, end int}
 */
func (nest *SubNest) BytesToIndex(src []byte, outBetweens [][]int) []int {
	result := nest.bytesToIndex(src, 1, outBetweens)

	if 0 != len(result) {
		return result[0]
	}

	return nil
}

/**
 *  to source data target all index
 *
 *  @param `src` source data
 *  @param `outBetweens` rule out between index, [0]=10 [1]=20, 10-20 position will rule out scanning
 *	@return all find result indexs [][]int{ []int{start index, end int}... }
 */
func (nest *SubNest) BytesToAllIndex(src []byte, outBetweens [][]int) [][]int {
	return nest.bytesToIndex(src, -1, outBetweens)
}

/**
 *  by index validate escape
 *
 *  @param src
 *  @param index index byte
 *	@return true is escape
 */
func isEscape(src []byte, index int) bool {
	result := false
	b := src[index]

	switch b {
	case '\\':
		fallthrough
	case '\'':
		fallthrough
	case '"':
		// 转义符需要匹配的基本包括 \ " ' 这3个字符。
		// 如果开始符号为这3个符号时就需要判断是否为转义符
		// 有可能会出现多个转义符号 string = " var a = \\\\\" 12\" "
		// 这时需要计算"\"的数量，双数则表示当前开始符号不是转义符，单数则是。
		// [a] [b] [\] ['] [e] [f]
		// [0] [1] [2] [3] [4] [5]
		// i = 3
		escapesNum := 0
		for j := index - 1; j >= 0; j-- {
			srcByte := src[j]
			if '\\' == srcByte {
				escapesNum++
			} else {
				break
			}
		}
		if 0 != escapesNum && 1 == escapesNum%2 {
			result = true
		}

	default:
	}
	return result
}

/**
 *  index 是否在排除的范围内
 *
 *	@param `outBetweens` 排除坐标的范围值
 *	@param `index` 检测是否排除的下标
 *	@return true 表示在排除范围内
 */
func isRuleOutIndex(outBetweens [][]int, index int) bool {
	result := false

	for i := 0; i < len(outBetweens); i++ {
		indexs := outBetweens[i]
		if 2 == len(indexs) {
			s := indexs[0]
			e := indexs[1]
			if index > s && index < e {
				result = true
				break
			}
		}
	}

	return result
}

//  The MIT License (MIT) - http://opensource.org/licenses/MIT
//
//  Copyright (c) 2014 slowfei
//
//  Create on 2014-12-16
//  Update on 2015-06-09
//  Email  slowfei(#)foxmail.com
//  Home   http://www.slowfei.com

//
//  interception util set
//  针对源数据进行特定的信息截取
//
package SFSubUtil

import (
	"bytes"
	"fmt"
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
	start   []byte
	end     []byte
	notNest bool // default false
}

/**
 *  new nest struct
 *
 *  @param `start`
 *  @param `end`
 */
func NewSubNest(start, end []byte) *SubNest {
	return &SubNest{start, end, false}
}

/**
 *  new not sub nest struct
 *
 *	find "(" content ")"
 *  find content "( 1 2 (3 4 5) 6 7 8 )" or "// temp http://www.slowfei.com [\n]"
 *
 *	find result "( 1 2 (3 4 5)" or "// temp http://www.slowfei.com [\n]"
 *
 */
func NewSubNotNest(start, end []byte) *SubNest {
	return &SubNest{start, end, true}
}

/**
 *  private bytes all index
 *
 *	@param `startFindIndex` src start find location index, will be sum to the results
 *  @param `number` 		find number -1 is all
 *	@param `outBetweens` 	rule out between index, [0]=10 [1]=20, 10-20 position will rule out scanning
 */
func (nest *SubNest) bytesToIndex(startFindIndex int, src []byte, number int, outBetweens [][]int) [][]int {
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
	notNest := nest.notNest

	if 0 == startLen || 0 == endLen {
		return nil
	}

	result := make([][]int, 0, number)
	startIndex := -1 // 起始坐标，封装用
	endIndex := -1   // 结尾坐标
	process := findPcsStart

	tempSrc := src[startFindIndex:srcLen]
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

		if findPcsEnd == process {

			ofEnd := bytes.Index(tempSrc[tempEndI:], end)

			if 0 > ofEnd {
				break
			}
			tempEndI += ofEnd

			//	查询到的结束符号需要判断是否是被过滤的，如果是被过滤的则跳过，继续下一个结束符的寻找。
			if !isEscape(tempSrc, tempEndI) && !isRuleOutIndex(outBetweens, tempEndI+startFindIndex, endLen) {
				balanceCount--
				endIndex = tempEndI

				// 看是否需要查询嵌套
				if !notNest {
					for {
						//	由于考虑到如果查询到的字符属于过滤的，则需要再次寻找，否则过滤的字符会累加在内（balanceCount++）。
						// n(循环周期)_n(balanceCount)
						// {1_1 {2_1 {3_1 }2_0 }3_0 }4_0
						// {1_1 {2_1 }2_0 {3_1 }3_0 {4_1 }4_0 }

						ofStart := bytes.Index(tempSrc[tempStartI:], start)
						tempStartI += ofStart

						if 0 > ofStart || tempEndI <= tempStartI {
							break
						}

						if !isEscape(tempSrc, tempStartI) && !isRuleOutIndex(outBetweens, tempStartI+startFindIndex, startLen) {
							balanceCount++
							break
						}

						//	递加一个index，继续寻找
						tempStartI++
					}
				}

				//	由于在for中条件内break跳出时没有递加index,所以这里进行一次递加。
				tempStartI++

			}

			tempEndI++

		} else if findPcsStart == process {

			of := bytes.Index(tempSrc[i:], start)
			if 0 > of {
				break
			}

			i += of

			if !isEscape(tempSrc, i) && !isRuleOutIndex(outBetweens, i+startFindIndex, startLen) {
				startIndex = i
				tempStartI = startIndex + 1
				tempEndI = tempStartI

				process = findPcsEnd
				balanceCount++
			}
		}

		if 0 == balanceCount && -1 != startIndex && -1 != endIndex {
			//	结尾index需要加上结束符的长度
			newEndIndex := endIndex + endLen
			result = append(result, []int{startIndex + startFindIndex, newEndIndex + startFindIndex})

			if len(result) == number {
				break
			}
			i = endIndex
			tempStartI = newEndIndex
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
 *	@param `startIndex` src start find location index, will be sum to the results
 *  @param `src` source data
 *  @param `outBetweens` rule out between index, [0]=10 [1]=20, 10-20 position will rule out scanning
 *	@return first find result index []int{start int, end int}
 */
func (nest *SubNest) BytesToIndex(startIndex int, src []byte, outBetweens [][]int) []int {
	result := nest.bytesToIndex(startIndex, src, 1, outBetweens)

	if 0 != len(result) {
		return result[0]
	}

	return nil
}

/**
 *  to source data target all index
 *
 *	@param `startIndex` src start find location index, will be sum to the results
 *  @param `src` source data
 *  @param `outBetweens` rule out between index, [0]=10 [1]=20, 10-20 position will rule out scanning
 *	@return all find result indexs [][]int{ []int{start index, end int}... }
 */
func (nest *SubNest) BytesToAllIndex(startIndex int, src []byte, outBetweens [][]int) [][]int {
	return nest.bytesToIndex(startIndex, src, -1, outBetweens)
}

/**
 *	get between rule out points
 *
 *	@param `src`   data source
 *	@param `nests` SubNest object, sequence has a certain efect
 *	@return data source points [0] is start point [1] is end point
 */
func GetOutBetweens(src []byte, nests ...*SubNest) [][]int {

	nestsLen := len(nests)
	srcLen := len(src)
	if 0 == nestsLen || 0 == srcLen {
		return nil
	}

	result := make([][]int, 0, 0)
	filterNestIndex := make([]int, nestsLen, nestsLen)

	for srcIndex := 0; srcIndex < srcLen; srcIndex++ {

		// 获取首位的SubNest对象
		tempIndex := -1
		tempFirstPoint := srcLen
		for i := 0; i < nestsLen; i++ {
			filterNestIndex[i] = -1
			nest := nests[i]
			ti := bytes.Index(src[srcIndex:], nest.start)
			if -1 != ti && tempFirstPoint > ti {
				tempFirstPoint = ti
				tempIndex = i
			}
		}
		if -1 == tempIndex {
			return result
		}
		tempFirstPoint += srcIndex
		firstNest := nests[tempIndex]
		filterNestIndex[0] = tempIndex

		firstIndex := firstNest.BytesToIndex(srcIndex, src, nil)

		if 2 != len(firstIndex) {
			// fmt.Println("temp_0_0:", firstIndex, result, tempFirstPoint, string(firstNest.start), srcIndex)
			// 如果走到这部表示 开始符号和结束符号 不对等，这能出现的问题就是开始符号的位置实在过滤的范围内。
			// 所以需要找到下一个开始符号，效验是否是被过滤的，如果不是则认为是规则的格式错误。
			var outBetweens [][]int = nil
			outBetweensLen := 0
			startLen := len(firstNest.start)
			newSrcIndex := tempFirstPoint + startLen
			tempSrcIndex := tempFirstPoint

			checkIndex := bytes.Index(src[newSrcIndex:], firstNest.start)
			if -1 == checkIndex {
				return result
			}
			checkIndex += newSrcIndex

			for {
				// fmt.Println("TODO:", checkIndex, filterNestIndex)
				checkBetweens := obCheckFilter(src, tempFirstPoint, checkIndex, nests, filterNestIndex)
				outLen := len(checkBetweens)
				// fmt.Println("TODO2:", checkBetweens)

				if 0 != outLen {
					// 记录效验获取的过滤参数最后一位的结尾数，避免是相同的过滤参数出现无限循环
					outEndIndex := checkBetweens[outLen-1][1]
					if 0 != outBetweensLen && outBetweens[outBetweensLen-1][1] >= outEndIndex {
						break
					}

					outBetweens = append(outBetweens, checkBetweens...)
					outBetweensLen = len(outBetweens)

					firstIndex = firstNest.BytesToIndex(tempSrcIndex, src, outBetweens)
					firstIndexLen := len(firstIndex)

					// 如果效验的index小于过滤的内容则需要再次验证寻找
					if 0 == firstIndexLen && outEndIndex < checkIndex {
						tempFirstPoint = outEndIndex
						continue
					}
				}

				if 0 == len(firstIndex) {
					// 如果出现多个嵌套符号，则需要继续验证下一个判断符，直到找到为止。如果格式错误会造成很大的资源浪费。
					sliceIndex := checkIndex + startLen
					nextCheckIndex := bytes.Index(src[sliceIndex:], firstNest.start)
					if -1 != nextCheckIndex {
						checkIndex = nextCheckIndex + sliceIndex
						tempFirstPoint = sliceIndex
						continue
					}
				}

				break
			}

		}

		if 2 == len(firstIndex) {
			firstStart := firstIndex[0]
			firestEnd := firstIndex[1]
			checkStart := firstStart
			tempOutEnd := -1
			var outBetweens [][]int = nil

			// 如果开始符号和结尾符号是相同的就直接添加，主要是相同的，嵌套的时候会判断转义符，会很少出现错误的验证。
			isSame := bytes.Equal(firstNest.start, firstNest.end)
			if isSame {
				result = append(result, []int{firstStart, firestEnd})
			}

			tempi := 0 // 防止无限循环
			for !isSame {
				tempi++
				if tempi >= 500 {
					fmt.Println("Accident: infinite for error.....")
					break
				}

				// fmt.Println("temp_0:", firstStart, firestEnd)
				checkBetweens := obCheckFilter(src, checkStart, firestEnd, nests, filterNestIndex)
				outLen := len(checkBetweens)

				// 如果存在过滤项则需要再次寻找
				if 0 != outLen {
					outBetweens = append(outBetweens, checkBetweens...)
					// 记录效验获取的过滤参数最后一位的结尾数，避免是相同的过滤参数出现无限循环
					outEnd := checkBetweens[outLen-1][1]

					// fmt.Println("temp_1:", outBetweens, firstStart, firestEnd)
					newFirstIndex := firstNest.BytesToIndex(srcIndex, src, outBetweens)
					if 2 != len(newFirstIndex) {
						if outEnd == tempOutEnd {
							break
						}
						// 此时就有可能是被下一行需要过滤的开头(firstNest.start)形成不对等的判断，导致查询异常，所以再次需要寻找开始符号。
						ti := bytes.Index(src[outEnd:], firstNest.start)
						if -1 == ti {
							break
						}

						firestEnd = ti + outEnd
						checkStart = outEnd
					} else {
						firestEnd = newFirstIndex[1]
						checkStart = outEnd
					}

					tempOutEnd = outEnd

					// fmt.Println("temp_2:", outBetweens, firstStart, firestEnd)
				} else {
					result = append(result, []int{firstStart, firestEnd})
					break
				}
			} // ene for {

			//
			srcIndex = firestEnd - len(firstNest.end)

		} else {
			break
		}
	}

	return result
}

/**
 *	效验过滤范围结果，如果在otherNest所查询的范围内则返回该下标
 *
 *	@param `src`
 *	@param `srcIndex`
 *	@param `checkIndex` 需要效验的源数据下标
 *	@param `otherNests`
 *	@param `filterIndex`
 *	@param `exceedAdd` 超出效验范围时，是否添加最后过滤项
 *	@return `[][]int` OutBetweens 子集的所有过滤项
 */
func obCheckFilter(src []byte, srcIndex, checkIndex int, otherNests []*SubNest, filterIndex []int) [][]int {
	var result [][]int = make([][]int, 0, 0)

	if len(src) <= srcIndex {
		return result
	}

	tempSrc := src[srcIndex:]
	tempIndex := -1
	// tempOutBetweens := make([][]int, 1, 1)
	tempFirstPoint := len(tempSrc)
	filterLen := len(filterIndex)

	for i := 0; i < len(otherNests); i++ {
		isFilter := false
		for j := 0; j < filterLen; j++ {
			fi := filterIndex[j]
			if i == fi {
				isFilter = true
				break
			}
		}

		if !isFilter {
			otNest := otherNests[i]
			ti := bytes.Index(tempSrc, otNest.start)
			if -1 != ti && tempFirstPoint > ti {
				tempFirstPoint = ti
				tempIndex = i
			}
		}
	}
	// fmt.Println("-0.0.0", tempIndex, srcIndex, filterIndex)
	if -1 == tempIndex {
		return nil
	}

	tempFirstPoint += srcIndex
	filterIndex = append(filterIndex, tempIndex)
	firstNest := otherNests[tempIndex]
	firstIndex := firstNest.BytesToIndex(tempFirstPoint, src, nil)

	if 2 != len(firstIndex) {
		// fmt.Println("-0.0:", result, firstIndex, checkIndex)

		newSrcIndex := tempFirstPoint + len(firstNest.start)

		ti := bytes.Index(src[newSrcIndex:], firstNest.start)
		if -1 == ti {
			return nil
		}
		ti += newSrcIndex

		outBetweens := obCheckFilter(src, tempFirstPoint, ti, otherNests, filterIndex)

		if 0 != len(outBetweens) {
			result = append(result, outBetweens...)
			firstIndex = firstNest.BytesToIndex(tempFirstPoint, src, result)
		}
		// fmt.Println("-0.1:", result, firstIndex, checkIndex)
	}
	// fmt.Println("-0:", result, firstIndex, checkIndex)
	if 2 == len(firstIndex) {

		firstStart := firstIndex[0]
		firestEnd := firstIndex[1]
		// fmt.Println("0: ", firstIndex)

		tempi := 0
		for {
			tempi++
			if tempi >= 500 {
				fmt.Println("Accident: check filter infinite for error.....")
				break
			}

			if checkIndex < firstStart {
				// fmt.Println("0_1: ", firstIndex, result)
				return result
			}

			// fmt.Println("1: ", firstIndex, firstStart, firestEnd, filterIndex)
			// 效验子集的结尾下标是否又是被过滤的，所以进行递归查询
			outBetweens := obCheckFilter(src, firstStart, firestEnd, otherNests, filterIndex)
			outBetweensLen := len(outBetweens)
			// 由于更改了该方法的放回参数，所以下面可能需要重构。
			// fmt.Println("2: ", firstIndex, outBetweens, len(outBetweens), result)
			// 效验无过滤时，则表示当前firstNest查找的结尾标识正确的，所以此时就需要判断是否效验下标是否在其范围内。
			if 0 != outBetweensLen {

				// 比较结尾的结果是否相同，如果相同则跳过，否则会出现无限循环
				resultLen := len(result)
				if 0 != resultLen {
					resultEnd := result[resultLen-1]
					outBetweenEnd := outBetweens[outBetweensLen-1]
					if resultEnd[0] == outBetweenEnd[0] && resultEnd[1] == outBetweenEnd[1] {
						break
					}
				}

				// 被过滤继续查找正确的子集
				result = append(result, outBetweens...)
				// fmt.Println("2_1: ", firstIndex, result)
				newFirstIndex := firstNest.BytesToIndex(firstStart, src, result)
				if 2 != len(newFirstIndex) {
					break
				}
				firestEnd = newFirstIndex[1]
			} else {

				// 如果效验下标在子集的范围内，则返回子集的坐标进行过滤

				//	如果效验的下标小于查询坐标起始位置的下标，表示效验的下标在当前子集的前面，则为正确的效验下标。
				//	此时返回nil则无需进行过滤
				if checkIndex < firstStart {
					// fmt.Println("3: ", firstIndex, result)
					break
				} else if checkIndex > firstStart && checkIndex < firestEnd {
					result = append(result, []int{firstStart, firestEnd}) // TODO 考虑这句是否是放置这里。
					// fmt.Println("4: ", firstIndex, result)
					break
				} else {

					result = append(result, []int{firstStart, firestEnd})
					// fmt.Println("5: ", firstIndex, result)
					// 继续寻找过滤其他子集
					/*
						{
							// 这里的注释标识第一个过滤的子集这里找不到大括号(需要记录)
							// 这里是第二个子集也找不到(需要记录，否则在使用BytesToIndex过滤时就只过滤下面一行的坐标)
							"} 大括号在这里，需要讲此双引号的坐标过滤，而BytesToIndex可能会找到此括号进行配对"

						}(这个是正确的大括号)
					*/

					// 记录此次的子集，因为这也是个正确的过滤集
					newFirstIndex := firstNest.BytesToIndex(firestEnd, src, nil)
					if 2 != len(newFirstIndex) {
						break
					}

					firstStart = newFirstIndex[0]
					firestEnd = newFirstIndex[1]
					// fmt.Println("5_1: ", newFirstIndex, result)
				}
			}
		}

	}

	return result
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
 *	@param `index` 		 检测是否排除的下标
 *	@param `symbolLen` 	 start or end symbol length. 避免相同的坐标排除
 *	@return true 表示在排除范围内
 */
func isRuleOutIndex(outBetweens [][]int, index int, symbolLen int) bool {
	result := false

	for i := 0; i < len(outBetweens); i++ {
		indexs := outBetweens[i]
		if 2 == len(indexs) {
			s := indexs[0]
			e := indexs[1]
			if index > s && index < e-symbolLen {
				result = true
				break
			}
		}
	}

	return result
}

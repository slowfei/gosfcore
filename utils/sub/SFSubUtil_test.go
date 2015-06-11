package SFSubUtil

import (
	"testing"
)

func Test_bytesToIndex(t *testing.T) {
	testSrc := []byte(`
{if{temp}{late}}

{456}

{789}

{ "if{123}}}}{{}}{}}}}{" }

"if{if456}"

"123\\"+"000"

"234"

"789"

`)

	var outIndexs [][]int = nil
	outSub := NewSubNest([]byte("\""), []byte("\""))
	outIndexs = outSub.BytesToAllIndex(0, testSrc, nil)

	sub := NewSubNest([]byte("{"), []byte("}"))
	indexs := sub.bytesToIndex(0, testSrc, -1, outIndexs)

	if 4 != len(indexs) {
		t.Fatal()
		return
	}

	for i := 0; i < len(indexs); i++ {
		s := indexs[i][0]
		e := indexs[i][1]
		t.Log(string(testSrc[s:e]))
	}
}

func Test_Debug(t *testing.T) {
	testStr := []byte(`
//  The MIT License (MIT) - http://opensource.org/licenses/MIT
//
//  Copyright (c) 2014 slowfei
//
//  Create on 2014-11-05
//  Update on 2015-05-07
//  Email  slowfei#foxmail.com
//  Home   http://www.slowfei.com

// golang implement parser
// temp
package main3


/**
 *	test pacakage
 *	test line2
 *	temp late3
 */
package main1

// test package
// temp line2
// temp3
package main2
`)

	var (
		SNBetweens = []*SubNest{
			NewSubNotNest([]byte("//"), []byte("\n")),
		}
	)

	outBetweens := make([][]int, 0, 0)
	for i := 0; i < len(SNBetweens); i++ {
		tempIndexs := SNBetweens[i].BytesToAllIndex(0, testStr, outBetweens)
		if 0 != len(tempIndexs) {
			outBetweens = append(outBetweens, tempIndexs...)
		}
	}
	t.Log(outBetweens)

	for i := 0; i < len(outBetweens); i++ {
		index := outBetweens[i]
		t.Log(string(testStr[index[0]:index[1]]))
	}
}

func TestGetOutBetweens(t *testing.T) {
	/*
		temp := "temp// //"
		{
		  1
		  // 1.1 {{}{}{}{{{}}}}
		  // 1.2 }}}}}{{}{}{}}}}}
		  {
		  	1.3
		  	{
		  		1.3.1
		  	}
		  }
		  // 1.4 }{{}}
		  1.5
		  // 1.6
		  // 1.7 {
		}
		"2"
		//  temp
		'
		  {
		  	"123"
			// ' temp '
		  }
		'
	*/

	// /* temp */

	/*
		{
			{
				1
			}
		    { 1.0 }
		    "1.1"
		    "1.11"
			// 1.3
		    // 2  {
		    // 3  {
		}

		' // 4'  temp
		' 5'

		' const Temp6 = "6" '

		// const Temp7 = "7"
	*/

	srcData := []byte(`
temp := "temp// //"
{
  1
  // 1.1 {{}{}{}{{{}}}}
  // 1.2 }}}}}{{}{}{}}}}}
  {
  	1.3
  	{
  		1.3.1
  	}
  }
  // 1.4 }{{}}
  1.5
  // 1.6
  // 1.7 {
}
"2"
//  temp
'
  {
  	"123"
	// temp 
  }
'
`)

	SNBetweens := []*SubNest{
		NewSubNest([]byte("/*"), []byte("*/")),
		NewSubNotNest([]byte("//"), []byte("\n")),
		NewSubNest([]byte(`"`), []byte(`"`)),
		NewSubNest([]byte(`'`), []byte(`'`)),
		NewSubNest([]byte("`"), []byte("`")),
		NewSubNest([]byte("{"), []byte("}")),
	}
	result := GetOutBetweens(srcData, SNBetweens...)

	for i := 0; i < len(result); i++ {
		temp := result[i]
		t.Log(string(srcData[temp[0]:temp[1]]))
	}

	if 5 != len(result) {
		t.Fatal(len(result))
		return
	}

}

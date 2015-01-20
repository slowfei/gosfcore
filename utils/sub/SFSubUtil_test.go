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
	outIndexs = outSub.BytesToAllIndex(testSrc, nil)

	sub := NewSubNest([]byte("{"), []byte("}"))
	indexs := sub.bytesToIndex(testSrc, -1, outIndexs)

	for i := 0; i < len(indexs); i++ {
		s := indexs[i][0]
		e := indexs[i][1]
		t.Log(string(testSrc[s:e]))
	}
}

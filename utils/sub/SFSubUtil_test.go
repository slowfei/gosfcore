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
(
	2345
	(test)
)

(890)
`)

	var (
		SNBetweens = []*SubNest{
			NewSubNest([]byte(`(`), []byte(`)`)),
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

	sub := NewSubNest([]byte("("), []byte(")"))
	subIndex := 0
	indexs := sub.bytesToIndex(0, testStr[subIndex:], -1, outBetweens)

	for i := 0; i < len(indexs); i++ {
		s := indexs[i][0]
		e := indexs[i][1]
		t.Log(string(testStr[s+subIndex : e+subIndex]))
	}

}

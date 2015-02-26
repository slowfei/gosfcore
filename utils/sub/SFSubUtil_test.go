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

const (
	
	{
		SFSubUtil.NewSubNest([]byte("/*"), []byte("*/")),
	}
)

const (
	Test2 = "2"
)


`)

	var (
		SNBraces   = NewSubNest([]byte("{"), []byte("}"))
		SNBetweens = []*SubNest{
			NewSubNest([]byte(`"`), []byte(`"`)),
			NewSubNest([]byte("`"), []byte("`")),
			NewSubNest([]byte(`'`), []byte(`'`)),
			SNBraces,
			NewSubNest([]byte("/*"), []byte("*/")),
			NewSubNest([]byte("//"), []byte("\n")),
		}
	)

	outBetweens := make([][]int, 0, 0)
	for i := 0; i < len(SNBetweens); i++ {
		tempIndexs := SNBetweens[i].BytesToAllIndex(testStr, outBetweens)
		if 0 != len(tempIndexs) {
			outBetweens = append(outBetweens, tempIndexs...)
		}
	}

	for i := 0; i < len(outBetweens); i++ {
		outIndex := outBetweens[i]
		t.Log(string(testStr[outIndex[0]:outIndex[1]]))
	}

	sub := NewSubNest([]byte("("), []byte(")"))
	subIndex := 2
	indexs := sub.bytesToIndex(testStr[subIndex:], -1, outBetweens)

	for i := 0; i < len(indexs); i++ {
		s := indexs[i][0]
		e := indexs[i][1]
		t.Log(string(testStr[s+subIndex : e+subIndex]))
	}

}

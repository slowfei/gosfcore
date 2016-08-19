package SFLocalize

import (
	"github.com/slowfei/gosfcore/utils/filemanager"
	"path"
	"testing"
)

var _testDataPath = path.Join(SFFileManager.GetCmdDir(), "test_data", "golocalize")

/**
 *	test 排序
 */
func TestCodeSort(t *testing.T) {

	loc, err := LoadLanguages("", _testDataPath)

	if nil != err {
		t.Error(err.Error())
		return
	}

	var ploc *localize
	ploc = loc.(*localize)

	if 0 == len(ploc.Languages) {
		t.Fatal("load languages fatal, Languages is zero.")
		return
	}

	for _, v := range ploc.Languages {
		t.Log(v.Code)
	}

}

/**
 *	测试 keystrings 文件读取
 */
func TestKeyStrings(t *testing.T) {

	loc, err := LoadLanguages("", _testDataPath)

	if nil != err {
		t.Error(err.Error())
		return
	}

	testCode := "zh-CN"
	key := "k1"
	result := testCode + "-keystrings"

	code, val, isbool := loc.KeyValue(testCode, key, "")

	if !isbool {
		t.Fatal(key, "is null")
	}

	if code != testCode {
		t.Fatal("testCode=", testCode, "; return code=", code)
	}

	if result != val {
		t.Fatal("value error:", "return value=", val, "; correct results=", result)
	}
	t.Log("return value=" + val + "; code=" + code)
}

/**
 *	测试获取文件路径
 */
func TestFilepath(t *testing.T) {
	loc, err := LoadLanguages("", _testDataPath)

	if nil != err {
		t.Error(err.Error())
		return
	}

	testCode := "zh-Hans"
	testfname := "index.html"

	full, fi := loc.FileInfo(testCode, testfname)

	if nil == fi {
		t.Fatal("get file fatal: test file name is", testfname)
	}

	if fi.Name() != testfname {
		t.Fatal("testfname != return file name.", "return file name=", fi.Name())
	}

	t.Log(full)

}

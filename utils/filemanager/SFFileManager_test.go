package SFFileManager

import (
	"fmt"
	"testing"
	"time"
)

func TestWirteFilepath(t *testing.T) {
	content := []byte("test content")

	err := WirteFilepath("/Users/slowfei/Downloads/test/testgodoc/doc/md/default/testgodoc/help/help.md", content)
	if nil != err {
		t.Error(err)
	}
}

//	测试获取命令行目录
func TestGetCmdDir(t *testing.T) {
	fmt.Println("GetCmdDir: " + GetCmdDir())
	fmt.Println("GetExecDir: " + GetExecDir())
}

//	测试重命名规则
func TestFileRenameRule(t *testing.T) {
	fileName := "file.txt"
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(1).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(2).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(3).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(4).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(5).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(6).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(7).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(8).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(9).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(10).txt" {
		t.Fail()
	}
	fileName = FileRenameRule(fileName)
	fmt.Println(fileName)
	if fileName != "file(11).txt" {
		t.Fail()
	}
	fileName = FileRenameRule("file(110).txt")
	fmt.Println(fileName)
	if fileName != "file(111).txt" {
		t.Fail()
	}
}

func BenchmarkFileRenameRule(b *testing.B) {
	fileName := "11111111111file.txt"
	start := time.Now()

	start = time.Now()
	FileRenameRule(fileName)
	fmt.Println(time.Now().Sub(start))

}

func TestFileCleanRenameRule(t *testing.T) {
	fileName := "file(1).txt"
	fmt.Println(FileCleanRenameRule(fileName))
}

package SFTextFormat

import (
	"github.com/slowfei/gosfcore/utils/filemanager"
	"path"
	"testing"
)

func TestRKeyValue(t *testing.T) {
	testDir := path.Join(SFFileManager.GetCmdDir(), "test_data")
	kvFilePath := path.Join(testDir, "key_value.txt")

	dataMap := RKeyValue(kvFilePath, false)

	if 3 != len(dataMap) {
		t.Fatal("read key value number not is 3")
	}

	if "value1" != dataMap["key1"] {
		t.Fatal("key1 != value1")
	}

	if "value2" != dataMap["key2"] {
		t.Fatal("key2 != value")
	}

	t.Log("dataMap=", dataMap)
}

func TestRKeyValueBlock(t *testing.T) {
	testDir := path.Join(SFFileManager.GetCmdDir(), "test_data")
	kvFilePath := path.Join(testDir, "key_value.txt")
	keyNum := 0

	RKeyValueBlock(kvFilePath, false, func(key, value string) bool {
		keyNum++

		if 1 == keyNum && "value1" != value && "key1" != key {
			t.Fatal("index-1: key1,value1 check fatal.")
		}
		if 2 == keyNum && "value2" != value && "key2" != key {
			t.Fatal("index-1: key2,value2 check fatal.")
		}
		t.Log("readfil:", key, "=", value)
		return true
	})

	if 3 != keyNum {
		t.Fatal("read key value file number not is 3")
	}
}

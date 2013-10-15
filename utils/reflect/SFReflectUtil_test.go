package SFReflectUtil

import (
	"fmt"
	"reflect"
	"testing"
)

type TestChildStruct struct {
	Tag string
}

type TestStruct struct {
	Name *string
	Tag  TestChildStruct
	Tags []*TestChildStruct
}

func TestSetBaseTypeValue(t *testing.T) {
	srt := TestStruct{}

	//	基础赋值
	srt.Name = new(string)
	SetBaseTypeValue(reflect.ValueOf(&srt).Elem().Field(0), "slowfei--")
	if *(srt.Name) != "slowfei--" {
		t.Error(fmt.Printf("(srt.Name) != slowfei-- : value: %v", *(srt.Name)))
	}

	// //	结构赋值
	setVal := TestChildStruct{Tag: "slowfei-tag"}
	SetBaseTypeValue(reflect.ValueOf(&srt).Elem().Field(1), setVal)
	if srt.Tag.Tag != "slowfei-tag" {
		t.Error(fmt.Printf("(srt.Tag.Tag) != slowfei-tag : value: %v", srt.Tag.Tag))
	}

	//	集合设值
	setVals := make([]*TestChildStruct, 2, 2)
	setVals[0] = &TestChildStruct{Tag: "slowfei-tag-1"}
	setVals[1] = &TestChildStruct{Tag: "slowfei-tag-2"}
	srt.Tags = make([]*TestChildStruct, 2, 2)
	srt.Tags[0] = new(TestChildStruct)
	srt.Tags[1] = new(TestChildStruct)
	SetBaseTypeValue(reflect.ValueOf(&srt).Elem().Field(2), setVals)
	for _, v := range srt.Tags {
		fmt.Println(v.Tag)
	}
	if setVals[0].Tag != srt.Tags[0].Tag || setVals[1].Tag != srt.Tags[1].Tag {
		t.Error(fmt.Printf("slice set fail : %v", setVals, srt.Tags))
	}

	//	指针值设值
	strintv := "slowfei"
	srt.Name = new(string)
	fmt.Printf("%p \n", srt.Name)
	SetBaseTypeValue(reflect.ValueOf(&srt).Elem().Field(0), &strintv)
	fmt.Println(*(srt.Name), srt.Name)

}

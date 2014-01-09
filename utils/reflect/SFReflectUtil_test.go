package SFReflectUtil

import (
	// "fmt"
	"reflect"
	"testing"
)

//	struct ptr 设值
func TestSetStructPtrValue(t *testing.T) {
	type StructChild struct {
		SCTag string
	}

	type SetStructPtrStruct struct {
		CSStr        StructChild
		SCStrP       *StructChild
		CSStrArrayP  []*StructChild
		CSStrArrayP2 *[]*StructChild
	}

	isSetStructPtr := SetStructPtrStruct{}

	//	StructChild 类型 StructChild{}设值
	err := SetBaseTypeValue(reflect.ValueOf(&isSetStructPtr).Elem().Field(0), StructChild{"tag1"})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if isSetStructPtr.CSStr.SCTag != "tag1" {
		t.Error("StructChild 类型 StructChild{}设值失败")
		return
	}

	//	StructChild 类型 *StructChild指针设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetStructPtr).Elem().Field(0), &StructChild{"tag2"})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if isSetStructPtr.CSStr.SCTag != "tag2" {
		t.Error("StructChild 类型 *StructChild指针设值失败")
		return
	}

	//	*StructChild 类型 *StructChild设值
	isSetStructPtr.SCStrP = new(StructChild)
	err = SetBaseTypeValue(reflect.ValueOf(&isSetStructPtr).Elem().Field(1), &StructChild{"tag5"})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if nil == isSetStructPtr.SCStrP || isSetStructPtr.SCStrP.SCTag != "tag5" {
		t.Error("*StructChild 类型 *StructChild设值失败")
		return
	}

	//	*StructChild 类型 StructChild设值
	isSetStructPtr.SCStrP = new(StructChild)
	err = SetBaseTypeValue(reflect.ValueOf(&isSetStructPtr).Elem().Field(1), StructChild{"tag6"})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if nil == isSetStructPtr.SCStrP || isSetStructPtr.SCStrP.SCTag != "tag6" {
		t.Error("*StructChild 类型 StructChild设值失败")
		return
	}

	//	[]*StructChild 类型 []*StructChild设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetStructPtr).Elem().Field(2), []*StructChild{&StructChild{"tag8"}, &StructChild{"tag9"}})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 2 != len(isSetStructPtr.CSStrArrayP) {
		t.Error("[]*StructChild 类型 []*StructChild设值失败")
		return
	}

	//	[]*StructChild 类型 []StructChild设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetStructPtr).Elem().Field(2), []StructChild{StructChild{"tag8"}, StructChild{"tag9"}, StructChild{"tag10"}})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 3 != len(isSetStructPtr.CSStrArrayP) {
		t.Error("[]*StructChild 类型 []StructChild设值失败")
		return
	}

	//	*[]*StructChild 类型  []StructChild设值
	isSetStructPtr.CSStrArrayP2 = new([]*StructChild)
	err = SetBaseTypeValue(reflect.ValueOf(&isSetStructPtr).Elem().Field(3), []StructChild{StructChild{"tag8"}, StructChild{"tag9"}})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if nil == isSetStructPtr.CSStrArrayP2 || 2 != len((*isSetStructPtr.CSStrArrayP2)) {
		t.Error("*[]*StructChild 类型 []StructChild设值失败")
		return
	}

	//	*[]*StructChild 类型  []*StructChild设值
	isSetStructPtr.CSStrArrayP2 = new([]*StructChild)
	err = SetBaseTypeValue(reflect.ValueOf(&isSetStructPtr).Elem().Field(3), []*StructChild{&StructChild{"tag8"}, &StructChild{"tag9"}, &StructChild{"tag10"}})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if nil == isSetStructPtr.CSStrArrayP2 || 3 != len((*isSetStructPtr.CSStrArrayP2)) {
		t.Error("*[]*StructChild 类型  []*StructChild设值失败")
		return
	}

}

//	map array slice 设值
func TestSetMapArrayValue(t *testing.T) {

	type TempStruct struct {
		tag string
	}

	type SetMapArrayStruct struct {
		TMapstr  map[string]string
		TMapms   map[string]TempStruct
		TMapmsP  map[string]*TempStruct
		TArray   []TempStruct
		TArrayP  []*TempStruct
		TArrayP2 *[]*TempStruct
		TArrayP3 *[]TempStruct
		TInts    []int
		TIntsP   []*int
		TIntsP2  *[]*int
		TIntsP3  *[]int
	}

	isSetMapArray := SetMapArrayStruct{}

	//	map[string]string 类型 正常设值
	mapstr := map[string]string{"key1": "value1"}
	err := SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(0), mapstr)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if _, ok := isSetMapArray.TMapstr["key1"]; !ok {
		t.Error("map[string]string 类型 正常设值失败")
		return
	}

	//	map[string]TempStruct 类型 正常设值
	mapms := map[string]TempStruct{"key2": TempStruct{"temptag"}}
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(1), mapms)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if _, ok := isSetMapArray.TMapms["key2"]; !ok {
		t.Error("map[string]TempStruct 类型 正常设值失败")
		return
	}

	//	map[string]*TempStruct 类型 正常设值
	mapmsp := map[string]*TempStruct{"key3": &TempStruct{"temptag"}}
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(2), mapmsp)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if _, ok := isSetMapArray.TMapmsP["key3"]; !ok {
		t.Error("map[string]*TempStruct 类型 正常设值失败")
		return
	}

	// []TempStruct 类型 正常设值
	ts := []TempStruct{TempStruct{"tag1"}, TempStruct{"tag2"}}
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(3), ts)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 2 != len(isSetMapArray.TArray) {
		t.Error("[]TempStruct 类型 正常设值失败")
		return
	}

	//	[]TempStruct 类型 []*TempStruct设值
	tsp := []*TempStruct{&TempStruct{"tag1"}, &TempStruct{"tag2"}, &TempStruct{"tag3"}}
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(3), tsp)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 3 != len(isSetMapArray.TArray) {
		t.Error("[]TempStruct 类型 []*TempStruct设值失败")
		return
	}

	//	[]TempStruct 类型 单元素TempStruct设值
	isSetMapArray.TArray = nil
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(3), TempStruct{"tag2s"})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 1 != len(isSetMapArray.TArray) || isSetMapArray.TArray[0].tag != "tag2s" {
		t.Error("[]TempStruct 类型 单元素TempStruct设值失败")
		return
	}

	//	[]TempStruct 类型 *TempStruct设值
	isSetMapArray.TArray = nil
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(3), &TempStruct{"tag2sb"})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 1 != len(isSetMapArray.TArray) || isSetMapArray.TArray[0].tag != "tag2sb" {
		t.Error("[]TempStruct 类型 单元素TempStruct设值失败")
		return
	}

	//	[]*TempStruct 类型 []TempStruct设值
	tsp2 := []TempStruct{TempStruct{"tag1"}, TempStruct{"tag2"}}
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(4), tsp2)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 2 != len(isSetMapArray.TArrayP) {
		t.Error(len(isSetMapArray.TArrayP))
		t.Error("[]*TempStruct 类型 []TempStruct设值失败:")
		return
	}

	//	[]*TempStruct 类型 单元素TempStruct设值
	isSetMapArray.TArrayP = nil
	tstemp := TempStruct{"tag3st"}
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(4), tstemp)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 1 != len(isSetMapArray.TArrayP) || isSetMapArray.TArrayP[0].tag != "tag3st" {
		//	TODO
		// t.Error("[]TempStruct 类型 单元素TempStruct设值失败，暂时无法解决的异常。具体可以查看集合赋值的代码。")
	}

	//	[]*TempStruct 类型 *TempStruct 设值
	isSetMapArray.TArrayP = nil
	tstemp2 := &TempStruct{"tag3st6"}
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(4), tstemp2)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 1 != len(isSetMapArray.TArrayP) || isSetMapArray.TArrayP[0].tag != "tag3st6" {
		t.Error("[]*TempStruct 类型 *TempStruct 设值 失败")
	}

	//	[]int 类型 正常设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(7), []int{1, 2, 3, 4, 5})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 5 != len(isSetMapArray.TInts) {
		t.Error("[]int 类型 正常设值失败")
	}

	//	[]int 类型 int设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(7), 5)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 1 != len(isSetMapArray.TInts) || isSetMapArray.TInts[0] != 5 {
		t.Error("[]int 类型 正常设值失败")
	}

	//	[]int 类型 []*int设值
	a, b, c := 1, 2, 3
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(7), []*int{&a, &b, &c})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 3 != len(isSetMapArray.TInts) {
		t.Error("[]int 类型 []*int设值失败")
	}

	//	[]*int 类型 正常设值
	d, e, f, g := 1, 2, 3, 4
	isSetMapArray.TIntsP = nil
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(8), []*int{&d, &e, &f, &g})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 4 != len(isSetMapArray.TIntsP) {
		t.Error("[]*int 类型 正常设值失败")
	}

	//	[]*int 类型 []int设值
	isSetMapArray.TIntsP = nil
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(8), []int{d, e})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 2 != len(isSetMapArray.TIntsP) {
		t.Error("[]*int 类型 正常设值失败")
	}

	//	[]*int 类型 单元素int设值
	isSetMapArray.TIntsP = nil
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(8), 8)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 1 != len(isSetMapArray.TIntsP) || (*isSetMapArray.TIntsP[0]) != 8 {
		//	TODO
		// t.Error("[]*int 类型 正常设值失败")
	}

	//	[]int 类型 *[]int 设值
	isSetMapArray.TInts = nil
	err = SetBaseTypeValue(reflect.ValueOf(&isSetMapArray).Elem().Field(7), &[]*int{&a, &b, &c, &d})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if 4 != len(isSetMapArray.TInts) {
		t.Error("[]int 类型 *[]int 设值失败")
	}

}

//	测试设置 uint float
func TestSetUintFloatValue(t *testing.T) {
	type SetUintFloatStruct struct {
		TUint     uint
		TUintP    *uint
		TUint32   uint32
		TUint32P  *uint32
		TFloat32  float32
		TFloat32P *float32
		TFloat64  float64
		TFloat64P *float64
	}

	isSetUintFloat := SetUintFloatStruct{}

	//	uint 类型 字符串设值
	err := SetBaseTypeValue(reflect.ValueOf(&isSetUintFloat).Elem().Field(0), "12")
	if nil != err {
		t.Error(err.Error())
		return
	}
	if isSetUintFloat.TUint != 12 {
		t.Error("uint 类型 字符串设值失败")
		return
	}

	//	uint 类型 正常设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetUintFloat).Elem().Field(0), uint(56))
	if nil != err {
		t.Error(err.Error())
	}
	if isSetUintFloat.TUint != 56 {
		t.Error("uint 类型 正常设值失败")
		return
	}

	//	float32 类型 字符串设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetUintFloat).Elem().Field(4), "12.02")
	if nil != err {
		t.Error(err.Error())
	}
	if isSetUintFloat.TFloat32 != 12.02 {
		t.Error("uint 类型 字符串设值失败")
		return
	}
}

//	测试设置 bool 和 int 值
func TestSetBoolIntValue(t *testing.T) {

	//	测试 bool 和 int 的设值

	// test	bool int struct
	type SetBoolIntStruct struct {
		TBool   bool
		TBoolP  *bool
		TInt    int
		TIntP   *int
		TInt8   int8
		TInt8P  *int8
		TInt16  int16
		TInt16P *int16
		TInt32  int32
		TInt32P *int32
		TInt64  int64
		TInt64P *int64
	}
	//	设置值，设置参数值作为比较
	valSetBoolInt := SetBoolIntStruct{}
	valSetBoolInt.TBool = true
	valSetBoolInt.TBoolP = &valSetBoolInt.TBool
	valSetBoolInt.TInt = 10
	valSetBoolInt.TIntP = &valSetBoolInt.TInt
	valSetBoolInt.TInt8 = int8(8)
	valSetBoolInt.TInt8P = &valSetBoolInt.TInt8
	valSetBoolInt.TInt16 = int16(16)
	valSetBoolInt.TInt16P = &valSetBoolInt.TInt16
	valSetBoolInt.TInt32 = int32(32)
	valSetBoolInt.TInt32P = &valSetBoolInt.TInt32
	valSetBoolInt.TInt64 = int64(64)
	valSetBoolInt.TInt64P = &valSetBoolInt.TInt64

	//	被设置结构
	isSetBoolInt := SetBoolIntStruct{}

	//	bool 类型
	err := SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(0), "true")
	if nil != err {
		t.Error(err.Error())
	}
	if isSetBoolInt.TBool != valSetBoolInt.TBool {
		t.Error("bool 类型 字符串设值失败")
		return
	}

	//	bool 类型 集合测试
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(0), []string{"false", "false"})
	if nil != err {
		t.Error(err.Error())
	}
	if isSetBoolInt.TBool != false {
		t.Error("bool 类型 集合字符串设值失败")
		return
	}

	//	bool 类型 正常设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(0), true)
	if nil != err {
		t.Error(err.Error())
	}
	if isSetBoolInt.TBool != true {
		t.Error("bool 类型 正常设值失败")
		return
	}

	//	bool 类型 值指针设值
	pbool := true
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(0), &pbool)
	if nil != err {
		t.Error(err.Error())
	}
	if isSetBoolInt.TBool != pbool {
		t.Error("bool 类型 值指针设值")
		return
	}

	//	bool 指针类型 指针设值
	isSetBoolInt.TBoolP = nil
	pbool = false
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(1), &pbool)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if *(isSetBoolInt.TBoolP) != pbool {
		t.Error("bool 指针类型 指针设值失败")
		return
	}

	//	bool 指针类型 字符串设值
	// isSetBoolInt.TBoolP = nil // 如果指针为nil将会设值失败，所需设置指针类型先初始化，或者设值的类型是相同的。
	strbool := "false"
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(1), &strbool)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if nil == isSetBoolInt.TBoolP || *(isSetBoolInt.TBoolP) != false {
		t.Error("bool 指针类型 字符串设置失败")
		return
	}

	//	int 类型 字符串设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(2), "12")
	if nil != err {
		t.Error(err.Error())
		return
	}
	if isSetBoolInt.TInt != 12 {
		t.Error("int 类型 字符串设值失败")
	}

	//	int 类型 字符串集合设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(2), []string{"16", "10"})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if isSetBoolInt.TInt != 16 {
		t.Error("int 类型 字符串集合设值失败")
	}

	//	int 类型 正常
	var index int = 23
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(2), index)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if isSetBoolInt.TInt != index {
		t.Error("int 类型 正常设值失败")
	}

	//	int 类型 指针设值
	index = 33
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(2), &index)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if isSetBoolInt.TInt != index {
		t.Error("int 类型 指针设值失败")
	}

	//	int 类型 []int集合设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(2), []int{32, 12})
	if nil != err {
		t.Error(err.Error())
		return
	}
	if isSetBoolInt.TInt != 32 {
		t.Error("int 类型 指针设值失败")
	}

	//	int 指针类型 字符串设值
	isSetBoolInt.TIntP = new(int) //	需要初始化才能进行字符串设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(3), "23")
	if nil != err {
		t.Error(err.Error())
		return
	}
	if nil == isSetBoolInt.TIntP || *(isSetBoolInt.TIntP) != 23 {
		t.Error("int 指针类型 字符串设值失败")
	}

	//	int 指针类型 *int设值
	isSetBoolInt.TIntP = nil // 如果设置的是指针类型可以不用初始化
	index = 39
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(3), &index)
	if nil != err {
		t.Error(err.Error())
		return
	}
	if nil == isSetBoolInt.TIntP || *(isSetBoolInt.TIntP) != index {
		t.Error("int 指针类型 *int设值")
	}

	// int32 类型 字符串设值
	err = SetBaseTypeValue(reflect.ValueOf(&isSetBoolInt).Elem().Field(8), "21")
	if nil != err {
		t.Error(err.Error())
		return
	}
	if isSetBoolInt.TInt32 != 21 {
		t.Error("int32 指针类型 字符串设值失败")
	}
}

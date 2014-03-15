//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-25
//  Update on 2013-10-17
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//	json 组件封装，基于encoding/json
//
package SFJson

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/slowfei/gosfcore/utils/reflect"
	"math"
	"reflect"
	"strconv"
	"strings"
)

var (
	//	我的理解好像作为一个存储区的作用，具体还不知道什么意思
	scratch [64]byte
)

//	设置一个struct主要为了标识是一个Json对象
type Json struct {
	rawdata     interface{}  //	原始数据
	analyzedBuf bytes.Buffer //	解析后的数据
	nullTag     string       //	空数据处理标识
	boolTag     string       //	bool处理标识
}

//	创建json解析对象
//	@v		 需要处理的数据对象，struct、map
//	@nullTag 空数据标识，如果数据key或字段为空的处理方案
//			e.g.:
//				默认 ""		- 如果空字符串不会将(空数据)字段或key添加到json解析的数据中
//				"null"		- 空数据解析: {"Key":null}
//				"nil"		- 空数据解析: {"Key":nil}
//	@boolTag 布尔类型标识
//			 e.g.:
//				默认 ""		- {"KeyBool":trye|false}
//				"YES|NO"	- {"KeyBool":YES|NO}
//				"1|0"		- {"KeyBool":1|0}
//
func NewJson(v interface{}, nullTag, boolTag string) (Json, error) {
	j := Json{rawdata: v, nullTag: nullTag, boolTag: boolTag}
	err := j.analyzeJson()
	return j, err
}

//	创建一个空类型的Json对象，输出的信息为{} or []
//
//	@isArray YES nil info = [], NO nil info = {}
//
func NewJsonNil(isArray bool) Json {
	var j Json

	if isArray {
		j = Json{nil, *bytes.NewBufferString("[]"), "", ""}
	} else {
		j = Json{nil, *bytes.NewBufferString("{}"), "", ""}

	}
	return j
}

func (j *Json) analyzeJson() error {

	if 0 < j.analyzedBuf.Len() {
		return nil
	}

	if nil == j.rawdata {
		return errors.New("Json raw data nil, can not be analyzed")
	}

	err := marshal(reflect.ValueOf(j.rawdata), j.nullTag, j.boolTag, &j.analyzedBuf)
	return err
}

func (j *Json) String() string {
	if 0 >= j.analyzedBuf.Len() {
		return ""
	}
	return string(j.analyzedBuf.Bytes())
}

func (j *Json) Byte() []byte {
	if 0 >= j.analyzedBuf.Len() {
		return nil
	}
	return j.analyzedBuf.Bytes()
}

func (j *Json) RawData() interface{} {
	return j.rawdata
}

//	将json解析为需要的结构或map
//	目前使用的是go内置的解析
//	@jsonByte
//	@v			需要地址引用传递，并且类型是slice
//				jsonMap 	[]map[string]string
//				jsonStruct	[]Struct
func Unmarshal(jsonByte []byte, v interface{}) error {
	return json.Unmarshal(jsonByte, v)
}

//	根据对象解析json
//	@v		  需要处理的数据对象，struct、map
//	@nullTag  空数据标识，如果数据key或字段为空的处理方案
//			e.g.:
//				默认 ""		- 如果空字符串不会将(空数据)字段或key添加到json解析的数据中
//				"null"		- 空数据解析: {"Key":null}
//				"nil"		- 空数据解析: {"Key":nil}
//	@boolTag 布尔类型标识
//			 e.g.:
//				默认 ""		- {"KeyBool":trye|false}
//				"YES|NO"	- {"KeyBool":YES|NO}
//				"1|0"		- {"KeyBool":1|0}
//
func Marshal(v interface{}, nullTag, boolTag string) ([]byte, error) {
	var buf bytes.Buffer
	err := marshal(reflect.ValueOf(v), nullTag, boolTag, &buf)
	return buf.Bytes(), err
}

func marshal(v reflect.Value, nullTag, boolTag string, buf *bytes.Buffer) error {
	var err error

	switch v.Kind() {
	case reflect.Bool:
		boolTrue := "true"
		boolFalse := "false"
		if "" != boolTag {
			bs := strings.Split(boolTag, "|")
			if 2 == len(bs) {
				boolTrue = bs[0]
				boolFalse = bs[1]
			}
		}
		bv := v.Bool()
		if bv {
			buf.WriteString(boolTrue)
		} else {
			buf.WriteString(boolFalse)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		b := strconv.AppendInt(scratch[:0], v.Int(), 10)
		buf.Write(b)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		b := strconv.AppendUint(scratch[:0], v.Uint(), 10)
		buf.Write(b)
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		if math.IsInf(f, 0) || math.IsNaN(f) {
			err = errors.New(strconv.FormatFloat(f, 'g', -1, v.Type().Bits()))
		}
		b := strconv.AppendFloat(scratch[:0], f, 'g', -1, v.Type().Bits())
		buf.Write(b)
	case reflect.String:
		str := v.String()
		if "" == str {
			buf.WriteString("\"\"")
		} else {
			buf.WriteString("\"" + str + "\"")
		}

	case reflect.Struct:
		vt := v.Type()
		numTField := vt.NumField()

		if 0 == numTField {
			if "" != nullTag && 0 < buf.Len() {
				buf.WriteString(nullTag)
			}
			break
		}

		buf.WriteByte('{')
		first := true

		for i := 0; i < numTField; i++ {
			childFieldT := vt.Field(i)
			childFieldV := v.Field(i)

			if "" == nullTag && SFReflectUtil.IsNullValue(childFieldV) {
				continue
			}

			if first {
				first = false
			} else {
				buf.WriteByte(',')
			}
			buf.WriteString("\"" + childFieldT.Name + "\"")
			buf.WriteByte(':')
			marshal(childFieldV, nullTag, boolTag, buf)
		}

		buf.WriteByte('}')
	case reflect.Map:
		if v.Type().Key().Kind() != reflect.String {
			err = errors.New("json: unsupported type: " + v.Type().String())
			break
		}
		if 0 == v.Len() {
			if "" != nullTag && 0 < buf.Len() {
				buf.WriteString(nullTag)
			}
			break
		}

		buf.WriteByte('{')
		first := true
		for _, k := range v.MapKeys() {
			kv := v.MapIndex(k)

			if "" == nullTag && SFReflectUtil.IsNullValue(kv) {
				continue
			}

			if first {
				first = false
			} else {
				buf.WriteByte(',')
			}
			buf.WriteString("\"" + k.String() + "\"")
			buf.WriteByte(':')
			marshal(kv, nullTag, boolTag, buf)
		}
		buf.WriteByte('}')

	case reflect.Array, reflect.Slice:
		if 0 == v.Len() {
			if "" != nullTag && 0 < buf.Len() {
				buf.WriteString(nullTag)
			}
			break
		}
		buf.WriteByte('[')
		first := true
		count := v.Len()
		for i := 0; i < count; i++ {
			v2 := v.Index(i)
			if "" == nullTag && SFReflectUtil.IsNullValue(v2) {
				continue
			}
			if first {
				first = false
			} else {
				buf.WriteByte(',')
			}
			marshal(v2, nullTag, boolTag, buf)
		}
		buf.WriteByte(']')
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			if "" != nullTag {
				buf.WriteString(nullTag)
			}
			break
		}
		marshal(v.Elem(), nullTag, boolTag, buf)
	default:
	}

	return err
}

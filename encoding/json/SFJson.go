//  Copyright 2013 slowfei And The Contributors All rights reserved.
//
//  Software Source Code License Agreement (BSD License)
//
//  Create on 2013-08-25
//  Update on 2014-09-25
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//  json 组件封装，基于encoding/json
package SFJson

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/slowfei/gosfcore/utils/filemanager"
	"github.com/slowfei/gosfcore/utils/reflect"
	"github.com/slowfei/gosfcore/utils/strings"
	"math"
	"reflect"
	"strconv"
	"strings"
)

var (
	// int uint float buffer
	scratch [64]byte

	// json data type
	// TypeInt          = DataType(1) TODO golang json parse no int type
	TypeFloat  = DataType(2)
	TypeBool   = DataType(3)
	TypeString = DataType(4)
	TypeArray  = DataType(5)
	TypeMap    = DataType(6)

	// json not nil data type
	TypeNotnilFloat  = DataTypeNotnil(2)
	TypeNotnilBool   = DataTypeNotnil(3)
	TypeNotnilString = DataTypeNotnil(4)
	TypeNotnilArray  = DataTypeNotnil(5)
	TypeNotnilMap    = DataTypeNotnil(6)
)

/**
 *  json struct data type
 *
 *  value type is int|float|string|array|map
 */
type DataType int8

/**
 *  json struct data not nil type
 *
 *  value is not null
 */
type DataTypeNotnil int8

/**
 *  Json struct object
 */
type Json struct {
	rawdata     interface{}  // 原始数据
	analyzedBuf bytes.Buffer // 解析后的数据
	nullTag     string       // 空数据处理标识
	boolTag     string       // bool处理标识
}

/**
 *  parse interface data new Json object
 *
 *  @param `v`          handle object struct or map or slice
 *  @param `nullTag`    nil tag, field or data is nil replace
 *          e.g.:
 *              default ""  - 如果空字符串不会将(空数据)字段或key添加到json解析的数据中
 *              "null"      - 空数据解析: {"Key":null}
 *              "nil"       - 空数据解析: {"Key":nil}
 *  @param `boolTag`    bool tag
 *          e.g.:
 *              default ""  - {"KeyBool":trye|false}
 *              "YES|NO"    - {"KeyBool":YES|NO}
 *              "1|0"       - {"KeyBool":1|0}
 *  @return Json
 *  @return error
 */
func NewJson(v interface{}, nullTag, boolTag string) (Json, error) {
	j := Json{rawdata: v, nullTag: nullTag, boolTag: boolTag}
	err := j.parseJson()
	return j, err
}

/**
 *  new nil Json struct object
 *
 *  @param isArray YES nil info = [], NO nil info = {}
 *
 */
func NewJsonNil(isArray bool) Json {
	var j Json

	if isArray {
		j = Json{nil, *bytes.NewBufferString("[]"), "", ""}
	} else {
		j = Json{nil, *bytes.NewBufferString("{}"), "", ""}

	}
	return j
}

/**
 *  parse json raw data
 *
 *  @return error
 */
func (j *Json) parseJson() error {

	if 0 < j.analyzedBuf.Len() {
		return nil
	}

	if nil == j.rawdata {
		return errors.New("Json raw data nil, can not be analyzed")
	}

	err := marshal(reflect.ValueOf(j.rawdata), j.nullTag, j.boolTag, &j.analyzedBuf)
	return err
}

/**
 *  to json string
 *
 *  @return
 */
func (j *Json) String() string {
	if 0 >= j.analyzedBuf.Len() {
		return ""
	}
	return string(j.analyzedBuf.Bytes())
}

/**
 *  to json bytes
 *
 *  @return
 */
func (j *Json) Bytes() []byte {
	if 0 == j.analyzedBuf.Len() {
		return nil
	}
	return j.analyzedBuf.Bytes()
}

/**
 *	to json format bytes
 *
 *	@return
 */
func (j *Json) BytesFormat() []byte {

	if 0 == j.analyzedBuf.Len() {
		return nil
	}

	var formatBuf bytes.Buffer
	err := json.Indent(&formatBuf, j.analyzedBuf.Bytes(), "", "\t")

	if nil != err {
		return nil
	}

	return formatBuf.Bytes()
}

/**
 *
 *	@param `path`	output file path
 *	@param `format` whether format output
 */
func (j *Json) WriteFilepath(path string, format bool) error {

	var data []byte = nil

	if format {
		data = j.BytesFormat()
	} else {
		data = j.Bytes()
	}

	return SFFileManager.WirteFilepath(path, data)
}

/**
 *  get json raw interface data
 *
 *  @return
 */
func (j *Json) RawData() interface{} {
	return j.rawdata
}

/**
 *  json bytes to struct or slice or map
 *
 *  @param `jsonByte`
 *  @param `v`       pointer object
 */
func Unmarshal(jsonByte []byte, v interface{}) error {
	return json.Unmarshal(jsonByte, v)
}

/**
 *  parse json bytes data
 *
 *  @param `v`          handle object struct or map or slice
 *  @param `nullTag`    nil tag, field or data is nil replace
 *          e.g.:
 *              default ""  - 如果空字符串不会将(空数据)字段或key添加到json解析的数据中
 *              "null"      - 空数据解析: {"Key":null}
 *              "nil"       - 空数据解析: {"Key":nil}
 *  @param `boolTag`    bool tag
 *          e.g.:
 *              default ""  - {"KeyBool":trye|false}
 *              "YES|NO"    - {"KeyBool":YES|NO}
 *              "1|0"       - {"KeyBool":1|0}
 *  @return []byte
 *  @return error
 */
func Marshal(v interface{}, nullTag, boolTag string) ([]byte, error) {
	var buf bytes.Buffer
	err := marshal(reflect.ValueOf(v), nullTag, boolTag, &buf)
	return buf.Bytes(), err
}

/**
 *  private marshal
 */
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

/**
 *  validate map type data
 *  can nesting validate
 *
 *  `validData` :
 *      {
 *          Field1:TypeInt,
 *          Field2:TypeNotnilString,
 *          Field3:map[string]interface{}{
 *                  childField1: TypeNotnilString
 *                  childField2: TypeFloat
 *          },
 *          Field4:[]interface{}{
 *              TypeString,
 *              TypeInt,
 *              TypeNotnilString
 *          }
 *      }
 *
 *  @param `jsonMap` json map struct
 *  @param `validData` valid data validate
 *  @return valid data is true
 */
func ValidateMap(jsonMap interface{}, validData map[string]interface{}) bool {

	if nil == validData || 0 == len(validData) {
		return true
	}

	switch v := jsonMap.(type) {
	case map[string]interface{}:
		for key, validObj := range validData {

			object, ok := v[key]

			if !ok {
				object, ok = v[SFStringsUtil.ToLower(key)]
			}

			if ok {

				if !validateObject(object, validObj) {
					return false
				}

			} else {

				switch validObj.(type) {
				case DataTypeNotnil, map[string]DataTypeNotnil, []DataTypeNotnil:
					return false
				case map[string]interface{}, []interface{}:
					//  保留此两个类型的验证，现在无法确定是否需要匹配。
				default:
				}
			}
		}
	default:
		return false
	}

	return true
}

/**
 *  validate array type data
 *  can nesting validate
 *
 *  `validData` :
 *      [
 *          TypeInt,
 *          TypeNotnilString,
 *          map[string]interface{}{
 *                  childField1: TypeNotnilString
 *                  childField2: TypeFloat
 *          },
 *          []interface{}{
 *              TypeString,
 *              TypeInt,
 *              TypeNotnilString
 *          }
 *      ]
 *
 *  @param `jsonArray` json array struct
 *  @param `validData` valid data validate
 *  @return valid data is true
 */
func ValidateArray(jsonArray interface{}, validDatas []interface{}) bool {
	validCount := len(validDatas)

	if nil == validDatas || 0 == validCount {
		return true
	}

	switch v := jsonArray.(type) {
	case []interface{}:
		arrayCount := len(v)
		for i := 0; i < arrayCount; i++ {
			if i < validCount {
				validObj := validDatas[i]
				object := v[i]

				if !validateObject(object, validObj) {
					return false
				}

			} else {
				//  TODO 考虑是否验证全部array, 根据规则验证合法数据长度相等的就可以了。
				break
			}
		}
	default:
		return false
	}

	return true
}

/**
 *  alone validate object type
 *
 *  @param `validData`
 */
func validateObject(object interface{}, validData interface{}) bool {

	switch objectValue := object.(type) {
	case bool:
		switch childValid := validData.(type) {
		case DataType:
			if childValid != TypeBool {
				return false
			}
		case DataTypeNotnil:
			if childValid != TypeNotnilBool {
				return false
			}
		default:
			return false
		}
	case float64:
		switch childValid := validData.(type) {
		case DataType:
			if childValid != TypeFloat {
				return false
			}
		case DataTypeNotnil:
			if childValid != TypeNotnilFloat {
				return false
			}
		default:
			return false
		}
	case string:
		switch childValid := validData.(type) {
		case DataType:
			if childValid != TypeString {
				return false
			}
		case DataTypeNotnil:
			if childValid != TypeNotnilString || 0 == len(objectValue) {
				return false
			}
		default:
			return false
		}
	case []interface{}:

		switch childValid := validData.(type) {
		case DataType:
			if childValid != TypeArray {
				return false
			}
		case DataTypeNotnil:
			if childValid != TypeNotnilArray || 0 == len(objectValue) {
				return false
			}
		case []interface{}:
			if !ValidateArray(objectValue, childValid) {
				return false
			}
		case []DataType:
			count := len(childValid)
			newValid := make([]interface{}, count, count)
			for i := 0; i < count; i++ {
				newValid[i] = childValid[i]
			}
			if !ValidateArray(objectValue, newValid) {
				return false
			}
		case []DataTypeNotnil:
			count := len(childValid)
			newValid := make([]interface{}, count, count)
			for i := 0; i < count; i++ {
				newValid[i] = childValid[i]
			}
			if !ValidateArray(objectValue, newValid) {
				return false
			}
		default:
			return false
		}

	case map[string]interface{}:

		switch chlidValid := validData.(type) {
		case DataType:
			if chlidValid != TypeMap {
				return false
			}
		case DataTypeNotnil:
			if chlidValid != TypeNotnilMap || 0 == len(objectValue) {
				return false
			}
		case map[string]interface{}:
			if !ValidateMap(objectValue, chlidValid) {
				return false
			}
		case map[string]DataType:
			newValid := make(map[string]interface{})
			for tk, tv := range chlidValid {
				newValid[tk] = tv
			}
			if !ValidateMap(objectValue, newValid) {
				return false
			}
		case map[string]DataTypeNotnil:
			newValid := make(map[string]interface{})
			for tk, tv := range chlidValid {
				newValid[tk] = tv
			}
			if !ValidateMap(objectValue, newValid) {
				return false
			}
		default:
			return false
		}

	default:
		//  TODO 进入这里可能是非有效的数据，或则golang更改了json的实现规则，暂时不做操作
		//  目前已知golang实现json有效的数据类型 float64、string、[]interface{}、map[string]interface{}
	}

	return true
}

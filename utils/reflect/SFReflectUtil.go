//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-24
//  Update on 2013-10-17
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//	reflect 辅助工具
//
package SFReflectUtil

import (
	"errors"
	// "fmt"
	// "io"
	// "os"
	"reflect"
	"strconv"
)

//	判断空值
func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

//	判断reflect.Value的null值，只有一些类型是 if null 的处理
func IsNullValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

//	根据字符串值设置反射对象的基本类型值
//	例如: int Float32 string...
func SetBaseTypeValue(refValue reflect.Value, setVal interface{}) error {
	var errTemp error = nil

	refValue = reflect.Indirect(refValue)
	if !refValue.CanSet() {
		errTemp = errors.New("failed to set value")
		return errTemp
	}

	switch refValue.Kind() {
	case reflect.Bool:
		switch v := setVal.(type) {
		case string:
			if b, err := strconv.ParseBool(v); err == nil {
				refValue.SetBool(b)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if b, err := strconv.ParseBool(v[0]); err == nil {
					refValue.SetBool(b)
				} else {
					errTemp = err
				}
			}
		case bool:
			refValue.SetBool(v)
		case *bool:
			refValue.SetBool(*v)
		}
	case reflect.Int:
		switch v := setVal.(type) {
		case string:
			if iv, err := strconv.Atoi(v); err == nil {
				refValue.Set(reflect.ValueOf(iv))
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if iv, err := strconv.Atoi(v[0]); err == nil {
					refValue.Set(reflect.ValueOf(iv))
				} else {
					errTemp = err
				}
			}
		case int:
			refValue.SetInt(int64(v))
		case *int:
			refValue.SetInt(int64(*v))
		}
	case reflect.Int8:
		switch v := setVal.(type) {
		case string:
			if i8, err := strconv.ParseInt(v, 10, 8); err == nil {
				refValue.SetInt(i8)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if i8, err := strconv.ParseInt(v[0], 10, 8); err == nil {
					refValue.SetInt(i8)
				} else {
					errTemp = err
				}
			}
		case int8:
			refValue.SetInt(int64(v))
		case *int8:
			refValue.SetInt(int64(*v))
		}
	case reflect.Int16:
		switch v := setVal.(type) {
		case string:
			if i16, err := strconv.ParseInt(v, 10, 16); err == nil {
				refValue.SetInt(i16)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if i16, err := strconv.ParseInt(v[0], 10, 16); err == nil {
					refValue.SetInt(i16)
				} else {
					errTemp = err
				}
			}
		case int16:
			refValue.SetInt(int64(v))
		case *int16:
			refValue.SetInt(int64(*v))
		}
	case reflect.Int32:
		switch v := setVal.(type) {
		case string:
			if i32, err := strconv.ParseInt(v, 10, 32); err == nil {
				refValue.SetInt(i32)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if i32, err := strconv.ParseInt(v[0], 10, 32); err == nil {
					refValue.SetInt(i32)
				} else {
					errTemp = err
				}
			}
		case int32:
			refValue.SetInt(int64(v))
		case *int32:
			refValue.SetInt(int64(*v))
		}

	case reflect.Int64:
		switch v := setVal.(type) {
		case string:
			if i64, err := strconv.ParseInt(v, 10, 64); err == nil {
				refValue.SetInt(i64)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if i64, err := strconv.ParseInt(v[0], 10, 64); err == nil {
					refValue.SetInt(i64)
				} else {
					errTemp = err
				}
			}
		case int64:
			refValue.SetInt(v)
		case *int64:
			refValue.SetInt(*v)
		}
	case reflect.Uint:
		switch v := setVal.(type) {
		case string:
			if u32, err := strconv.ParseUint(v, 10, 32); err == nil {
				refValue.SetUint(u32)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if u32, err := strconv.ParseUint(v[0], 10, 32); err == nil {
					refValue.SetUint(u32)
				} else {
					errTemp = err
				}
			}
		case uint:
			refValue.SetUint(uint64(v))
		case *uint:
			refValue.SetUint(uint64(*v))
		}
	case reflect.Uint8:
		switch v := setVal.(type) {
		case string:
			if u8, err := strconv.ParseUint(v, 10, 8); err == nil {
				refValue.SetUint(u8)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if u8, err := strconv.ParseUint(v[0], 10, 8); err == nil {
					refValue.SetUint(u8)
				} else {
					errTemp = err
				}
			}
		case uint8:
			refValue.SetUint(uint64(v))
		case *uint8:
			refValue.SetUint(uint64(*v))
		}
	case reflect.Uint16:
		switch v := setVal.(type) {
		case string:
			if u16, err := strconv.ParseUint(v, 10, 16); err == nil {
				refValue.SetUint(u16)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if u16, err := strconv.ParseUint(v[0], 10, 16); err == nil {
					refValue.SetUint(u16)
				} else {
					errTemp = err
				}
			}
		case uint16:
			refValue.SetUint(uint64(v))
		case *uint16:
			refValue.SetUint(uint64(*v))
		}
	case reflect.Uint32:
		switch v := setVal.(type) {
		case string:
			if u32, err := strconv.ParseUint(v, 10, 32); err == nil {
				refValue.SetUint(u32)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if u32, err := strconv.ParseUint(v[0], 10, 32); err == nil {
					refValue.SetUint(u32)
				} else {
					errTemp = err
				}
			}
		case uint32:
			refValue.SetUint(uint64(v))
		case *uint32:
			refValue.SetUint(uint64(*v))
		}
	case reflect.Uint64:
		switch v := setVal.(type) {
		case string:
			if u64, err := strconv.ParseUint(v, 10, 64); err == nil {
				refValue.SetUint(u64)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if u64, err := strconv.ParseUint(v[0], 10, 64); err == nil {
					refValue.SetUint(u64)
				} else {
					errTemp = err
				}
			}
		case uint64:
			refValue.SetUint(v)
		case *uint64:
			refValue.SetUint(*v)
		}
	case reflect.Float32:
		switch v := setVal.(type) {
		case string:
			if f32, err := strconv.ParseFloat(v, 32); err == nil {
				refValue.SetFloat(f32)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if f32, err := strconv.ParseFloat(v[0], 32); err == nil {
					refValue.SetFloat(f32)
				} else {
					errTemp = err
				}
			}
		case float32:
			refValue.SetFloat(float64(v))
		case *float32:
			refValue.SetFloat(float64(*v))
		}
	case reflect.Float64:
		switch v := setVal.(type) {
		case string:
			if f64, err := strconv.ParseFloat(v, 64); err == nil {
				refValue.SetFloat(f64)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if f64, err := strconv.ParseFloat(v[0], 64); err == nil {
					refValue.SetFloat(f64)
				} else {
					errTemp = err
				}
			}
		case float64:
			refValue.SetFloat(v)
		case *float64:
			refValue.SetFloat(*v)
		}
	case reflect.String:

		switch v := setVal.(type) {
		case string:
			refValue.SetString(v)
		case []string:
			if 0 != len(v) {
				refValue.SetString(v[0])
			}
		case *string:
			refValue.SetString(*v)
		}
	case reflect.Complex64, reflect.Complex128:
	case reflect.Map:
		setValue := reflect.ValueOf(setVal)
		setValueInd := reflect.Indirect(setValue)
		if 0 < setValueInd.Len() {
			if refValue.Type() == setValueInd.Type() {
				if setValue.Kind() == reflect.Ptr {
					refValue.Set(setValue.Elem())
				} else {
					refValue.Set(setValue)
				}
			}
		}
	case reflect.Array, reflect.Slice:
		setValue := reflect.ValueOf(setVal)
		setValueInd := reflect.Indirect(setValue)
		if 0 < setValueInd.Len() {
			if refValue.Type() == setValueInd.Type() {
				if setValue.Kind() == reflect.Ptr {
					refValue.Set(setValue.Elem())
				} else {
					refValue.Set(setValue)
				}
			} else if refValue.Type().Elem() == setValueInd.Type() {
				//	如果类型不相等，表示设置的值不是[]的类型，所以直接赋值到[0]中。
				makeVal := reflect.MakeSlice(refValue.Type(), 1, 1)
				if setValue.Kind() == reflect.Ptr {
					makeVal.Index(0).Set(setValue.Elem())
				} else {
					makeVal.Index(0).Set(setValue)
				}
				refValue.Set(makeVal)
			}
		}
	case reflect.Struct:
		setValue := reflect.ValueOf(setVal)
		if refValue.Type() == setValue.Type() {
			refValue.Set(setValue)
		}
	case reflect.Ptr:
	case reflect.Uintptr, reflect.Chan, reflect.Func, reflect.UnsafePointer:
	default:
		errTemp = errors.New("unknown type")
	}

	return errTemp
}

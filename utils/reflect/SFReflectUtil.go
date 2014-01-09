//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-8-24
//  Update on 2014-01-06
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//
//	reflect 辅助工具
//
package SFReflectUtil

import (
	"errors"
	"fmt"
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

//
//	基本类型的反射设值
//
//	@param refValue	被设置的反射类型
//	@param setVal	需要设置的值
//
//	@return	error 错误信息
func SetBaseTypeValue(refValue reflect.Value, setVal interface{}) error {
	var errTemp error = nil

	// refValue = reflect.Indirect(refValue)
	if !refValue.CanSet() {
		errTemp = errors.New(fmt.Sprintf("set value(%v) can not be changed", refValue))
		return errTemp
	}

	intBitSize := -1

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

		//----int----------------------------------
	case reflect.Int:
		if -1 >= intBitSize {
			intBitSize = 0
		}
		fallthrough
	case reflect.Int8:
		if -1 >= intBitSize {
			intBitSize = 8
		}
		fallthrough
	case reflect.Int16:
		if -1 >= intBitSize {
			intBitSize = 16
		}
		fallthrough
	case reflect.Int32:
		if -1 >= intBitSize {
			intBitSize = 32
		}
		fallthrough
	case reflect.Int64:
		if -1 >= intBitSize {
			intBitSize = 64
		}
		switch v := setVal.(type) {
		case string:
			if i64, err := strconv.ParseInt(v, 10, intBitSize); err == nil {
				refValue.SetInt(i64)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if i64, err := strconv.ParseInt(v[0], 10, intBitSize); err == nil {
					refValue.SetInt(i64)
				} else {
					errTemp = err
				}
			}
		case int:
			//	不用得fallthrough关键字，只有这样了。
			refValue.SetInt(int64(v))
		case int8:
			refValue.SetInt(int64(v))
		case int16:
			refValue.SetInt(int64(v))
		case int32:
			refValue.SetInt(int64(v))
		case []int:
			if 0 != len(v) {
				refValue.SetInt(int64(v[0]))
			}
		case []int8:
			if 0 != len(v) {
				refValue.SetInt(int64(v[0]))
			}
		case []int16:
			if 0 != len(v) {
				refValue.SetInt(int64(v[0]))
			}
		case []int32:
			if 0 != len(v) {
				refValue.SetInt(int64(v[0]))
			}
		case *int:
			refValue.SetInt(int64(*v))
		case *int8:
			refValue.SetInt(int64(*v))
		case *int16:
			refValue.SetInt(int64(*v))
		case *int32:
			refValue.SetInt(int64(*v))
		case int64:
			refValue.SetInt(v)
		case []int64:
			if 0 != len(v) {
				refValue.SetInt(v[0])
			}
		case *int64:
			refValue.SetInt(*v)
		}

		//----uint----------------------------------
	case reflect.Uint:
		if -1 >= intBitSize {
			intBitSize = 32
		}
		fallthrough
	case reflect.Uint8:
		if -1 >= intBitSize {
			intBitSize = 8
		}
		fallthrough
	case reflect.Uint16:
		if -1 >= intBitSize {
			intBitSize = 16
		}
		fallthrough
	case reflect.Uint32:
		if -1 >= intBitSize {
			intBitSize = 32
		}
		fallthrough
	case reflect.Uint64:
		if -1 >= intBitSize {
			intBitSize = 64
		}
		switch v := setVal.(type) {
		case string:
			if u64, err := strconv.ParseUint(v, 10, intBitSize); err == nil {
				refValue.SetUint(u64)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if u64, err := strconv.ParseUint(v[0], 10, intBitSize); err == nil {
					refValue.SetUint(u64)
				} else {
					errTemp = err
				}
			}
		case uint:
			refValue.SetUint(uint64(v))
		case uint8:
			refValue.SetUint(uint64(v))
		case uint16:
			refValue.SetUint(uint64(v))
		case uint32:
			refValue.SetUint(uint64(v))
		case []uint:
			if 0 != len(v) {
				refValue.SetUint(uint64(v[0]))
			}
		case []uint8:
			if 0 != len(v) {
				refValue.SetUint(uint64(v[0]))
			}
		case []uint16:
			if 0 != len(v) {
				refValue.SetUint(uint64(v[0]))
			}
		case []uint32:
			if 0 != len(v) {
				refValue.SetUint(uint64(v[0]))
			}
		case *uint:
			refValue.SetUint(uint64(*v))
		case *uint8:
			refValue.SetUint(uint64(*v))
		case *uint16:
			refValue.SetUint(uint64(*v))
		case *uint32:
			refValue.SetUint(uint64(*v))
		case uint64:
			refValue.SetUint(v)
		case []uint64:
			if 0 != len(v) {
				refValue.SetUint(v[0])
			}
		case *uint64:
			refValue.SetUint(*v)
		}

		//----float----------------------------------
	case reflect.Float32:
		if -1 >= intBitSize {
			intBitSize = 32
		}
		fallthrough
	case reflect.Float64:
		if -1 >= intBitSize {
			intBitSize = 64
		}
		switch v := setVal.(type) {
		case string:
			if f64, err := strconv.ParseFloat(v, intBitSize); err == nil {
				refValue.SetFloat(f64)
			} else {
				errTemp = err
			}
		case []string:
			if 0 != len(v) {
				if f64, err := strconv.ParseFloat(v[0], intBitSize); err == nil {
					refValue.SetFloat(f64)
				} else {
					errTemp = err
				}
			}
		case float32:
			refValue.SetFloat(float64(v))
		case []float32:
			if 0 != len(v) {
				refValue.SetFloat(float64(v[0]))
			}
		case *float32:
			refValue.SetFloat(float64(*v))
		case float64:
			refValue.SetFloat(v)
		case []float64:
			if 0 != len(v) {
				refValue.SetFloat(v[0])
			}
		case *float64:
			refValue.SetFloat(*v)
		}

		//----string----------------------------------
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
	case reflect.Map:
		setValue := reflect.ValueOf(setVal)
		setValueInd := reflect.Indirect(setValue)

		if refValue.Type() == setValueInd.Type() {
			refValue.Set(setValueInd)
		}

	case reflect.Array, reflect.Slice:
		setValue := reflect.ValueOf(setVal)

		switch setValue.Kind() {
		case reflect.Ptr:
			//	进入此处的可能性
			//	[]int  	*[]int
			//	[]*int 	*[]*int
			//	[]int 	*[]*int
			//	[]int 	*int
			//	[]*int	*[]int
			//	[]*int 	*int
			setValue = setValue.Elem()
			fallthrough
		default:
			//	进入此处的可能性
			//	[]int	[]int
			//	[]*int	[]*int
			//	[]int 	[]*int
			//	[]int 	int
			//	[]*int	[]int
			//	[]*int	int
			refValueType := refValue.Type()
			setValueType := setValue.Type()

			switch setValueType.Kind() {
			case reflect.Slice, reflect.Array:
				if refValueType == setValueType {
					//	[]int	[]int
					//	[]*int	[]*int
					refValue.Set(setValue)
				} else {
					//	[]*int	[]int
					//	[]int 	[]*int
					refValueTypeElem := refValueType.Elem() //	elem = *int int
					setValueTypeElem := setValueType.Elem() //	elem = int *int

					if refValueTypeElem.Kind() == reflect.Ptr && refValueTypeElem.Elem() == setValueTypeElem {
						//	if []*int []int -> int == int
						count := setValue.Len()
						refValue.SetLen(0)
						if 0 != count {
							refValue.Set(reflect.MakeSlice(refValueType, 0, count))
							for i := 0; i < count; i++ {
								sv := setValue.Index(i)
								refValue.Set(reflect.Append(refValue, sv.Addr()))
							}
						}
					} else if setValueTypeElem.Kind() == reflect.Ptr && refValueTypeElem == setValueTypeElem.Elem() {
						//	if []int []*int -> int == int
						count := setValue.Len()
						refValue.SetLen(0)
						if 0 != count {
							refValue.Set(reflect.MakeSlice(refValueType, 0, count))
							for i := 0; i < count; i++ {
								sv := setValue.Index(i)
								refValue.Set(reflect.Append(refValue, sv.Elem()))
							}
						}

					}

				}

			default:
				// []int 	int
				// []*int	int
				refValueTypeElem := refValueType.Elem() //	elem = *int int

				if refValueTypeElem == setValueType {
					//	[]int 	int
					makeVal := reflect.MakeSlice(refValueType, 1, 1)
					makeVal.Index(0).Set(setValue)
					refValue.Set(makeVal)
				} else if refValueTypeElem.Kind() == reflect.Ptr && refValueTypeElem.Elem() == setValueType {
					//	[]*int int

					if setValue.CanAddr() {
						makeVal := reflect.MakeSlice(refValueType, 1, 1)
						makeVal.Index(0).Set(setValue.Addr())
						refValue.Set(makeVal)
					} else {
						//	如果设值的类型结构无法获取到地址则无法进行赋值，可能只直接传递结构的原因，暂时无法解决。
						//
						//	TODO 如果传递的是值类型，则获取不到addr地址信息，则会抛出reflect.Value.Addr of unaddressable value危机
						//	待解决
						// fmt.Println(setValue.Type())
					}

				}

			}
		}

	case reflect.Struct:
		setValue := reflect.ValueOf(setVal)
		switch setValue.Kind() {
		case reflect.Ptr:
			setVElem := setValue.Elem()
			if refValue.Type() == setVElem.Type() {
				refValue.Set(setVElem)
			}
		default:
			if refValue.Type() == setValue.Type() {
				refValue.Set(setValue)
			}
		}

	case reflect.Ptr:
		setValue := reflect.ValueOf(setVal)
		if refValue.Type() == setValue.Type() {
			refValue.Set(setValue)
		} else {
			indRefVel := reflect.Indirect(refValue)
			if reflect.Invalid != indRefVel.Kind() {
				return SetBaseTypeValue(indRefVel, setVal)
			}
		}
	case reflect.Complex64, reflect.Complex128:
	case reflect.Uintptr, reflect.Chan, reflect.Func, reflect.UnsafePointer:
	default:
		errTemp = errors.New("unknown set type")
	}

	return errTemp
}

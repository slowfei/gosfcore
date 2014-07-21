//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-11-01
//  Update on 2014-07-21
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//	debug tool
package SFDebug

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

const (
	HorizeontalTab = "   "
)

//	break
func Break(format bool, v ...interface{}) {

	fmt.Println("Variables:")
	Fdump(os.Stdout, format, v...)

	var pc, file, line, ok = runtime.Caller(1)
	if ok {
		fmt.Printf("\n\n[Debug] %s() [%s:%d]\n[Stack]\n", runtime.FuncForPC(pc).Name(), file, line)
	}

	for i := 2; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc).Name()
		if 0 != len(file) {
			//	/usr/local/go/src/pkg/runtime/proc.c:1223 (0x173d0)
			fmt.Printf("%s(...)\n%s:%d (0x%x)\n", fn, file, line, pc)
		} else {
			// 	runtime.goexit(...)
			// 	L1223: runtime.goexit(...) (0x173d0)
			fmt.Printf("L%d: %s(...) (0x%x)\n", line, fn, pc)
		}
	}

	fmt.Println("\nPress Enter to Continue")
	fmt.Scanln()
	fmt.Println("Break End...")
}

//	打印参数的结构信息
//	@out
//	@v		需要打印的参数
func Fdump(out io.Writer, format bool, v ...interface{}) {

	count := len(v)
	if 0 == count {
		return
	}
	outBuf := bytes.NewBufferString("")
	outBuf.WriteByte('{')
	outBuf.WriteByte('\n')
	for i := 0; i < count; i++ {
		rev := reflect.ValueOf(v[i])
		ts := rev.Type().String()
		outBuf.WriteString(HorizeontalTab + "\"" + strconv.Itoa(i+1) + "." + ts + "\":")
		printReflectValue(format, 2, outBuf, rev)
		if i != count-1 {
			outBuf.WriteByte(',')
		}
		outBuf.WriteByte('\n')
	}
	outBuf.WriteByte('}')
	outBuf.WriteByte('\n')
	outBuf.WriteTo(out)
}

//	直接使用os.Stdout 进行参数结构信息的打印
func Dump(v ...interface{}) {
	Fdump(os.Stdout, true, v...)
}

//	print value
func printReflectValue(format bool, level int, outBuf *bytes.Buffer, v reflect.Value) {
	if !v.IsValid() {
		fmt.Fprint(outBuf, "null")
		return
	}

	switch v.Kind() {
	case reflect.Bool:
		fmt.Fprint(outBuf, v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprint(outBuf, v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprint(outBuf, v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprint(outBuf, v.Float())
	case reflect.String:
		outBuf.WriteString("\"" + v.String() + "\"")
	case reflect.Struct:
		vt := v.Type()
		outBuf.WriteByte('{')
		tab := ""

		fcount := v.NumField()
		for i := 0; i < fcount; i++ {
			ft := vt.Field(i)
			name := ft.Name
			tab = ""
			if format {
				for j := 0; j < level; j++ {
					tab += HorizeontalTab
				}
				outBuf.WriteString("\n" + tab)
			}
			//	格式化包含类型字符串
			// fmt.Fprint(outBuf, "\"("+ft.Type.String()+")"+name+"\":")
			outBuf.WriteString("\"" + name + "\":")
			printReflectValue(format, level+1, outBuf, v.Field(i))
			if i != fcount-1 {
				outBuf.WriteByte(',')
			}

		}
		if format {
			reTabIndex := len(tab) - len(HorizeontalTab)
			if len(tab) < reTabIndex {
				reTabIndex = 0
			}
			if -1 >= reTabIndex {
				reTabIndex = len(tab)
			}
			outBuf.WriteString("\n" + tab[0:reTabIndex] + "}")
		} else {
			outBuf.WriteByte('}')
		}
	case reflect.Map:
		outBuf.WriteByte('{')
		tab := ""

		mapKeys := v.MapKeys()
		kcount := len(mapKeys)
		for i := 0; i < kcount; i++ {
			mv := v.MapIndex(mapKeys[i])
			tab = ""
			if format {
				for j := 0; j < level; j++ {
					tab += HorizeontalTab
				}
				outBuf.WriteString("\n" + tab)
			}
			printReflectValue(format, level+1, outBuf, mapKeys[i])
			outBuf.WriteByte(':')
			printReflectValue(format, level+1, outBuf, mv)
			if i != kcount-1 {
				outBuf.WriteByte(',')
			}

		}
		if format {
			reTabIndex := len(tab) - len(HorizeontalTab)
			if len(tab) < reTabIndex {
				reTabIndex = 0
			}
			if -1 >= reTabIndex {
				reTabIndex = len(tab)
			}
			if 0 == kcount {
				outBuf.WriteByte('}')
			} else {
				outBuf.WriteString("\n" + tab[0:reTabIndex] + "}")
			}
		} else {
			outBuf.WriteByte('}')
		}
	case reflect.Slice:
		if v.IsNil() {
			outBuf.WriteString("null")
			break
		}
		if v.Type().Elem().Kind() == reflect.Uint8 {
			s := v.Bytes()
			outBuf.WriteByte('"')
			if len(s) < 1024 {
				// for small buffers, using Encode directly is much faster.
				dst := make([]byte, base64.StdEncoding.EncodedLen(len(s)))
				base64.StdEncoding.Encode(dst, s)
				outBuf.Write(dst)
			} else {
				// for large buffers, avoid unnecessary extra temporary
				// buffer space.
				enc := base64.NewEncoder(base64.StdEncoding, outBuf)
				enc.Write(s)
				enc.Close()
			}
			outBuf.WriteByte('"')
			break
		}
		// Slices can be marshalled as nil, but otherwise are handled
		// as arrays.
		fallthrough
	case reflect.Array:
		outBuf.WriteByte('[')
		tab := ""
		n := v.Len()
		for i := 0; i < n; i++ {
			tab = ""
			if format {
				for j := 0; j < level; j++ {
					tab += HorizeontalTab
				}
				outBuf.WriteString("\n" + tab)
			}

			printReflectValue(format, level+1, outBuf, v.Index(i))

			if i != n-1 {
				outBuf.WriteByte(',')
			}
		}
		if format {
			reTabIndex := len(tab) - len(HorizeontalTab)
			if len(tab) < reTabIndex {
				reTabIndex = 0
			}
			if -1 >= reTabIndex {
				reTabIndex = len(tab)
			}
			if 0 == n {
				outBuf.WriteByte(']')
			} else {
				outBuf.WriteString("\n" + tab[:reTabIndex] + "]")
			}

		} else {
			outBuf.WriteByte(']')
		}
	case reflect.Ptr:
		if v.IsNil() {
			outBuf.WriteString("null")
			return
		}
		printReflectValue(true, level, outBuf, v.Elem())
	case reflect.Interface:
		if v.IsNil() {
			outBuf.WriteString("null")
			return
		}
		printReflectValue(true, level, outBuf, v.Elem())
	case reflect.Invalid:
		outBuf.WriteString("invalid")
	default:
		outBuf.WriteString("unknow")
	}
	return
}

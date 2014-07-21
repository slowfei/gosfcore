package SFDebug

import (
	"bytes"
	"fmt"
	// "reflect"
	"testing"
)

type T4 struct {
	F7 string
	f8 string
}

type T3 struct {
	F5 string
	f6 string
	T4
}

type T2 struct {
	F3 string
	f4 string
	T3
}

type Tstruct struct {
	F1 string
	F2 string
	T2
	te2 *T2
}

func TestDegub(t *testing.T) {
	fmt.Println("start:...？")
	v := Tstruct{}
	v.F1 = "vf1"
	v.F2 = "vf2"
	Break(false, "v1", "v2", v)

	fmt.Println("end...")
}

//	测试打印信息
func TestPrintValue(t *testing.T) {

	buf := bytes.NewBufferString("")

	v := Tstruct{}
	v.F1 = "vf1"
	v.F2 = "vf2"
	v.T2.F3 = "vf3"
	v.T2.f4 = "vf4"
	v.T3.F5 = "vf5"
	v.T3.f6 = "vf6"
	v.T4.F7 = "vf7"
	v.T4.f8 = "vf8"
	v.te2 = &v.T2

	ss := make([]string, 2)
	ss[0] = "array1"
	ss[1] = "array2"

	si := make([]interface{}, 3)
	si[0] = &v.T2
	si[1] = v.T3
	si[2] = ss

	m3 := make(map[string]interface{})
	m3["k7"] = "temp7"
	m3["k9"] = make([]interface{}, 0)

	m2 := make(map[string]interface{})
	m2["k6"] = "temp"
	m2["km3"] = m3

	m := make(map[string]interface{})
	m["k133"] = v.T2
	// m["k2"] = v.T3
	// m["k3"] = v.T4
	// m["k5"] = "12"
	// m["km2"] = m2
	// m["ka6"] = ss
	// m["ks7"] = si

	// printReflectValue(true, 1, buf, reflect.ValueOf(m))

	Dump(m, si, m3)

	fmt.Println(buf.String())
}

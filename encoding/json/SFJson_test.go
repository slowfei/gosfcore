package SFJson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type StructType struct {
	TypeName string
	TypeUUID string
}

type StructJson struct {
	Name     string
	Sex      int
	Money    float64
	Type     StructType
	Married  bool
	ArrayMap []map[string]string
	Array    []string
	UID      *struct {
		UIDStr string
	}
}

func TestValidateMap(t *testing.T) {

	jsonObject := `
		{
			"Name"	: "slowfei",
			"Sex"	: 1,
			"Money"	: 88888.888,
			"Type"	: {
				"TypeName"	: "sl_type",
				"TypeUUID"	: "sl_typeuuid"
			},
			"Married":false,
			"ArrayMap":[
				{ "AKey_1" : "AValue_1" },
				{ "AKey_2" : "AValue_2" }
			],
			"Array":[ "arr_1", "arr_2", "arr_3"],
			"UID":null,
			"Mapnil":{}
		}`

	// jsonArray := `
	// ["Name1","Name2"]
	// `

	var object interface{} = nil
	err := Unmarshal([]byte(jsonObject), &object)
	if nil != err {
		t.Error(err)
		return
	}

	result := ValidateMap(object, map[string]interface{}{
		"Name":  TypeNotnilString,
		"Sex":   TypeFloat,
		"Money": TypeFloat,
		"Type": map[string]interface{}{
			"TypeName": TypeNotnilString,
			"TypeUUID": TypeString,
		},
		"Married": TypeBool,
		"ArrayMap": []interface{}{
			map[string]interface{}{"AKey_1": TypeNotnilString},
			map[string]interface{}{"AKey_2": TypeString},
		},
		"Array": []DataTypeNotnil{
			TypeNotnilString,
		},
		"UID":    TypeString,
		"Mapnil": TypeMap,
	})

	if !result {
		t.Fatal("json validate fatal.")
	}
}

//	测试创建空json对象
func TestJsonNewNil(t *testing.T) {
	arrayNil := NewJsonNil(true)
	if arrayNil.String() != "[]" {
		t.Fatal("new array json nil != []")
		return
	}

	mapNil := NewJsonNil(false)
	if mapNil.String() != "{}" {
		t.Fatal("new map json nil != {}")
		return
	}

}

func TestMarshal(t *testing.T) {
	var sj StructJson
	sj.Name = "slowfei"
	sj.Sex = 25
	sj.Money = 88888.888
	sj.Type.TypeName = "sl_type"
	sj.Type.TypeUUID = "sl_typeuuid"
	sj.Married = false
	sj.ArrayMap = make([]map[string]string, 2, 2)
	sj.ArrayMap[0] = map[string]string{"AKey_1": "AValue_1"}
	sj.ArrayMap[1] = map[string]string{"AKey_2": "AValue_2"}
	sj.Array = []string{"arr_1", "arr_2", "arr_3"}

	jb, err := Marshal(sj, "null", "true|false")
	if nil != err {
		fmt.Println(err)
	}
	if string(jb) != `{"Name":"slowfei","Sex":25,"Money":88888.888,"Type":{"TypeName":"sl_type","TypeUUID":"sl_typeuuid"},"Married":false,"ArrayMap":[{"AKey_1":"AValue_1"},{"AKey_2":"AValue_2"}],"Array":["arr_1","arr_2","arr_3"],"UID":null}` {
		t.Fail()
	}
	// fmt.Println(string(jb))

}

func TestJson(t *testing.T) {
	var sj StructJson
	sj.Name = "slowfei"
	sj.Sex = 25
	sj.Money = 88888.888
	sj.Type.TypeName = "sl_type"
	sj.Type.TypeUUID = "sl_typeuuid"
	sj.Married = false
	sj.ArrayMap = make([]map[string]string, 2, 2)
	sj.ArrayMap[0] = map[string]string{"AKey_1": "AValue_1"}
	sj.ArrayMap[1] = map[string]string{"AKey_2": "AValue_2"}
	sj.Array = []string{"arr_1", "arr_2", "arr_3"}

	json, _ := NewJson(sj, "null", "")
	if json.String() != `{"Name":"slowfei","Sex":25,"Money":88888.888,"Type":{"TypeName":"sl_type","TypeUUID":"sl_typeuuid"},"Married":false,"ArrayMap":[{"AKey_1":"AValue_1"},{"AKey_2":"AValue_2"}],"Array":["arr_1","arr_2","arr_3"],"UID":null}` {
		t.Fail()
	}
	// fmt.Println(json.String())

}

func Benchmark_SFMarshal(b *testing.B) {
	var sj StructJson
	sj.Name = "slowfei"
	sj.Sex = 25
	sj.Money = 88888.888
	sj.Type.TypeName = "sl_type"
	sj.Type.TypeUUID = "sl_typeuuid"
	sj.Married = false
	sj.ArrayMap = make([]map[string]string, 2, 2)
	sj.ArrayMap[0] = map[string]string{"AKey_1": "AValue_1", "AKey_1_1": "AValue_1_1"}
	sj.ArrayMap[1] = map[string]string{"AKey_2": "AValue_2", "AKey_2_2": "AValue_2_2"}
	sj.Array = []string{"arr_1", "arr_2", "arr_3"}

	for i := 0; i < b.N; i++ {
		_, err := Marshal(sj, "null", "")
		if nil != err {
			break
		}
	}
}
func Benchmark_GolangMarshal(b *testing.B) {
	var sj StructJson
	sj.Name = "slowfei"
	sj.Sex = 25
	sj.Money = 88888.888
	sj.Type.TypeName = "sl_type"
	sj.Type.TypeUUID = "sl_typeuuid"
	sj.Married = false
	sj.ArrayMap = make([]map[string]string, 2, 2)
	sj.ArrayMap[0] = map[string]string{"AKey_1": "AValue_1", "AKey_1_1": "AValue_1_1"}
	sj.ArrayMap[1] = map[string]string{"AKey_2": "AValue_2", "AKey_2_2": "AValue_2_2"}
	sj.Array = []string{"arr_1", "arr_2", "arr_3"}

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(sj)
		if nil != err {
			break
		}
	}
}

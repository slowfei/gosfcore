package SFJson

import (
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

func TestMarshal(t *testing.T) {
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

	jb, err := Marshal(sj, "null", "true|false")
	if nil != err {
		fmt.Println(err)
	}
	if string(jb) != `{"Name":"slowfei","Sex":25,"Money":88888.888,"Type":{"TypeName":"sl_type","TypeUUID":"sl_typeuuid"},"Married":false,"ArrayMap":[{"AKey_1":"AValue_1","AKey_1_1":"AValue_1_1"},{"AKey_2":"AValue_2","AKey_2_2":"AValue_2_2"}],"Array":["arr_1","arr_2","arr_3"],"UID":null}` {
		t.Fail()
	}
	fmt.Println(string(jb))

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
	sj.ArrayMap[0] = map[string]string{"AKey_1": "AValue_1", "AKey_1_1": "AValue_1_1"}
	sj.ArrayMap[1] = map[string]string{"AKey_2": "AValue_2", "AKey_2_2": "AValue_2_2"}
	sj.Array = []string{"arr_1", "arr_2", "arr_3"}

	json, _ := NewJson(sj, "null", "")
	if json.String() != `{"Name":"slowfei","Sex":25,"Money":88888.888,"Type":{"TypeName":"sl_type","TypeUUID":"sl_typeuuid"},"Married":false,"ArrayMap":[{"AKey_1":"AValue_1","AKey_1_1":"AValue_1_1"},{"AKey_2":"AValue_2","AKey_2_2":"AValue_2_2"}],"Array":["arr_1","arr_2","arr_3"],"UID":null}` {
		t.Fail()
	}
	fmt.Println(json.String())

}

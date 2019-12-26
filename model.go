package gson

import (
	"reflect"
	"strings"
	"encoding/json"
	"fmt"
)

const (
	JsonTypeString  = "String"
	JsonTypeNumber  = "Number"
	JsonTypeArray   = "Array"
	JsonTypeBoolean = "Boolean"
	JsonTypeStruct  = "Struct"
)

type JsonObj struct {
	Type        string
	value       interface{}
	arrayValue  []*JsonObj
	structValue map[string]*JsonObj
}
func (obj *JsonObj) GetInt(key string) int {
	return 0
}
func (obj *JsonObj) Get(key string) *JsonObj {
	val, ok := obj.structValue[key]
	if !ok {
		return nil
	}
	return val
}

func (obj *JsonObj) GetItem(i int) *JsonObj {
	if obj.Type != JsonTypeArray {
		panic("iilggbxhcn not Array JsonObj")
	}
	if i > len(obj.arrayValue)-1 {
		panic("iilggbxhcn index out range")
	}
	return nil
}

func NewJsonObj(obj interface{}) *JsonObj {
	result := &JsonObj{}
	kind := reflect.TypeOf(obj).Kind()
	switch kind {
	case reflect.String:
		str := obj.(string)
		if strings.HasPrefix(str, "[") {
			result.Type = JsonTypeArray
			//todo
			err := json.Unmarshal([]byte(str), &result.arrayValue)
			if err != nil {
				fmt.Println("z3mb4ixchn", "NewJsonObj", err)
				return nil
			}

		} else if strings.HasPrefix(str, "{") {
			result.Type = JsonTypeStruct
			//todo
			err := json.Unmarshal([]byte(str), &result.structValue)
			if err != nil {
				fmt.Println("0hsxpy1qr5", "NewJsonObj", err)
				return nil
			}
		} else { // if string not start with [ or { ,recognized as string
			result.Type = JsonTypeString
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		result.Type = JsonTypeNumber
	case reflect.Bool:
		result.Type = JsonTypeBoolean
	}
	result.value = obj
	return result
}

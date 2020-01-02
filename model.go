package gson

import (
	"reflect"
	"fmt"
	"strings"
	"encoding/json"
	"strconv"
)

const (
	JsonTypeString  = "String"
	JsonTypeNumber  = "Number"
	JsonTypeArray   = "Array"
	JsonTypeBoolean = "Boolean"
	JsonTypeStruct  = "Struct"
)

var isDebug bool = false

type JsonObj struct {
	Type        string
	value       interface{}
	structValue map[string]*JsonObj
	arrayValue  []*JsonObj
}

func (obj *JsonObj) Int() (int, bool) {
	intValue, ok := obj.value.(int)
	if !ok && isDebug {
		fmt.Println("obj not int")
	}
	return intValue, ok
}
func (obj *JsonObj) Float64() (float64, bool) {
	floatValue, ok := obj.value.(float64)
	if !ok && isDebug {
		fmt.Println("obj not float64")
	}
	return floatValue, ok
}
func (obj *JsonObj) String() (value string, ok bool) {
	value, ok = obj.value.(string)
	if !ok && isDebug {
		fmt.Println("obj not string")
	}
	return
}
func (obj *JsonObj) Bool() (value bool, ok bool) {
	value, ok = obj.value.(bool)
	if !ok && isDebug {
		fmt.Println("obj not bool")
	}
	return
}
func (obj *JsonObj) Get(key string) (value *JsonObj) {
	return obj.structValue[key]
}
func (obj *JsonObj) GetInt(key string) (value int, ok bool) {
	sObj := obj.Get(key)
	if sObj == nil {
		if isDebug {
			fmt.Println("key", key, "not found")
		}
		return
	}
	return sObj.Int()
}
func (obj *JsonObj) GetFloat64(key string) (value float64, ok bool) {
	sObj := obj.Get(key)
	if sObj == nil {
		if isDebug {
			fmt.Println("key", key, "not found")
		}
		return
	}
	return sObj.Float64()
}
func (obj *JsonObj) GetString(key string) (value string, ok bool) {
	sObj := obj.Get(key)
	if sObj == nil {
		if isDebug {
			fmt.Println("key", key, "not found")
		}
		return
	}
	return sObj.String()
}
func (obj *JsonObj) GetBool(key string) (value bool, ok bool) {
	sObj := obj.Get(key)
	if sObj == nil {
		if isDebug {
			fmt.Println("key", key, "not found")
		}
		return
	}
	return sObj.Bool()
}
func (obj *JsonObj) GetArray(key string) []*JsonObj {
	return obj.arrayValue
}
func (obj *JsonObj) GetItem(index int) *JsonObj {
	if obj.Type != JsonTypeArray {
		panic("not Array JsonObj")
	}
	if index > len(obj.arrayValue)-1 {
		panic("index out range")
	}
	return obj.arrayValue[index]
}
func (obj *JsonObj) Interface() interface{} {
	return obj.value
}
func (obj *JsonObj) ToJsonString() string {
	switch obj.Type {
	case JsonTypeString:
		return obj.value.(string)
	case JsonTypeNumber:
		//todo int and float
		//return fmt.Sprint(obj.value)
	case JsonTypeBoolean:
		return strconv.FormatBool(obj.value.(bool))
	case JsonTypeArray:
		//todo
	case JsonTypeStruct:
		//todo
	}
	return ""
}
func NewJsonObj(obj interface{}) *JsonObj {
	result := &JsonObj{}
	kind := reflect.TypeOf(obj).Kind()
	switch kind {
	case reflect.String:
		str := obj.(string)
		if strings.HasPrefix(str, "[") {
			result.Type = JsonTypeArray
			var tmpSlice []interface{}
			err := json.Unmarshal([]byte(str), &tmpSlice)
			if err != nil {
				fmt.Println("z3mb4ixchn", "NewJsonObj", err)
				return nil
			}
			for _, one := range tmpSlice {
				result.arrayValue = append(result.arrayValue, NewJsonObj(one))
			}
		} else if strings.HasPrefix(str, "{") {
			result.Type = JsonTypeStruct
			result.structValue= map[string]*JsonObj{}
			var tmpObj map[string]interface{}
			err := json.Unmarshal([]byte(str), &tmpObj)
			if err != nil {
				fmt.Println("0hsxpy1qr5", "NewJsonObj", err)
				return nil
			}
			for key, one := range tmpObj {
				result.structValue[key] = NewJsonObj(one)
			}
		} else { // if string not start with [ or { ,recognized as string
			result.Type = JsonTypeString
		}
		result.value = obj
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result.Type = JsonTypeNumber
		val, _ := strconv.Atoi(fmt.Sprintln(obj))
		result.value = val
	case reflect.Float32, reflect.Float64:
		result.Type = JsonTypeNumber
		val, _ := strconv.ParseFloat(fmt.Sprintln(obj), 64)
		result.value = val
	case reflect.Bool:
		result.Type = JsonTypeBoolean
		result.value = obj
	case reflect.Interface:
		fmt.Println("????????")
	default:
		//todo obj to jsonObj
		fmt.Println("un suupprot")
	}
	return result
}

func SetDebugModule() {
	isDebug = true
}
func SetReleaseModule() {
	isDebug = false
}

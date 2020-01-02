package gson

import (
	"fmt"
	"strconv"
)

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
		log("key", key, "not found")
		return
	}
	return sObj.Int()
}
func (obj *JsonObj) GetFloat64(key string) (value float64, ok bool) {
	sObj := obj.Get(key)
	if sObj == nil {
		log("key", key, "not found")
		return
	}
	return sObj.Float64()
}
func (obj *JsonObj) GetString(key string) (value string, ok bool) {
	sObj := obj.Get(key)
	if sObj == nil {
		log("key", key, "not found")
		return
	}
	return sObj.String()
}
func (obj *JsonObj) GetBool(key string) (value bool, ok bool) {
	sObj := obj.Get(key)
	if sObj == nil {
		log("key", key, "not found")
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
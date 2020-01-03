package gson

import (
	"fmt"
	"reflect"
	"encoding/json"
)

func (obj *JsonObj) Int() (int, bool) {
	var intValue int
	float64Value, ok := obj.value.(float64)
	if !ok {
		intValue, ok = obj.value.(int)
		if !ok && isDebug {
			fmt.Println("obj not int", reflect.TypeOf(obj.value))
		}
	} else {
		intValue = int(float64Value)
	}
	return intValue, ok
}
func (obj *JsonObj) Float64() (float64, bool) {
	floatValue, ok := obj.value.(float64)
	if !ok && isDebug {
		fmt.Println("obj not float64", reflect.TypeOf(obj.value))
	}
	return floatValue, ok
}
func (obj *JsonObj) String() (value string, ok bool) {
	value, ok = obj.value.(string)
	if !ok && isDebug {
		fmt.Println("obj not string", reflect.TypeOf(obj.value))
	}
	return
}
func (obj *JsonObj) Bool() (value bool, ok bool) {
	value, ok = obj.value.(bool)
	if !ok && isDebug {
		fmt.Println("obj not bool", reflect.TypeOf(obj.value))
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
	sObj := obj.Get(key)
	if sObj == nil {
		log("key", key, "not found")
		return nil
	}
	return sObj.arrayValue
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
	bSlice, err := json.Marshal(obj.value)
	if err != nil {
		log("JsonObj.ToJsonString", "json.Marshal", err)
	}
	//todo kill alloc
	return string(bSlice)
}


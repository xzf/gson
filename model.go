package gson

/*
 * auth :xzf
 * email:xzf12315@gmail.com
 */

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


type JsonObj struct {
	Type        string
	value       interface{}
	structValue map[string]*JsonObj
	arrayValue  []*JsonObj
}


func NewJsonObj(obj interface{}) *JsonObj {
	result := &JsonObj{
		value: obj,
	}
	kind := reflect.TypeOf(obj).Kind()
	switch kind {
	case reflect.String:
		str := obj.(string)
		// json string for struct
		if strings.HasPrefix(str, "{") {
			result.Type = JsonTypeStruct
			result.structValue = map[string]*JsonObj{}
			var tmpObj map[string]interface{}
			//todo find a way to kill this json.Unmarshal alloc
			err := json.Unmarshal([]byte(str), &tmpObj)
			if err != nil {
				log("NewJsonObj", JsonTypeStruct, obj, err)
				return nil
			}
			for key, one := range tmpObj {
				result.structValue[key] = NewJsonObj(one)
			}
			return result
		}
		// json string for array
		if strings.HasPrefix(str, "[") {
			result.Type = JsonTypeArray
			var tmpSlice []interface{}
			//todo find a way to kill this json.Unmarshal alloc
			err := json.Unmarshal([]byte(str), &tmpSlice)
			if err != nil {
				log("NewJsonObj", JsonTypeArray, obj, err)
				return nil
			}
			for _, one := range tmpSlice {
				result.arrayValue = append(result.arrayValue, NewJsonObj(one))
			}
			return result
		}
		// if string not start with [ or {, recognized as string
		result.Type = JsonTypeString
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result.Type = JsonTypeNumber
		//todo find a way to kill this fmt.Sprint alloc
		val, _ := strconv.Atoi(fmt.Sprint(obj)) //all type cast to int
		result.value = val
	case reflect.Float32, reflect.Float64: //Float32 cast to Float64
		result.Type = JsonTypeNumber
		//todo find a way to kill this fmt.Sprint alloc
		val, _ := strconv.ParseFloat(fmt.Sprint(obj), 64)
		result.value = val
	case reflect.Bool:
		result.Type = JsonTypeBoolean
		result.value = obj
	//case reflect.Interface:
	//	fmt.Println("????????")
	default:
		result.value = obj
		//todo sorry,This solution is so stupid
		str, err := json.Marshal(obj)
		if err != nil {
			log("NewJsonObj", reflect.Array)
			return nil
		}
		return NewJsonObj(string(str))
	}
	return result
}

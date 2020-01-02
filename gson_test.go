package gson

import (
	"testing"
	"fmt"
)

func TestGson(t *testing.T) {
	SetDebugModule()
	testSlice := []interface{}{
		`{
			"D":[1,2,3]
		}`,
		`{
			"A":1,
			"B":"2",
			"C":false,
			"D":[1,2,3],
			"E":{"E1":1,"E2":"2","E3":true}
		}`,
		struct {
			A int
			B string
			C bool
			D []int
			E struct {
				E1 int
				E2 string
				E3 bool
			}
		}{
			A: 1,
			B: "2",
			C: false,
			D: []int{1, 2, 3},
			E: struct {
				E1 int
				E2 string
				E3 bool
			}{E1: 1, E2: "2", E3: true},
		},
		1,
		"2",
		false,
		"[1,2,3]",
		`[
	{"A":1,"B":"2","C":false},
	{"A":2,"B":"3","C":true},
	{"A":3,"B":"4","C":false}
]`,
	}
	for index, one := range testSlice {
		obj := NewJsonObj(one)
		if obj == nil {
			fmt.Println(index, one)
			panic("no pass")
		}
		switch index {
		case 0:
			//a, _ := obj.Get("A").Int()
			//panicIfFalse(obj.Get("A").Type == JsonTypeNumber)
			//panicIfFalse(a == 1)
			//b, _ := obj.Get("B").String()
			//panicIfFalse(obj.Get("B").Type == JsonTypeString)
			//panicIfFalse(b == "2")
			//c, ok := obj.Get("C").Bool()
			//panicIfFalse(obj.Get("C").Type == JsonTypeBoolean)
			//panicIfFalse(c == false && ok == true)
			array := obj.GetArray("D")
			panicIfFalse(obj.Get("D").Type == JsonTypeArray)
			fmt.Println(len(array),array[0].value)
			fmt.Println(len(array))
			fmt.Println(len(array))
			fmt.Println(len(array))
			fmt.Println(len(array))
			fmt.Println(len(array))
			fmt.Println(len(array))
			fmt.Println(len(array))
			panicIfFalse(len(array) == 3)
			d0, _ := array[0].Int()
			panicIfFalse(d0 == 1)
			d1, _ := array[1].Int()
			panicIfFalse(d1 == 2)
			d2, _ := array[2].Int()
			panicIfFalse(d2 == 3)
			panicIfFalse(obj.Get("E") != nil)
			panicIfFalse(obj.Get("E").Type == JsonTypeStruct)
			panicIfFalse(obj.Get("E").Get("E1").Type == JsonTypeNumber)
			panicIfFalse(obj.Get("E").Get("E1") != nil)
			ee1, _ := obj.Get("E").Get("E1").Int()
			panicIfFalse(ee1 == 1)

			panicIfFalse(obj.Get("E").Get("E2").Type == JsonTypeString)
			panicIfFalse(obj.Get("E").Get("E2") != nil)
			ee2, _ := obj.Get("E").Get("E2").String()
			panicIfFalse(ee2 == "2")
			panicIfFalse(obj.Get("E").Get("E3").Type == JsonTypeBoolean)
			panicIfFalse(obj.Get("E").Get("E3") != nil)
			ee3, _ := obj.Get("E").Get("E3").Bool()
			panicIfFalse(ee3 == true)
		case 1:
		case 2:
		case 3:
		case 4:
		case 5:
		case 6:
		}
	}
}
func panicIfFalse(isNoPanic bool) {
	if !isNoPanic {
		panic("no pass")
	}
}

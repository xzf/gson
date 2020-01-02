package gson

import "fmt"

func log(slice ...interface{}) {
	if isDebug {
		fmt.Println(slice...)
	}
}

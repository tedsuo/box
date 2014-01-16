package box

import (
	"fmt"
	"reflect"
	"strconv"
)

func Equals(item1, item2 interface{}) bool {
	return reflect.DeepEqual(item1, item2)
}

func Int(input interface{}) int {
	switch num := input.(type) {
	case int:
		return num
	case int32:
		return int(num)
	case int64:
		return int(num)
	case string:
		intVal, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		return intVal
	default:
		panic(fmt.Sprintf("unable to convert %s to int", num))
	}
}

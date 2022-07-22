package list

import (
	"fmt"
	"reflect"
)

func RemoveIndex(l []interface{}, i int) ListResult {
	elementType := reflect.TypeOf(l).Elem()
	fmt.Println(elementType)

	return ListResult{}
}

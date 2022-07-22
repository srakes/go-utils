package helper

import (
	"strconv"
	"strings"
)

func RemoveIndexFromInts(l []int, i int) []int {
	if len(l) == 1 {
		return []int{}
	}
	var newL []int
	newL = append(newL, l[:i]...)
	newL = append(newL, l[i+1:]...)

	return newL
}

func RemoveIndexFromStrings(l []string, i int) []string {
	if len(l) == 1 {
		return []string{}
	}

	var newL []string
	newL = append(newL, l[:i]...)
	newL = append(newL, l[i+1:]...)

	return newL
}

func RemoveString(l []string, w string, t int) []string {
	if t == 0 {
		return l
	}

	for i, word := range l {
		if word == w {
			var result []string
			result = append(result, RemoveIndexFromStrings(l, i)...)
			if t == -1 {
				return RemoveString(result, w, -1)
			} else {
				return RemoveString(result, w, t-1)
			}
		}
	}

	return l
}

func RemoveLetterFromString(w, l string, t int) string {
	ss := strings.Split(w, "")
	return strings.Join(RemoveString(ss, l, t), "")
}

func RemoveIntFromString(s string) string {
	ss := strings.Split(s, "")
	var result []string
	for i, l := range ss {
		_, err := strconv.Atoi(l)
		if err != nil {
			continue
		} else {
			result = append(result, ss[:i]...)
			result = append(result, ss[i+1:]...)
			return strings.Join(result, "")
		}
	}

	return s
}

// todo: maybe turn this into a function that will just copy the slice
// func getSliceType(s interface{}) reflect.Type {
// 	if reflect.TypeOf(s).Kind() != reflect.Slice {
// 		fmt.Println("Please provide a slice")
// 		return nil
// 	}
// 	elementType := reflect.TypeOf(s).Elem()
// 	v := reflect.ValueOf(s)
// 	ll := v.Len()
// 	lc := (v.Cap() + 1) * 2
// 	newL := reflect.MakeSlice(reflect.SliceOf(elementType), ll, lc)
// 	reflect.Copy(newL, v)

// 	return elementType
// }

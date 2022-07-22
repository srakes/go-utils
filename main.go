package main

import (
	"fmt"

	"github.com/srakes/go-utils/pkg/helper"
)

func main() {
	t := []int{1, 2}
	fmt.Println(helper.RemoveIndexFromInts(t, 0))

	s := []string{"test", "scott", "test", "rakes", "test"}
	fmt.Println(helper.RemoveString(s, "test", -1))

	m := "helpp"
	fmt.Println(helper.RemoveLetterFromString(m, "p", 1))

	z := "ab0c4"
	fmt.Println(helper.RemoveIntFromString(z))
}

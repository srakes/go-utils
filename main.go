package main

import (
	"fmt"
)

func main() {
	t := []int{1, 2, 3, 4, 5}
	newL := srList.RemoveIndex(t, 1)
	fmt.Println(newL)
}

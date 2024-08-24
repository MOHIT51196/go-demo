package main

import (
	"fmt"
	"packages/search"
)

func main() {
	var num = 22
	var list = []int{-1, 0, 3, 5, 11, 21, 22, 41, 99924}
	fmt.Printf("Binary serach (%d) = %d\n", num, search.BinarySearch(list, num))
}

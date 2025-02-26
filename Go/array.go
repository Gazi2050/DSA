package main

import (
	"fmt"
	"slices"
)

var arr [5]int
var arr2 = [5]int{1, 2, 3, 4, 5}
var arr3 = [...]string{"a", "b", "c"}
var arr4 = [...]bool{true}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [...]int{10, 20, 30}
	a[4] = 99

	fmt.Println(a[4])
	fmt.Println(b)
	fmt.Println(arr)
	fmt.Println(arr2)
	// Convert arr3 (array) to a slice=> arr3[:] before using slices.Index
	index := slices.Index(arr3[:], "b")
	fmt.Println(index)
	fmt.Println(len(arr4))
}

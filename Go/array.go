package main

import "fmt"

//In Go, arrays cannot be constants. The const keyword in Go is used for compile-time constants, and arrays are not allowed to be constants because their size or elements can vary at runtime.Arrays hold values that can be modified, so Go does not allow arrays to be constants.
var arr [5]int
var arr2 = [5]int{1, 2, 3, 4, 5}
var arr3 = [...]string{"a", "b", "c"}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [...]int{10, 20, 30}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

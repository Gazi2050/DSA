package main

import "fmt"

var arr [5]int
var arr2 = [...]int{1, 2, 3, 4, 5}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [...]int{10, 20, 30}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(arr)
	fmt.Println(arr2)
}

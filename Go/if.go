package main

import "fmt"

func main() {
	x := 1
	if x < 10 {
		fmt.Println(true)
	} else if x > 10 {
		fmt.Println("working")
	} else {
		fmt.Println(false)
	}
}

package main

import "fmt"

//Go has only one looping construct: the for loop. Unlike other languages that have while and do-while, Go simplifies iteration by handling all looping needs with for.

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
}

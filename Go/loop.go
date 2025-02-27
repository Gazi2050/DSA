package main

import "fmt"

//Go has only one looping construct: the for loop. Unlike other languages that have while and do-while, Go simplifies iteration by handling all looping needs with for.

func main() {
	/*normal for loop*/
	// for i := 1; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	/*Infinite loop*/
	// for {
	// 	fmt.Println("this is Infinite loop")
	// }

	/*for as a while Loop*/
	// j := 1
	// for j < 5 {
	// 	fmt.Println(j)
	// 	j++
	// }

	/*for as a do-while Loop*/
	k := 1
	for {
		fmt.Println("count:", k)
		k++
		if k > 10 {
			break
		}
	}
}

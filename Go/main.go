package main

import "fmt"

/*variable*/
// Global variable declaration using `var` (Can be changed)
var number = 100
var singleCharacter = 'G'
var name = "gazi"
var bool = true

// Constant declaration using `const` (Cannot be changed)
const id = 9

/*Types*/

func main() {
	// Local variable declaration using `:=` (Shorthand syntax, inferred type)
	funcVar := "hello world"
	// Printing values
	fmt.Println(number)
	fmt.Println(singleCharacter)
	fmt.Println(name)
	fmt.Println(bool)
	fmt.Println(id)
	fmt.Println(funcVar)
}

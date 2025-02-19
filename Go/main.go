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
var str1 = "hello" //inferred type(dynamic type)
//static type
var str2 string = "world"
var num int = 6
var num2 int8 = 99

var one, two, three int = 1, 2, 3 //many variabe in one variable keyword

func main() {
	// Local variable declaration using `:=` (Shorthand syntax, inferred type)
	funcVar := "hello world"
	var ()
	// Printing values
	fmt.Println(number)
	fmt.Println(singleCharacter)
	fmt.Println(name)
	fmt.Println(bool)
	fmt.Println(id)
	fmt.Println(funcVar)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
}

package main

import "fmt"

/*variable*/
// Global variable declaration using `var` (Can be changed)
var number = 100
var singleCharacter = 'G'
var name = "gazi"
var bool1 = true

// Constant declaration using `const` (Cannot be changed)
const id = 9

/*Types*/
//inferred type(dynamic type)
var str1 = "hello"
var num = 100
var isTrue = true

//static type
var str2 string = "world"
var num1 int = 6
var num2 int8 = 99
var num3 float64 = 77.7

var one, two, three int = 1, 2, 3 //many variabe in one variable keyword
// group variable
var (
	num4    int
	str3    string
	isFalse bool
)

func main() {
	// Local variable declaration using `:=` (Shorthand syntax, inferred type)
	// can not stay unused a variable in function
	funcString := "hello world"
	funcNumber := 99
	funcBool := true
	var ()
	// Printing values (print in block line)
	fmt.Println(funcString)
	fmt.Println(funcNumber)
	//print in same line
	fmt.Print(funcBool)
	fmt.Printf(funcString) // only print sting

}

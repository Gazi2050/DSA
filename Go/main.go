package main

import "fmt"

/* Variable Declarations */
// Global variables (Can be changed)
var number = 100
var singleCharacter = 'G'
var name = "gazi"
var bool1 = true

// Constant (Cannot be changed)
const id = 9

/* Type Declarations */
// Inferred type (dynamic typing)
var str1 = "hello"
var num = 100
var isTrue = true

// Explicit (static typing)
var str2 string = "world"
var num1 int = 6
var num2 int8 = 99
var num3 float64 = 77.7
var isNum bool = true

// Multiple variables in one statement
var one, two, three int = 1, 2, 3

// Grouped variable declaration
var (
	num4    int
	str3    string
	isFalse bool
)

/* Operators */
// the type of keyword (var or const) used for values must match the keyword type used for the result to avoid errors in Go.
var value1 int = 385
var value2 int = 45
var result int = value1 + value2
var result2 int = value1 - value2
var result3 int = value1 * value2
var result4 int = value1 / value2

const value4 int = 350
const value5 int = 45
const result5 int = value4 + value5

func main() {
	// cannot use the shorthand operator (:=,+=,-=,*=,%=,++,--) outside of a function in Go.
	funcNumber := 99
	funcNumber++
	funcNumber--
	num := 20
	num += 10
	num -= 5
	num *= 2
	num /= 5
	num %= 3
	fmt.Println(num)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(funcNumber)
	// Local variable declaration using shorthand `:=`
	// funcString := "hello world"
	// funcBool := true
	// Printing values
	// fmt.Println(funcString) // Print with a new line
	// fmt.Println(funcNumber)
	// Print in the same line
	// fmt.Print(funcBool)
	//Print in the same line in Printf()
	// fmt.Printf("%s", funcString)  // ✅ Correct (prints a string)
	// fmt.Printf("%2d", funcNumber) // ✅ Corrected (uses %d for integers)
	// fmt.Printf("%t", funcBool)    // ✅ Corrected (uses %t for boolean values)

}

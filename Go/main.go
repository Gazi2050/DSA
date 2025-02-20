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

func main() {
	// Local variable declaration using shorthand `:=`
	funcString := "hello world"
	funcNumber := 99
	funcBool := true

	// Printing values
	fmt.Println(funcString) // Print with a new line
	fmt.Println(funcNumber)

	// Print in the same line
	fmt.Print(funcBool)
	//Print in the same line in Printf()
	fmt.Printf("%s", funcString)  // ✅ Correct (prints a string)
	fmt.Printf("%2d", funcNumber) // ✅ Corrected (uses %d for integers)
	fmt.Printf("%t", funcBool)    // ✅ Corrected (uses %t for boolean values)

}

package main

import "fmt"

func main() {
	// Logical operators
	a, b, c := true, false, true

	// AND (&&)
	fmt.Println("a && b:", a && b) // false
	fmt.Println("a && c:", a && c) // true

	// OR (||)
	fmt.Println("a || b:", a || b) // true
	fmt.Println("b || c:", b || c) // true

	// NOT (!)
	fmt.Println("!a:", !a) // false
	fmt.Println("!b:", !b) // true

	// Comparison operators
	x, y := 10, 5
	z := 10

	// Greater than (>)
	fmt.Println("x > y:", x > y) // true

	// Less than (<)
	fmt.Println("x < y:", x < y) // false

	// Equal to (==)
	fmt.Println("x == y:", x == y) // false
	fmt.Println("x == z:", x == z) // true

	// Not equal to (!=)
	fmt.Println("x != y:", x != y) // true

	// Strict equal (===) does not exist in Go, use `==`
	// Go comparison is simple with `==` for primitives and struct types, no `===`.

	// Combining logical and comparison operators
	fmt.Println("(x > y) && (a || b):", (x > y) && (a || b))   // true
	fmt.Println("(x < y) || !(a && c):", (x < y) || !(a && c)) // false

	// Complex condition with multiple comparisons
	fmt.Println("(x > y) && (x == z):", (x > y) && (x == z)) // true
}

package main

import "fmt"

func user(name string) {
	fmt.Println("name:", name)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	user("gazi")
	fmt.Println(add(2, 5))
}

package main

import "fmt"

//In Go, there is no concept of "objects" in the same way as object-oriented languages.Go supports struct types, which can be used to define structured data types similar to objects, along with methods to operate on them.

//In Go, we can define methods on structs. Unlike object-oriented languages where methods belong to objects, in Go, methods are simply functions with a receiver.

type User struct {
	name string
	age  int
}

func main() {
	o1 := User{name: "gazi", age: 23}
	fmt.Println(o1)
}

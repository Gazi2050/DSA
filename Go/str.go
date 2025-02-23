package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello world"
	// var LongStr string = "A Long String"
	var strIndex = "String"
	fmt.Printf("%c", strIndex[0])
	fmt.Println(len(str))
	fmt.Println(strings.Contains(str, "world"))

}

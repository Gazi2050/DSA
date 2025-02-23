package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello world"
	// var LongStr string = "A Long String"
	// var strIndex = "String"
	// fmt.Printf("%c", strIndex[0])
	// fmt.Println(len(str))
	// fmt.Println(strings.Contains(str, "world"))
	// fmt.Println(strings.ContainsAny(str, "asn"))
	// fmt.Println(strings.HasPrefix(str, "He"))
	// fmt.Println(strings.HasSuffix(str, "ld"))
	// fmt.Println(strings.Index(str, "d"))
	// fmt.Println(strings.IndexAny(str, "asd"))
	fmt.Println(strings.LastIndex(str, "l"))
	// fmt.Println(strings.ToLower(str))

}

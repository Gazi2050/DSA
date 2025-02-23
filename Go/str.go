package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = " Hello world "
	// var LongStr string = "A Long String"
	// var strIndex = "String"
	// fmt.Printf("%c", strIndex[0])
	// fmt.Println(len(str))
	/*Searching & Checking*/
	// fmt.Println(strings.Contains(str, "world"))
	// fmt.Println(strings.ContainsAny(str, "asn"))
	// fmt.Println(strings.HasPrefix(str, "He"))
	// fmt.Println(strings.HasSuffix(str, "ld"))
	// fmt.Println(strings.Index(str, "d"))
	// fmt.Println(strings.IndexAny(str, "asd"))
	// fmt.Println(strings.LastIndex(str, "l"))
	/*Modifying Strings*/
	// fmt.Println(strings.ToLower(str))
	// fmt.Println(strings.ToUpper(str))
	// fmt.Println(strings.Replace(str, " ", "f", 3))
	// fmt.Println(strings.ReplaceAll(str, " ", "f"))
	// fmt.Println(strings.Trim(str, " H"))
	// fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.Repeat(str, 9))

}

package main

import "fmt"

func main() {
	x := 1
	if x < 10 {
		fmt.Println(true)
	} else if x > 10 {
		fmt.Println("working")
	} else {
		fmt.Println(false)
	}

	day := 's'
	switch day {
	case 's':
		fmt.Println("holiday")
	case 'f':
		fmt.Println("pray day")
	case 'm':
		fmt.Println("normal day")
	default:
		fmt.Println("good day")
	}
}

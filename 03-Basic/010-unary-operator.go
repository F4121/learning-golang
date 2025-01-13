package main

import "fmt"

func main() {
	var a = 1

	a++
	fmt.Println(a) // prints 2
	a++
	fmt.Println(a) // prints 3

	a--
	fmt.Println(a) // prints 2

	// Negative value
	var b = -2
	b--
	fmt.Println(b) // prints -3
	b++
	fmt.Println(b) // prints -2
}

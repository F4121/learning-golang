package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// Augmenthed assignments
	// ex. a = a + 10
	a += 10
	fmt.Println(a)

	// ex. a = a - 10
	a -= 5
	fmt.Println(a)

	// ex. a = a * 10
	a *= 2
	fmt.Println(a)

	// ex. a = a / 10
	a /= 2
	fmt.Println(a)

	// ex. a = a % 10
	a %= 3
	fmt.Println(a)
}

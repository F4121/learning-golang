package main

import "fmt"

func main() {
	var on bool = true
	var off bool = false

	// && / And
	fmt.Println(on && off) // false

	// || / OR
	fmt.Println(on || off) // true

	// ! / Not
	fmt.Println(!on) // false

}

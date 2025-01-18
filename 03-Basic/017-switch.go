package main

import "fmt"

func main() {
	// Otherway to handle condition

	name := "Faizi"
	switch name {
	case "Faizi":
		fmt.Println("Hello Faizi")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hello Stranger")
	}

	// Short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is longer than 5")
	case false:
		fmt.Println("Name is shorter than 5")
	}

	// Switch without condition
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Name is longer than 5")
	case length < 5:
		fmt.Println("Name is shorter than 5")
	}

}

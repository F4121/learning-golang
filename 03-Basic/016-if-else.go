package main

import "fmt"

func main() {
	// Condition
	name := "Faizi"

	if name == "Faizi" { // watch here, value return true / false
		fmt.Println("Hello Faizi")
	} else if name == "Jhon" {
		fmt.Println("Hello Jhon")
	} else {
		fmt.Println("Hello Stranger")
	}

	// Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Your name is long")
	} else {
		fmt.Println("Good Name!")
	}

}

package main

import "fmt"

func main() {
	// Initial variable with data type
	var name string

	name = "Jhon"
	fmt.Println(name)

	name = "Faizi"
	fmt.Println(name)

	// Direct initialisation without data type
	var lastname = "Skyler"
	fmt.Println(lastname)

	// initialisation variable without var
	hobby := "Swimming"
	fmt.Println(hobby)

	// initial bulk variable
	var (
		firstName = "Jhon"
		lastName  = "Skyler"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}

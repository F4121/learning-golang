package main

import "fmt"

func main() {
	// const = declare variable but can't change value
	const firstName = "Joni"
	const lastName = "Janji"

	// got error "cannot assign to firstName (neither addressable nor a map index expression)"
	firstName = "Lala"
	lastName = "Lapo"

	// Declare bulk constant
	const (
		hobby = "Swimming"
		age   = 20
	)
	fmt.Println(hobby)
	fmt.Println(age)

}

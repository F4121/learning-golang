package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "John"
	middleName = "Doe"
	lastName = "Developer"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println("Name :", a, b, c)
}

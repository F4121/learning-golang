package main

import "fmt"

func getFullName() (string, string) {
	return "Jhon", "Doe"
}

// jika ingin menghiraukan value return misal lastName, maka ganti lastName dengan underscore

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

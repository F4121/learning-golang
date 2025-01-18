package main

import "fmt"

func main() {
	blacklist := func(name string) bool { //create function on variable
		return name == "dog"
	}
	registerUser("Jhon", blacklist)

	// OR

	registerUser("dog", func(name string) bool { //Direct add function without declare
		return name == "dog"
	})

}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User is blacklisted")
	} else {
		fmt.Println("Welcome", name)
	}
}

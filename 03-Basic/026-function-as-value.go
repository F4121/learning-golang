package main

import "fmt"

// Function adalah first class citizen
// Function juga merupakan tipe data, dan bisa disimpan dalam variable

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Jhon"))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

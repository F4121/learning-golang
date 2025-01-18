package main

import "fmt"

func main() {
	result := getHello("Jhon")
	fmt.Println(result)
}

func getHello(name string) string { //function should return string
	return "Hello, " + name
}

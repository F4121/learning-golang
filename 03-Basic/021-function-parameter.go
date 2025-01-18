package main

import "fmt"

func main() {
	sayHelloTo("Jaka", "Sembung")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

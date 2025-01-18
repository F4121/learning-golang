package main

import "fmt"

// Function tidak hanya bisa kita simpan dalam value
// tapi juga bisa kita gunakan sebagai parameter untuk function lain

func sayHelloWithFilter(name string, filter func(string) string) { // func filter menandakan harus punya 1 parameter dengan tipe string dan return data string
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Dog" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Jhon", spamFilter)

	// OR

	filter := spamFilter
	sayHelloWithFilter("Dog", filter)

	sayHelloWithFilter2("Doe", spamFilter2)

}

// Type Declaration
// Type declaration adalah cara kita membuat type baru dari type yang sudah ada (ALIAS)
type Filter func(string, int) string

func sayHelloWithFilter2(name string, filter Filter) { // func filter menandakan harus punya 1 parameter dengan tipe string dan return data string
	fmt.Println("Hello ", filter(name, 20))
}

func spamFilter2(name string, age int) string {
	if name == "Dog" {
		return "..."
	} else {
		return fmt.Sprintf("Hi i am %s and i am %d years old", name, age)
	}
}

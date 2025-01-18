package main

import "fmt"

// Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama

func main() {
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++ // counter adalah variable yang berada di scope function main
	}

	increment()
	increment()
	fmt.Println(counter) // counter adalah variable yang berada di scope function main
}

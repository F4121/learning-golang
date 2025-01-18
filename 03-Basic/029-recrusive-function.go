package main

import "fmt"

// Recursive function adalah function yang memanggil function dirinya sendiri

// Sample function without recursive
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// Sample function using recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	// Have a same result with
	fmt.Println(factorialRecursive(10))

}

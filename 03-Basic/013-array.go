package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Jhon"
	names[1] = "Doe"
	names[2] = "Developer"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Will got error "index 3 out of bounds [0:3]" because array only have 3 capacity (index 0, 1, 2)
	names[3] = "Hello"
	fmt.Println(names)

	// Declare array with values
	var values = [3]int{
		1, 2, 3,
	}

	// if the value like this
	// var values = [3]int{
	// 	1, 2,
	// }
	// array still got 3 value but with the default value = 0

	fmt.Println(values)

	// Function on Array
	// - len(array) get length of array
	fmt.Println(len(values))
	// - array[index] get value base on index
	fmt.Println(values[0])
	// - array[index] = value set value base on index
	fmt.Println("Before change value")
	fmt.Println(values)
	fmt.Println("After change value")
	values[0] = 10
	fmt.Println(values)

	// Declare array without specific length
	var longestArray = [...]int{
		1, 2, 3, 4, 5,
	}
	fmt.Println(longestArray)
	fmt.Println(len(longestArray))
	fmt.Println(longestArray[4])
	longestArray[0] = 10
	fmt.Println(longestArray)

}

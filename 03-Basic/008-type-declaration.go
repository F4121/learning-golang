package main

import "fmt"

func main() {

	// type is aliasing data type, ex. userID is same with data type string
	type userID string

	var johnID userID = "11111111"
	fmt.Println(johnID)

	var example string = "22222222222"
	var exampleID userID = userID(example)

	fmt.Println(exampleID)

}

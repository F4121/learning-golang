package main

import (
	"fmt"
	"strings"
)

func main() {
	/**
	PACKAGE strings
	- Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data string
	- ada banyak sekali function yang bisa kita gunakan

	reference
	https://pkg.go.dev/strings
	*/

	// Sample Trim
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	/**
	Output:
	Hello, Gophers
	*/

	// Sample ToLower
	fmt.Println(strings.ToLower("Gopher"))
	/**
	Output:
	gopher
	*/

	// ample ToUpper
	fmt.Println(strings.ToUpper("Gopher"))
	/**
	Output:
	GOPHER
	*/

	// Sample Split
	fmt.Println(strings.Split("hello,world", ","))
	/**
	Output:
	[hello world]
	*/

	// Sample Contains
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	/**
	Output:
	false
	true
	true
	true
	false
	false
	*/

	// Sample ReplaceAll
	fmt.Println(strings.ReplaceAll("hello,world", ",", "-"))
	/**
	Output:
	hello-world
	*/

}

package main

import (
	"container/list"
	"fmt"
)

func main() {
	/**
	Package container / list
	- Package container / list adalah implementasi struktur data double linked list digolang

	reference
	https://pkg.go.dev/container/list
	*/

	/**
	Sample container / list
	*/

	data := list.New()
	data.PushBack("A")
	data.PushBack("B")
	data.PushBack("C")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // A

	next := head.Next()
	fmt.Println(next.Value) // B

	next = next.Next()
	fmt.Println(next.Value) // C

	// Get value with loops
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

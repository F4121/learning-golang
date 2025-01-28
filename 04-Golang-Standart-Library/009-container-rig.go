package main

import (
	"container/ring"
	"fmt"
)

func main() {
	/**
	Package container/ring
	- Package container/ring adalah implementasi struktur data circular list
	- Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)

	reference
	https://pkg.go.dev/container/ring
	*/

	var data *ring.Ring = ring.New(5)

	// Manual assign value to ring
	// data.Value = "Value 1"
	// // data.Next()
	// // data.Value = "Value 2"
	// // OR
	// data.Next().Value = "Value 2"
	// data.Next().Next().Value = "Value 3"
	// data.Next().Next().Next().Value = "Value 4"
	// data.Next().Next().Next().Next().Value = "Value 5"

	// Assign value to ring using loops
	for i := 0; i < data.Len(); i++ {
		data.Value = fmt.Sprintf("Value %d", i+1)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})

}

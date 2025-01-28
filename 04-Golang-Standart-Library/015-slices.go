package main

import (
	"fmt"
	"slices"
)

func main() {
	/**
	Package slices
	- di golang versi terbaru terdapat fitur bernama Generic
	- Fitur Generic ini membuat kita bisa membuat parameter dengan tipe yang bisa berubah-ubah, tanpa harus menggunakan interaface kosong / any
	- Salah satu package yang menggunakan fitur Generic ini adalah package slices
	- Package slices ini digunakan untuk memanipulasi data di slice

	reference
	https://pkg.go.dev/slices
	*/

	names := []string{"Budi", "Jhon", "Tono", "Joko"}
	values := []int{100, 200, 300, 80, 20}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Joko"))
	fmt.Println(slices.Index(names, "Tono"))

}

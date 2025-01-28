package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Not Found Error")
)

func main() {
	/**
	PACKAGE ERRORS
	- Sebelumnya kita sudah pernah membahas tentang interface error yang merupakan representasi dari error digolang dan membuat error menggunakan function errors.New()
	- Sebenarnya masih banyak yang bisa kita lakukan menggunakan package errors, contohnya ketika kita ingin membuat beberapa value error yang berbeda

	reference
	https://pkg.go.dev/errors
	*/

	/**
	Mengecek jenis error
	- Misal kita membuat jenis error sendiri, lalu kita ingin mengecek jenis errornya
	- Kita bisa menggunakan function errors.Is() untuk mengecek jenis type errornya
	*/

	// Sample Errors
	err := GetById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unknown Error")
		}
	}
}

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Faizi" {
		return NotFoundError
	}

	return nil

}

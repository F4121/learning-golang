package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) { // function akan mereturn (nilai, hasil error)
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0") // return error
	} else {
		return nilai / pembagi, nil //return error as nil jika tidak ada error
	}
}

func main() {
	/**
	ERROR

	** ERROR INTERFAFCE
	- Golang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya dalah error

	type error interface {
		Error() string
	}

	** MEMBUAT ERROR
	- Untuk membuat error, kita tida perlu manual.
	- Golang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (Package bawaan golang)
	*/

	fmt.Println(Pembagian(12, 2))
	fmt.Println(Pembagian(12, 0))
}

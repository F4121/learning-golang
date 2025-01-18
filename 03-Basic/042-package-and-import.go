package main

import (
	"fmt"
	"learning-golang/03-Basic/helper"
)

func main() {
	/**
	PACKAGE
	- Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang telah kita buat di golang
	- Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
	- Package sendiri sebenernya hanya direktori folder di operasi kita
	*/

	/**
	IMPORT
	- Secara standart, file golang hanya bisa mengakses file golang lainnya yang berada dalam package yang sama
	- Jika kita ingim mengakses file golang yang berada diluar package, maka kita bisa menggunakan import
	*/

	result := helper.SayHello("Faizi")
	fmt.Println(result)
}

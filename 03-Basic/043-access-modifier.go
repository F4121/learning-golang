package main

import (
	"fmt"
	"learning-golang/03-Basic/helper"
)

func main() {
	/**
	ACCESS MODIFIER
	- di bahasa pemrograman lain, biasanya ada kata kunci yang bisa kita gunakan untuk menentukan access modifier terhadap suatu function atau variable
	- di golang, untuk menentukan access modifier, cukup dengan nama function atau variable
	- jika namanya diawali dengan huruf besar, maka artinya bisa diakses dari package lain, jika dimulai dengan huruff kecil artinya tidak bisa diakses dari dari package lain
	*/

	// Untuk Sample lihat di file helper :)

	result := helper.SayHello("Faizi")
	fmt.Println(result)

	fmt.Println(helper.sayGoodBye("Faizi")) // .\043-access-modifier.go:21:21: undefined: helper.sayGoodBye
	fmt.Println(helper.version)             // .\043-access-modifier.go:22:21: undefined: helper.version

}

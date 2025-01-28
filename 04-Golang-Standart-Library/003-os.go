package main

import (
	"fmt"
	"os"
)

func main() {
	/**
	PACKAGE OS
	- Golang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
	- Package os berisikan fungsionalitas untuk mengakses itur sistem operasi secara independen (bisa digunakan disemua sistem operasi)
	- Package os juga berisikan fungsionalitas untuk mengakses direktori dan file


	reference
	https://pkg.go.dev/os
	*/

	// Sample os.Args
	/**
	Args adalah parameter yang dikirim ketika kita menjalankan command di terminal. contoh ketika kita run applikasi golang
	- go run main.go arg1 arg2 arg3
	- go run main.go -h
	- go run main.go --help
	- go run main.go -v
	- go run main.go faizi belajar golang
	- go run main.go "faizi belajar golang" to print single string
	diatas adalah contoh Args dan Args bisa kita tangkap mengggunakan package os
	*/

	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	// Sample Hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(hostname)
	}

}

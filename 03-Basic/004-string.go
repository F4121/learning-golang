package main

import "fmt"

func main() {
	fmt.Println("This is string")

	// ada beberapa func string seperti
	// len("string") digunakan untuk menghitung jumlah karakter di string
	// "string"[number] digunakan untuk mengambil karakter diddalam string (value dalam bentuk byte)

	fmt.Println(len("Hello"))
	fmt.Println("Hello"[1])
}

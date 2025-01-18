package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/**
	OPERATOR NEW
	- Sebelumnya untuk membuat pointer dengan menggaunakan operator &
	- Golang juga memilki function new yang bisa kita gunakan untuk membuat pointer
	- Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
	*/

	// Sample without new
	alamat1 := &Address{}

	// Sample with operator new
	alamat2 := new(Address)
	alamat3 := alamat2

	alamat3.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
	fmt.Println(alamat3)

}

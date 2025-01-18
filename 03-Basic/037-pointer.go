package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/**
	Pass by Value
	- Secara default di golang semua variable itu di passing by value, bukan by reference
	- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi valuenya
	*/
	// address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// address2 := address1 //Copy value

	// address2.City = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	/**
	Pointer
	- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
	- Sederahananya, dengan kemampuan pointer kita bisa membuat pass by reference

	Operator &
	- Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variable nya
	*/
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

}

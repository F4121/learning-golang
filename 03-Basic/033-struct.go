package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	/**
	Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan.
	Struct biasanya representasi data dalam program aplikasi yang kita buat.
	Data struct disimpan didalam field
	sederhananya struct adalah kumpulan dari field
	*/

	// Struct adalah template data / prototype data
	// Struct tidak bisa langsung digunakan
	// Namun bisa membuat data / object dari struct yang kita buat

	var jhon Customer
	jhon.Name = "Jhon"
	jhon.Address = "Jalan Jalan"
	jhon.Age = 25

	fmt.Println(jhon)
	fmt.Println(jhon.Name)
	fmt.Println(jhon.Address)
	fmt.Println(jhon.Age)

	// Oke!! sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct

	// Alt 1
	nana := Customer{
		Name:    "Nana",
		Address: "Jalan Nana",
		Age:     30,
	}
	fmt.Println(nana)
	fmt.Println(nana.Name)
	fmt.Println(nana.Address)
	fmt.Println(nana.Age)

	// Alt 2
	miya := Customer{"Miya", "Jalan Miya", 27}
	fmt.Println(miya)
	fmt.Println(miya.Name)
	fmt.Println(miya.Address)
	fmt.Println(miya.Age)

	// STRUCT METHOD
	miya.sayHello("Agus")

}

func (customer Customer) sayHello(name string) { //this is struct method
	fmt.Println("Hello", name, "my name is", customer.Name)
}

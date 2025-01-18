package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func ChangeAddressToIndonesiaWithPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	/**
	Pointer di Function
	- Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan dicopy lalu dikirim ke function tersebut
	- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
	- Hal ini membuat variable jadi aman, karena tidak akan bisa diubah
	- Namun kadang kita ingin membuat function yan g bisa mengubah data asli parameter tersebut
	- Untuk melakukan ini, kita juga bisa menggunakan pointer di function
	- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya
	*/
	address := Address{"Bekasi", "Jawa Barat", ""}
	ChangeAddressToIndonesia(address)
	fmt.Println(address) // {Bekasi Jawa Barat }

	/**
	value yang dikirimkan ke ChangeAddressToIndonesiaWithPointer haruslah pointer
	contoh:
	*/
	var address2 *Address = &Address{}
	ChangeAddressToIndonesiaWithPointer(address2)
	fmt.Println(address2)

	/**
	namun jika varible kita bukan address saat didefine diawal
	kita bisa menambahkan & pada saat mengirimkan parameter di function seperti contoh dibawah
	*/
	var address3 Address = Address{}
	ChangeAddressToIndonesiaWithPointer(&address3)
	fmt.Println(address3) // {Bekasi Jawa Barat Indonesia}

}

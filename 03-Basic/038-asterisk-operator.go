package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/**
	Operator *
	- Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
	- Semua variable yang mengacu ke data yang sama tidak akan berubah
	- Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *
	*/
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"
	fmt.Println(address1) // {Bandung DKI Jakarta Indonesia}
	fmt.Println(address2) // &{Bandung DKI Jakarta Indonesia

	address2 = &Address{"Bekasi", "Jawa Barat", "Indonesia"} // digunakan untuk reasign address2 ke data baru / address baru sehingga address 2 sudah tidak referer ke address1
	fmt.Println(address1)                                    // {Bandung DKI Jakarta Indonesia}
	fmt.Println(address2)                                    // &{Bekasi Jawa Barat Indonesia}

	// Sample asterisk
	/**
	sebagai contoh pada variable address1 dan address2.

	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1

	maka memory di address1 dan address2 adalah sama, mengacu ke data Jakarta

	nah astersik ini gunakan untuk reasign value kepada semua variable yang memiliki data memory yang sama

	*address2 = Address{"Bekasi", "Jawa Barat", "Indonesia"}

	maka yang terjadi adalah data di address1 akan diubah menjadi data Bekasi

	fmt.Println(address1) // Hasilnya {"Bekasi", "Jawa Barat", "Indonesia"}
	fmt.Println(address2) // Hasilnya &{"Bekasi", "Jawa Barat", "Indonesia"}
	*/

}

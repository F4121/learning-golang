package main

import "fmt"

func main() {
	// For
	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++ // increment
	}

	// For with statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("For Statement :", counter)
	}

	// For bisa digunakan untuk melakukan itegrasi terhadap semua data collection
	// Data collection contohnya Array, Slice, Map

	// Manual access value using loop
	names := []string{"Jhon", "Lebron", "Akamsi"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Using For Range
	// For Range akan mengambil index dan value secara otomatis
	// Jadi kita tidak perlu lagi membuat variabel index
	for index, name := range names {
		fmt.Println("Index :", index, "Name :", name)
	}
	// Jika tidak butuh menggunakan index maka bisa mengganti index dengan underscore (_)
	// for _, name := range names {
	// 	fmt.Println("Name :", name)
	// }
}

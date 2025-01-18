package main

import "fmt"

// Parameter yang berada diposisi akhir, memiliki kemampuan dijadikan sebuah varargs
// Varargs artinya datanya bisa menerima lebbih dari satu input, atau anggap saja semacam array
// apa bedanya dengan parameter biasa dengan tipe data array ??
// jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
// jika parameter menggunakan varargs, kita bisa langsung mengirim datanya jika lebih dari satu, cukup gunakan tanda koma

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10}
	totalWithSlice := sumAllSlice(numbers...)
	fmt.Println(totalWithSlice)
}

// Slice bisa digunakan sebagai varargs
func sumAllSlice(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

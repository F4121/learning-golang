package main

import "fmt"

type Man struct {
	Name string
}

/*
*
Jika tanpa * maka akan mereturn Faizi, karena parameter bertipe value yang mana akan diduplikasi tanpa mengubah data aslinya
Jika dengan * maka akan mereturn &Faizi, karena parameter bertipe pointer yang mana akan mengembalikan alamat memori Faizi
dan hasil dari faizi.Married() adalah Mr. Faizi
*/
func (man *Man) Married() { // Function Married akan mejadi bagian dari struct Man
	man.Name = "Mr. " + man.Name
}

func main() {
	/**
	Pointer di Method
	- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses didalam method adalah pass by value
	- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
	*/
	faizi := Man{"Faizi"}
	faizi.Married()

	fmt.Println(faizi)
}

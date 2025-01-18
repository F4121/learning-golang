package main

import "fmt"

func main() {
	// Map addalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
	// sederhananya, Map adalah tipe data kumpulan key-value, dimana kata kuncinya bersifat unik, tidak boleh sama
	// berbeda dengan array dan slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kkata kuncinya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

	// Cara 1
	var person map[string]string = map[string]string{}
	person["name"] = "John"
	person["age"] = "25"

	// Cara 2
	person2 := map[string]string{
		"name": "Lala",
		"age":  "24",
	}
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person)

	fmt.Println(person2["name"])
	fmt.Println(person2["age"])
	fmt.Println(person2)

	// apabila mengakses key yang tidak didefine dialam map maka akan ditampilkan default value sesuai dengan tipe data. misal tipe data string maka default valuenya adalah string kosong / ""

	// Function Map
	// 1. make(map[TypeKey]TypeValue)
	// untuk membuat map baru
	book := make(map[string]string)
	book["title"] = "Golang"
	book["author"] = "John"
	fmt.Println(book)

	// 2. len(map)
	// untuk mengetahui berapa banyak data yang ada di dalam map
	fmt.Println(len(book))

	// 3. map[key]
	// untuk mengambil data di map dengan key
	fmt.Println(book["title"])

	// 4. map[key] = value
	// untuk mengubah data di map dengan key
	book["title"] = "Python"
	fmt.Println(book)

	// 5. delete(map, key)
	// untuk menghapus data di dalam map
	delete(book, "author")
	fmt.Println(book)

}

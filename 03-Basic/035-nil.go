package main

import "fmt"

func main() {
	/**
	Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
	Berbeda dengan golang, di golang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuat default valuenya
	namun di golang ada data nil, yaitu data kosong
	!!! nil sendiri hanya bisa digunakan di beberapa tipe data, seperti inteface, function, map, slice pointer dan channel
	*/

	data := NewMap("Jhon")
	// data := NewMap("")
	if data == nil {
		fmt.Println("Data map kosong")
	} else {
		fmt.Println(data["name"])
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	/**
	Type assertion merupakkan kemampuan merubah tipe data menjadi tipe data yang diinginkan
	Fitur ini sering sekali digunakan ketika kita bertemui dengan data interface kosong
	*/
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int) // panic!! karna data tidak sesuai. didalam func random terdapat return string, jika kita mengkonversinya kedalam int maka akan terjadi panic
	// fmt.Println(resultInt)

	/**
	TYPE ASSERTION MENGGUANAKAN SWITCH
	ini adalah cara aman untuk menghandle type assertion ketika tidak tau tipe data yang akan di konversi
	*/
	result2 := random()
	switch value := result2.(type) {
	case string:
		fmt.Println("String ", value)
	case int:
		fmt.Println("Int ", value)
	default:
		fmt.Println("Unknown ", value)
	}

}

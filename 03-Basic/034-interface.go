package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello2(hasName HasName) {
	fmt.Println("Hello, my name is", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func upsError() interface{} {
	// return 1
	return true
}

// Di version golang terbaru kita menggunakan any sebagai pengganti interface kosong contoh seperti berikut
func upsError2() any { //
	// return 1
	return true
}

func main() {
	/**
	Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
	sebuah interface berisikan definisi-definisi method
	biasanya interface digunakan sebagai kontrak


	**** implementasi interface
	setiap tipe data yang sesuai dengan kontrak interface, secara otomatis diangap sebagai interface tersebut.
	sehingga kita tidak perlu mengimplementasi interface secara manual
	hal ini agak berbeda dengan bahasa pemgrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana.
	*/

	person := Person{Name: "Jhon"}
	sayHello2(person)

	animal := Animal{Name: "Kucing"}
	sayHello2(animal)

	/**
	****** Interface Kosong
	Golang bukanlah bahasa pemrograman yang berorientasi objek
	Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
	contoh di Java ada java.lang.Object
	untuk menangani kasus seperti ini, di Golang kita bisa menggunakan interface kosong
	Interface kosong adalah interface yang tidak memilki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
	Interface kosong, juga memiliki type alias bernama any


	Berikut beberapa contoh penggunaan interface kosong digolang :

	fmt.Prinln(a ...interface{})
	panic(v inteface{})
	recover() interface{}
	dll.
	*/

	var kosong any = upsError()
	fmt.Println(kosong)

}

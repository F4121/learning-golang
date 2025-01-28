package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"20"` // Include struct tag
}

type Person struct {
	Name    string `required:"true" max:"20"` // Include struct tag
	Address string `required:"true" max:"20"` // Include struct tag
	Email   string `required:"true" max:"20"` // Include struct tag
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("========Type Name ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, " with type ", structField.Type)
		fmt.Println("is Required ", structField.Tag.Get("required"))
		fmt.Println("=============Max value ", structField.Tag.Get("max"))
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface() //Interface adalah datanya
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	/**
	Package reflect
	- Dalam bahasa pemrograman biasanya ada fitur reflection, dimana kita bisa melihat struktur code kita pada saat aplikasi sedang berjalan
	- Hal ini bisa dilakukan digolang dengan menggunakan package reflect
	- Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, anda bisa mengeksplorasi package reflect ini secara otodidak
	- Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
	reference
	https://pkg.go.dev/reflect
	*/

	// Sample
	readField(Sample{"Jhon"})
	// Person
	readField(Person{"Jhon", "Jl. Jawa 29 Blok AC", "jhondoe210@gmail.com"})

	// with validation
	person := Person{
		Name:    "ada",
		Address: "ada",
		Email:   "ada",
	}
	fmt.Println("IS VALID :", isValid(person))
}

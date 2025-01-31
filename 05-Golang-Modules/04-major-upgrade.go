package main

import (
	"fmt"

	say_hello "github.com/F4121/module-say-hello/v2"
)

func main() {
	/**
	MAJOR UPGRADE
	- Ada kalanya ketika module memiliki perubahan yang cukup signifikan kita akan melakukan upgrade major version misal v1 ke v2
	- Nah hal tersebut juga bisa lakukan digolang module. caranya adalah
		- kita cukup ubah nama module di go.mod
		- contoh pada repo https://github.com/F4121/module-say-hello kita buka file go.mod
		- setelah itu kita akan menemukan sytax  "module github.com/F4121/module-say-hello"
		- kita ubah menjadi "module github.com/F4121/module-say-hello/{version yang diinginkan}
		- contoh "module github.com/F4121/module-say-hello/v2"
		- setalah itu repo module perlu ditag ulang misal v2.0.0
		- setelah itu kita harus mengupdate go.mod di project yang menggunakan module tersebut
		- contoh "github.com/F4121/module-say-hello" menjadi
		- "github.com/F4121/module-say-hello/v2"



	** SAMPLE
	pada repo module say hello saya sudah melakukan update code di versi 2 dengan menambahkan parameter name

	*/

	fmt.Println(say_hello.SayHello("John"))

}

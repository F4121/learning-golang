package main

import (
	"fmt"

	say_hello "github.com/F4121/module-say-hello"
)

func main() {
	/**
	Membuat Golang Modules
	- Untuk membuat module baru, kita bisa menggunakan perintah: go mod init nama-module
	- Golang akan secara otomatis membuat file go.mod yang berisikan informasi nama module dan juga versi golang yang kita gunakan

	Rilis Module
	- Golang terintegrasi baik dengan Git
	- Untuk merilis module kita hanya perlu membuat Tag di Git


	*** SAMPLE
	saya telah membuat module bernama say hello di git
	https://github.com/F4121/module-say-hello
	yang mana module tersebut akan saya coba panggil di module ini
	menggunakan syntax go get {url-git}

	code dibawah adalah hasil dari call module dari repo "github.com/F4121/module-say-hello"
	*/

	fmt.Println(say_hello.SayHello())

}

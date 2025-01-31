package main

import (
	"fmt"

	say_hello "github.com/F4121/module-say-hello"
)

func main() {
	/**
	UPGRADE MODULE
	- untuk melakukan upgrade module kita hanya perlu tag ulang pada repository module setelah itu kita modify version di go.mod.
	- setelah ubah version di go mod perlu menjalankan sytax go get agar golang mendownload versi module yang dimaksud


	** SAMPLE
	pada module say_hello tag v1.0.1 hanya me-return value hello
	kita akan ubah version module say_hello ke version v1.0.2 kita akan merubah version yang ada di gomod
	setelah itu kita jalankan ulang go get agar module ter-update
	*/

	fmt.Println(say_hello.SayHello())

	// Success print "Hello World!" sesuai dengan code yang telah di ubah dari v1.0.1 -> v1.0.2
}

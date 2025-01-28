package main

func main() {
	/**
	Package io
	- IO adalah singkatan dari Input Output, merupakan fitur di Golang yang digunakan sebagai standart untuk proses Input Output
	- di golang semua mekanisme input output pasti mengukuti standart package io

	reference
	https://pkg.go.dev/io
	*/

	/**
	*** READER
	- Reader digunakan untuk membaca input, golang menggunakan kontrak interface bernama Reader yang terdapat di package io
		type Reader interface {
			Read(p []byte) (n int, err Error)
		}
	*/

	/**
	*** WRITER
	- Writer digunakan untuk menuli ke output, golang menggunakan kontrak interface bernama Writer yang terdapat di package io
		type Writer interface {
			Write(p []byte) (n int, err Error)
		}
	*/

	/**
	*** IMPLEMENTASI IO
	- Penggunaan dari IO sendiri di golang terdapat dibanyak package, sebelumnya contoh kita menggunakan CSV Reader dan CSV Writer
	- Karena package io sebenarnya hanya kontrak untuk IO, untuk implementasinya kita harus lakukan sendiri
	- Tapi untungnya golang juga menyediakan package untuk mengimplementasikan IO secara mudah, yaitu menggunakan package bufio
	*/

}

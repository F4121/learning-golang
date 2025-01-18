package main

import (
	"fmt"
	"learning-golang/03-Basic/database"
	_ "learning-golang/03-Basic/internal" // ini adalah contoh blank identifier yang mana ketika package diimport namun hanya ingin menjalankan init nya
)

func main() {
	/**
	PACKAGE INITIALIZATION
	- Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
	- Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function intisialisasi untuk membuka koneksi ke database
	- Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init
	*/
	// Sample lihat di package database file mysql
	fmt.Println(database.GetDatabase())

	/**
	BLANK IDENTIFIER
	- Kadang kita hanya inigin menjalankan init function dipackage tanpa harus mengeksekusi salah satu function yang ada di package
	- Secara default, golang akan komplen ketika ada package yang diimport namun tidak digunakan
	- Untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import
	*/
}

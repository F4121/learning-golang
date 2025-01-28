package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	/**
	Package path
	- Package path digunakan untuk memanipulasi data path seperti path di URL atau path di File System
	- Secara default Package path menggunakan slash sebagai karakter path nya, oleh karena itu cocok untuk data URL
	- Namun jika ingin menggunakan untuk memanipulasi path di file system, karena windows menggunakan backslash, maka khusus untuk file system. perlu menggunakan package path/filepath

	reference
	https://pkg.go.dev/path
	https://pkg.go.dev/path/filepath
	*/

	// Sample path
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))

	// Sample filepath
	fmt.Println(filepath.Dir("hello/world.go"))
	fmt.Println(filepath.Base("hello/world.go"))
	fmt.Println(filepath.Ext("hello/world.go"))
	fmt.Println(filepath.IsAbs("hello/world.go"))   //absolute disini maksudnya path ditulis lengkap ex. di windows ditulis hingga lokasi drive contoh D:\Learning\Golang\
	fmt.Println(filepath.IsLocal("hello/world.go")) //local adalah kebalikan dari absolute yang mana path tidak dituliskan dari depan / tidak lengkap
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}

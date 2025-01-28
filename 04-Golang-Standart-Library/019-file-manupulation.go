package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	/**
	File Management
	- Di package os, terdapat File Management, namun sengaja ditunda dulu pembahasannya, karena kita perlu tahu dulu terkait IO
	- Saat kita membuat atau membaca file menggunakan Package os, struct File merupakan implementasi dari io.Reader dan io.Writer
	- Oleh karena itu, kita bisa melakukan baca dan tulis terhadap File tersebut menggunakan Package io / bufio
	*/

	/**
	*** OPEN FILE
	- Untuk membuat / membaca file kita bisa menggunakan os.OpenFile(name, flag, permission)
	- name berisikan nama ffile, bisa absolute atau relative / local
	- flag merupakan penanda file, apakah untuk membaca, menulis dan lain-lain
	- permission merupakan permission yang diperlukan ketika membaut file bisa kita sumulasikan disini

	reference
	https://chmod-calculator.com/
	*/

	/**

	*** File Flag di Package os
	// flags may be implemented on a given system.
	const (
		// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.0_WRONLY // open the file write-only.
		O_RDWR int = syscall.0_RDWR // open the file read-write.
		// The remaining values may be or'ed in to control behavior.
		0_APPEND int = syscall.0_APPEND // append data to the file when writing.
		O_CREATE int = syscall.0_CREAT // create a new file if none exists.
		0_EXCL int = syscall.0_EXCL // used with
		0_CREATE, file must not exist.
		0_SYNC int = syscall.0_SYNC // open for synchronous I/O.
		0_TRUNC int = syscall.0_TRUNC // truncate regular writable file when opened.
	)
	*/

	// Samplle Buat File Baru
	createNewFile("sample.log", "this is sample log")

	// Sample readd file
	result, _ := readFile("sample.log")
	fmt.Println(result)

	// Add to File
	addToFile("sample.log", "\nthis is additionals log")

}

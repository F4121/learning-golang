package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	PACKAGE strconv
	- Sebelumnya kita sudah belajar cara konversi tipe data misal dari int32 ke int34
	- Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? misal dari int ke string, atau sebaliknya
	- Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)

	reference
	https://pkg.go.dev/strconv
	*/

	/**
	Berikut adalah contoh function di package strconv yang sering digunakan

	Atoi (string to int) and Itoa (int to string).
	i, err := strconv.Atoi("-42")
	s := strconv.Itoa(-42)

	atau bisa menggunakan parse

	strconv.parseBool(string)		mengubah string ke bool
	strconv.parseFloat(string)		mengubah string ke float
	strconv.parseInt(string)		mengubah string ke int64
	strconv.FormatBool(bool)		mengubah bool ke string
	strconv.FormatFloat(float, ...)	mengubah float4 ke string
	strconv.FormatInt(int, ...)		mengubah int64 ke string
	*/

	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultInt)
	}

}

package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	PACKAGE math
	- Package math merupakan package yang berisikan constant dan fungsi matematika

	reference
	https://pkg.go.dev/math
	*/

	/**
	Berikut adalah contoh function di package strconv yang sering digunakan

	math.Round(float64)				Membulatkan float64 ke atas atau kebawah sesuai dengan yang paling dekat
	math.Floor(float64)				Membulatkan float64 ke bawah
	math.Ceil(float64)				Membulatkan float64 ke atas
	math.Max(float64, float64)		Mengambil nilai float64 yang paling besar
	math.Min(float64, float64)		Mengambil nilai float64 yang paling kecil

	*/

	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.40))
	fmt.Println(math.Round(1.6))
	fmt.Println(math.Max(1.40, 2.40))
	fmt.Println(math.Min(1.40, 2.40))

}

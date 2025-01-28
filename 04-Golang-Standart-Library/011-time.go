package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	Package time
	- Package time adalah package yang berisikan fungsionalitas untuk management waktu di golang
	reference
	https://golang.org/pkg/time
	*/

	/**
	Berikut adalah beberapa function yang ada di package time
	time.Now()						Untuk mendapatkan waktu saat ini
	time.Date(...)					Untuk membuat waktu
	time.Parse(layout, string)		Untuk parsing waktu dari string
	time.Sleep()					Untuk membuat program menunggu sebelum melanjutkan
	*/

	// var now time.Time = time.Now()
	// fmt.Println(now)

	// var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	// fmt.Println(utc)
	// fmt.Println(utc.Local())

	formatter := "2007-01-02 15:04:05"
	value := "2020-10-10 10:10:10"

	// formatter := "2006-01-02 15:04:05"
	// value := "2020-10-10 10:10:10"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valueTime)
	}

}

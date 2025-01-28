package main

import (
	"fmt"
	"regexp"
)

func main() {
	/**
	Package regexp
	- Package regexp adalah utilitas di golang untuk melakukan pencarian regula expression
	- Regular expression di golang menggunakan library C yang dibuat google bernama RE2
	- RE2 adalah implementasi dari regular expression yang lebih cepat dan lebih kecil dari regex

	reference
	https://github.com/google/re2/wiki/Syntax
	https://pkg.go.dev/regexp

	*** Beberapa Function di Package regexp
	regexp.MustCompile(string) 			Membuat regexp
	regexp.MatchString(string) bool		Mengecek apakah regexp match dengan string
	regexp.FindAllString(string, max)	Mencari string yang match dengan maximum jumlah hasil
	*/

	var regex *regexp.Regexp = regexp.MustCompile(`j([a-z])n`)
	fmt.Println(regex.MatchString("jon"))  //true
	fmt.Println(regex.MatchString("Loki")) //false

	fmt.Println(regex.FindAllString("jin jan jun Jan Jin Jun Loki Leko Lesu jon Wek Wok", 10)) // parameter terakhir menandakan max data yang ingin dicaris
}

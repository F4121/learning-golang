package main

import "fmt"

func main() {
	// Break digunakan untuk menghentikan seluruh perulangan
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Loop :", i)
	}

	//Continue digunakan untuk menghentikan perulangan yang berjalan dan langsung melanjutkan ke perulangan selanjutnya
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("Loop 2 :", i)
	}

}

package goroutine

import (
	"fmt"
	"testing"
	"time"
)

/*
Pengenalan goroutine
- Goroutine adalah sebuah thread ringan yang dikelola oleh Go Runtime
- Ukuran goroutine sangat kecil, sekitar 2kb, jauh lebih kecil dibandingkan Thread yang bisa sampai 1mb atau 1000kb
- Namun tidak seperti thread yang berjalan parallel, goroutine berjalan secara concurrent


Cara kerja goroutine
- Sebenarnya, goroutine dijalankan oleh Go Scheduler dalam thread, dimana jumlah threadnya sebanyak GOMAXPROCS (biasanya sejumlah core CPU)
- Jadi sebenarnya tidak bisa dibilang Goroutine itu pengganti thread, karena Goroutine sendiri berjalan di atas Thread
- Namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua sudah diatur oleh Go Scheduler


Membuat Goroutine
- Untuk membuat goroutine di golang sangatlah sederhana
- Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine
- Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous, artinya tidak akan ditunggu sampai function tersebut selesasi
- Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai


*** Note
- Goroutine tidak cocok digunakan untuk menjalankan function yang me-return nilai, karena nilai tersebut tidak akan bisa diakses oleh aplikasi yang menjalankan goroutine


Goroutine sangat ringan
- Speerti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
- Kita bisa membuat ribuan, bahkan sampai jutaan goroutin tanpa takut boros memory
- Tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan

**/

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(Number int) {
	fmt.Println("Display", Number)
}

func TestCreateWithoutGoroutine(t *testing.T) {
	// running synchronous
	RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)

	//result will be :
	//Hello World
	//Ups

}

func TestCreateGoroutine(t *testing.T) {
	// running asynchronous because using goroutine add go keyword
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)

	//result will be :
	//Ups
	//Hello World
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

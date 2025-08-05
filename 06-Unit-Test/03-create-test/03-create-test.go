package main

func main() {
	/**
	CREATE TEST
	- Aturan unit test
	  - golang memiliki aturan cara membuat file khusus untuk unit test
	  - Untuk membuat file unit test, kita harus menggunakan akhiran _test
	  - Jika misal kita membuat file hello_world.go, artinya untuk membuat unit testnya, kita harus membuat file hello_world_test.go

	- Aturan function unit test
	 - selain aturan nama file, di golang juga sudah diatur unntuk nama function unit test
	 - nama function untuk unit test harus diawali dengan nama Test (case sensitive harus T besar)
	 - Misal jika kita ingin mengetest function HelloWorld, maka kita akanmembuat function unit test dengan nama TestHelloWorld
	 - Tidak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penting adalah harus diawali dengan Test
	 - Selanjutnya harus memiliki parameter (t *testing.T) dan tidak mengembalikan / return value

	 untuk sample lihat file hello-world-test.go

	 - Menjalankan Unit Test
	  - untuk menjalankan unit test kita bisa menggunakan perintah
	  	go test
	  - jika kita ingin melihat lebnih detail function test apa saja yang sudah dirunning, kita bisa gunakan perintah
	  	go test -v
	  - dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah
	  	go test -v -run TestNamaFunction / go test -v -run=TestNamaFunction

	notes: jika kita ingin menjalankan spesifik function perlu diketahui bahwa run spesifik function digolang bersifat prefix misal
	kita memilki 2 function test TestHelloWorld dan TestHelloWorldJoko. jika kita melakukan test dengan syntax
	go test -v -run TestHelloWorld
	maka akan menjalankan semua function test yang memiliki prefix TestHelloWorld, yang mana TestHelloWorldJoko juga akan di running

	!!!!!!!!!
	dan juga ketika kita menjalankan go test, secara default golang akan menjalankan go test hanya dipath active saja
	misal path active / terminal kita sedang didalam folter helper. maka golang hanya mencari file test didalam folder helper.
	namun kita juga bisa melakukan test pada semua folder diproject menggunakan syntax
	go test -v ./...

	*/
}

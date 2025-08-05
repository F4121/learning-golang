package benchmark

/*
#1
Benchmark
- Selain unit test, golang testing package juga mendukung untuk melakukan benchmark
- Benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
- Benchmark di golang dilakukan dengan cara otomatis melakukan iterasi kode yang kita panggil berkali kali sampai waktu tertentu
- Kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh testing.B bawaan dari testing package

testing.B
- testing.B adalah struct yang digunakan untuk melakukan benchmark
- testing.B mirip dengan testing.T, terdapat function Fail(), FailNow(), Error(), Fatal() dan lain-lain.
- yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark
- salah satunya adalah attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark

Cara kerja benchmark
- Cara kerja benchmark di golang sangat sederahana
- Gimana kita hanya perlu membuat perulangan sejumlah N attribute
- Nanti secara otomatis golang akan melkukan eksekusi sejumlah perulangan yang di tentukan secara otomatis, lalu mendeteksi berapa lama proses tersebut berjalan, dan disimpulkan pperforma benchmarknya dalam waktu



#2
Benchmark function
- Mirip seperti unit test, untuk benchmark pun di golang sudah ditentukan nama functionnya, harus diawalli dengan kata Benchmark, misal BenchmarkHelloWolrd, BenchmarkXxx
- Selain itu, harus memiliki parameter (b *testing.B)
- dan tidak boleh mengembalikan return value
- untuk nama file benchmark, sama seperti unit test, diakhiri dengan _test, misal hello_world_test.go

Menjalankan Benchmark
- untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test namun ditambahkan parameter bench:
go test -v -bench=.
- jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah:
go test -v -run=NotMatchUnitTest -bench=.
- Kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika kita hanya ingin menjalakankan benchmark tertentu, kita bisa gunakan:
go test -v -run=NotMatchUnittest -bench=BenchmarkTest
- Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah:
go test -v -bench=. ./..




#3
Sub Benchmark
- Sama seperti testing.T, di testing.B juga kita bisa membuat sub benchmark mengunakan function Run()

Menjalankan hanya sub benchmark
- Saat kita menjalankan benchmark function, maka semua benchmark akan berjalan
- Namun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah
go test -v -bench=BenchmarkNama/NamaSub



#4
Table Benchmark
- Sama seperti di unit test, programmer golang terbiasa membuat table benchmark juga
- Ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat banyak benchmark function


**/

func HelloWorld(name string) string {
	return "Hello, " + name
}

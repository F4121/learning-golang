package main

func main() {
	/**
	MENGGAGALKAN UNIT TEST
	- menggagalkan unit test menggunakan panic bukanlah hal yang bagus
	- golang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
	- terdapat function

		- Fail() ==> akan menggagalkan unit test namun tetap melanjutkan eksekusi unit test (kode program selanjutnya. bukan test selanjutnya) , namun diakhir ketika selesai maka unit test tersebut dianggap gagal,
		- FailNow() ==> akan menggagalkan unit test saat ini juga tanpa melanjutkan eksekusi unit test,
		- Error() ==> function lebih seperti melakukan log(print) error, namun setelah melukan log error, dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal. namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai ,
		- Fatal() ==> Mirip dengan Error(), hanya saja setalah melakukan log error, dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti

	  jika kita ingin menggagalkan unit test
	*/
}

package main

import "fmt"

func main() {
	// Tipe data slice adalah potongan data dari Array
	// Sama seperti array, tapi tidak memiliki ukuran tetap

	// Tipe data slice memiliki 3 data yaitu pointer, length, capacity
	// Pointer adalah alamat memori dari data
	// Length adalah jumlah data yang ada di dalam slice
	// Capacity adalah jumlah data yang dapat ditambahkan ke dalam slice, dimana length tidak boleh lebih dari capacity

	// Membuat slice dari array
	// array[low:high] ===> Membuat slice dari array dimulai index low sampai index sebelum high
	// array[low] ===> Membuat slice dari array dimulai dari index low sampai index akhir di array
	// array[:high] ===> membuat slice dari array dimulai dari index 0 sampai index sebelum high
	/// array[:] ===> membuat slice dari array dimulai index 0 sampai index akhir di array

	// SAMPLE
	// Month := ['Januari', 'Febuari', 'Maret', 'April', 'Mei', 'Juni', 'Juli', 'Agustus', 'September', 'Oktober', 'November', 'Desember']
	// Slice1 := array[4:7]
	// Slice1 akan berisikan ['Mei', 'Juni', 'Juli']
	// dan Slice1 reference ke Pointer = 4 (diambil dari low slice), Length = 3 (total data yang diambil), Capacity = 8 (dihitung dari low sampai end of array)

	names := [...]string{"Jhon", "Layla", "Miya", "Bruno", "Nana", "Lancelot"}
	slice := names[4:6]
	fmt.Println(slice) // [Nana Lancelot]

	slice2 := names[:3]
	fmt.Println(slice2) // [Jhon Layla Miya]

	slice3 := names[3:]
	fmt.Println(slice3) // [Bruno Nana Lancelot]

	slice4 := names[:]
	fmt.Println(slice4) // [Jhon Layla Miya Bruno Nana Lancelot]

	// define slice manual
	var slice5 []string = names[:]
	fmt.Println(slice5) // [Jhon Layla Miya Bruno Nana Lancelot]

	// Function di slice
	// len(slice) untuk mendapatkan panjang slice
	// cap(slice) untuk mendapatkan capacity slice
	// append(slice, value) untuk menambahkan value ke dalam slice
	// make([]TypeData, length, capacity) untuk membuat slice baru
	// copy(slice, array) untuk mengcopy data dari array ke dalam slice

	// Sample Append
	days := [...]string{
		"Senin", "Selasa", "Rabu", "Kamis",
		"Jumat", "Sabtu", "Minggu",
	}

	daySlice1 := days[5:]
	fmt.Println(daySlice1) // [Sabtu Minggu]

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice2) // [Sabtu Minggu Libur Baru]
	fmt.Println(days)

	// Sample MAKE Slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Apple"
	newSlice[1] = "Windows"
	fmt.Println(newSlice)      // [Apple Windows]
	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	// newSlice[2] = "Unix" // got error index out of range [2] with length 2 (should be use append)

	newSlice2 := append(newSlice, "Unix")
	fmt.Println(newSlice2) // [Apple Windows Unix]

	// Jika melakukan modifikasi pada newSlice2[0] = "Android" maka hasilnya newSlice[0] akan menjadi "android" juga karna
	// capacity masih tersisa, namun jika kita capacity tidak tersedia, slice akan membuat array baru yang mana itu mendandakan bahwa
	// array newSlice sudah tidak terhubung dengan newSlice2

	// Sample Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// !!! Perlu hati2 dalam membuat slice karna mirip seperti pembuatan array
	// iniArray := [...]int{1,2,3,4,5}
	// iniArray := [5]int{1,2,3,4,5}
	// iniSlice := []int{1,2,3,4,5}

	// !!!! Dalam real case pembuatan aplikasi, di golang akan sedikit menemukan case data array, kebanyakan akan menggunakan slice

}

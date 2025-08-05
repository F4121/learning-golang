package mock

/*
Mock
- Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil dia akan menghasilkan data yang sudah kita program diawal
- Mock adalah salah satu tenik dalam unit test, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing
- Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil API Call ke third party service, hal ini sangat sulit untuk di test, karena unit testing kita harus selalu memanggil third party service, dan belum tentu responsenya sesuai dengan apa yang kita mau
- Pada kasus seperti ini , cocok sekali untuk menggunakan mock object


Testify Mock
- Untuk membuat mock object tidak ada fitur bawaan golang, namun kita bisa menggunakan library testify yang sebelumnya kita gunakan untuk assertion
- Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
- Namun, perlu diperhatikanm, jika desain kode program kita jelek, akan sulit untuk melakukan mocking, jadi pastikan kita melkukan pembuatan desain kode program kita dengan baik


// Contoh kasus
Aplikasi query ke database
- kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke database
- dimana kita akan buat layer service sebagai business logic, dan layer repository sebagai jembatan ke database
- agar kode kita muddah untuk di test, disarankan agar membuat kontrak berupa Interface

**/

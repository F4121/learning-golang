package channel

/*
Pengenalan Channel
- Channel adalah tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine
- di Channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda
- Saat melakukan pengiriman data ke Channel, goroutin akan ter-block, sampai ada yang menerima data tersebut
- Maka dari itu channel disebut sebagai alat komunikasi synchronous (blocking)
- Channel cocok sekali sebagai alternative seperti mekanisme async await yang terdapat di beberapa bahasa pemrograman lain


Diagram channel
==============================
go -> channel -> go
==============================
- channel adalah tempat penyimpanan data yang bisa digunakan oleh goroutine
- goroutine yang mengirim data ke channel akan ter-block sampai ada yang menerima data tersebut
- goroutine yang menerima data dari channel akan ter-block sampai ada data yang dikirimkan


Karakteristik Channel
- Secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan lagi, kita harus menunggu data yang ada di channel diambil
- Channel hanya bisa menerima satu jenis data
- Channel bisa diambil dari lebih dari satu goroutine
- Channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak


Membuat Channel
- Channel di golang direpresentasikan dengan tidak data chan
- Untuk membuat channel sangatlah mudah, kita bisa menggunakan make(), mirip ketika membuat map
- Namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan ke dalam channel tersebut


Mengirim dan Menerima Data dari Channel
- Seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
- Untuk mengirim data, kita bisa gunakan kode: channel <- data
- Sedangkan untuk menerima data, bisa gunakan kode: data <- channel
- Jika selesai, jangan lupa untuk menutup channel menggunakan function close()


Channel Sebagai Parameter
- Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
- Sebelumnya kita tahu bahwa di golang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer (agar pass by reference)
- Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut


Channel In dan Out
- Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut
- Kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk mengirim data, atau hanya dapat digunakan untuk menerima data
- Hal ini bisa kita lakukan di parameter dengan cara menandai apakan channel ini digunakan untuk in (mengirimn data) atau out (menerima data)



Buffered Channel
- Seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
- Artinya jika kita menambahkan data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
- Kadang-kadang ada kkasus dimana pengirim lebih cepat dibandingkan penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
- Untungnya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampng data antrian di Channel

Buffer Capacity
- Kita bebas memasukkan bebera jumlah kapasitas antrian dalam buffer
- Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer
- Jika kita mengirimkan data ke 6, maka kita diminta untuk menunggu sampai buffernya ada yang kosong
- Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirimkan datas

Diagram Channel Buffer
=================================
go ->     channel       -> go
	 [1] [2] [3] [4] [5]
| | | |||||buffer||||| | | | |
=================================



Range Channel
- Kadang kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
- Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
- Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
- Ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
- Ini lebih sederhana dari pada kita melkukan pengecekan channel secara manual



Select Channel
- Kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
- Lalu kita ingin mendapatkan data dari semua channel tersebut
- Untuk melakukan hal tersebut, kita bisa menggunakan select channel dari golang
- Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random


Default Select
- Apa yang terjadi jika kita melakukan select terhadap channel yang ternyata tidak ada datanya
- Maka kita akan menunggu sampai ada datanya
- Kadang mungkin kita ingin melakukan sesuatu jika misal semua channel tidak ada datanya ketika kita melakukan select channel
- Dalam select, kita bisa menambahkan default, dimana ini akan dieksekusi jika memang di semua channel yang kita select tidak ada datanya

**/

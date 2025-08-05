package concurrencyandparallelgo

/*
Pengenalan Parallel Programming
- Saat ini kita hidup di era multicore, dimana jarang sekali kita menggunakan prosesor yang single core
- Semakin canggih hardware, maka software pun akan mengikuti, dimana sekarang kita bisa dengan mudah membuat proses parallel diaplikasi
- Parallel programming sederhananya adalah memecah suatu masalah dengan cara membaginya menjadi yang lebih kecil, dan dijalankan secara bersamaan pada waktu yang bersamaan pula

Contoh parallel
- Menjalankan beberapa aplikasi sekaligus di sistem operasi kita (office, editor, browser, dan lain-lain)
- Beberapa koki memyiapkan makanan di restoran, dimana koki membuat makanan masing-masing
- Antrian di bank, dimana tiap teller melayani nasabannya masing-masing

Saat belajar parallel programming kita akan mengenal Process dan Thread

Process
- Process adalah sebuah eksekusi program
- Process mengkonsumsi memory besar
- Process saling terisolasi dengan process lain
- process lama untuk dijalankan dihentikan

contoh process
- seperti kita membuka chrome, spotify, dan lain-lain

Thread
- Thread adalah segmen dari process
- Thread menggunakan memory kecil
- Thread bisa saling berhubungan jika dalam process yang lama
- Thread cepat untuk dijalankan dan dihentikan

contoh thread
- misal didalam chrome kita membuka beberapa thread


Parallel vs sConcurrency
- Berbeda dengan paralel (menjalan beberapa pekerjaan secara bersamaan), concurrencry adalah menjalankan beberapa pekerjaan secara bergantian
- Dalam parallel kita biasanya membutuhkan banyak Thread, sedangkan dalam concurrency, kita hanya membutuhkan sedikit Thread


Contoh concurrency
- Saat kita makan di cafe, kita bisa makan, lalu ngobrol, lalu minum, makan lagi, ngobrol lagi, minum lagi dan seterusnya
tetapi kita tidak bisa pada saat yang bersamaan minum, makan dan ngobrol, hanya bisa melakukkan satu hal pada saat waktu, namun bisa bergantian kapanpun kita mau

CPU-bound
- Banyak algoritma dibuat yang hanya membutuhkan CPU untuk menjalankannya. Algoritma jenis ini biasanya sangat tergantung dengan kecepatan CPU
- Contoh paling populer adalah Machine Learning, oleh karena itu sekarang banyak sekali teknologi Machine Learning yang banyak menggunakan GPU karena memiliki core yang lebih banyak dibandingkan CPU biasanya
- Jenis algoritma seperti ini tidak ada benefitnya menggunakan concurrency programming, namun bisa dibantu dengan implementasi Parallel Programming.


I/O-Bound
- I/O-bound adalah kebalikan dari sebelumnya, dimana biasanya algoritma atau aplikasi sangat tergantung dengan kkecepatan input output device yang digunakan
- Contoh aplikasi seperti membaca data dari file, database dan lain-lain.
- Kebanyakan saat ini biasanya kita akan membuat aplikasi jenis seperti ini
- Aplikasi jenis I/O-bound, walaupun bbisa terbantu dengan implementasi parallel programming, tapi benefitnya akan lebih baik jika menggunakan concurrency programming.
- Bayangkan kita membaca data dari database, dan Thread harus menunggu 1 detik untuk mendapatkan balasan dari database, padahal waktu 1 detik itu jika menggunakan conccurency programming, bisa digunakan untuk melakukan hal lain lagi




**/

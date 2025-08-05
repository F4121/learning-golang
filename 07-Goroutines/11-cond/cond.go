package cond

/*
Cond
- Cond adalah implementasi locking berbasis kondisi
- Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi lockingnya,
namun berbeda dengan Locker biasanya. di Cond terdapat function Wait() untuk menunggu apakah perlu menunggu atau tidak
- Function Signal() bisa digunakan untuk membberi tahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan function Broadcast() digunakan untuk
memeberi tahu semua goroutine agar tidak perlu menunggu lagi
- Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)
**/

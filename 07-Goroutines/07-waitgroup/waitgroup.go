package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

/*
WaitGroup
- WaitGroup adalah fitur yang bisa digunakan untuk menuggu sebuah proses selesai dilakukan
- Hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine,
tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
- Kasus seperti ini bisa menggunakan WaitGroup
- Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setelah proses goroutine selesai,
kita bisa gunakan method Done()
- Untuk menunggu semua proses selesai, kita bisa menggunakan method Wait()

**/

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)

	/*
		jangan lupa untuk menambbahkan group.Done() karena jika lupa maka function akan wait atau terjadi deadlock
	*/

}

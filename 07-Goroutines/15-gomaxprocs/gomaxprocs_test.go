package gomaxprocs

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/*
GOMAXPROCS
- Sebelumnya diawal kita sudah bahas bahwa goroutine itu dijalankan didalam Thread
- Pertanyaannya, sbeerapa banyak Thread yang ada di golang ketika aplikasi kita berjalan ?
- Untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function
didalam package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah thread
- Secara default, jumlah thread di golang itu sebanyak jumlah CPU di komputer kita
- Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()


* jika ingin melihat jumlah thread gunakan function runtime.GOMAXPROCS(-1) wajib dibawah 0 karena jika di atas 0
maka kita akan mengubah jumlah thread yang digunakan

**/

func TestGetGomaxprocs(t *testing.T) {
	// Simulasi goroutine
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU :", totalCpu)

	// ingat , GOMAXPROCS(-1) untuk melihat jumlah thread
	totalThread := runtime.GOMAXPROCS(-1)

	// sample ubah jumlah thread menjadi 20
	// runtime.GOMAXPROCS(20)

	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine :", totalGoroutine)

	group.Wait()
}

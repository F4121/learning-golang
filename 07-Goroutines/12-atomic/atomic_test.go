package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
Atomic
- Golang memiliki package yang bernama sync/atomic
- Atomic merupakan package yang digunankan untuk menggunakan data primitive secara aman pada proses concurrent
- Contohnya sbebelumnya kita telah mengguakan Mutex untuk melakukan locking ketika ingin menaikkan angka di counter.
hal ini sebenernya bisa digunakan menggunakan Atomic package
- Ada banyak sekali function di atomic package, kita bisa explore sendiri dihalaman dokumentasinya
- https://golang.org/pkg/sync/atomic/


**/

func TestRaceConditionAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				// Ketika menggunakan atomic maka tidak akan terjadi race condition
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Done = ", x)
}

package racecondition

import (
	"fmt"
	"testing"
	"time"
)

/*
Masalah Dengan Goroutine
- Saat kita menggunakan goroutine, dia tidak hanya berjalan secara conccurent, tapi bisa parallel juga
karena bisa ada beberapa thread yang berjalan secara parallel
- Hal ini sangat berbahaya ketika kita melkukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan
- Hal ini bisa menyebabkan masalah yang namanya Race Condition

Untuk mengatasi race condition, bisa lanjut ke materi Mutex :)
**/

// Sample Race Condition
func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x) //result should be 100000, because race condition value will be randomly under 100000. x will be assign with many goroutine before other goroutine already done assign x value
}

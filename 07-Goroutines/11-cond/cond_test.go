package cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitConditiion(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

/*
Ketika kita menggunakan cond.Signal() maka akan jalan 1 per 1
berbeda jika kita menggunakan cond.Broadcast() maka akan jalan semua tanpa kondisi wait


notes:
kondisi tidak berurutan pada hasil print terjadi karena aplikasi dijalankan secara concurrent dan parallel
jadi itu normal

**/

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitConditiion(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// Sample Broadcast
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}

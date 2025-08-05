package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Ketika kita menggunakan data pada pool setelah kita Get() wajib melakukan Put(data)
jika kita tidak melakukan Put / menggembalikan data ke pool maka pool akan kosong atau nil
*/

func TestPool(t *testing.T) {
	// pool := sync.Pool{}

	/*
		contoh pool dengan default value return New
		pool := sync.Pool{
			New: func() interface{} {
				return "New"
			},
		}
		**/
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("John")
	pool.Put("Doe")
	pool.Put("Programmer")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("Data ", data)
			// time sleep untuk simulasi datanya habis
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Complete")

}

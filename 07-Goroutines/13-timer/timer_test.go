package timer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Timer
- Timer adalah representasi satu kejadian (misal nya, beritahu saya ketika 10 detik kemudian)
- Ketika waktu timer sudah expire, maka event akan dikirimkan ke dalam channel
- Untuk membuat Timer kita bisa menggunakan time.NewTimer(duration)


time.After()
- Kadang kita hanya butuh channelnya saja, tidak membutuhkan data Timer nya
- untuk melakukan hal itu kita bisa menggunakan function time.After(duration)


time.AfterFunc()
- Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
- Kita bisa memanfaatkan Timer dengan menggunakan ffunction time.AfterFunc()
- Kita tidak perlu lagi menggunakan channelnya, cukup kirimkan function yang akan dipanggil ketika Timer mengirim kejadian

**/

func TestTimer(t *testing.T) {
	timer := time.NewTimer(10 * time.Second)
	fmt.Println("Timer 1 ==", time.Now())

	time := <-timer.C
	fmt.Println("Timer 2 ==", time)

	// Sample output akan terjadi perbbedaan 10 detik karena kita menggunakan timer
	// 2025-08-05 23:29:01.1618553 +0700 +07 m=+0.001648001
	// 2025-08-05 23:29:11.1618553 +0700 +07 m=+10.001648001

}

func TestTimerAfter(t *testing.T) {
	channel := time.After(10 * time.Second)
	fmt.Println("Timer 1 ==", time.Now())

	time := <-channel
	fmt.Println("Timer 2 ==", time)
}

func TestTimerAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(10*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()
}

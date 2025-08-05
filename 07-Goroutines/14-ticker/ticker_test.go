package ticker

import (
	"fmt"
	"testing"
	"time"
)

/*
time.Ticker
- Ticker adalah representasi kejadian yang berulang
- Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
- Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
- Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()


time.Tick()
- Kadang kita tidak butuh data Ticker nya, kita hanya butuh channelnya saja
- Jika demikian, kita bisa menggunakan function timerTick(duration), function ini tidak akan mengembalikan Ticker, hanya mengembalikan channel timernya saja

**/

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}

}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}

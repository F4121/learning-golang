package channel

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	// mengirim data
	channel <- "Faiz"

	// menerima data
	data := <-channel
	fmt.Println(data) // Faiz

	defer close(channel)
}

func TestCreateChannel2(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Faiz Programming"
		fmt.Println("Done send data to channel")
	}()

	// Jika channel diatas tidak diambil datanya makan akan lock dan otomatis printLn dibawah tidak akan ditampilan
	// Begitu juga jika data dibawah diambil dari channel tapi tidak ada data yang tersimpan dichannel maka akan lock juga

	data := <-channel
	fmt.Println("Show Data from Channel", data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Faiz Go Programming"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println("Show Data from Channel", data)
	time.Sleep(5 * time.Second)

}

// Send Only
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Faiz Go Programming"
}

// Receive Only
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("Show Data from Channel: ", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 5) // buffered channel with capacity 5
	defer close(channel)
	fmt.Println(cap(channel)) // melihat panjang buffer: 5
	fmt.Println(len(channel)) // melihat jumlah data di buffer: 0

	go func() {
		channel <- "Go"
		channel <- "Lang"
		channel <- "Programming"
		channel <- "Language"
		channel <- "Faizi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	/*
		Sebagai contoh kita mendeklarasikan channel dengan buffer 5
		Simulasi
		1. Jika kita mendeklarasikan channel tanpa parameter buffer, ketika function dijalankan akan error
		   deadlock karena kita mengirim data ke channel tapi tidak ada penerimanya sehingga terjadi kondisi lock
		2. Begitu juga ketika kita mendeklarasikan channel misal buffernya 1 tapi kita assign 2 data ke channel maka terjadi lock juga
		3. Adapula ketika kita salah misal buffer cap ada 3 tapi kita berusaha untuk menggambil data buffer ke 4 akan menyebabkan lock
		**/

	time.Sleep(2 * time.Second)
	fmt.Println("Done")

}

// Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Loop " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Data :", data)
	}

	fmt.Println("Done")
}

// Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2 ", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2 ", data)
			counter++
		default:
			fmt.Println("Waiting Data")
		}

		if counter == 2 {
			break
		}
	}
}

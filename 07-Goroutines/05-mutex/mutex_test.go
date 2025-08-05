package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Mutex (Mutual Exclusion)
- Untuk mengatasi masalah race condition tersebut di golang terdapat sebuah struct bernama sync.Mutex
- Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking terhadap mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock
- Dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang diperbolehkan, setelah goroutine tersebut melakukan unlock, baru goroutine selanjutnya diperbolehkan melakukan lock lagi
- Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi


RWMutex (Read Write Mutex)
- Kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data tapi juga membaca data
- Kita sebenarnya bisa menggunakan mutex saja, namun masalahnya nanti akan rebutan antara proses membaca dan mengubah
- Digolang telah disediaan struct RWMutex (Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write


**/

// Sample Handle Race Condition Using Mutex
func TestRaceConditionHandleWithMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x) //result 100000, no race condition because handle with Mutex lock and unlock
}

// Sample RWMutex
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	/*
		hasil dari loopingan diatas harusnya 10000 dan race condition telah aman dihandle dengan RWMutex pada AddBalance dan
		looping GetBalance pun berurutan berkan RLock sehingga tidak terjadi race condition
	**/

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance:", account.GetBalance())
}

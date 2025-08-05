package deadlock

import (
	"fmt"
	"testing"
	"time"
)

/*
func TestDeadlock ini adalah contoh deadlock dimana user1 dan user2 saling menunggu karena terjadi lock di function Transfer.
di goroutine 1 user2.Lock() tidak dapat di execute karena user2 sedang melakukan lock pada goroutine 2. dan kondisi ini yang dinamakan deadlock karena
tidak ada func yang jalan atau saling menunggu
**/

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "User1",
		Balance: 400,
	}
	user2 := UserBalance{
		Name:    "User2",
		Balance: 500,
	}

	//goroutine 1
	go Transfer(&user1, &user2, 200)
	//goroutine 2
	go Transfer(&user2, &user1, 100)

	time.Sleep(3 * time.Second)

	fmt.Println("User ", user1.Name, "Balance: ", user1.Balance)
	fmt.Println("User ", user2.Name, "Balance: ", user2.Balance)
}

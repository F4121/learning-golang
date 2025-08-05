package beforeandaftertest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
BEFORE dan After Test
- Biasanya dalam unit test, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi
- Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antar unit test function, maka membuat manual di unit test functionnya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya
- untungnya di golang terdapat fitur yang bernama testing.M
-fitur ini bernama main, dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa digunakan untuk melakuan before dan after tests


testing.M
- Untukk mengatur esekusi unit test, kita cukup membuat sebuah function bernama TestMain dengan parameter testing.M
- Jika terdapat function TestMain tersebut, makka secara otomatis golang akan mengeksekusi function ini tiap kali akan menjalankan unit test di sebuah package
- Dengan ini kita bisa mengatur Before dan After unit test sesuai dengan yang kita mau
- Ingat, function TestMain itu dieksekusi hanya seali per Golang package, bukan per tiap function unit test
**/

func TestMain(m *testing.M) {

	// Before Test
	fmt.Println("Before test")

	m.Run()

	// After test
	fmt.Println("After Test")

}

func HelloWorldTest(t *testing.T) {
	result := HelloWorld("Faiz")
	require.Equal(t, "Hello Faiz", result)
}

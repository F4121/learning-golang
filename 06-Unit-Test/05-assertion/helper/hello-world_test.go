package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Assertion
/*
- Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan
- Apalagi jika result data yang harus di cek itu banyak
- Oleh karena itu, sangat disarankan untuk menggunakan assertion untuk melakukan pengecekan
- Sayangnya, Go-Lang tidak menyediakan package untuk assertion, sehingga butuh menambahkan library untuk melakukan assertion ini

- Salah satu library assertion yang paling populer gi Golang adalah Testify
- Kita bisa menggunakan library ini untuk melakukan assertion terhadap result data di unit test
- https://github.com/stretchr/testify
- Kita bisa menambahkannya di Go Module kita:
  go get github.com/stretchr/testify


  Assert VS require
  - Testify menyediakan dua packae untuk assertion yaitu assert dan require
  - saat kita menggunakan assert, jika pengecekan gagal, maka assert akan memanggil Fail(), artinya eksekusi unit test akan tetap dilanjutkan
  - sedangkan jika kita menggunakan require, jika pengecekan gagal, maka require akan memanggil FailNow(), artinya eksekusi unit test tidak akan dilanjutkan
*/

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorldAssert("Budi")
	assert.Equal(t, "Hello Joko", result, "Result must be 'Hello Joko'")
	fmt.Println("TestHelloWorld Assert Done")
}

// Contoh require
func TestHelloWorldAssertRequire(t *testing.T) {
	result := HelloWorldAssert("Budi")
	require.Equal(t, "Hello Joko", result, "Result must be 'Hello Joko'")
	fmt.Println("TestHelloWorld Assert Success")
}

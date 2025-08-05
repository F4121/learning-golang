package subtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Menjalankan hanya sub test
- Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa menggunakan perintah go test -run TestNamaFunction
- Jika kita ingin menjalankan hanya salah satu sub test, kita bbisa gunakan perintah go test -run TestNamaFunction/NamaSubTest
- atau untuk menjalankan semua test semua sub test di semua function kita bisa gunakan perintah go test -run /NamaSubTest

**/

func TestHelloWorld(t *testing.T) {
	t.Run("Joko", func(t *testing.T) {
		results := HelloWorld("Joko")
		require.Equal(t, "Hello, Joko", results, "Result must be 'Hello, Joko'")
	})

	t.Run("Joki", func(t *testing.T) {
		results := HelloWorld("Joki")
		require.Equal(t, "Hello, Joki", results, "Result must be 'Hello, Joki'")
	})
}

package skiptest

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
SKIP TEST
- Kadang dalam keadaan tertentu, kita ingin membabtalkan eksekusi unit test
- di Golang kita juga bisa membatalkan eksekusi test jika kita mau
- Untuk membatalkan unit test kita bisa menggunakan function Skip()
**/

// contoh kode skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unit test tidak berjalan di Windows")
	}

	result := HelloWorld("Faiz")
	require.Equal(t, "Hello Faiz", result)
}

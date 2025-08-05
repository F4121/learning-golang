package benchmark

import "testing"

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Layla")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Miya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Miya")
		}
	})

	b.Run("Odete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Odete")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Bruno",
			request: "Bruno",
		},
		{
			name:    "Gatot Kaca",
			request: "Gatot Kaca",
		},
		{
			name:    "Balmond bin Abu Lahab",
			request: "Balmond bin Abu Lahab",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

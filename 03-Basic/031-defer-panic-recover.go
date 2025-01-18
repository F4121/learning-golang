package main

import "fmt"

func logging() {
	fmt.Println("Succes loggin")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups error!!")
	}
}

func endApp() {
	fmt.Println("End App")
}

func main() {
	// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
	// Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
	runApplication()

	// Panic function adalah function yang bisa kita gunakan untuk menghentikan program
	// Panic function biasanya dipanggil ketika terjadi panic pada saat prgoram kita berjalan
	// Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
	runApp(true)

	// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
	// Dengan recover proses panic akan terhentu, sehingga program akan tetap berjalan
	runAppWithRecover(true)
}

func endAppWithRecover() {
	fmt.Println("End App With Recover")
	message := recover()
	fmt.Println("Got Error ", message)

}

func runAppWithRecover(error bool) {
	defer endAppWithRecover()
	if error {
		panic("Error!!")
	}

}

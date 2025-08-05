package subtest

/*
Sub Test
- Golang mendukung fitur pembuatan function unit test didalam function unit test
- fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test bahasa pemrograman lain
- untuk membuat sub test, kita bisa menggunakan function Run()
**/

func HelloWorld(name string) string {
	return "Hello, " + name
}

package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("John")

	if result != "Hello John" {
		// error
		panic("Result is not Hello John")
	}
}

func TestHelloWorldJoko(t *testing.T) {
	result := HelloWorld("Joko")

	if result != "Hello Joko" {
		// error
		panic("Result is not Hello Joko")
	}
}

package helper

import (
	"fmt"
	"testing"
)

// untuk mengetahui perbedaanya jalan test 1 per 1

func TestHelloWorldJohn(t *testing.T) {
	result := HelloWorld("John")

	if result != "Hello John" {
		// error
		panic("Result is not Hello John")
	}
}

func TestHelloWorldJoko(t *testing.T) {
	result := HelloWorld("Fail")

	if result != "Hello Joko" {
		// error
		t.Fail()
	}
	fmt.Println("Test Hello World Joko Done")
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Fail")

	if result != "Hello Joko" {
		// error
		t.FailNow()
	}
	fmt.Println("Test Hello World Budi Done") // tidak akan tereksekusi jika menggunakan t.FailNow()

}

func TestHelloWorldJoko2(t *testing.T) {
	result := HelloWorld("Fail")

	if result != "Hello Joko" {
		// error
		t.Error("result must be 'Hello Joko' ")
	}
	fmt.Println("Test Hello World Joko Done")
}

func TestHelloWorldBudi2(t *testing.T) {
	result := HelloWorld("Fail")

	if result != "Hello Joko" {
		// error
		t.Fatal("result must be 'Hello Fail'")
	}
	fmt.Println("Test Hello World Budi Done") // tidak akan tereksekusi jika menggunakan t.Fatal()

}

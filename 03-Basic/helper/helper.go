package helper

var version = "1.0.0"      // Tidak bisa diakses dari luar
var Application = "Golang" //Bisa diakses dari luar

func sayGoodBye(name string) string { // ini juga tidak bisa diakses dari luar
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	/**
	PACKAGE flag
	- Package flag berisikan fungsionalitas untuk memparsing command line argument

	reference
	https://pkg.go.dev/flag
	*/

	var username *string = flag.String(
		"username",          // Flag name
		"root",              // Default value
		"database username", // Description
	)
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username", *username) // karna balikan dari flag adalah pointer maka untuk mendapatkan valuenya kita gunakan asterik operator jika tidak makan balikannya adalah pointer
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

	// run with
	// go run 004-flag.go -username="jhon" -password="p@sword" -host="123.3.3.3" -port="3000"

}

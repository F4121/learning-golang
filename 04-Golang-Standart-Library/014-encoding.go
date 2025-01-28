package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/**
	Package encoding
	- Golang menyediakan package encoding untuk melakukan encode dan decode
	- Golang menyediakan berbadagai macam algoritma untuk encoding, contoh populer adalah base64, csv dan json

	reference
	https://pkg.go.dev/encoding
	*/

	var value string = "Jhon Developer Handal"
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Println(string(decoded)) //string because convert from slice byte to string
	}

	// SAMPLE CSV READER
	csvString := "Jhon, Developer, Handal \n" +
		"Lala, Suka, Olahraga \n" +
		"Budi, Suka, Mengaji"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF { //EFO  = End of File
			break
		}
		fmt.Println(record)
	}

	// SAMPLE CSV WRITER
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Jhon", "Developer", "Handal"})
	_ = writer.Write([]string{"Jhon", "Developer", "GO"})
	_ = writer.Write([]string{"Jhon", "Developer", "Python"})

	writer.Flush()

}

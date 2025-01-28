package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/**
	Package bufio
	- package bufio atau singkatan dari buffered io
	- package ini digunakan untuk membuat data IO seperti Reader dan Writer

	reference
	https://pkg.go.dev/bufio
	*/

	// Sample Reader from string
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	// Sample Writer
	writer := bufio.NewWriter(os.Stdout) //os.Stdout adalah tujuan yang mana akan dituliskan diterminal
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("welcome\n")

	writer.Flush()

}

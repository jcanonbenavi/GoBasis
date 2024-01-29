package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	text := "Hello World"
	// Create a Reader from a string
	reader := strings.NewReader(text)
	// Create a buffer to store the data
	buffer := make([]byte, 8)

	for {
		// Read from the reader into the buffer
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		// Print the buffer as a string
		fmt.Println(string(buffer[:n]))
	}

}

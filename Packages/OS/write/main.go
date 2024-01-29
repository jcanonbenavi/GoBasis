package main

import (
	"fmt"
	"os"
)

func main() {
	// O_CREATE: create the file if it doesn't exist
	// O_WRONLY: open the file for writing only
	// O_APPEND: append data to the file when writing
	f, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println(err)
		return
	}
	// write data to the file using the Write() method. It's write a slice of bytes
	f.Write([]byte("Hello World, I'm Alexandra!\n"))
}

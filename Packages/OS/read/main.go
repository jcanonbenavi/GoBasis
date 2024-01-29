package main

import (
	"fmt"
	"os"
)

func main() {
	// open the file for reading only
	f, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// close the file when the function returns
	defer f.Close()
	// read data from the file using the Read() method. It's read a slice of bytes
	// with length 8
	buffer := make([]byte, 8)
	for {
		// read data from the file
		n, err := f.Read(buffer)
		if err != nil {
			// check if the end of file is reached
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			return
		}
		// print the data read from the file
		//:n is to print only the bytes read
		fmt.Println(string(buffer[:n]))
	}
}

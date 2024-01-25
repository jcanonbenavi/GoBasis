package internal

import (
	"fmt"
	"io"
	"os"
)

func Readfile(file string) (bytes []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic: %v", r)
		}
	}()
	data, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	bytes, err = io.ReadAll(data)
	data.Close()
	return

}

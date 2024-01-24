package internal

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// IS
func ErrorExample() {
	errBase := errors.New("error base")
	errWrap := fmt.Errorf("error wrap: %w", errBase)
	fmt.Println(errors.Is(errWrap, errBase)) //true
}

//UnWrap

func ErrorExampleUnWrap() {
	errBase := errors.New("error base")
	errWrap := fmt.Errorf("error wrap: %w", errBase)
	if errBase == errors.Unwrap(errWrap) {
		fmt.Println("error base")
	}
}

// As
func ErrorExampleAs() {
	_, err := os.Open("no-file.txt")
	var pathError *fs.PathError
	if errors.As(err, &pathError) {
		fmt.Println(pathError.Err)
	}
}

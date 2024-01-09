package main

import (
	"fmt"

	internal "github.com/jcanonbenavi/app/internal/exercises"
)

func main() {
	result := internal.Load(20000)
	fmt.Println(result)
}

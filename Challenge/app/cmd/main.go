package main

import (
	"fmt"

	"github.com/jcanonbenavi/app/internal/tickets"
)

func main() {
	total, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(total)

}

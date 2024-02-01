package tickets

import (
	"fmt"

	"github.com/jcanonbenavi/app/internal"
	"github.com/jcanonbenavi/app/internal/repository"
)

type TicketsMap struct {
	db map[int]internal.Ticket
}

func NewTicketsMap() *TicketsMap {
	return &TicketsMap{
		db: make(map[int]internal.Ticket),
	}
}

// ejemplo 1
func GetTotalTickets(destination string) (total int, err error) {
	values, err := repository.LoadDataFromCSV()
	if err != nil {
		err = fmt.Errorf("error loading data from csv: %w", err)
		return
	}
	for _, value := range values {
		if value.Destination == destination {
			total++
		}
	}
	return
}

// ejemplo 2
func GetMornings(time string) (int error) {
	return
}

// ejemplo 3
func AverageDestination(destination string, total int) (int error) {
	return
}

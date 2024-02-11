package tickets

import (
	"fmt"
	"time"

	"github.com/jcanonbenavi/app/internal"
)

type TicketsMap struct {
	db map[int]internal.Ticket
}

func NewTicketsMap(db map[int]internal.Ticket) *TicketsMap {
	return &TicketsMap{
		db: db,
	}
}

// ejemplo 1
func (t *TicketsMap) GetTotalTickets(destination string) (total int, err error) {
	for _, value := range t.db {
		if value.Destination == destination {
			total++
		}
	}
	return
}

func StringToTime(ticketsMap map[int]internal.Ticket) (hours []time.Time, err error) {
	layout := "15:04"
	for _, value := range ticketsMap {
		parsehours, parseErr := time.Parse(layout, value.Flight_time)
		if parseErr != nil {
			err = fmt.Errorf("error parsing time: %w", err)
			return
		}

		hours = append(hours, parsehours)

	}
	return

}

// ejemplo 2
func (t *TicketsMap) GetMornings(timeString string) (total int, err error) {
	hours, err := StringToTime(t.db)
	if err != nil {
		err = fmt.Errorf("error getting hours: %w", err)
	}
	morningStart, _ := time.Parse("15:04", "0:00")
	morningEnd, _ := time.Parse("15:04", "12:00")
	afternoonEnd, _ := time.Parse("15:04", "18:00")
	nightEnd, _ := time.Parse("15:04", "24:00")

	for _, value := range hours {
		if timeString == "morning" && value.After(morningStart) && value.Before(morningEnd) {
			total++
		}
		if timeString == "afternoon" && value.After(morningEnd) && value.Before(afternoonEnd) {
			total++
		}
		if timeString == "night" && value.After(afternoonEnd) && value.Before(nightEnd) {
			total++
		}

	}
	return
}

// ejemplo 3
func (t *TicketsMap) AverageDestination(destination string, total int) (average float64, err error) {
	totalByDestination := 0
	for _, value := range t.db {
		if value.Destination == destination {
			totalByDestination++
		}
	}

	average = (float64(totalByDestination) / float64(len(t.db))) * 100
	return
}

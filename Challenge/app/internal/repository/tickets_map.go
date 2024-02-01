package repository

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/jcanonbenavi/app/internal"
)

func LoadDataFromCSV() (ticketsMap map[int]internal.Ticket, err error) {
	// Open the file
	file, err := os.Open("tickets.csv")
	if err != nil {
		err = errors.New("Error opening file")
		return
	}
	defer file.Close()
	//initialize the map
	ticketsMap = make(map[int]internal.Ticket)
	// Read the file
	rd := csv.NewReader(file)
	// Read the records
	record, err := rd.ReadAll()
	if err != nil {
		err = errors.New("Error reading file")
		return
	}
	// Iterate through the records
	for _, value := range record {
		price, err := strconv.Atoi(value[5])
		if err != nil {
			return nil, errors.New("Error converting price")
		}
		ID, err := strconv.Atoi(value[0])
		if err != nil {
			return nil, errors.New("Error converting ID")
		}

		ticket := internal.Ticket{
			Id:          ID,
			Name:        value[1],
			Email:       value[2],
			Destination: value[3],
			Flight_time: value[4],
			Price:       price,
		}
		// Add the ticket to the map

		ticketsMap[ID] = ticket

	}
	return
}

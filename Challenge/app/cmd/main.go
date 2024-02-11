package main

import (
	"fmt"

	"github.com/jcanonbenavi/app/internal/repository"
	"github.com/jcanonbenavi/app/internal/tickets"
)

func main() {
	values, err := repository.LoadDataFromCSV()
	if err != nil {
		fmt.Println(err)
		return
	}

	ticketsRepository := tickets.NewTicketsMap(values)
	country := "Brazil"
	totalByDestination, err := ticketsRepository.GetTotalTickets(country)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total tickets to", country, ":", totalByDestination)

	totalMorningTickets, err := ticketsRepository.GetMornings("morning")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total tickets", ":", totalMorningTickets)

	AveragePerDestination, err := ticketsRepository.AverageDestination(country, len(values))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Average tickets to", country, ":", AveragePerDestination)

}

package tickets_test

import (
	"testing"

	"github.com/jcanonbenavi/app/internal"
	"github.com/jcanonbenavi/app/internal/tickets"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("Success: GetTotalTickets", func(t *testing.T) {

		// Arrange
		// Database
		st := tickets.NewTicketsMap(
			map[int]internal.Ticket{
				1: {
					Id:          1,
					Name:        "John Doe",
					Email:       "gmail.com",
					Destination: "London",
					Flight_time: "6:00",
					Price:       100,
				},
				2: {
					Id:          2,
					Name:        "Jane Doe",
					Email:       "hotmail.com",
					Destination: "Paris",
					Flight_time: "8:00",
					Price:       200,
				},
				3: {
					Id:          3,
					Name:        "John Doe",
					Email:       "gmail.com",
					Destination: "London",
					Flight_time: "6:00",
					Price:       100,
				},
			},
		)
		// Act
		destination, err := st.GetTotalTickets("London")
		if err != nil {
			t.Error(err)
		}
		// assert
		expected := 2
		require.Equal(t, expected, destination)
		require.NoError(t, err)
	})

	t.Run("GetMornings", func(t *testing.T) {
		// Arrange
		// Database
		st := tickets.NewTicketsMap(
			map[int]internal.Ticket{
				1: {
					Id:          1,
					Name:        "John Doe",
					Email:       "gmail.com",
					Destination: "London",
					Flight_time: "6:00",
					Price:       100,
				},
				2: {
					Id:          2,
					Name:        "Jane Doe",
					Email:       "hotmail.com",
					Destination: "Paris",
					Flight_time: "15:00",
					Price:       200,
				},
				3: {
					Id:          3,
					Name:        "John Doe",
					Email:       "gmail.com",
					Destination: "London",
					Flight_time: "6:00",
					Price:       100,
				},
			},
		)
		// Act
		destination, err := st.GetMornings("morning")
		if err != nil {
			t.Error(err)
		}
		// assert
		expected := 2
		require.Equal(t, expected, destination)
		require.NoError(t, err)
	})

	t.Run("AverageDestination", func(t *testing.T) {
		// Arrange
		// Database
		db := map[int]internal.Ticket{
			1: {
				Id:          1,
				Name:        "John Doe",
				Email:       "gmail.com",
				Destination: "London",
				Flight_time: "6:00",
				Price:       100,
			},
			2: {
				Id:          2,
				Name:        "Jane Doe",
				Email:       "hotmail.com",
				Destination: "Paris",
				Flight_time: "15:00",
				Price:       200,
			},
			3: {
				Id:          3,
				Name:        "John Doe",
				Email:       "gmail.com",
				Destination: "London",
				Flight_time: "6:00",
				Price:       100,
			},
		}
		st := tickets.NewTicketsMap(db)
		// Act
		destination, err := st.AverageDestination("London", len(db))
		if err != nil {
			t.Error(err)
		}
		// assert
		expected := 66.66666666666666
		require.Equal(t, expected, destination)
		require.NoError(t, err)
	})
}

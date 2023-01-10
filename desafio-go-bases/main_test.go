package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockBookings = []*Ticket{
	{
		1,
		"pedro",
		"pedro@email.com",
		"Colombia",
		10800,
		500,
	},
	{
		2,
		"juan",
		"juan@email.com",
		"Argentina",
		14400,
		500,
	},
	{
		3,
		"carlos",
		"carlos@email.com",
		"Argentina",
		14400,
		500,
	},
}

func TestGetTotalTickets_Success(t *testing.T) {
	var expectedResult = 1

	result := GetTotalTickets("Colombia", &mockBookings)

	assert.Equal(t, expectedResult, result)
}

func TestGetBookingsByTimeRange_Success1(t *testing.T) {
	var expectedResult = 3

	result, err := GetBookingsByTimeRange("3:00", "4:00", &mockBookings)

	assert.Equal(t, expectedResult, result)
	assert.NoError(t, err)
}

func TestGetBookingsByTimeRange_Success2(t *testing.T) {
	var expectedResult = 1

	result, err := GetBookingsByTimeRange("3:00", "3:30", &mockBookings)

	assert.Equal(t, expectedResult, result)
	assert.NoError(t, err)
}

func TestAverageDestionation_Success(t *testing.T) {
	var expectedResult = 1.0 / 3
	result := AverageDestination("Colombia", &mockBookings)

	assert.Equal(t, expectedResult, result)
}

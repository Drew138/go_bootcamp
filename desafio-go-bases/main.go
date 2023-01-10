package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	id      int
	name    string
	email   string
	country string
	time    int
	price   int
}

func timeStampToSeconds(timestamp string) (int, error) {
	tmp := strings.Split(timestamp, ":")
	if len(tmp) != 2 {
		return 0, errors.New("invalid timestamp provided")
	}
	total := 0
	val, err := strconv.Atoi(tmp[0])
	if err != nil {
		return 0, err
	}
	total += val * 3600
	val, err = strconv.Atoi(tmp[1])
	if err != nil {
		return 0, err
	}
	total += val * 60
	return total, nil
}

func readData(file string) (*[]*Ticket, error) {

	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(f)
	tickets := []*Ticket{}
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		ticket, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return &tickets, nil
}

func parseLine(line []string) (*Ticket, error) {
	if len(line) != 6 {
		return nil, errors.New("line must contain six entries")
	}
	id, err := strconv.Atoi(line[0])
	if err != nil {
		return nil, err
	}
	name := line[1]
	email := line[2]
	country := line[3]
	time, err := timeStampToSeconds(line[4])
	if err != nil {
		return nil, err
	}
	price, err := strconv.Atoi(line[5])
	if err != nil {
		return nil, err
	}
	return &Ticket{id, name, email, country, time, price}, nil
}

func main() {
	tickets, err := readData("tickets.csv")

	if err != nil {
		panic(err)
	}
	destination := "Finland"
	fmt.Printf("Total number of tickets for %s: %d\n", destination, GetTotalTickets(destination, tickets))
	start := "r:00"
	end := "6:00"
	total, err := GetBookingsByTimeRange(start, end, tickets)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total number of tickets between %s and %s: %d\n", start, end, total)
	average := AverageDestination(destination, tickets)
	fmt.Printf("Average number of tickets for %s in days: %f\n", destination, average)
}

func GetTotalTickets(destination string, tickets *[]*Ticket) int {
	total := 0
	for _, ticket := range *tickets {
		if ticket.country == destination {
			total++
		}
	}
	return total
}

func GetBookingsByTimeRange(start, end string, tickets *[]*Ticket) (int, error) {
	startSeconds, err := timeStampToSeconds(start)
	if err != nil {
		return 0, err
	}
	endSeconds, err := timeStampToSeconds(end)
	if err != nil {
		return 0, err
	}
	total := 0
	for _, ticket := range *tickets {
		if ticket.time >= startSeconds && ticket.time <= endSeconds {
			total++
		}
	}
	return total, nil
}

func AverageDestination(destination string, tickets *[]*Ticket) float64 {
	tmp := GetTotalTickets(destination, tickets)
	return float64(tmp) / float64(len(*tickets))
}

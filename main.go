package main

import "fmt"

// application entrypoint
func main() {
	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference" // short way
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T, conferenceTicktes is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	// conferenceName is string, conferenceTicktes is int, remainingTickets is uint

	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	var userName string
	var userTickets int

	userName = "John"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
}

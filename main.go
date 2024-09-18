package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your name:")
		fmt.Scanln(&firstName) // update the value by reference

		fmt.Println("Enter your last name:")
		fmt.Scanln(&lastName)

		fmt.Println("Enter your Email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		fmt.Printf("These are all our bookings: %v\n", bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("First names: %v\n", firstNames)
	}
}

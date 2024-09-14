package main

import "fmt"

func main() {
	const conferenceTickets = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	var name string
	var email string
	var userTickets uint

	fmt.Println("Enter you name:")
	fmt.Scanln(&name) // update the value by reference

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, name)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", name, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v\n", bookings)
}

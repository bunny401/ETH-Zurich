package main

import "fmt"

func main() {
	var conferenceName string= "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and ", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")


	var bookings []string
	var firstName string
	var lastName string
	var email string
	var userTicket int

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter Email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to buy")
	fmt.Scan(&userTicket)

	remainingTickets=remainingTickets - userTicket
	bookings= append(bookings, firstName +" " + lastName)




	fmt.Println("Thank you", firstName, lastName, "for booking", userTicket, "tickets. You will receive the confirmation at", email)
	fmt.Println(remainingTickets," tickets are remaining for the ",conferenceName)

}

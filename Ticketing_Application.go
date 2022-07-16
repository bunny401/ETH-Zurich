package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and ", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	for {
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

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicketNumber bool = userTicket > 0 && userTicket <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTicket
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Println("Thank you", firstName, lastName, "for booking", userTicket, "tickets. You will receive the confirmation at :", email)
			fmt.Println(remainingTickets, " tickets are remaining for the ", conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of the bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our confernce is fully booked. Come back next year")
				break
			}

		} else {
			fmt.Println("YOUR INPUT DATA IS INVALID. TRY AGAIN!")

		}

	}

}

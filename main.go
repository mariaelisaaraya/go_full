package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\nWe have total of %v tickets and %v are still available\nGet your tickets here to attend",
		conferenceName, conferenceTickets, remainingTickets)

	for {
		var firtsName string
		var lastName string
		var email string
		var userTickets uint

		//asking for user input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firtsName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of ticket: ")
		fmt.Scan(&userTickets)

		//ifvalidName := len(firtsName) >= 2 && len(lastName) >= 2

		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firtsName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firtsName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Out conference is booked out. Come back next year")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so yo can't book %v tickets\n", remainingTickets, userTickets)

		}
	}

}

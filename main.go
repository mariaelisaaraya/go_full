package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{} //now: list of map

func main() {

	greetUsers()

	for {
		firtsName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firtsName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firtsName, lastName, email)

			firtsNames := getFirstNames()
			fmt.Printf("The first name of bookings are: %v\n", firtsNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Out conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email adress  you entered doesn't contain @ sign")

			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")

			}
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking aplication\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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
	return firtsName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firtsName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// this create a map for user
	bookings = append(bookings, firtsName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firtsName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const ConferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", ConferenceTickets, remainingTickets)
	fmt.Println("Please enter your Name and Tickets Needed")

	var userName string //defining type
	var UserAge uint    // defining the int type. so, user can't book in negative
	fmt.Scan(&userName, &UserAge)
	fmt.Printf("the user %v booked %v tickets", userName, UserAge)
}

// For Loop

package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const ConferenceTickets = 50
	remainingTickets := 50
	var booking []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", ConferenceTickets, remainingTickets)

	for {

		var userName string //defining type
		var UserAge uint    // defining the int type. so, user can't book in negative

		fmt.Println("Please enter your Name and Tickets Needed")
		fmt.Scan(&userName, &UserAge)
		fmt.Printf("the user %v booked %v tickets\n", userName, UserAge)
		booking = append(booking, userName)
		fmt.Printf("total index %v\n", booking)
		fmt.Printf("total index %v\n", booking[0])
		fmt.Printf("total index type %T\n", booking)
		fmt.Printf("total index type %v\n", len(booking))
	}

}

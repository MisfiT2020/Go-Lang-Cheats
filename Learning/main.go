// If else

package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const ConferenceTickets = 50
	var remainingTickets uint = 50
	var booking []string
	var userTicksets uint

	for {

		var userName string //defining type
		//var UserAge uint    // defining the int type. so, user can't book in negative

		fmt.Printf("Welcome to %v booking application\n", conferenceName)

		fmt.Println("Please enter your Name and Tickets Needed")
		fmt.Scan(&userName, &userTicksets)
		fmt.Printf("the user %v booked %v tickets\n", userName, userTicksets)

		if userTicksets <= remainingTickets {
			remainingTickets := remainingTickets - userTicksets
			fmt.Printf("these tickets are available %v", remainingTickets)
			fmt.Printf("We have total of %v tickets and %v are still available\n", ConferenceTickets, remainingTickets)

			booking = append(booking, userName)
			fmt.Printf("total index %v\n", booking)
			fmt.Printf("total index %v\n", booking[0])
			fmt.Printf("total index type %T\n", booking)
			fmt.Printf("total index type %v\n", len(booking))

			if remainingTickets == 0 {
				fmt.Println("Our Conference Tickets has ended...")
				break
			}

		} else {
			fmt.Printf("you can't choose more than available tickets.... available tickets: %v", remainingTickets)
		}

	}

}

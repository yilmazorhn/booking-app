package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	conferenceTickets := 50
	var remainingTickets int = 30

	fmt.Printf("conferenceName %T, conferenceTickets is %T, remainingTickets is %T", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)

}

package main

import (
	"booking-app/helper"
	"booking-app/model"
	"fmt"
	"sync"
)

var conferenceName = "Go Conference"

const conferenceTickets uint32 = 4

var remainingTickets uint32 = conferenceTickets

func main() {
	var wg sync.WaitGroup
	bookings := make([]model.UserDetail, 0)
	fmt.Println("Muruga")
	greetUser()

Loop:
	for {
		var selection int16
		fmt.Println("Choose your option ")
		fmt.Println("1. Book Tickets")
		fmt.Println("2. Check Availability")
		fmt.Println("3. View Booked Tickets")
		fmt.Println("4. Exit")
		fmt.Print("Enter (1/2/3/4) : ")
		fmt.Scan(&selection)
		switch selection {
		case 1:
			ticket, err := bookTicket()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				remainingTickets -= ticket.NumberOfTickets
				bookings = append(bookings, ticket)
				wg.Add(1)
				go helper.SendEmail(&ticket.TicketHolders, &wg)
			}
		case 2:
			fmt.Printf("Number of remaining tickets %d\n", remainingTickets)
		case 3:
			helper.PrintBookedTickes(bookings)
		case 4:
			break Loop
		default:
			fmt.Println("Invalid option. Please try again")
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines completed")
	helper.PrintBookedTickes(bookings)
}

func greetUser() {
	fmt.Println("Welcome to", conferenceName, "booking application !!")
	fmt.Printf("We have total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)

}

func bookTicket() (model.UserDetail, error) {
	var firstName string
	var lastName string
	var email string
	var numberOfTickets uint32
	var ticketHolders = make(map[string][2]string) //key : string, value : array of string

	fmt.Print("Enter first name ")
	fmt.Scan(&firstName)
	fmt.Print("Enter last name ")
	fmt.Scan(&lastName)
	fmt.Print("Enter email ")
	fmt.Scan(&email)
	for {
		fmt.Print("Enter numberOfTickets ")
		fmt.Scan(&numberOfTickets)
		if numberOfTickets <= 0 || numberOfTickets > 5 {
			fmt.Println("Ticket count not in range")
		} else {
			break
		}
	}

	for i := 0; i < int(numberOfTickets); i++ {
		var name string
		var emailOfTicketHolder string
		fmt.Printf("Enter name of user %v : ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Enter email of user %v : ", i+1)
		fmt.Scan(&emailOfTicketHolder)

		var emailStatus [2]string
		emailStatus[0] = emailOfTicketHolder
		emailStatus[1] = "false" // email not sent yet
		ticketHolders[name] = emailStatus
	}

	user := model.UserDetail{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: numberOfTickets,
		TicketHolders:   ticketHolders}

	err := helper.ValidateUser(user, remainingTickets)
	if err != nil {
		return model.UserDetail{}, err
	}
	return user, nil
}

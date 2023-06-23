package helper

import (
	"booking-app/model"
	"errors"
	"fmt"
	"sync"
	"time"
)

func ValidateUser(user model.UserDetail, remainingTickets uint32) error {

	if len(user.FirstName) == 0 {
		return errors.New("first name cannot be empty")
	}

	if len(user.LastName) == 0 {
		return errors.New("last name cannot be empty")
	}

	if len(user.Email) == 0 {
		return errors.New("last name cannot be empty")
	}

	if user.NumberOfTickets > remainingTickets {
		msg := fmt.Sprintf("Cannot book %d tickets. Max available tickets is %v", user.NumberOfTickets, remainingTickets)
		return errors.New(msg)
	}
	return nil
}

func PrintBookedTickes(bookings []model.UserDetail) {
	header := "Name\tEmail\tNumber Of Tickets\tEmail Status"
	fmt.Println(header)

	for _, booking := range bookings {
		emailStatusMsg := ""
		for name, value := range booking.TicketHolders {
			emailStatusMsg += fmt.Sprintf("%s - %s, ", name, value[1])
		}
		row := fmt.Sprintf("%s\t%s\t%d\t%s", booking.FirstName, booking.Email, booking.NumberOfTickets, emailStatusMsg)
		fmt.Println(row)
	}

}

func SendEmail(ticketHolder *map[string][2]string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(10 * time.Second)
	for name, info := range *ticketHolder {
		info[1] = "true"
		time.Sleep(10 * time.Second)
		(*ticketHolder)[name] = info // Update the modified array back in the map
	}
}

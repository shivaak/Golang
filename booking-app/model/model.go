package model

type UserDetail struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint32
	TicketHolders   map[string][2]string
}

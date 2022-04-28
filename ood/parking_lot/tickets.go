package parking_lot

type TicketStatus int

type Ticket struct {
	TicketNumber string
	TicketStatus TicketStatus
}

func NewTicket(vehicleType VehicleType) *Ticket {
	return &Ticket{}
}

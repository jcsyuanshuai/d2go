package parking_lot

type VehicleBehavior interface {
	AssignTicket(t *Ticket)
}

type VehicleType int

const (
	Car       VehicleType = 1
	Truck     VehicleType = 2
	Electric  VehicleType = 3
	Van       VehicleType = 4
	Motorbike VehicleType = 5
)

type Vehicle struct {
	VehicleNumber string
	VehicleType   VehicleType
	Ticket        *Ticket
}

func (v *Vehicle) AssignTicket(t *Ticket) {
	v.Ticket = t
}

func NewVehicle(number string, vehicleType VehicleType) VehicleBehavior {
	return &Vehicle{
		VehicleNumber: number,
		VehicleType:   vehicleType,
	}
}

package parking_lot

type UserStatus int

const (
	UserStatusActive     UserStatus = 1
	UserStatusBlocked    UserStatus = 2
	UserStatusBanned     UserStatus = 3
	UserStatusCompromise UserStatus = 4
	UserStatusArchived   UserStatus = 5
	UserStatusUnknown    UserStatus = 6
)

type User struct {
	UserName string
	Password string
	Email    string
	Phone    string
	Address  string
	Status   UserStatus
}

type AdminBehavior interface {
	AddParkingFloor(floor ParkingFloor)
	AddParkingSpot(floor ParkingFloor, spot ParkingSpot)
	AddParkingDisplayBoard(floor ParkingFloor, board DisplayBoard)
	AddCustomerPanel(floor ParkingFloor, panel CustomerPanel)
	AddEntrancePanel(panel EntrancePanel)
	AddExitPanel(panel ExitPanel)
}

type Admin User

func (a *Admin) AddParkingFloor(floor ParkingFloor) {
	panic("implement me")
}

func (a *Admin) AddParkingSpot(floor ParkingFloor, spot ParkingSpot) {
	panic("implement me")
}

func (a *Admin) AddParkingDisplayBoard(floor ParkingFloor, board DisplayBoard) {
	panic("implement me")
}

func (a *Admin) AddCustomerPanel(floor ParkingFloor, panel CustomerPanel) {
	panic("implement me")
}

func (a *Admin) AddEntrancePanel(panel EntrancePanel) {
	panic("implement me")
}

func (a *Admin) AddExitPanel(panel ExitPanel) {
	panic("implement me")
}

func NewAdmin() AdminBehavior {
	return &Admin{}
}

type ParkingAttendantBehavior interface {
	ProcessTicket(ticketNumber string)
}

type ParkingAttendant User

func (p *ParkingAttendant) ProcessTicket(ticketNumber string) {
	panic("implement me")
}

func NewParkingAttendant() ParkingAttendantBehavior {
	return &ParkingAttendant{}
}

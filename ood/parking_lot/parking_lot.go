package parking_lot

import "sync"

type ParkingLotBehavior interface {
	IsFull(vehicleType VehicleType) bool
	IncrementSpotCount(vehicleType VehicleType)
	NewParkingTicket(vehicle *Vehicle) *Ticket
	AddParkingFloor(floor *ParkingFloor)
	AddEntrancePanel(panel *EntrancePanel)
	AddExitPanel(panel *ExitPanel)
}

type parkingLot struct {
	mu             sync.Mutex
	Name           string
	Address        string
	ParkingRate    float64
	SpotCount      map[ParkingSpotType]int
	MaxSpotCount   map[ParkingSpotType]int
	EntrancePanels []*Panel
	ExitPanels     []*Panel
	ParkingFloors  []*ParkingFloor
	ActiveTickets  map[string]*Ticket
}

func (p *parkingLot) IsFull(vehicleType VehicleType) bool {
	switch vehicleType {
	case Van:
	case Car:
		count := p.SpotCount[ParkingSpotTypeCompact] + p.SpotCount[ParkingSpotTypeLarge]
		maxCount := p.MaxSpotCount[ParkingSpotTypeCompact] + p.MaxSpotCount[ParkingSpotTypeLarge]
		return count >= maxCount
	case Truck:
		return p.SpotCount[ParkingSpotTypeLarge] >= p.MaxSpotCount[ParkingSpotTypeLarge]
	case Motorbike:
		return p.SpotCount[ParkingSpotTypeMotorbike] >= p.MaxSpotCount[ParkingSpotTypeMotorbike]
	case Electric:
		count := p.SpotCount[ParkingSpotTypeCompact] + p.SpotCount[ParkingSpotTypeLarge] + p.SpotCount[ParkingSpotTypeElectric]
		maxCount := p.MaxSpotCount[ParkingSpotTypeCompact] + p.MaxSpotCount[ParkingSpotTypeLarge] + p.MaxSpotCount[ParkingSpotTypeElectric]
		return count >= maxCount
	}
	panic("vehicle type not support")
}

func (p *parkingLot) IncrementSpotCount(vehicleType VehicleType) {
	switch vehicleType {
	case Truck:
	case Van:
		p.SpotCount[ParkingSpotTypeLarge] += 1
	case Motorbike:
		p.SpotCount[ParkingSpotTypeMotorbike] += 1
	case Car:
		if p.SpotCount[ParkingSpotTypeCompact] < p.MaxSpotCount[ParkingSpotTypeCompact] {
			p.SpotCount[ParkingSpotTypeCompact] += 1
		} else {
			p.SpotCount[ParkingSpotTypeLarge] += 1
		}
	case Electric:
		if p.SpotCount[ParkingSpotTypeElectric] < p.MaxSpotCount[ParkingSpotTypeElectric] {
			p.SpotCount[ParkingSpotTypeElectric] += 1
		} else if p.SpotCount[ParkingSpotTypeCompact] < p.MaxSpotCount[ParkingSpotTypeCompact] {
			p.SpotCount[ParkingSpotTypeCompact] += 1
		} else {
			p.SpotCount[ParkingSpotTypeLarge] += 1
		}
	}
}

func (p *parkingLot) NewParkingTicket(vehicle *Vehicle) *Ticket {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.IsFull(vehicle.VehicleType) {
		return nil
	}
	ticket := NewTicket(vehicle.VehicleType)
	vehicle.AssignTicket(ticket)
	p.IncrementSpotCount(vehicle.VehicleType)
	p.ActiveTickets[ticket.TicketNumber] = ticket
	return ticket
}

func (p *parkingLot) AddParkingFloor(floor *ParkingFloor) {
	panic("implement me")
}

func (p *parkingLot) AddEntrancePanel(panel *EntrancePanel) {
	panic("implement me")
}

func (p *parkingLot) AddExitPanel(panel *ExitPanel) {
	panic("implement me")
}

var p ParkingLotBehavior

func init() {
	p = &parkingLot{}
}

func IsFull(vehicleType VehicleType) bool {
	return p.IsFull(vehicleType)
}

func IncrementSpotCount(vehicleType VehicleType) {
	p.IncrementSpotCount(vehicleType)
}

func NewParkingTicket(vehicle *Vehicle) *Ticket {
	return p.NewParkingTicket(vehicle)
}

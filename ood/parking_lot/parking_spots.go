package parking_lot

type ParkingSpotBehavior interface {
	IsFree() bool
	AssignVehicle(vehicle *Vehicle)
	RemoveVehicle()
}

type ParkingSpotType string

const (
	ParkingSpotTypeHandicapped ParkingSpotType = "Handicapped"
	ParkingSpotTypeCompact     ParkingSpotType = "Compact"
	ParkingSpotTypeLarge       ParkingSpotType = "Large"
	ParkingSpotTypeMotorbike   ParkingSpotType = "Motorbike"
	ParkingSpotTypeElectric    ParkingSpotType = "Electric"
)

type ParkingSpot struct {
	SpotNumber      int
	Free            bool
	Vehicle         *Vehicle
	ParkingSpotType ParkingSpotType
}

func (p *ParkingSpot) IsFree() bool {
	return p.Free
}

func (p *ParkingSpot) AssignVehicle(vehicle *Vehicle) {
	p.Vehicle = vehicle
	p.Free = false
}

func (p *ParkingSpot) RemoveVehicle() {
	p.Vehicle = nil
	p.Free = true
}

func NewParkingSpot(spotType ParkingSpotType) ParkingSpotBehavior {
	return &ParkingSpot{
		ParkingSpotType: spotType,
	}
}

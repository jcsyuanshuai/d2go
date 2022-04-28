package parking_lot

type ParkingFloorBehavior interface {
	AddParkingSpot(spot *ParkingSpot)
	AssignVehicleToSpot(vehicle *Vehicle, spot *ParkingSpot)
	UpdateDisplayBoard(spot *ParkingSpot)
	FreeSpot(spot *ParkingSpot)
}

type SpotTable map[int]*ParkingSpot

type ParkingFloor struct {
	FloorId               int64
	FloorName             string
	TotalParkingSpotTable map[ParkingSpotType]SpotTable
	FreeParkingSpotCount  map[ParkingSpotType]int
	DisplayBoard          *DisplayBoard
}

func (p *ParkingFloor) AddParkingSpot(spot *ParkingSpot) {
	p.TotalParkingSpotTable[spot.ParkingSpotType][spot.SpotNumber] = spot
}

func (p *ParkingFloor) AssignVehicleToSpot(vehicle *Vehicle, spot *ParkingSpot) {
	spot.AssignVehicle(vehicle)
	p.UpdateDisplayBoard(spot)
}

func (p *ParkingFloor) UpdateDisplayBoard(spot *ParkingSpot) {
	freeSpot := p.DisplayBoard.FreeSpot[spot.ParkingSpotType]
	spotTable := p.TotalParkingSpotTable[spot.ParkingSpotType]
	if freeSpot.SpotNumber == spot.SpotNumber {
		for _, val := range spotTable {
			if val.IsFree() {
				freeSpot = val
			}
		}
		p.DisplayBoard.ShowEmptySpotNumber()
	}
}

func (p *ParkingFloor) FreeSpot(spot *ParkingSpot) {
	spot.RemoveVehicle()
	p.FreeParkingSpotCount[spot.ParkingSpotType] += 1
}

func NewParkingFloor() ParkingFloorBehavior {
	return &ParkingFloor{}
}

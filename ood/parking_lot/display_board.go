package parking_lot

import "fmt"

type DisplayBoardBehavior interface {
	ShowEmptySpotNumber()
}

type DisplayBoard struct {
	Id       int64
	FreeSpot map[ParkingSpotType]*ParkingSpot
}

func (d *DisplayBoard) ShowEmptySpotNumber() {
	for k, v := range d.FreeSpot {
		if v.IsFree() {
			fmt.Printf("[%s] is free.\n", k)
		} else {
			fmt.Printf("[%s] is full.\n", k)
		}
	}
}

func NewDisplayBoard() DisplayBoardBehavior {
	return &DisplayBoard{}
}

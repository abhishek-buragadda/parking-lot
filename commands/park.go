package commands

import (
	"fmt"
	"parkinglot/src/types"
)

type Park struct{}

func (p *Park) Execute(parkingLot *types.ParkingLot, args []string) {
	if len(args) < 2 {
		fmt.Printf("Not enough arguemnts. Required %d Given %d\n", 2, len(args))
		return
	}
	vehicle := types.NewVehicle(args[0], args[1])
	slot := parkingLot.GetFreeSlot()
	if slot != nil {
		slot.Park(vehicle)
		fmt.Printf("Allocated slot number: %d\n", slot.GetSlotId())
	} else {
		fmt.Println("Sorry, parking lot is full")
	}
}

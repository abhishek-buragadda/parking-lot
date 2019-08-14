package commands

import (
	"fmt"
	"parkinglot/src/types"
	"strconv"
)

type Leave struct{}

func (l *Leave) Execute(parkingLot *types.ParkingLot, args []string) {
	if len(args) < 1 {
		fmt.Printf("Required slot number")
		return
	}
	slotNumber, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		fmt.Printf("Not able to convert input to number\n")
	}
	slot := parkingLot.GetSlot(slotNumber)
	if slot != nil {
		if slot.IsFreeSlot(){
			fmt.Printf("Slot number %d already free\n", slot.GetSlotId())
			return
		}
		slot.UnPark()
		fmt.Printf("Slot number %d is free\n", slot.GetSlotId())
		return
	}
	fmt.Println("Slot number to Leave Not Found")
}

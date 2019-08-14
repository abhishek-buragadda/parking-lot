package commands

import (
	"fmt"
	"parkinglot/src/types"
	"strconv"
)

type Status struct{}

func (s *Status) Execute(parkingLot *types.ParkingLot, args []string){
	fmt.Println("Slot No.    Registration No    Colour")
	for _, slot := range parkingLot.GetSlots(){
		if !slot.IsFreeSlot(){
			slotId := strconv.FormatUint(slot.GetSlotId(),10)
			fmt.Printf("%-8s    %-15s    %-s\n", slotId,slot.GetVehicle().GetRegNumber(), slot.GetVehicle().GetColor()  )

		}
	}
}
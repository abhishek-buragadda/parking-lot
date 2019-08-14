package commands

import (
	"fmt"
	"parkinglot/src/types"
	"strconv"
)

type CreateParkingLot struct{}

func (p *CreateParkingLot) Execute(parkingLot *types.ParkingLot,args []string) {
	if len(args) >= 1 {
		size, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			fmt.Println("Parking Capacity should be a number")
			return
		}
		parkingLot.Initialize(size)
		fmt.Printf("Created a parking lot with %d slots\n", size)
	}else{
		fmt.Println("1 argument is required")
	}
}

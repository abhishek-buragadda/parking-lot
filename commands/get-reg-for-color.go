package commands

import (
	"fmt"
	"parkinglot/src/types"
	"strings"
)

type GetRegForColor struct{}

func (r *GetRegForColor) Execute(parkingLot *types.ParkingLot, args []string) {
	if len(args) < 1 {
		fmt.Println("Require 1 argument")
		return
	}
	vehicleRegNumbers := parkingLot.GetRegNumberByColor(args[0])
	if len(vehicleRegNumbers) == 0 {
		fmt.Println("Not found")
		return
	}
	fmt.Println(strings.Join(vehicleRegNumbers, ", "))
}

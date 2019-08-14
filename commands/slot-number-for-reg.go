package commands

import (
	"fmt"
	"parkinglot/src/types"
	"strings"
)

type GetSlotForReg struct {

}
func (s *GetSlotForReg) Execute(parkingLot *types.ParkingLot, args []string){
	if len(args) < 1 {
		fmt.Println("Require 1 argument")
		return
	}
	slots := parkingLot.GetSlotsByRegNumber(args[0])
	if len(slots) == 0 {
		fmt.Println("Not found")
		return
	}
	fmt.Println(strings.Join(slots,", "))
}
package commands

import (
	"fmt"
	"parkinglot/src/types"
)

type Command interface {
	Execute(parkingLot *types.ParkingLot, args []string)
}

var commands = map[string]Command{
"park": &Park{},
"leave": &Leave{},
"registration_numbers_for_cars_with_colour": &GetRegForColor{},
"slot_numbers_for_cars_with_colour": &GetSlotForColor{},
"slot_number_for_registration_number": &GetSlotForReg{},
"status": &Status{},
"create_parking_lot": &CreateParkingLot{},
}

func GetCommandByName(name string) Command{
	if command := commands[name]; command == nil {
		fmt.Println("command not found ")
		return nil
	} else {
		 return command
	}
}

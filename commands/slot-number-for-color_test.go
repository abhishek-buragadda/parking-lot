package commands

import (
	"parkinglot/src/types"
)

func ExampleGetSlotForColorIsPresent() {
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(4)
	vehicle1 := types.NewVehicle("1234","White")
	vehicle2 := types.NewVehicle("1264","Blue")
	parkinglot.GetSlot(1).Park(vehicle1)
	parkinglot.GetSlot(2).Park(vehicle2)
	c := &GetSlotForColor{}
	args := []string{"Blue"}
	c.Execute(parkinglot, args)
	// Output:
	// 2
}

func ExampleGetSlotForColorIsNotPresent() {
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(4)
	vehicle1 := types.NewVehicle("1234","White")
	vehicle2 := types.NewVehicle("1264","Blue")
	parkinglot.GetSlot(1).Park(vehicle1)
	parkinglot.GetSlot(2).Park(vehicle2)
	c := &GetSlotForColor{}
	args := []string{"Red"}
	c.Execute(parkinglot, args)
	// Output:
	// Not found
}

func ExampleGetSlotForColorWhenNoInput() {
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(4)
	vehicle1 := types.NewVehicle("1234","White")
	vehicle2 := types.NewVehicle("1264","Blue")
	parkinglot.GetSlot(1).Park(vehicle1)
	parkinglot.GetSlot(2).Park(vehicle2)
	c := &GetSlotForColor{}
	args := []string{}
	c.Execute(parkinglot, args)
	// Output:
	// Require 1 argument
}

func ExampleGetSlotForColorWhenPLNotInitialized(){
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(0)
	c := &GetSlotForColor{}
	args := []string{"Blue"}
	c.Execute(parkinglot, args)
	// Output:
	// Not found
}
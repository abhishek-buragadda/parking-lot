package commands

import "parkinglot/src/types"

func ExampleGetRegForColorIsPresent() {
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(4)
	vehicle1 := types.NewVehicle("1234","White")
	vehicle2 := types.NewVehicle("1264","Blue")
	parkinglot.GetSlot(1).Park(vehicle1)
	parkinglot.GetSlot(2).Park(vehicle2)
	c := &GetRegForColor{}
	args := []string{"Blue"}
	c.Execute(parkinglot, args)
	// Output:
	// 1264
}

func ExampleGetRegForColorIsNotPresent() {
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(4)
	vehicle1 := types.NewVehicle("1234","White")
	vehicle2 := types.NewVehicle("1264","Blue")
	parkinglot.GetSlot(1).Park(vehicle1)
	parkinglot.GetSlot(2).Park(vehicle2)
	c := &GetRegForColor{}
	args := []string{"Red"}
	c.Execute(parkinglot, args)
	// Output:
	// Not found
}

func ExampleGetRegForColorWhenNoInput() {
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(4)
	vehicle1 := types.NewVehicle("1234","White")
	vehicle2 := types.NewVehicle("1264","Blue")
	parkinglot.GetSlot(1).Park(vehicle1)
	parkinglot.GetSlot(2).Park(vehicle2)
	c := &GetRegForColor{}
	args := []string{}
	c.Execute(parkinglot, args)
	// Output:
	// Require 1 argument
}

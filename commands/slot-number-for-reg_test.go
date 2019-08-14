package commands

import "parkinglot/src/types"

func ExampleGetSlotForRegIsPresent() {
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(4)
	vehicle1 := types.NewVehicle("1234","White")
	vehicle2 := types.NewVehicle("1264","Blue")
	parkinglot.GetSlot(1).Park(vehicle1)
	parkinglot.GetSlot(2).Park(vehicle2)
	c := &GetSlotForReg{}
	args := []string{"1264"}
	c.Execute(parkinglot, args)
	// Output:
	// 2
}

func ExampleGetSlotForRegIsNotPresent() {
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(4)
	vehicle1 := types.NewVehicle("1234","White")
	vehicle2 := types.NewVehicle("1264","Blue")
	parkinglot.GetSlot(1).Park(vehicle1)
	parkinglot.GetSlot(2).Park(vehicle2)
	c := &GetSlotForReg{}
	args := []string{"4231"}
	c.Execute(parkinglot, args)
	// Output:
	// Not found
}

func ExampleGetSlotForRegWhenNoInput() {
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(4)
	vehicle1 := types.NewVehicle("1234","White")
	vehicle2 := types.NewVehicle("1264","Blue")
	parkinglot.GetSlot(1).Park(vehicle1)
	parkinglot.GetSlot(2).Park(vehicle2)
	c := &GetSlotForReg{}
	args := []string{}
	c.Execute(parkinglot, args)
	// Output:
	// Require 1 argument
}

func ExampleGetSlotForRegWhenPLNotInitialized(){
	parkinglot := types.GetParkingLotInstance()
	parkinglot.Initialize(0)
	c := &GetSlotForReg{}
	args := []string{"3245"}
	c.Execute(parkinglot, args)
	// Output:
	// Not found
}
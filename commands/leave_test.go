package commands

import (
	"github.com/stretchr/testify/assert"
	"parkinglot/src/types"
	"testing"
)


func TestLeaveWhenVehiclePresent(t *testing.T) {
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(4)
	vehicle := types.NewVehicle("1234","White")
	parkingLot.GetSlot(1).Park(vehicle)
	c := &Leave{}
	args := []string{"1"}
	c.Execute(parkingLot, args )
	assert.Equal(t, parkingLot.GetSlot(1).IsFreeSlot(), true )
}

func TestLeaveWhenVehicleNotPresent(t *testing.T) {
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(4)
	c := &Leave{}
	args := []string{"1"}
	c.Execute(parkingLot, args )
	assert.Equal(t, parkingLot.GetSlot(1).IsFreeSlot(), true )
}


func ExampleLeaveVehiclePresent() {
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(4)
	vehicle := types.NewVehicle("1234","White")
	parkingLot.GetSlot(1).Park(vehicle)
	c := &Leave{}
	args := []string{"1"}
	c.Execute(parkingLot, args )
	// Output:
	// Slot number 1 is free
}

func ExampleLeaveVehicleNotPresent() {
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(4)
	c := &Leave{}
	args := []string{"1"}
	c.Execute(parkingLot, args )
	// Output:
	// Slot number 1 already free
}

func ExampleLeaveWhenParkingNotInitialized() {
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(0)
	c := &Leave{}
	args := []string{"1"}
	c.Execute(parkingLot, args )
	// Output:
	// Slot number to Leave Not Found
}


package commands

import (
	"github.com/stretchr/testify/assert"
	"parkinglot/src/types"
	"testing"
)

func TestCreateParkingLotValidSize(t *testing.T) {
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(0)
	c := &CreateParkingLot{}
	args := []string{"2"}
	c.Execute(parkingLot, args)
	assert.Equal(t, parkingLot.GetSize() , uint64(2) )
}

func TestCreateParkingLotInValidSize(t *testing.T) {
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(0)
	c := &CreateParkingLot{}
	args := []string{"a"}
	c.Execute(parkingLot, args)
	assert.Equal(t, parkingLot.GetSize() , uint64(0) )
}


func ExampleCreateParkingLotInvalidSize(){
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(0)
	c := &CreateParkingLot{}
	args := []string{"a"}
	c.Execute(parkingLot, args )
	// Output:
	// Parking Capacity should be a number
}

func ExampleCreateParkingLotValidSize(){
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(0)
	c := &CreateParkingLot{}
	args := []string{"4"}
	c.Execute(parkingLot, args )
	// Output:
	// Created a parking lot with 4 slots
}

func ExampleCreateParkingLotNoArgs(){
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(0)
	c := &CreateParkingLot{}
	args := []string{}
	c.Execute(parkingLot, args )
	// Output:
	// 1 argument is required
}
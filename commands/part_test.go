package commands

import (
	"github.com/stretchr/testify/assert"
	"parkinglot/src/types"
	"testing"
)


func TestParkInAFreeSlot(t *testing.T){
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(4)
	c := &Park{}
	args := []string{"1234","white"}
	c.Execute(parkingLot, args)
	assert.Equal(t, parkingLot.GetSlot(1).GetVehicle().GetRegNumber(), "1234")
	assert.Equal(t, parkingLot.GetSlot(1).GetVehicle().GetColor(), "white")

}
func ExampleParkWithaFreeSlot() {
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(4)
	c := &Park{}
	args := []string{"1234","white"}
	c.Execute(parkingLot, args)
	// Output:
	// Allocated slot number: 1
}


func ExampleParkWhenSlotsAreFull(){
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(1)
	vehicle := types.NewVehicle("1234","White")
	parkingLot.GetSlot(1).Park(vehicle)
	c := &Park{}
	args := []string{"1234","white"}
	c.Execute(parkingLot, args)
	// Output:
	// Sorry, parking lot is full
}

func ExampleParkWhenParkingLotNotInitialized(){
	parkingLot := types.GetParkingLotInstance()
	parkingLot.Initialize(0)
	c := &Park{}
	args := []string{"1234","white"}
	c.Execute(parkingLot, args)
	// Output:
	// Sorry, parking lot is full
}

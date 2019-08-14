package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParkingLot_GetFreeSlot(t *testing.T) {
	parkingLot:= GetParkingLotInstance()
	parkingLot.Initialize(3)

	assert.Equal(t, parkingLot.GetFreeSlot(), parkingLot.GetSlot(1))
	vehicle1 := NewVehicle("1234","White")
	vehicle2 := NewVehicle("1264","Blue")
	parkingLot.GetSlot(1).Park(vehicle1)
	parkingLot.GetSlot(2).Park(vehicle2)
	parkingLot.GetSlot(1).UnPark()
	// when the vehicle is removed from the slot 1, the next free slot will be 1 not 3
	assert.Equal(t, parkingLot.GetSlot(1),  parkingLot.GetFreeSlot())
}


func TestParkingLot_GetSlotsByColor(t *testing.T) {
	parkingLot:= GetParkingLotInstance()
	parkingLot.Initialize(3)

	assert.Equal(t, []string(nil),  parkingLot.GetSlotsByColor("Blue"))
	vehicle1 := NewVehicle("1234","White")
	vehicle2 := NewVehicle("1264","Blue")
	parkingLot.GetSlot(1).Park(vehicle1)
	parkingLot.GetSlot(2).Park(vehicle2)
	assert.Equal(t, []string{"2"}, parkingLot.GetSlotsByColor("Blue"))
}


func TestParkingLot_GetRegNumberByColor(t *testing.T) {
	parkingLot:= GetParkingLotInstance()
	parkingLot.Initialize(3)

	assert.Equal(t, []string(nil),  parkingLot.GetRegNumberByColor("Blue"))
	vehicle1 := NewVehicle("1234","White")
	vehicle2 := NewVehicle("1264","Blue")
	parkingLot.GetSlot(1).Park(vehicle1)
	parkingLot.GetSlot(2).Park(vehicle2)
	assert.Equal(t, []string{"1264"}, parkingLot.GetRegNumberByColor("Blue"))
}


func TestParkingLot_GetSlotsByRegNumber(t *testing.T) {
	parkingLot:= GetParkingLotInstance()
	parkingLot.Initialize(3)

	assert.Equal(t, []string(nil),  parkingLot.GetSlotsByRegNumber("1234"))
	vehicle1 := NewVehicle("1234","White")
	vehicle2 := NewVehicle("1264","Blue")
	parkingLot.GetSlot(1).Park(vehicle1)
	parkingLot.GetSlot(2).Park(vehicle2)
	assert.Equal(t, []string{"2"}, parkingLot.GetSlotsByRegNumber("1264"))
}

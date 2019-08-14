package types

import (
	"strconv"
	"sync"
)

var parkingLot *ParkingLot
var once sync.Once
type ParkingLot struct {
	slots []Slot
	size uint64
}

func GetParkingLotInstance() *ParkingLot {
	once.Do(func() {
		parkingLot = &ParkingLot{nil, 0}
	})
	return parkingLot
}

func (p *ParkingLot) Initialize(n uint64){
 	p.slots = make([]Slot, n)
	p.size = n
	for i := uint64(0); i < n; i++ {
		p.slots[i] = Slot{
			id:      i+1,
			vehicle: nil,
		}
	}

}

func (p *ParkingLot) GetSlot(slotId uint64 )  *Slot {
	for i:=0; i < len(p.slots) ; i++ {
		if p.slots[i].GetSlotId() == slotId {
			return &p.slots[i]
		}
	}
	return nil
}

func (p *ParkingLot) GetFreeSlot() *Slot {
	for  i:=0; i< len(p.slots) ; i++  {
		if p.slots[i].IsFreeSlot(){
			return &p.slots[i]
		}
	}
	return nil
}

func(p *ParkingLot) GetSlotsByColor(color string ) []string {
	var results []string = nil
	for i:=0; i< len(p.slots) ; i++ {
		if !p.slots[i].IsFreeSlot() && p.slots[i].GetVehicle().GetColor() == color{
			results = append(results, strconv.FormatUint(p.slots[i].GetSlotId(), 10))
		}
	}
	return results
}

func (p *ParkingLot) GetSlotsByRegNumber( regNumber string ) []string {
	var results []string
	for i:=0; i < len(p.slots); i++ {
		if !p.slots[i].IsFreeSlot()  && p.slots[i].GetVehicle().GetRegNumber() == regNumber {
			results = append(results, strconv.FormatUint(p.slots[i].GetSlotId(),10))
		}
	}
	return results
}

func (p *ParkingLot) GetRegNumberByColor( color string ) []string{
	var results []string
	for i:=0; i < len(p.slots); i++ {
		if !p.slots[i].IsFreeSlot() && p.slots[i].GetVehicle().GetColor() == color {
			results = append(results, p.slots[i].GetVehicle().GetRegNumber())
		}
	}
	return results
}


func(p *ParkingLot ) GetSize() uint64 {
	return p.size
}

func (p *ParkingLot) GetSlots() []Slot {
	return p.slots
}
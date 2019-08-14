package types


type Slot struct{
	id uint64
	vehicle *Vehicle

}


func (s *Slot) IsFreeSlot() bool{
	return s.vehicle == nil
}

func (s *Slot) Park(vehicle *Vehicle){
	s.vehicle = vehicle
}

func (s *Slot) GetSlotId() uint64{
	return s.id
}

func (s *Slot) GetVehicle() *Vehicle{
	return s.vehicle
}
func (s *Slot) UnPark(){
	s.vehicle = nil
}
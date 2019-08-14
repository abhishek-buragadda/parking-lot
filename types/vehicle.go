package types


type Vehicle struct {
	regNumber string
	color string
}

func (v *Vehicle) GetRegNumber() string {
	return v.regNumber
}

func (v *Vehicle) GetColor() string {
	return  v.color
}

func  NewVehicle(regNumber string , color string ) *Vehicle{
	return &Vehicle{ regNumber: regNumber , color: color  }
}
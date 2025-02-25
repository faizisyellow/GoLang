package vehicle

import (
	"fmt"

	"example.com/car/suv"
)

type Vehicle struct {
	suv.Suv
	category string
}

func New(category string, newSuv suv.Suv) *Vehicle {
	return &Vehicle{
		category: category,
		Suv:      newSuv,
	}
}

func (v *Vehicle) DisplayVehicleWithSuv() {
	fmt.Printf("Vehicle with category %v has suv named: %v\n", v.category, v.Suv.Name)
}

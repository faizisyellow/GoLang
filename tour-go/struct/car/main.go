package main

import (
	"example.com/car/suv"
	"example.com/car/vehicle"
)

type cars interface {
	Display()
}

func displayCar(c cars) {
	c.Display()
}

func main() {
	newSuv := suv.New("rubicon", "jeep")
	newVehicle := vehicle.New("car", *newSuv)

	newVehicle.DisplayVehicleWithSuv()
}

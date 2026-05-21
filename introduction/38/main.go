package main

import "fmt"

type Car struct {
	Brand string
	Year  int
}

func NewCar(brand string, year int) *Car {
	car := Car{
		Brand: brand,
		Year:  year,
	}

	return &car
}

func main() {
	car := NewCar("car", 2026)

	fmt.Println("Car before change: ", car)

	car.Year = 2000

	fmt.Println("Car after change: ", car)
}

package main

import "fmt"

type Car struct {
	carType string
	fuelin  float64
}

func (c Car) distance() float64 {
	return c.fuelin / 1.5
}

func main() {
	c := Car{
		carType: "SUV",
		fuelin:  10.5,
	}
	fmt.Println("Car Type : ", c.carType)
	fmt.Println("est. Distance : ", c.distance())
}

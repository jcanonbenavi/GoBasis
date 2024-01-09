package internal

import "fmt"

type Vehicle interface {
	Start()
	DisplayInfo()
}

type Engine struct {
	Horsepower int
	Type       string
}

type Chasis struct {
	Material string
}

type Bodywork struct {
	Color string
}

type Car struct {
	Engine
	Chasis
	Bodywork
	NumberOfDoors int
}

type Motorcycle struct {
	Engine
	Chasis
	Bodywork
}

func (C Car) Start() {
	fmt.Println("The car is starting with a roar!")
}

func (c Car) DisplayInfo() {
	fmt.Printf("This car has %d doors, a %s body, a %s chasis, and a %d horsepower %s engine. \n",
		c.NumberOfDoors, c.Bodywork.Color, c.Chasis.Material, c.Engine.Horsepower, c.Engine.Type)
}

func (C Motorcycle) Start() {
	fmt.Println("The motorcycle is starting with a vroom!")
}

func (c Motorcycle) DisplayInfo() {
	fmt.Printf("This motorcycle has a %s body, a %s chasis, and a %d horsepower %s engine. \n",
		c.Bodywork.Color, c.Chasis.Material, c.Engine.Horsepower, c.Engine.Type)
}

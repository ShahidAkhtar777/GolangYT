package main

import "fmt"

type MotorVehicle interface {
	Milege() float64
}

type BMW struct {
	distance     float64
	fuel         float64
	averageSpeed string
}

type Audi struct {
	distance float64
	fuel     float64
}

func (b BMW) Milege() float64 {
	return b.distance / b.fuel
}

func (a Audi) Milege() float64 {
	return a.distance / a.fuel
}

func (b BMW) AverageSpeed() string {
	return b.averageSpeed
}

func typeAssertion(m MotorVehicle) {
	// Type assertion helps us extract the value of underlying type
	// fmt.Println(m.averageSpeed) --> It won't be able to access
	vehicle := m.(BMW)
	avSpeed := vehicle.averageSpeed

	fmt.Printf("Average speed of this BMW: %v\n", avSpeed)
}

func totalMilege(m []MotorVehicle) {
	totalMilege := 0.0
	for _, v := range m {
		totalMilege += v.Milege()
	}
	fmt.Printf("Total Milege of cars %f Km/L\n", totalMilege)
}

func main() {
	// Interface provides method signature for similar types of object
	// To implement an interface we need to implement all methods declared in that interface
	// Go interfaces are implemented implicitly. Unlike java we don't need to specify implements keyword.

	b1 := BMW{
		distance:     169.7,
		fuel:         45.6,
		averageSpeed: "80",
	}

	a1 := Audi{
		distance: 176.5,
		fuel:     56.8,
	}

	vehicles := []MotorVehicle{a1, b1}
	totalMilege(vehicles)

	// We can find dynamic value of interface with syntax i.(Type)
	// where i is variable of type interface and type is the type that implements interface
	typeAssertion(b1)
}

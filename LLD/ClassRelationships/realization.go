// Realization is an "implements" relationship where a class fulfills a contract defined by an interface.

/*
Realization v/s Inheritance
- Realization Models Capability
- Inheritance Models Identity
*/

//For ex : Flyable(an interface that define contract) connects unrelated things (Bird, Airplane, Drone) that share a capability.
//Animal(struct) connects related things (Dog, Cat, Bird) that share an identity.

package main

import (
	"fmt"
)

// Flyable interface
type Flyable interface {
	Fly()
	GetFlightInfo() string
}

// Bird struct
type Bird struct {
	species  string
	wingSpan float64
}

func NewBird(species string, wingSpan float64) *Bird {
	return &Bird{species: species, wingSpan: wingSpan}
}

func (b *Bird) Fly() {
	fmt.Println(b.species + " flaps its wings and takes off.")
}

func (b *Bird) GetFlightInfo() string {
	return fmt.Sprintf("%s (wingspan:) %f m, powered by muscle", b.species, b.wingSpan)
}

// Airplane struct
type Airplane struct {
	model       string
	maxAltitude int
}

func NewAirplane(model string, maxAltitude int) *Airplane {
	return &Airplane{model: model, maxAltitude: maxAltitude}
}

func (a *Airplane) Fly() {
	fmt.Println(a.model + " engines roar as it accelerates down the runway.")
}

func (a *Airplane) GetFlightInfo() string {
	return fmt.Sprintf("%s (max altitude: %d ft, powered by jet engines)", a.model, a.maxAltitude)
}

// Drone struct
type Drone struct {
	batteryLevel int
	maxRange     float64
}

func NewDrone(batteryLevel int, maxRange float64) *Drone {
	return &Drone{batteryLevel: batteryLevel, maxRange: maxRange}
}

func (d *Drone) Fly() {
	fmt.Printf("Drone propellers spin up. Battery at %d%%.\n", d.batteryLevel)
}

func (d *Drone) GetFlightInfo() string {
	return fmt.Sprintf("Drone (range: %f km, battery: %d %%)", d.maxRange, d.batteryLevel)
}

func testRealization() {
	flyingThings := []Flyable{
		NewBird("Eagle", 2.3),
		NewAirplane("Boeing 737", 41000),
		NewDrone(85, 10.0),
	}

	for _, flyer := range flyingThings {
		fmt.Println(flyer.GetFlightInfo())
		flyer.Fly()
		fmt.Println()
	}
}

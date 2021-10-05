package pattern7

import "fmt"

// Используется для:
// обеспечения возможности изменения логики или порядка работы алгоритма в процессе работы;
// кастомизации.

func Start() {
	v := NewVehicle()
	fg := NewFirstGear()
	sg := NewSecondGear()
	tg := NewThirdGear()
	fmt.Printf("Odometer %d km\n", v.getOdometer())
	v.SetGear(fg)
	v.moveOneHour()
	fmt.Printf("Odometer %d km\n", v.getOdometer())
	v.SetGear(sg)
	v.moveOneHour()
	fmt.Printf("Odometer %d km\n", v.getOdometer())
	v.SetGear(tg)
	v.moveOneHour()
	fmt.Printf("Odometer %d km\n", v.getOdometer())
	v.SetGear(fg)
	v.moveOneHour()
	fmt.Printf("Odometer %d km\n", v.getOdometer())
}

type Gear interface {
	moveOneHour(*Vehicle)
}

type FirstGear struct {
	maxSpeed int
}

func NewFirstGear() Gear {
	return &FirstGear{
		maxSpeed: 25,
	}
}

func (g *FirstGear) moveOneHour(v *Vehicle) {
	v.km += g.maxSpeed
}

type SecondGear struct {
	maxSpeed int
}

func NewSecondGear() Gear {
	return &SecondGear{
		maxSpeed: 50,
	}
}

func (g *SecondGear) moveOneHour(v *Vehicle) {
	v.km += g.maxSpeed
}

type ThirdGear struct {
	maxSpeed int
}

func NewThirdGear() Gear {
	return &ThirdGear{
		maxSpeed: 70,
	}
}

func (g *ThirdGear) moveOneHour(v *Vehicle) {
	v.km += g.maxSpeed
}

type Vehicle struct {
	km   int
	gear Gear
}

func NewVehicle() *Vehicle {
	return &Vehicle{
		km: 0,
	}
}

func (v *Vehicle) SetGear(g Gear) {
	v.gear = g
}

func (v *Vehicle) moveOneHour() {
	v.gear.moveOneHour(v)
}

func (v *Vehicle) getOdometer() int {
	return v.km
}

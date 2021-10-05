package pattern6

import "fmt"

// Используется для:
// отделения логики создания объектов от основной логики.

func Start() {
	kr := Factory("Rio")
	fmt.Println(kr.get())
	vp := Factory("Polo")
	fmt.Println(vp.get())
}

func Factory(m string) Vehicle {
	if m == "Rio" {
		return NewKiaRio()
	} else if m == "Polo" {
		return NewVWPolo()
	}
	return nil
}

type Vehicle interface {
	get() string
}

type Car struct {
	brand     string
	model     string
	bhp       int
	modelYear int
}

func (c *Car) get() string {
	return fmt.Sprintf("This is %s of %d model year, produced by %s with %d break horse power", c.model,
		c.modelYear, c.brand, c.bhp)
}

type KiaRio struct {
	Car
}

func NewKiaRio() Vehicle {
	return &KiaRio{
		Car{
			brand:     "Kia",
			model:     "Rio",
			modelYear: 2021,
			bhp:       123,
		},
	}
}

type VWPolo struct {
	Car
}

func NewVWPolo() Vehicle {
	return &VWPolo{
		Car{
			brand:     "Volkswagen",
			model:     "Polo",
			modelYear: 2021,
			bhp:       110,
		},
	}
}

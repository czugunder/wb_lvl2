package pattern2

import "fmt"

// Используется для:
// уменьшения размера конструктора создаваемого объекта;
// создания немного отличающихся в значениях, но одинаковых в конструкции объектов.

func Start() {
	carBuilder := NewVehicleBuilder("car", "gray", 220)
	plant := NewPlant(carBuilder)

	car1 := plant.buildVehicle()
	fmt.Println("Car #1 params:", car1.color, car1.speed)

	carBuilder.setSpeed(180)
	car2 := plant.buildVehicle()
	fmt.Println("Car #2 params:", car2.color, car2.speed)

	planeBuilder := NewVehicleBuilder("plane", "white", 810)
	plant.changeBuilder(planeBuilder)

	plane1 := plant.buildVehicle()
	fmt.Println("Plane #1 params:", plane1.color, plane1.speed)

	planeBuilder.setColor("green")
	plane2 := plant.buildVehicle()
	fmt.Println("Plane #2 params:", plane2.color, plane2.speed)

	plant.changeBuilder(carBuilder)
	carBuilder.setColor("white")
	car3 := plant.buildVehicle()
	fmt.Println("Car #3 params:", car3.color, car3.speed)
}

type Plant struct {
	builder VehicleBuilder
}

func NewPlant(builder VehicleBuilder) *Plant {
	return &Plant{
		builder: builder,
	}
}

func (p *Plant) changeBuilder(builder VehicleBuilder) {
	p.builder = builder
}

func (p *Plant) buildVehicle() Vehicle {
	return p.builder.produceVehicle()
}

type Vehicle struct {
	color string
	speed int
}

type VehicleBuilder interface {
	setColor(string)
	setSpeed(int)
	produceVehicle() Vehicle
}

func NewVehicleBuilder(vType, color string, speed int) VehicleBuilder {
	if vType == "car" {
		return newCarBuilder(color, speed)
	} else if vType == "plane" {
		return newPlaneBuilder(color, speed)
	}
	return nil
}

func newCarBuilder(color string, speed int) *car {
	return &car{
		color: color,
		speed: speed,
	}
}

type car struct {
	color string
	speed int
}

func (c *car) setColor(color string) {
	c.color = color
}

func (c *car) setSpeed(speed int) {
	c.speed = speed
}

func (c *car) produceVehicle() Vehicle {
	return Vehicle{
		color: c.color,
		speed: c.speed,
	}
}

func newPlaneBuilder(color string, speed int) *plane {
	return &plane{
		color: color,
		speed: speed,
	}
}

type plane struct {
	color string
	speed int
}

func (p *plane) setColor(color string) {
	p.color = color
}

func (p *plane) setSpeed(speed int) {
	p.speed = speed
}

func (p *plane) produceVehicle() Vehicle {
	return Vehicle{
		color: p.color,
		speed: p.speed,
	}
}

package pattern3

import "fmt"

// Используется для:
// увеличения или изменения функциональности структур без изменения самих структур.

func Start() {
	c := NewCar("blue", 205, 150)
	p := NewPlane("red", 840, 12000)

	sr := NewSpeedRadar()
	c.accept(sr)
	fmt.Printf("SpeedRadar got  that speed of the %s car is %d\n", c.getColor(), sr.speed)
	p.accept(sr)
	fmt.Printf("SpeedRadar got that speed of the %s plane is %d\n", p.getColor(), sr.speed)

	ud := NewUniqeData()
	c.accept(ud)
	fmt.Printf("UniqeData got that brake horse power of the %s car is %d\n", c.getColor(), ud.intData)
	p.accept(ud)
	fmt.Printf("UniqeData got that altitude of the %s plane is %d\n", p.getColor(), ud.intData)
}

type Vehicle interface {
	getColor() string
	accept(Visitor)
}

type car struct {
	color string
	speed int
	bhp   int
}

func NewCar(color string, speed, bhp int) *car {
	return &car{
		color: color,
		speed: speed,
		bhp:   bhp,
	}
}

func (c *car) getColor() string {
	return c.color
}

func (c *car) accept(v Visitor) {
	v.visitCar(c)
}

type plane struct {
	color string
	speed int
	alt   int
}

func NewPlane(color string, speed, alt int) *plane {
	return &plane{
		color: color,
		speed: speed,
		alt:   alt,
	}
}

func (p *plane) getColor() string {
	return p.color
}

func (p *plane) accept(v Visitor) {
	v.visitPlane(p)
}

type Visitor interface {
	visitCar(*car)
	visitPlane(*plane)
}

type SpeedRadar struct {
	speed int
}

func NewSpeedRadar() *SpeedRadar {
	return &SpeedRadar{}
}

func (s *SpeedRadar) visitCar(c *car) {
	s.speed = c.speed
}

func (s *SpeedRadar) visitPlane(p *plane) {
	s.speed = p.speed
}

type UniqueData struct {
	intData int
}

func NewUniqeData() *UniqueData {
	return &UniqueData{}
}

func (u *UniqueData) visitCar(c *car) {
	u.intData = c.bhp
}

func (u *UniqueData) visitPlane(p *plane) {
	u.intData = p.alt
}

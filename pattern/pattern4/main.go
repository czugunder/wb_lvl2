package pattern4

import "fmt"

// Используется для:
// унификации обработчика у инвокеров;
// для отложенного выполнения запроса так как сама команда содержит субъект.

func Start() { // client
	v := NewVehicle()                    // receiver
	gasCommand := NewGasCommand(v)       // command1
	brakeCommand := NewBrakeCommand(v)   // command2
	gasPedal := NewPedal(gasCommand)     // invoker 1
	brakePedal := NewPedal(brakeCommand) // invoker 2

	fmt.Println("Is car accelerating:", v.acceleration)
	gasPedal.push()
	fmt.Println("Is car accelerating:", v.acceleration)
	brakePedal.push()
	fmt.Println("Is car accelerating:", v.acceleration)

}

type Vehicle struct {
	acceleration bool
}

func NewVehicle() *Vehicle {
	return &Vehicle{}
}

func (v *Vehicle) gas() {
	v.acceleration = true
}

func (v *Vehicle) brake() {
	v.acceleration = false
}

type command interface {
	execute()
}

type pedal struct {
	com command
}

func NewPedal(com command) *pedal {
	return &pedal{
		com: com,
	}
}

func (k *pedal) push() {
	k.com.execute()
}

type gasCom struct {
	car *Vehicle
}

func NewGasCommand(v *Vehicle) *gasCom {
	return &gasCom{
		car: v,
	}
}

func (ic *gasCom) execute() {
	ic.car.gas()
}

type brakeCom struct {
	car *Vehicle
}

func NewBrakeCommand(v *Vehicle) *brakeCom {
	return &brakeCom{
		car: v,
	}
}

func (sc *brakeCom) execute() {
	sc.car.brake()
}

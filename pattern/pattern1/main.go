package pattern1

import "fmt"

// Используется для:
// представления конструктора сервиса в лаконичном виде, скрывая за фасадом сложность инициализации отдельных служб.

func Start() {
	f := NewFacade("Roman")
	fmt.Println(f.GetVals())
	f.InvertState()
	fmt.Println(f.GetVals())
	f.SetNewName("Nikitin")
	fmt.Println(f.GetVals())
}

type facade struct {
	s1 *service1
	s2 *service2
}

func NewFacade(name string) *facade {
	return &facade{
		s1: newService1(),
		s2: newService2(name),
	}
}

func (f *facade) SetNewName(name string) {
	f.s2.setName(name)
}

func (f *facade) InvertState() {
	f.s1.invert()
}

func (f *facade) GetVals() (bool, string) {
	return f.s1.state, f.s2.name
}

type service1 struct {
	state bool
}

func newService1() *service1 {
	return &service1{
		state: true,
	}
}

func (s *service1) invert() {
	s.state = !s.state
}

type service2 struct {
	name string
}

func newService2(name string) *service2 {
	return &service2{
		name: name,
	}
}

func (s *service2) setName(name string) {
	s.name = name
}

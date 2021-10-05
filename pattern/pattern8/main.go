package pattern8

import (
	"errors"
	"fmt"
)

// Используется для:
// смены состояний объекта после, например, запроса, то есть применим к логике, где у объекта есть конкретные состояния;
// смены состояний основываясь на текущем состоянии, а не на условиях.

func Start() {
	grechka := NewGrechka(500)
	if err := grechka.currentState.plant(); err != nil {
		fmt.Println(err)
	}
	if err := grechka.currentState.harvest(); err != nil {
		fmt.Println(err)
	}
	if err := grechka.currentState.boil(); err != nil {
		fmt.Println(err)
	}
	if err := grechka.currentState.eat(); err != nil {
		fmt.Println(err)
	}
	grechka.setGrams(250)
	if err := grechka.currentState.eat(); err != nil {
		fmt.Println(err)
	}
}

type Grechka struct {
	seedState   state
	groundState state
	foodState   state

	currentState state

	gramms int
}

func (g *Grechka) setGrams(gr int) {
	g.gramms = gr
}

func NewGrechka(gr int) *Grechka {
	g := Grechka{}
	s1 := seedState{
		grechka: &g,
	}
	s2 := groundState{
		grechka: &g,
	}
	s3 := foodState{
		grechka: &g,
	}
	g.seedState = &s1
	g.groundState = &s2
	g.foodState = &s3
	g.currentState = g.seedState
	g.setGrams(gr)
	return &g
}

type state interface {
	plant() error   // семена
	harvest() error // в земле
	boil() error    // семена
	eat() error     // еда
}

type seedState struct {
	grechka *Grechka
}

func (s *seedState) plant() error {
	s.grechka.currentState = s.grechka.groundState
	fmt.Printf("%d gramms of buckwheat were planted to the soil!\n", s.grechka.gramms)
	return nil
}

func (s *seedState) harvest() error {
	return errors.New("you have seeds already")
}

func (s *seedState) boil() error {
	s.grechka.gramms *= 2
	s.grechka.currentState = s.grechka.foodState
	fmt.Printf("%d gramms of buckwheat are ready to be eaten!\n", s.grechka.gramms)
	return nil
}

func (s *seedState) eat() error {
	return errors.New("seeds are bad as food")
}

type groundState struct {
	grechka *Grechka
}

func (s *groundState) plant() error {
	return errors.New("seeds are in soil already")
}

func (s *groundState) harvest() error {
	s.grechka.gramms *= 2
	s.grechka.currentState = s.grechka.seedState
	fmt.Printf("%d gramms of buckwheat were harvested form the soil!\n", s.grechka.gramms)
	return nil
}

func (s *groundState) boil() error {
	return errors.New("seeds are in soil, you cant boil them")
}

func (s *groundState) eat() error {
	return errors.New("seeds are in soil, you cant eat them")
}

type foodState struct {
	grechka *Grechka
}

func (s *foodState) plant() error {
	return errors.New("you cant plant boiled seeds")
}

func (s *foodState) harvest() error {
	return errors.New("you cant harvest seeds, they aren't in the ground")
}

func (s *foodState) boil() error {
	return errors.New("seeds are boiled already")
}

func (s *foodState) eat() error {
	s.grechka.gramms = 0
	s.grechka.currentState = s.grechka.seedState
	fmt.Printf("%d gramms of buckwheat left, you've eaten it all!\n", s.grechka.gramms)
	return nil
}

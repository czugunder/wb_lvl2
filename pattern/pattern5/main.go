package pattern5

import (
	"fmt"
)

// Используется для:
// последовательной обработки запроса.

func Start() {
	m1 := NewMug(false, false, false)
	MakeTeaWithLemon(m1)
	fmt.Println()
	m2 := NewMug(true, false, false)
	MakeTeaWithLemon(m2)
	fmt.Println()
	m3 := NewMug(false, true, false)
	MakeTeaWithLemon(m3)
	fmt.Println()
	m4 := NewMug(false, false, true)
	MakeTeaWithLemon(m4)
	fmt.Println()
	m5 := NewMug(true, true, false)
	MakeTeaWithLemon(m5)
	fmt.Println()
	m6 := NewMug(false, true, true)
	MakeTeaWithLemon(m6)
	fmt.Println()
	m7 := NewMug(true, false, true)
	MakeTeaWithLemon(m7)
	fmt.Println()
	m8 := NewMug(true, true, true)
	MakeTeaWithLemon(m8)
	fmt.Println()
}

type Mug struct {
	isWaterIn bool
	isBrewIn  bool
	isLemonIn bool
}

func NewMug(w, b, l bool) *Mug {
	return &Mug{
		isLemonIn: l,
		isWaterIn: w,
		isBrewIn:  b,
	}
}

func MakeTeaWithLemon(mug *Mug) {
	cb := NewCuttingBoard()
	b := NewBrew(cb)
	p := NewPot(b)
	p.do(mug)
}

type action interface {
	do(*Mug)
}

type Pot struct {
	next action
}

func NewPot(nextAction action) *Pot {
	return &Pot{
		next: nextAction,
	}
}

func (p *Pot) do(mug *Mug) {
	if mug.isWaterIn {
		fmt.Println("Boiled water is already in the mug!")
	} else {
		mug.isWaterIn = true
		fmt.Println("Boiled water was poured into the mug.")
	}
	p.next.do(mug)
}

type Brew struct {
	next action
}

func NewBrew(nextAction action) *Brew {
	return &Brew{
		next: nextAction,
	}
}

func (b *Brew) do(mug *Mug) {
	if mug.isBrewIn {
		fmt.Println("Brew is already in the mug!")
	} else {
		mug.isBrewIn = true
		fmt.Println("Brew was added to the mug.")
	}
	b.next.do(mug)
}

type CuttingBoard struct {
	next action
}

func NewCuttingBoard() *CuttingBoard {
	return &CuttingBoard{}
}

func (c *CuttingBoard) do(mug *Mug) {
	if mug.isLemonIn {
		fmt.Println("Lemon is already in the mug! Tea is ready, enjoy!")
	} else {
		mug.isLemonIn = true
		fmt.Println("Lemon was dropped into the mug. Tea is ready, enjoy!")
	}
}

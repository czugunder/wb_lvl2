package dev07

import (
	"fmt"
	"reflect"
	"time"
)

func Start() {
	mymain()
}

func mymain() {
	//var or func(channels ... <-chan interface{}) <-chan interface {}
	var or = ORChan2

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v\n", time.Since(start))

}

// ORChan1 - по одной горутине на каждый входящий канал, если канал закрывается, закрывается и общий
func ORChan1(channels ...<-chan interface{}) <-chan interface{} {
	comChan := make(chan interface{})
	for _, t := range channels {
		go func(c <-chan interface{}) {
			select {
			case _, opened := <-c:
				if !opened {
					close(comChan)
				}
			}
		}(t)
	}
	return comChan
}

// ORChan2 - без горутин, но на reflect, есть reflect.Select который проводит select над массивом типа SelectCase,
// в который загоняются входящие каналы, если один из каналов закроется, закроется и основной канал
func ORChan2(channels ...<-chan interface{}) <-chan interface{} {
	var sel []reflect.SelectCase
	comChan := make(chan interface{})
	for _, c := range channels {
		sel = append(sel,
			reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
	}
	_, _, opened := reflect.Select(sel)
	if !opened {
		close(comChan)
	}
	return comChan
}

// Что лучше не тестировал, но по сути, что создавать по одной дополнительной горутине на канал, что
// использовать рефлект - все замедляет.

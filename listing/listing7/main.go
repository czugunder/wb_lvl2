package listing7

import (
	"fmt"
	"math/rand"
	"time"
)

func Start() {
	mymain()
}

func asChan(vs ...int) <-chan int { // только извлечение из канала
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v                                                        // пишет в канал
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // задержка
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int { // только извлечение из канала
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a: // v, ok := <- a вместе с обработчиком ok решило бы вопрос c нулями
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func mymain() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

// 1, 2, 3, 4, 5, 6, 7, 8, 0............

// Числа возвращаются последовательно благодаря задержке.
// Нули возвращаются потому что при прослушивании закрытого канала будут возвращаться нулевые значения
// типа данного канала.
// Каналы a и b закрываются в анонимной горутине фунции merge, а select в свою очередь не проверят во время чтения,
// открыт ли канал.

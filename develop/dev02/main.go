package dev02

import (
	"errors"
	"fmt"
)

// Start функция для запуска программы
func Start() {
	s := `\\`
	fmt.Println(Unpack(s))
}

// Unpack главная функция, выполняющая распаковку
func Unpack(s string) (string, error) {
	var letterBuffer rune
	var gotLetter bool
	var isScreen bool
	var returnBuffer []rune
	var err error
	for _, r := range s {
		if d, isD := isDigit(r); isD { // цифра
			if !isScreen { // если она не экранирована
				if gotLetter { // если перед цифрой стоял символ
					fillBuffer(&returnBuffer, letterBuffer, d) // запись в буфер этого символа сколько нужно раз
					gotLetter = false                          // символ отработали, флаг нужно сбросить
				} else {
					err = errors.New("incorrect input string") // ошибка - два числа подряд
				}
			} else { // если цифра экранирована
				letterBuffer = r // заносится в буфер как символ
				isScreen = false // сброс флага экрана
				gotLetter = true // флаг нахождения символа поднять
			}
		} else if r == 92 { // экран
			if !isScreen { // если флаг экрана опущен
				isScreen = true // поднять этот флаг
				if gotLetter {
					fillBuffer(&returnBuffer, letterBuffer, 1)
					gotLetter = false
				}
			} else { // если флаг экрана поднят
				letterBuffer = r // записать слэш как символ
				isScreen = false // флаг экрана снять
				gotLetter = true // флаг нахождения символа поднять
			}
		} else { // все остальное
			if !gotLetter { // если в буфере нет пойманного символа
				letterBuffer = r // символ в него записывается, ожидается либо число, либо след символ
				gotLetter = true // флаг нахождения символа поднять
			} else { // если в буфере уже есть символ, то
				fillBuffer(&returnBuffer, letterBuffer, 1) // тот что в буфере записывается в одном экземпляре
				letterBuffer = r                           // в буфер записывается текущий
			}
		}
	}
	if gotLetter {
		fillBuffer(&returnBuffer, letterBuffer, 1) // проверка последнего одиночного символа
	}
	return string(returnBuffer), err
}

func fillBuffer(s *[]rune, r rune, t int) {
	for i := 0; i < t; i++ {
		*s = append(*s, r)
	}
}

func isDigit(r rune) (int, bool) {
	if r > 47 && r < 58 { // тогда число (ASCII приколы из ассемблера)
		return getDigit(r), true
	}
	return 0, false
}

func getDigit(r rune) int {
	return int(r) - 48
}

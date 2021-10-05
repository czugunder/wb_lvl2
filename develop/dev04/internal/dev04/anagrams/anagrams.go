package anagrams

import (
	"fmt"
	"strings"
	"wb_lvl2/develop/dev04/internal/dev04/anagramHelper"
)

// WordToMap преобразует слово в характеризующую хеш-таблицу
func WordToMap(word string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range word {
		if _, inMap := m[v]; inMap {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	return m
}

// Anagrams центровая функция, находит анаграммы
func Anagrams(words *[]string) map[string]*[]string {
	ah := anagramhelper.NewAnagramHelper()
	for _, word := range *words {
		curWord := stringToLower(word)
		curWordMap := WordToMap(curWord)
		if ad := ah.FindByHeaderMap(curWordMap); ad != nil {
			ad.Add(curWord)
		} else {
			ah.Add(curWord, curWordMap)
		}
	}
	ah.DeleteUniqueAndSort()
	return ah.FormMap()
}

func stringToLower(s string) string {
	return strings.ToLower(s)
}

// PrintAnagrams выводит анаграммы в stdout
func PrintAnagrams(m map[string]*[]string) {
	for k, v := range m {
		fmt.Printf("%s: %v\n", k, *v)
	}
}

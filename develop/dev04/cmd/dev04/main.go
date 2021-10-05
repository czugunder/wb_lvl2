package main

import (
	"wb_lvl2/develop/dev04/internal/dev04/anagrams"
)

func main() {
	var words = []string{"пятак", "листок", "природа", "пятка", "столик",
		"тяпка", "слиток"}
	anagrams.PrintAnagrams(anagrams.Anagrams(&words))
}

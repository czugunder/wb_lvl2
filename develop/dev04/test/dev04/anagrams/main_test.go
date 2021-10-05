package anagrams_test

import (
	"testing"
	"wb_lvl2/develop/dev04/internal/dev04/anagramHelper"
	"wb_lvl2/develop/dev04/internal/dev04/anagrams"
)

func TestCompareWordMaps(t *testing.T) {
	m1 := map[rune]int{'п': 1, 'р': 2, 'и': 1, 'о': 1, 'д': 1, 'а': 1}
	m2 := map[rune]int{'о': 1, 'д': 1, 'а': 1, 'п': 1, 'р': 2, 'и': 1}   // такая же, как m1
	m3 := map[rune]int{'п': 999, 'р': 2, 'и': 1, 'о': 1, 'д': 1, 'а': 1} // руны те же, кол-во другое
	m4 := map[rune]int{'в': 1, 'р': 2, 'и': 1, 'о': 1, 'д': 1, 'а': 1}   // другой набор рун, длина та же
	m5 := map[rune]int{'п': 1, 'р': 2, 'и': 1}                           // не та длина

	r1 := anagramHelper.CompareWordMaps(m1, m2) // true
	r2 := anagramHelper.CompareWordMaps(m1, m3) // false
	r3 := anagramHelper.CompareWordMaps(m1, m4) // false
	r4 := anagramHelper.CompareWordMaps(m1, m5) // false

	if r1 != true {
		t.Fatalf("Comparing of equal maps (m1:%v, m2:%v) went wrong, should be true, got false\n",
			m1, m2)
	}
	if r2 != false {
		t.Fatalf("Comparing of maps (m1:%v, m3:%v) with equal runes "+
			"but different counters went wrong, should be false, got true\n",
			m1, m3)
	}
	if r3 != false {
		t.Fatalf("Comparing of maps (m1:%v, m4:%v) with different runes went wrong, should be false, got true\n",
			m1, m4)
	}
	if r4 != false {
		t.Fatalf("Comparing of maps (m1:%v, m5:%v) with different length went wrong, should be false, got true\n",
			m1, m5)
	}
}

func TestWordToMap(t *testing.T) {
	word := "природа"
	wordMap := map[rune]int{'п': 1, 'р': 2, 'и': 1, 'о': 1, 'д': 1, 'а': 1}

	testMap := anagrams.WordToMap(word)

	areSame := anagramHelper.CompareWordMaps(wordMap, testMap)
	if !areSame {
		t.Fatalf("Incorrent word to map convertion on word %s, expected: %v, got: %v\n", word, wordMap, testMap)
	}
}

func TestAnagrams(t *testing.T) {
	var words = []string{"пЯтак", "Листок", "приРода", "пяТка", "СТОЛИК", "тяпка", "слиток"}
	var anargamGroup1 = []string{"пятка", "тяпка"}
	var anargamGroup2 = []string{"слиток", "столик"}
	anagramMap := make(map[string]*[]string)
	anagramMap["пятак"] = &anargamGroup1
	anagramMap["листок"] = &anargamGroup2

	testMap := anagrams.Anagrams(&words)

	if len(testMap) != len(anagramMap) {
		t.Fatalf("Returned map has unexpected length, expected: %d, got:%d\n", len(anagramMap), len(testMap))
	}
	for k, v := range anagramMap {
		if tv, found := testMap[k]; !found {
			t.Fatalf("Returned map hasn't got expected key %s\n", k)
		} else {
			if len(*v) != len(*tv) {
				t.Fatalf("Length of value array %v is not equal to length of expected array %v, expected: %d,"+
					" got: %d\n", *tv, *v, len(*v), len(*tv))
			} else {
				for i := range *v {
					if (*v)[i] != (*tv)[i] {
						t.Fatalf("Wrong order or content of value array, expected: %v, got: %v\n", *v, *tv)
					}
				}
			}
		}
	}
}

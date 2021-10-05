package anagramHelper

import (
	"sort"
)

type anagramData struct {
	header    string
	headerMap map[rune]int
	tail      []string
}

func (ad *anagramData) Add(word string) {
	ad.tail = append(ad.tail, word)
}

func NewAnagramData(header string, headerMap map[rune]int) *anagramData {
	return &anagramData{
		header:    header,
		headerMap: headerMap,
	}
}

type anagramHelper struct {
	data []*anagramData
}

func NewAnagramHelper() *anagramHelper {
	return &anagramHelper{}
}

func (ah *anagramHelper) Add(header string, headerMap map[rune]int) {
	ah.data = append(ah.data, NewAnagramData(header, headerMap))
}

func (ah *anagramHelper) FindByHeaderMap(m map[rune]int) *anagramData {
	for _, v := range ah.data {
		if CompareWordMaps(v.headerMap, m) {
			return v
		}
	}
	return nil
}

func (ah *anagramHelper) DeleteUniqueAndSort() {
	var newData []*anagramData
	for _, v := range ah.data {
		if len(v.tail) != 0 {
			SortArray(&v.tail)
			newData = append(newData, v)
		}
	}
	ah.data = newData
}

func (ah *anagramHelper) FormMap() map[string]*[]string {
	m := make(map[string]*[]string)
	for _, v := range ah.data {
		m[v.header] = &v.tail
	}
	return m
}

func SortArray(arr *[]string) {
	sort.Strings(*arr)
}

func CompareWordMaps(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, inMap2 := m2[k1]; inMap2 {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

package anagramhelper

import (
	"sort"
)

// AnagramData вспомогательный тип для организации групп анаграмм, где header первое найденное вхождение
// headerMap - хеш-таблица характеризующая данную группу, tail - слайс с остальными вхождениями
type AnagramData struct {
	header    string
	headerMap map[rune]int
	tail      []string
}

// Add добавляет новое вхождение
func (ad *AnagramData) Add(word string) {
	ad.tail = append(ad.tail, word)
}

// NewAnagramData создает экземпляр AnagramData
func NewAnagramData(header string, headerMap map[rune]int) *AnagramData {
	return &AnagramData{
		header:    header,
		headerMap: headerMap,
	}
}

// AnagramHelper вспомогательный тип, хранит слайс AnagramData
type AnagramHelper struct {
	data []*AnagramData
}

// NewAnagramHelper создает экземпляр AnagramHelper
func NewAnagramHelper() *AnagramHelper {
	return &AnagramHelper{}
}

// Add добавляет новую группу анаграмм
func (ah *AnagramHelper) Add(header string, headerMap map[rune]int) {
	ah.data = append(ah.data, NewAnagramData(header, headerMap))
}

// FindByHeaderMap ищет группу анаграмм по характеризующий ее хеш-таблице
func (ah *AnagramHelper) FindByHeaderMap(m map[rune]int) *AnagramData {
	for _, v := range ah.data {
		if CompareWordMaps(v.headerMap, m) {
			return v
		}
	}
	return nil
}

// DeleteUniqueAndSort удаляет множества из одного элемента и сортирует выдачу
func (ah *AnagramHelper) DeleteUniqueAndSort() {
	var newData []*AnagramData
	for _, v := range ah.data {
		if len(v.tail) != 0 {
			SortArray(&v.tail)
			newData = append(newData, v)
		}
	}
	ah.data = newData
}

// FormMap формирует хеш-таблицу, где ключ первое вхождение анаграммы, а значение - слайс из остальных вхождений
func (ah *AnagramHelper) FormMap() map[string]*[]string {
	m := make(map[string]*[]string)
	for _, v := range ah.data {
		m[v.header] = &v.tail
	}
	return m
}

// SortArray сортирует слайс строк
func SortArray(arr *[]string) {
	sort.Strings(*arr)
}

// CompareWordMaps сравнивает две хеш-таблицы, нужно для поиска анаграмм по характеризующим хем-таблицам
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

package sort

import (
	"log"
	"math"
	"strconv"
	"strings"
)

// Strings - тип соответствующий интерфейсу Sorter
type Strings struct {
	st         []string
	customLess func(int, int) bool
	flags      *Flags
}

// NewSortString создает экземпляр Strings
func NewSortString(f *Flags) *Strings {
	s := Strings{}
	s.flags = f
	s.setCustomLess()
	return &s
}

func (s *Strings) Len() int {
	return len(s.st)
}

func (s *Strings) Swap(i, j int) {
	s.st[i], s.st[j] = s.st[j], s.st[i]
}

func (s *Strings) Less(i, j int) bool {
	return s.customLess(i, j)
}

func (s *Strings) getCustomLess() func(int, int) bool {
	return s.customLess
}

func (s *Strings) setCustomLess() {
	if s.flags.M {
		s.customLess = s.monthLess
	} else if s.flags.H {
		s.customLess = s.suffixLess
	} else {
		if s.flags.K == "" {
			s.customLess = s.classicLess
		} else {
			s.customLess = s.kFlagLess
		}
	}
}

// GetStrings возвращает внутренний буфер (слайс) сортируемых строк
func (s *Strings) GetStrings() []string {
	return s.st
}

// SetStrings заполняет внутренний буфер (слайс) сортируемых строк
func (s *Strings) SetStrings(str []string) {
	s.st = str
}

func (s *Strings) append(str string) {
	s.st = append(s.st, str)
}

// RemoveDuplicates удаляет повторяющиеся строки во внутреннем буфере сортируемых строк
func (s *Strings) RemoveDuplicates() {
	keys := make(map[string]bool)
	var newStrings []string
	for _, entry := range s.st {
		if _, value := keys[entry]; !value { // детектирование наличия в мапе по значению, можно и через подавленное
			keys[entry] = true
			newStrings = append(newStrings, entry)
		}
	}
	s.st = newStrings
}

func (s *Strings) classicLess(i, j int) bool {
	eq, less := s.isIntLess(len(s.st[i]), len(s.st[j]))
	if !eq && !s.flags.B {
		if !less {
			return false
		}
		return true
	}
	strI, strJ := s.st[i], s.st[j]
	bufI, bufJ := []rune(strI), []rune(strJ)
	for cnt := 0; cnt < len(s.st[i]); cnt++ {
		eq, less = s.isRuneLess(bufI[cnt], bufJ[cnt])
		if !eq {
			if less {
				return true
			}
		}
	}
	return false
}

// CheckSuffix находит суффикс из входной хэш-таблицы sufs во входной строке str и возвращает его
func CheckSuffix(sufs map[string]int, str string) (suf string) {
	for k := range sufs {
		if strings.HasSuffix(str, k) {
			suf = k
			return
		}
	}
	return
}

func (s *Strings) suffixLess(i, j int) bool {
	suffixes := map[string]int{"n": -9, "mi": -6, "m": -3, "K": 3, "M": 6, "G": 9}
	var str1, str2 string
	var suf1, suf2 string
	if s.flags.K != "" {
		str1, str2 = ProcessKFlag(s.st[i], s.flags), ProcessKFlag(s.st[j], s.flags)

	} else {
		str1, str2 = s.st[i], s.st[j]
	}
	suf1 = CheckSuffix(suffixes, str1)
	suf2 = CheckSuffix(suffixes, str2)
	str1 = strings.TrimSuffix(str1, suf1)
	str2 = strings.TrimSuffix(str2, suf2)
	val1, err1 := strconv.ParseFloat(str1, 32)
	if err1 != nil {
		log.Fatalln("incorrect number", err1)
	}
	val2, err2 := strconv.ParseFloat(str2, 32)
	if err2 != nil {
		log.Fatalln("incorrect number", err2)
	}
	if suf1 != "" {
		val1 *= math.Pow10(suffixes[suf1])
	}
	if suf2 != "" {
		val2 *= math.Pow10(suffixes[suf2])
	}
	_, less := s.isFloatLess(val1, val2)
	return less
}

func (s *Strings) monthLess(i, j int) bool {
	months := []string{"JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"}
	str1, str2 := "", ""
	if s.flags.K != "" {
		str1, str2 = ProcessKFlag(s.st[i], s.flags), ProcessKFlag(s.st[j], s.flags)
		if len(str1) != 3 || len(str2) != 3 {
			log.Fatalln("incorrect -k flag usage with flag -M")
		}
	} else {
		str1, str2 = s.st[i], s.st[j]
		if len(str1) != 3 || len(str2) != 3 {
			log.Fatalln("incorrect file input with flag -M")
		}
	}
	var iFound, jFound = -1, -1
	for mi, m := range months {
		if m == str1 {
			iFound = mi
		}
		if m == str2 {
			jFound = mi
		}
	}
	if iFound < 0 || jFound < 0 {
		log.Fatalln("incorrect file input with flag -M")
	}
	if iFound < jFound {
		return true
	}
	return false
}

func (s *Strings) kFlagLess(i, j int) bool {
	strI, strJ := ProcessKFlag(s.st[i], s.flags), ProcessKFlag(s.st[j], s.flags)
	eq, less := s.isIntLess(len(strI), len(strJ))
	if !eq {
		if !less {
			return false
		}
		return true
	}
	bufI, bufJ := []rune(strI), []rune(strJ)
	for cnt := 0; cnt < len(bufI); cnt++ {
		eq, less = s.isRuneLess(bufI[cnt], bufJ[cnt])
		if !eq {
			if less {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

// ProcessKFlag на основе дешифрованного флага -k возвращает ключ для сортировки из строки
func ProcessKFlag(s string, f *Flags) string {
	var buf string
	sl := strings.Split(s, " ")
	if len(sl) < f.KFlag.F2 {
		log.Fatalln("file doesn't suit flag -k (column quantity mismatch)")
	} else {
		var internalC1, internalF2, internalC2 = f.KFlag.C1, f.KFlag.F2, f.KFlag.C2
		if internalF2 == 0 {
			internalF2 = len(sl)
		}
		if internalC1 == 0 {
			internalC1 = 1
		}
		if internalC2 == 0 {
			internalC2 = len(sl[internalF2-1])
		}
		if len(sl[f.KFlag.F1-1]) < internalC1 || len(sl[internalF2-1]) < internalC2 {
			log.Fatalln("file doesn't suit flag -k (position in column quantity mismatch)")
		} else {
			if f.KFlag.F1 == internalF2 {
				buf = sl[internalF2-1][internalC1-1 : internalC2]
			} else {
				buf += sl[f.KFlag.F1-1][internalC1-1:]
				for i := f.KFlag.F1; i < internalF2-1; i++ {
					buf += sl[i]
				}
				buf += sl[internalF2-1][:internalC2]
			}
		}
	}
	return buf
}

func (s *Strings) isRuneLess(r1, r2 rune) (eq, less bool) {
	if r1 < r2 {
		less = true
	}
	if r1 == r2 {
		eq = true
	}
	if s.flags.R {
		less = !less
	}
	return
}

func (s *Strings) isIntLess(i1, i2 int) (eq, less bool) {
	if i1 < i2 {
		less = true
	}
	if i1 == i2 {
		eq = true
	}
	if s.flags.R {
		less = !less
	}
	return
}

func (s *Strings) isFloatLess(i1, i2 float64) (eq, less bool) {
	if i1 < i2 {
		less = true
	}
	if i1 == i2 {
		eq = true
	}
	if s.flags.R {
		less = !less
	}
	return
}

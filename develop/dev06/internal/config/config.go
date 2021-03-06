package config

import (
	"flag"
	"strconv"
	"strings"
)

// Config - тип для хранения конфигурации
type Config struct {
	F      string
	D      string
	S      bool
	Ranges []*FRange
}

// NewConfig создает экземпляр Config
func NewConfig() *Config {
	return &Config{}
}

// SetFlags считывает флаги
func (c *Config) SetFlags() {
	flag.StringVar(&c.F, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&c.D, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&c.S, "s", false, "выводить только строки с разделителем")
	flag.Parse()
}

// DecodeFlagF декодирует флаг -f
func (c *Config) DecodeFlagF() error {
	ranges := strings.Split(c.F, ",")
	var decodedRanges []*FRange
	for i := range ranges {

		if v, err := DecodeRange(ranges[i]); err != nil {
			return err
		} else {
			if CheckValidRange(v) {
				decodedRanges = append(decodedRanges, v)
			} else {
				return NewRangeError()
			}
		}
	}
	c.Ranges = decodedRanges
	return nil
}

// FRange - тип, описывающий флаг -f
type FRange struct {
	start     int
	end       int
	fromStart bool
	toEnd     bool
}

// GetStart возвращает начальное поле
func (fr *FRange) GetStart() int {
	return fr.start
}

// GetEnd возвращает конечное поле
func (fr *FRange) GetEnd() int {
	return fr.end
}

// GetFromStart возвращает булево значение - является ли нулевое поле начальным
func (fr *FRange) GetFromStart() bool {
	return fr.fromStart
}

// GetToEnd возвращает булево значение - является ли последнее поле конечным
func (fr *FRange) GetToEnd() bool {
	return fr.toEnd
}

// NewSingleRange возвращает промежуток из одного поля [i]
func NewSingleRange(start int) *FRange {
	return &FRange{
		start: start,
		end:   start,
	}
}

// NewLongRange возвращает промежуток из нескольких полей [i:j]
func NewLongRange(start, end int) *FRange {
	return &FRange{
		start: start,
		end:   end,
	}
}

// NewToEndRange возвращает промежуток, концом которого является последнее поле [i:]
func NewToEndRange(start int) *FRange {
	return &FRange{
		start: start,
		toEnd: true,
	}
}

// NewFromStartRange возвращает промежуток, началом которого является нулевое поле [:j]
func NewFromStartRange(end int) *FRange {
	return &FRange{
		end:       end,
		fromStart: true,
	}
}

// CheckValidRange проверяет, может ли существовать такой промежуток
func CheckValidRange(fr *FRange) bool {
	if fr.fromStart {
		if !fr.toEnd && fr.start == 0 && fr.end != 0 {
			return true
		}
	}
	if fr.toEnd {
		if !fr.fromStart && fr.start != 0 && fr.end == 0 {
			return true
		}
	}
	if !fr.fromStart && !fr.toEnd {
		if fr.start != 0 && fr.start <= fr.end {
			return true
		}
	}
	return false
}

// DecodeRange декодирует строку в FRange
func DecodeRange(rawRange string) (*FRange, error) {
	r := strings.ReplaceAll(rawRange, " ", "")
	dashCount := strings.Count(r, "-")
	if dashCount > 1 {
		return nil, NewRangeError()
	} else {
		dashIndex := strings.Index(r, "-")
		if dashIndex >= 0 { // если есть тире
			if len(r) > 1 { // если есть что-то кроме тире
				if dashIndex == 0 { // если тире в самом начале
					if v, err := strconv.Atoi(r[1:]); err != nil {
						return nil, err
					} else { // если после него число
						return NewFromStartRange(v), nil // возвращается промежуток S-
					}
				} else if dashIndex == len(r)-1 { // если тире в самом конце
					if v, err := strconv.Atoi(r[:len(r)-1]); err != nil {
						return nil, err
					} else { // если перед ним число
						return NewToEndRange(v), nil // возвращается промежуток -F
					}
				} else { // если тире где-то посередине
					if v1, err1 := strconv.Atoi(r[:dashIndex]); err1 != nil {
						return nil, err1
					} else { // если перед ним число
						if v2, err2 := strconv.Atoi(r[dashIndex+1:]); err2 != nil {
							return nil, err2
						} else { // и если после него число
							return NewLongRange(v1, v2), nil // возвращается промежуток S-F
						}
					}
				}
			} else { // если там только тире
				return nil, NewRangeError()
			}
		} else { // если тире нет
			if v, err := strconv.Atoi(r); err != nil {
				return nil, err
			} else {
				return NewSingleRange(v), nil // возвращается одинарный промежуток (S-S)
			}
		}
	}
}

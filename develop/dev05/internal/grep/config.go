package grep

import (
	"flag"
)

type Config struct {
	FlagA   int
	FlagB   int
	FlagC   int
	Flagc   bool
	Flagi   bool
	Flagv   bool
	FlagF   bool
	Flagn   bool
	Pattern string
	Files   []string
}

func NewConfig() *Config {
	return &Config{}
}
func (f *Config) SetConfig() error {
	flag.IntVar(&f.FlagA, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&f.FlagB, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&f.FlagC, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&f.Flagc, "c", false, "печатать количество строк")
	flag.BoolVar(&f.Flagi, "i", false, "игнорировать регистр")
	flag.BoolVar(&f.Flagv, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&f.FlagF, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&f.Flagn, "n", false, "печатать номер строки")
	flag.Parse()
	if err := f.GetRest(flag.Args()); err != nil {
		return err
	}
	return nil
}

func (f *Config) GetRest(a []string) error {
	if len(a) < 1 {
		return NewNoPatternError()
	} else {
		f.Pattern = a[0]
	}
	for i := 1; i < len(a); i++ {
		f.Files = append(f.Files, a[i])
	}
	return nil
}

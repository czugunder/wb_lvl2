package sort

import (
	"io"
	"os"
)

// Flags флаги конфигурации
type Flags struct {
	K      string
	N      bool
	R      bool
	U      bool
	M      bool
	B      bool
	C      bool
	H      bool
	KFlag  *KFlag
	Writer io.Writer
}

// NewFlags функция создания экземпляра Flags
func NewFlags() *Flags {
	return &Flags{
		KFlag:  NewKFlag(),
		Writer: os.Stdout,
	}
}

// SetWriter устанавливает поток вывода данных
func (f *Flags) SetWriter(w io.Writer) {
	f.Writer = w
}

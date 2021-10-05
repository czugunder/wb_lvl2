package sort

import (
	"flag"
)

// Sorting - тип сортировки, считывает аргументы и запускает сортировку по файлам
type Sorting struct {
	flags     *Flags
	filePaths []string
}

// NewSorting создает экземпляр sorting
func NewSorting() *Sorting {
	return &Sorting{
		flags: NewFlags(),
	}
}

// SetFlags устанавливает флаги
func (s *Sorting) SetFlags(f *Flags) error {
	if err := f.KFlag.DecodeKFlag(f.K); err != nil {
		return err
	}
	s.flags = f
	return nil
}

// SetFiles устанавливает пути к файлам для сортировки
func (s *Sorting) SetFiles(paths []string) {
	s.filePaths = paths
}

// ReadFlags считывает флаги
func (s *Sorting) ReadFlags() error {
	flag.StringVar(&s.flags.K, "k", "", "указание колонки для сортировки")
	flag.BoolVar(&s.flags.N, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&s.flags.R, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&s.flags.U, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&s.flags.M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&s.flags.B, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&s.flags.C, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&s.flags.H, "h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.Parse()
	err := s.flags.KFlag.DecodeKFlag(s.flags.K)
	if err != nil {
		return err
	}
	return nil
}

// ReadFilePaths пути к файлам для сортировки
func (s *Sorting) ReadFilePaths() {
	s.filePaths = flag.Args()
}

// SortFiles запускает считывание файла, запускает сортировку, запускает запись результата в файл
func (s *Sorting) SortFiles() error {
	var err error
	for _, fp := range s.filePaths {
		tf := NewTextFile(fp, s.flags)
		err = tf.Read()
		if err != nil {
			return err
		}
		tf.SortInternal()
		err = tf.Write()
		if err != nil {
			return err
		}
	}
	return nil
}

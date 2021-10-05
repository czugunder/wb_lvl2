package sort

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// TextFile тип, взаимодействующий с файлом для сортировки
type TextFile struct {
	path    string
	strings *Strings
	flags   *Flags
}

// NewTextFile возвращает экземпляр TextFile
func NewTextFile(path string, fl *Flags) *TextFile {
	return &TextFile{
		path:    path,
		strings: NewSortString(fl),
		flags:   fl,
	}
}

// GetStrings возвращает внутренний буфер (слайс) сортируемых строк из экземпляра структуры Strings
func (t *TextFile) GetStrings() []string {
	return t.strings.GetStrings()
}

// SetStrings заполняет внутренний буфер (слайс) сортируемых строк внутри экземпляра структуры Strings
func (t *TextFile) SetStrings(s *Strings) {
	t.strings = s
}

// Read считывает файл и записывает его построчно в Strings
func (t *TextFile) Read() error {
	var s string
	f, err := os.Open(t.path)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err = r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return err
			}
			if len(s) > 0 {
				t.strings.append(s)
			}
			break
		} else {
			if len(s) > 0 {
				t.strings.append(s[:len(s)-1])
			}
		}
	}
	return nil
}

// Write записывает файл используя строки из Strings
func (t *TextFile) Write() error {
	f, err := os.Create(t.path) // не Open т.к. в файл нужно записать
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for i, s := range t.GetStrings() {
		if i == t.strings.Len()-1 {
			_, err = w.WriteString(s)
		} else {
			_, err = w.WriteString(s + "\n")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// SortInternal обрабатывает внутри флаги -u и -c, а так же запускает сортировку считанных из файлов строк
func (t *TextFile) SortInternal() {
	if t.flags.U {
		t.strings.RemoveDuplicates()
	}
	if t.flags.C {
		if sort.SliceIsSorted(t.GetStrings(), t.strings.getCustomLess()) {
			t.PrintLine("Input is sorted")
		} else {
			t.PrintLine("Input is NOT sorted")
		}
	} else {
		sort.Stable(t.strings)
	}
}

// PrintLine выводит строку в поток указанный в конфигурации
func (t *TextFile) PrintLine(line string) {
	if _, err := fmt.Fprintln(t.flags.Writer, line); err != nil {
		log.Fatalln(err)
	}
}

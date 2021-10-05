package grep

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
)

type Source struct {
	Config *Config
	Path   string
	wg     *sync.WaitGroup
	lines  []string
}

func NewSource(c *Config, wg *sync.WaitGroup, p string) *Source {
	return &Source{
		Config: c,
		Path:   p,
		wg:     wg,
	}
}

func (s *Source) CheckLine(line string) bool {
	target := true
	if s.Config.Flagv { // обработка флага -v
		target = false
	}
	patternLine := s.Config.Pattern
	compareLine := line
	if s.Config.FlagF { // обработка флага -F
		if s.Config.Flagi { // обработка флага -i
			patternLine = stringToLower(patternLine)
			compareLine = stringToLower(compareLine)
		}
		if stringContains(compareLine, patternLine) {
			return target
		}
	} else {
		if s.Config.Flagi { // обработка флага -i
			patternLine = "(?i)" + patternLine
		}
		if match, _ := regexp.MatchString(patternLine, compareLine); match {
			return target
		}
	}
	return !target
}

func (s *Source) Run() {
	defer s.wg.Done()
	if s.Path == "" {
		if err := s.StdioRun(os.Stdout, os.Stdin); err != nil {
			log.Fatalln(err)
		}
	} else {
		if err := s.FileRun(os.Stdout); err != nil {
			log.Fatalln(err)
		}
	}
}

func (s *Source) FileRun(w io.Writer) error {
	if err := s.readFile(); err != nil {
		return err
	}
	if s.Config.Flagc {
		var counter int
		for _, v := range s.lines {
			if s.CheckLine(v) {
				counter++
			}
		}
		if err := s.PrintCounter(w, counter); err != nil {
			return err
		}
	} else {
		for i, v := range s.lines {
			if s.CheckLine(v) {
				if s.Config.FlagB > 0 || s.Config.FlagC > 0 { // обработка флагов -B и -C для файлов
					if s.Config.FlagB > 0 { // флаг -B имеет больший приоритет, чем -С в оригинальном grep
						if err := s.printPrevious(w, s.Config.FlagB, i); err != nil {
							return err
						}
					} else {
						if err := s.printPrevious(w, s.Config.FlagC, i); err != nil {
							return err
						}
					}
				}
				if err := s.PrintLine(w, i); err != nil {
					return err
				}
				if s.Config.FlagA > 0 || s.Config.FlagC > 0 { // обработка флагов -A и -C для файлов
					if s.Config.FlagA > 0 { // флаг -A имеет больший приоритет, чем -С в оригинальном grep
						if err := s.printFurther(w, s.Config.FlagA, i); err != nil {
							return err
						}
					} else {
						if err := s.printFurther(w, s.Config.FlagC, i); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

func (s *Source) StdioRun(w io.Writer, r io.Reader) error { // флаг -c в оригинальном grep не работает с STDIO
	scanner := bufio.NewScanner(r)
	var index int
	var line string
	furtherMap := make(map[int]int)
	for scanner.Scan() {
		line = scanner.Text()
		s.AddLine(line)
		if s.Config.FlagA > 0 || s.Config.FlagC > 0 { // обработка флагов -A и -C для STDIO
			if v, ok := furtherMap[index]; ok {
				for i := 0; i < v; i++ {
					if err := s.PrintLine(w, index); err != nil {
						return err
					}
				}
			}
		}
		if s.CheckLine(line) {
			if s.Config.FlagB > 0 || s.Config.FlagC > 0 { // обработка флагов -B и -C для STDIO
				if s.Config.FlagB > 0 { // флаг -B имеет больший приоритет, чем -С в оригинальном grep
					if err := s.printPrevious(w, s.Config.FlagB, index); err != nil {
						return err
					}
				} else {
					if err := s.printPrevious(w, s.Config.FlagC, index); err != nil {
						return err
					}
				}
			}
			if err := s.PrintLine(w, index); err != nil {
				return err
			}
			if s.Config.FlagA > 0 || s.Config.FlagC > 0 { // обработка флагов -A и -C для STDIO
				if s.Config.FlagA > 0 { // флаг -A имеет больший приоритет, чем -С в оригинальном grep
					setFurtherMap(furtherMap, s.Config.FlagA, index)
				} else {
					setFurtherMap(furtherMap, s.Config.FlagC, index)
				}
			}
		}
		index++
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}

func (s *Source) readFile() error {
	if v, err := ReadFile(s.Path); err != nil {
		return err
	} else {
		s.lines = v
	}
	return nil
}

func (s *Source) printPrevious(w io.Writer, n, i int) error {
	for j := i - n; j < i; j++ {
		if err := s.PrintLine(w, j); err != nil {
			return err
		}
	}
	return nil
}

func (s *Source) printFurther(w io.Writer, n, i int) error {
	for j := i + 1; j < n+i+1; j++ {
		if err := s.PrintLine(w, j); err != nil {
			return err
		}
	}
	return nil
}

func (s *Source) PrintLine(w io.Writer, i int) error {
	if len(s.lines) > i && i >= 0 {
		line := s.lines[i]
		if s.Config.Flagn { // обработка флага -n
			line = strconv.Itoa(i+1) + ":" + line
		}
		prefix := "stdio:"
		if s.Path != "" {
			prefix = s.Path + ":"
		}
		if _, err := fmt.Fprintln(w, prefix+line); err != nil {
			return err
		}
	}
	return nil
}

func (s *Source) PrintCounter(w io.Writer, cnt int) error {
	prefix := "stdio:"
	if s.Path != "" {
		prefix = s.Path + ":"
	}
	if _, err := fmt.Fprintln(w, prefix+"COUNT:"+strconv.Itoa(cnt)); err != nil {
		return err
	}
	return nil
}

func (s *Source) AddLine(line string) {
	s.lines = append(s.lines, line)
}

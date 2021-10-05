package cut

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"wb_lvl2/develop/dev06/internal/config"
)

// Cut - основной тип в программе
type Cut struct {
	config *config.Config
	writer io.Writer
	reader io.Reader
}

// NewCut создает экземпляр Cut
func NewCut() *Cut {
	return &Cut{}
}

// SetConfig записывает новую конфигурацию в Cut
func (c *Cut) SetConfig(cfg *config.Config) {
	c.config = cfg
}

// SetWriter задает новый поток вывода
func (c *Cut) SetWriter(w io.Writer) {
	c.writer = w
}

// SetReader задает новый поток ввода
func (c *Cut) SetReader(r io.Reader) {
	c.reader = r
}

// Configure инициализирует конфигурацию
func (c *Cut) Configure() error {
	cfg := config.NewConfig()
	cfg.SetFlags()
	if err := cfg.DecodeFlagF(); err != nil {
		return err
	}
	c.SetConfig(cfg)
	c.SetWriter(os.Stdout)
	c.SetReader(os.Stdin)
	return nil
}

// Run запускает программу
func (c *Cut) Run() error {
	scanner := bufio.NewScanner(c.reader)
	for scanner.Scan() {
		line := scanner.Text()
		rLine, sepFound := c.FormatString(line)
		if !c.config.S || (c.config.S && sepFound) { // обработка флага -s
			if _, err := fmt.Fprintln(c.writer, rLine); err != nil {
				return err
			}
		}
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}

// FormatString обрабатывает строку, делит на колонки
func (c *Cut) FormatString(line string) (string, bool) {
	sLine := strings.Split(line, c.config.D)
	var rLine string
	if len(sLine) == 1 {
		return line, false
	}
	pattern := c.MakePattern(len(sLine))
	for i, v := range pattern {
		if v {
			rLine += sLine[i] + c.config.D
		}
	}
	rLine = rLine[:len(rLine)-len(c.config.D)]
	return rLine, true
}

// MakePattern обрабатывает флаг -f
func (c *Cut) MakePattern(sLen int) []bool {
	pattern := make([]bool, sLen, sLen)
	var last int
	for _, r := range c.config.Ranges {
		if r.GetFromStart() {
			if r.GetEnd() > sLen {
				last = sLen
			} else {
				last = r.GetEnd()
			}
			for i := 0; i < last; i++ {
				pattern[i] = true
			}
		} else if r.GetToEnd() {
			for i := r.GetStart() - 1; i < sLen; i++ {
				pattern[i] = true
			}
		} else {
			if r.GetEnd() > sLen {
				last = sLen
			} else {
				last = r.GetEnd()
			}
			for i := r.GetStart() - 1; i < last; i++ {
				pattern[i] = true
			}
		}
	}
	return pattern
}

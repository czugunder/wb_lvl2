package grep

import "sync"

// Grep основной тип фильтрации
type Grep struct {
	config *Config
	wg     *sync.WaitGroup
}

// NewGrep создает экземпляр Grep
func NewGrep() *Grep {
	return &Grep{
		config: NewConfig(),
		wg:     &sync.WaitGroup{},
	}
}

// SetConfig записывает новый Config в Grep
func (g *Grep) SetConfig(c *Config) {
	g.config = c
}

// Run запускает фильтрацию, если указано несколько файлов, то каждый обрабатывается в отдельной рутине
func (g *Grep) Run() {
	if len(g.config.Files) == 0 { // обработка stdin, если не указано ни одного файла
		s := NewSource(g.config, g.wg, "")
		g.wg.Add(1)
		s.Run()
		g.wg.Wait()
	} else { // обработка файлов
		for _, v := range g.config.Files {
			s := NewSource(g.config, g.wg, v)
			g.wg.Add(1)
			s.Run()
		}
		g.wg.Wait()
	}
}

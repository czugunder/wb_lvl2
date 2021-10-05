package grep

import "sync"

type grep struct {
	config *Config
	wg     *sync.WaitGroup
}

func NewGrep() *grep {
	return &grep{
		config: NewConfig(),
		wg:     &sync.WaitGroup{},
	}
}

func (g *grep) SetConfig(c *Config) {
	g.config = c
}

func (g *grep) Run() {
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

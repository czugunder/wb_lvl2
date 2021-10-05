package main

import (
	"log"
	"wb_lvl2/develop/dev08/internal/dev08/shell"
)

func main() {
	s := shell.NewShell()
	s.Configure("roman", "ubuntu")
	if err := s.Run(); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"
	"wb_lvl2/develop/dev05/internal/grep"
)

func main() {
	conf := grep.NewConfig()
	err := conf.SetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	g := grep.NewGrep()
	g.SetConfig(conf)
	g.Run()
}

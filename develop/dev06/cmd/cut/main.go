package main

import (
	"log"
	"wb_lvl2/develop/dev06/internal/cut"
)

func main() {
	c := cut.NewCut()
	if err := c.Configure(); err != nil {
		log.Fatalln(err)
	}
	if err := c.Run(); err != nil {
		log.Fatalln(err)
	}
}

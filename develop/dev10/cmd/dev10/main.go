package main

import (
	"wb_lvl2/develop/dev10/internal/dev10/config"
	"wb_lvl2/develop/dev10/internal/dev10/telnetUtil"
)

func main() {
	c := config.Configure()
	t := telnetUtil.NewTelnet(c)
	t.Run()
}

package main

import "wb_lvl2/develop/dev11/internal/dev11/server"

func main() {
	c := server.DefaultConfig()
	//c := server.NewConfig("localhost:8888")
	s := server.NewSever(c)
	s.Run()
}

package main

import (
	"log"
	"wb_lvl2/develop/dev03/internal/dev03/sort"
)

func main() {
	s := sort.NewSorting()
	if err := s.ReadFlags(); err != nil {
		log.Fatalln(err)
	}
	s.ReadFilePaths()
	if err := s.SortFiles(); err != nil {
		log.Fatalln(err)
	}
}

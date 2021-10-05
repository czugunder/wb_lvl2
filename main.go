package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	//assignment9.Start()
}

func createFolders() {
	for i := 1; i < 12; i++ {
		err := os.Mkdir("/home/roman/GolandProjects/wb_lvl2/develop/dev"+strconv.Itoa(i), 0777)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

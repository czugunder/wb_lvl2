package main

import (
	"flag"
	"log"
	"wb_lvl2/develop/dev09/internal/dev09/wget"
)

func main() {
	var url, path string
	flag.StringVar(&url, "u", "", "URL адрес сайта")
	flag.StringVar(&path, "p", "", "путь к папке, где будет сохранен сайт")
	if url == "" {
		log.Fatalln("directory is not set")
	}
	if path == "" {
		log.Fatalln("save path is not set")
	}
	w := wget.NewWget()
	w.SetSaveDirectory(path)
	w.SetDomain(url)
	w.SaveSite()
}

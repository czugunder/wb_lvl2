package config

import (
	"flag"
	"log"
	"time"
)

type Config struct {
	Timeout time.Duration
	Host    string
	Port    string
}

func Configure() *Config {
	c := &Config{}
	flag.DurationVar(&c.Timeout, "timeout", 0, "telnet timeout")
	flag.Parse()
	left := flag.Args()
	if len(left) != 2 {
		log.Fatalln("incorrect host and port")
	}
	c.Host = left[0]
	c.Port = left[1]
	return c
}

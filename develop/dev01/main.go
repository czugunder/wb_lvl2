package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

func main() {
	if time, err := ntp.Time("0.bee1vik-ntp.pool.ntp.org"); err != nil {
		l := log.New(os.Stderr, "TIMEERR ", 0)
		l.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(time)
	}
}

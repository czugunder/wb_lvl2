package telnetUtil

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wb_lvl2/develop/dev10/internal/dev10/config"
)

type telnetUtil struct {
	config  *config.Config
	conn    net.Conn
	errorCh chan error
}

func NewTelnet(config *config.Config) *telnetUtil {
	return &telnetUtil{
		config:  config,
		errorCh: make(chan error),
	}
}

func (t *telnetUtil) Run() {
	t.connect()
	sigint := make(chan os.Signal)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	go t.receive()
	go t.send()
	select {
	case err := <-t.errorCh:
		log.Println(err)
		t.disconnect()
		return
	case <-sigint:
		t.disconnect()
		return
	}
}

func (t *telnetUtil) connect() {
	conn, err := net.DialTimeout("tcp4", t.config.Host+":"+t.config.Port, t.config.Timeout)
	if err != nil {
		time.Sleep(t.config.Timeout) // при подключении к несуществующему серверу, программа должна завершаться через timeout
		log.Fatalln("connection error")
	}
	t.conn = conn
}

func (t *telnetUtil) disconnect() {
	t.conn.Close()
}

func (t *telnetUtil) receive() {
	r := bufio.NewReader(t.conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil { // закрытие сокета
			t.errorCh <- err
			return
		}
		fmt.Print(line)
	}
}

func (t *telnetUtil) send() {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if err != nil { // ctrl+d (EOF)
			t.errorCh <- err
			return
		}
		_, err = t.conn.Write([]byte(line))
		if err != nil {
			t.errorCh <- err
			return
		}
	}
}

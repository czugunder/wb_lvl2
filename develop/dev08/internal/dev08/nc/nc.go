package nc

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type netcat struct {
	reader io.Reader
	writer io.Writer
}

func NewNC(r io.Reader) *netcat {
	return &netcat{
		reader: r,
	}
}

func (nc *netcat) Run(req []string) error {
	if len(req) != 3 { // первый аргумент -t или -u для определения типа протокола, дальше адрес и порт
		return &ArgsError{}
	}
	if req[0] == "-u" {
		return nc.UDP(req[1] + ":" + req[2])
	} else if req[0] == "-t" {
		return nc.TCP(req[1] + ":" + req[2])
	} else {
		return &NoConnectionType{}
	}
}

func (nc *netcat) TCP(url string) error {
	if s, err := net.ResolveTCPAddr("tcp4", url); err != nil { // это чтоб localhost в 127.0.0.1 красиво превращался
		return err
	} else {
		if c, errD := net.DialTCP("tcp4", nil, s); err != nil { // создание соединения
			return errD
		} else {
			defer c.Close() // отложенное закрытие соединения
			nc.writer = c
			if err = nc.session(); err != nil { // передача данных
				return err
			}
		}
	}
	return nil
}

func (nc *netcat) UDP(url string) error { // тут все аналогично TCP()
	if s, err := net.ResolveUDPAddr("udp4", url); err != nil {
		return err
	} else {
		if c, errD := net.DialUDP("udp4", nil, s); err != nil {
			return errD
		} else {
			defer c.Close()
			nc.writer = c
			if err = nc.session(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (nc *netcat) session() error {
	r := bufio.NewScanner(nc.reader)
	for r.Scan() { // считывание из ридера
		line := r.Text()
		if line == "exit" { // обработка выхода
			break
		}
		if _, err := fmt.Fprint(nc.writer, line+"\n"); err != nil { // отправка данных
			return err
		}
	}
	return nil
}

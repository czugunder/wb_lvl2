package sort

import (
	"errors"
	"strconv"
)

// KFlag тип декодированного из строки флага -k
type KFlag struct {
	F1 int
	C1 int
	F2 int
	C2 int
}

// NewKFlag создает экземпляр KFlag
func NewKFlag() *KFlag {
	return &KFlag{}
}

// DecodeKFlag принимает на вход строку флага -k и декодирует ее
func (k *KFlag) DecodeKFlag(flag string) error {
	var err error
	var buf [4]string
	var j, l int
	for i, v := range flag {
		if v == '.' {
			buf[j] = flag[l:i]
			l = i + 1
			j++
		}
		if v == ',' {
			buf[j] = flag[l:i]
			l = i + 1
			j = 2
		}
		if i == len(flag)-1 {
			buf[j] = flag[l : i+1]
		}
	}
	if buf[0] != "" {
		k.F1, err = strconv.Atoi(buf[0])
		if err != nil {
			return err
		}
		if k.F1 < 1 {
			return errors.New("-k params origin is 1")
		}
	}
	if buf[1] != "" {
		k.C1, err = strconv.Atoi(buf[1])
		if err != nil {
			return err
		}
		if k.C1 < 1 {
			return errors.New("-k params origin is 1")
		}
	}
	if buf[2] != "" {
		k.F2, err = strconv.Atoi(buf[2])
		if err != nil {
			return err
		}
		if k.F2 < 1 {
			return errors.New("-k params origin is 1")
		}
	}
	if buf[3] != "" {
		k.C2, err = strconv.Atoi(buf[3])
		if err != nil {
			return err
		}
		if k.C2 < 1 {
			return errors.New("-k params origin is 1")
		}
	}
	return nil
}

package shell

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"wb_lvl2/develop/dev08/internal/dev08/shell"
)

// test pwd
func TestShell_Run1(t *testing.T) {
	readBuff := bytes.Buffer{}
	writeBuff := bytes.Buffer{}
	s := shell.NewShell()
	s.SetUserName("test")
	s.SetSystemName("test")
	s.SetReader(&readBuff)
	s.SetWriter(&writeBuff)
	currentDirectory, errPrep := os.Getwd()
	if errPrep != nil {
		t.Fatal("can't run test, current directory is not resolved")
	}
	exp := []string{currentDirectory, ""}
	shellPrefix := "\u001B[31m[test@test "

	readBuff.WriteString("pwd")
	err := s.Run()
	resBuff := writeBuff.String()
	res := strings.Split(resBuff, "\n")

	if err != nil {
		t.Fatal("incorrect work of pwd command")
	}
	if len(exp) != len(res) {
		t.Fatalf("incorrect length of output, expected: %d, got: %d", len(exp), len(res))
	}
	for i := range exp {
		if strings.Index(res[i], shellPrefix) != 0 {
			t.Fatalf("shell prefix is missing in string with %d index: %s", i, res[i])
		}
		if !strings.Contains(res[i], exp[i]) {
			t.Fatalf("output: <%s> doesn't contain: <%s>", res[i], exp[i])
		}
	}
}

// test echo
func TestShell_Run2(t *testing.T) {
	readBuff := bytes.Buffer{}
	writeBuff := bytes.Buffer{}
	s := shell.NewShell()
	s.SetUserName("test")
	s.SetSystemName("test")
	s.SetReader(&readBuff)
	s.SetWriter(&writeBuff)
	exp := []string{"123", ""}
	shellPrefix := "\u001B[31m[test@test "

	readBuff.WriteString("echo 123")
	err := s.Run()
	resBuff := writeBuff.String()
	res := strings.Split(resBuff, "\n")

	if err != nil {
		t.Fatal("incorrect work of echo command")
	}
	if len(exp) != len(res) {
		t.Fatalf("incorrect length of output, expected: %d, got: %d", len(exp), len(res))
	}
	for i := range exp {
		if strings.Index(res[i], shellPrefix) != 0 {
			t.Fatalf("shell prefix is missing in string with %d index: %s", i, res[i])
		}
		if !strings.Contains(res[i], exp[i]) {
			t.Fatalf("output: <%s> doesn't contain: <%s>", res[i], exp[i])
		}
	}
}

// test conveyor
func TestShell_Run3(t *testing.T) {
	readBuff := bytes.Buffer{}
	writeBuff := bytes.Buffer{}
	s := shell.NewShell()
	s.SetUserName("test")
	s.SetSystemName("test")
	s.SetReader(&readBuff)
	s.SetWriter(&writeBuff)
	currentDirectory, errPrep := os.Getwd()
	if errPrep != nil {
		t.Fatal("can't run test, current directory is not resolved")
	}
	exp := []string{currentDirectory, "123 " + currentDirectory, ""}

	readBuff.WriteString("pwd | echo 123")
	err := s.Run()
	resBuff := writeBuff.String()
	res := strings.Split(resBuff, "\n")

	if err != nil {
		t.Fatal("incorrect work of pwd command")
	}
	if len(exp) != len(res) {
		t.Fatalf("incorrect length of output, expected: %d, got: %d", len(exp), len(res))
	}
	for i := range exp {
		if !strings.Contains(res[i], exp[i]) {
			t.Fatalf("output: <%s> doesn't contain: <%s>", res[i], exp[i])
		}
	}
}

package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mitchellh/go-ps"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	netcat "wb_lvl2/develop/dev08/internal/dev08/nc"
)

type shell struct {
	reader     io.Reader
	writer     io.Writer
	userName   string
	systemName string
	run        bool
	pipeMode   bool
	pipeBuffer *bytes.Buffer
}

func NewShell() *shell {
	return &shell{
		run: true,
	}
}

func (s *shell) Configure(username, sysname string) {
	s.SetReader(os.Stdin)
	s.SetWriter(os.Stdout)
	s.SetSystemName(sysname)
	s.SetUserName(username)
}

func (s *shell) SetReader(r io.Reader) {
	s.reader = r
}

func (s *shell) SetWriter(w io.Writer) {
	s.writer = w
}

func (s *shell) SetUserName(un string) {
	s.userName = un
}

func (s *shell) SetSystemName(sn string) {
	s.systemName = sn
}

func (s *shell) Run() error {
	if err := s.readLines(); err != nil {
		return err
	}
	return nil
}

func (s *shell) readLines() error {
	b := bufio.NewScanner(s.reader)
	if err := s.printPrefix(); err != nil {
		return err
	}
	var line string
	for b.Scan() {
		line = strings.TrimSuffix(b.Text(), "\n")
		if err := s.forkHandler(line); err != nil {
			return err
		}
		if s.run != true {
			break
		}
		if err := s.printPrefix(); err != nil {
			return err
		}
	}
	if b.Err() != nil {
		return b.Err()
	}
	return nil
}

func (s *shell) forkHandler(line string) error {
	calls := strings.Split(line, "&")
	if len(calls) == 1 { // форка нет
		if err := s.pipeHandler(calls[0]); err != nil {
			return err
		}
	} else { // форк есть
		for _, v := range calls {
			if isParent, _, errS := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0); errS != 0 {
				s.exitNonZero()
			} else {
				if isParent == 0 { // потомок
					if err := s.pipeHandler(v); err != nil {
						s.exitNonZero()
					}
					s.exitZero()
				}
			}
		}
	}
	return nil
}

func (s *shell) pipeHandler(call string) error {
	assignments := strings.Split(call, "|")
	if len(assignments) > 1 {
		s.pipeBuffer = &bytes.Buffer{}
		s.pipeMode = true
		for _, v := range assignments {
			if err := s.selector(v + s.pipeBuffer.String()); err != nil {
				return err
			}
		}
		s.pipeMode = false
		s.pipeBuffer.Reset()
	} else {
		if err := s.selector(call); err != nil {
			return err
		}
	}
	return nil
}

func (s *shell) selector(l string) error {
	request := strings.Fields(l)
	if len(request) > 0 {
		switch request[0] {
		case "cd":
			if len(request) != 2 {
				if err := s.printLine("incorrect cd syntax"); err != nil {
					return err
				}
			} else {
				if err := s.cd(request[1]); err != nil {
					return err
				}
			}
		case "pwd":
			if path, err := s.pwd(); err != nil {
				return err
			} else {
				if err = s.printLine(path); err != nil {
					return err
				}
			}
		case "echo":
			if err := s.echo(request[1:]); err != nil {
				return err
			}
		case "kill":
			if errS := s.kill(request[1:]); errS != nil {
				for _, e := range errS {
					if err := s.printLine(e.Error()); err != nil {
						return err
					}
				}
			}
		case "ps":
			if prc, err := s.ps(); err != nil {
				return err
			} else {
				if err = s.printLine("PID\tCMD"); err != nil {
					return err
				}
				for _, r := range prc {
					if err = s.printLine(fmt.Sprintf("%d\t%s", r.Pid(), r.Executable())); err != nil {
						return err
					}
				}
			}
		case "nc":
			if err := s.netcat(request[1:]); err != nil {
				return err
			}
		case "exec":
			if err := s.exec(request[1:]); err != nil {
				return err
			}
		case "exit":
			s.exit()
		}
	}
	return nil
}

func (s *shell) printLine(l string) error {
	if _, errP := fmt.Fprint(s.writer, "["+strconv.Itoa(syscall.Getpid())+"] "+l+"\n"); errP != nil {
		return errP
	}
	if s.pipeMode {
		if _, errP := fmt.Fprint(s.pipeBuffer, " "+l); errP != nil {
			return errP
		}
	}
	return nil
}

func (s *shell) printPrefix() error {
	if currentDir, errPWD := s.pwd(); errPWD != nil {
		return errPWD
	} else {
		currentDirShards := strings.Split(currentDir, "/")
		currentDir = currentDirShards[len(currentDirShards)-1]
		prefix := "\033[31m" + "[" + s.userName + "@" + s.systemName + " " + currentDir + "]$ \033[37m"
		if _, errP := fmt.Fprint(s.writer, prefix); errP != nil {
			return errP
		}
	}
	return nil
}

func (s *shell) cd(path string) error {
	if err := os.Chdir(path); err != nil {
		return err
	}
	return nil
}

func (s *shell) pwd() (string, error) {
	if dir, err := os.Getwd(); err != nil {
		return "", err
	} else {
		return dir, nil
	}
}

func (s *shell) echo(args []string) error {
	if err := s.printLine(strings.Join(args, " ")); err != nil {
		return err
	}
	return nil
}

func (s *shell) kill(args []string) []error {
	var errs []error
	for _, potentialPID := range args {
		if PID, err := strconv.Atoi(potentialPID); err != nil {
			errs = append(errs, err)
		} else {
			if err = syscall.Kill(PID, syscall.SIGKILL); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return nil
}

func (s *shell) ps() ([]ps.Process, error) {
	return ps.Processes()
}

func (s *shell) netcat(args []string) error {
	nc := netcat.NewNC(s.reader)
	if err := nc.Run(args); err != nil {
		return err
	}
	return nil
}

func (s *shell) exec(args []string) error {
	if len(args) > 0 {
		name := args[0]
		arg := strings.Join(args[1:], " ")
		cmd := exec.Command(name, arg)
		if err := cmd.Run(); err != nil {
			return err
		}
	} else {
		return &NoExec{}
	}
	return nil
}

func (s *shell) exit() {
	s.run = false
}

func (s *shell) exitZero() {
	os.Exit(0)
}

func (s *shell) exitNonZero() {
	os.Exit(2)
}
